/*
Copyright 2021 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

var registryTemplate = `// Code generated by "register.go" DO NOT EDIT.
package main

import (
	_ "github.com/flant/addon-operator/sdk"
{{ range $value := . }}
	_ "{{ $value }}"
{{- end }}
)
`

func cwd() string {
	_, f, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get caller")
	}

	dir, err := filepath.Abs(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ { // ../../
		dir = filepath.Dir(dir)
	}

	// If deckhouse repo directory is symlinked (e.g. to /deckhouse), resolve the real path.
	// Otherwise, filepath.Walk will ignore all subdirectories.
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		panic(err)
	}

	return dir
}

type edition struct {
	Name       string `yaml:"name,omitempty"`
	ModulesDir string `yaml:"modulesDir,omitempty"`
}

type editions struct {
	Editions []edition `yaml:"editions,omitempty"`
}

func searchHooks(hookModules *[]string, dir, workDir string) error {
	files := make(map[string]interface{})

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f != nil && f.IsDir() {
			if f.Name() == "internal" {
				return filepath.SkipDir
			}
			if f.Name() == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		trimDir := workDir
		moduleName := filepath.Join(
			deckhouseModuleName,
			filepath.Dir(
				strings.TrimPrefix(path, trimDir),
			),
		)
		files[moduleName] = struct{}{}
		return nil
	})

	for hook := range files {
		*hookModules = append(*hookModules, hook)
	}

	return err
}

const deckhouseModuleName = "github.com/deckhouse/deckhouse/"

func main() {
	workDir := cwd()

	var (
		output string
		stream = os.Stdout
	)
	flag.StringVar(&output, "output", "", "output file for generated code")
	flag.Parse()

	if output != "" {
		var err error
		stream, err = os.Create(output)
		if err != nil {
			panic(err)
		}

		defer stream.Close()
	}
	content, err := os.ReadFile(workDir + "/editions.yaml")
	if err != nil {
		panic(fmt.Sprintf("cannot read editions file: %v", err))
	}

	e := editions{}
	err = yaml.Unmarshal(content, &e)
	if err != nil {
		panic(fmt.Errorf("cannot unmarshal editions file: %v", err))
	}

	for i, ed := range e.Editions {
		if ed.Name == "" {
			panic(fmt.Sprintf("name for %d index is empty", i))
		}
	}

	var hookModules []string
	if err := searchHooks(&hookModules, filepath.Join(workDir, "global-hooks"), workDir); err != nil {
		panic(err)
	}

	moduleDirs := make([]string, 0)
	for _, ed := range e.Editions {
		additionalModuleDirs, err := filepath.Glob(filepath.Join(workDir, fmt.Sprintf("%s/*/hooks", ed.ModulesDir)))
		if err != nil {
			panic(err)
		}
		moduleDirs = append(moduleDirs, additionalModuleDirs...)
	}

	moduleDirs = append(moduleDirs, requirementCheckDirs(workDir, e)...)

	for _, dir := range moduleDirs {
		if err := searchHooks(&hookModules, dir, workDir); err != nil {
			panic(err)
		}
	}

	sort.Strings(hookModules)

	t := template.New("registry")
	t, err = t.Parse(registryTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(stream, hookModules)
	if err != nil {
		panic(err)
	}
}

func requirementCheckDirs(workDir string, e editions) []string {
	moduleDirs := make([]string, 0)
	for _, ed := range e.Editions {
		additionalModuleDirs, err := filepath.Glob(filepath.Join(workDir, fmt.Sprintf("%s/*/requirements", ed.ModulesDir)))
		if err != nil {
			panic(err)
		}
		moduleDirs = append(moduleDirs, additionalModuleDirs...)
	}

	return moduleDirs
}
