#!/bin/bash

# Copyright 2024 Flant JSC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

for i in $(find . -regex '.*.md' -print); do
  grep -q "^---" $i
  if [ $? -gt 0 ]; then continue; fi
  cat $i | tr -d '\n' | grep -lv "^---.*permalink: .*---" &> /dev/null
  if [ $? -eq 0 ]; then
    # permalink is absent, add permalink
    PERMALINK="/$(echo $i | sed -E 's#(modules_)(en|ru)/#\2/modules/#' | sed 's#docs/##g'| tr '[:upper:]' '[:lower:]' | sed 's#\.md$#.html#' | sed 's#^\.\/##' | sed 's#readme\.html$##' )"
    sed -i "1apermalink: $PERMALINK" $i
  fi
done
