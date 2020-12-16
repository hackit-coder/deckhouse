/*

User-stories:
1. There are module settings. They must be exported via Secret d8-node-manager-cloud-provider.
2. There are applications which must be deployed — cloud-controller-manager, csi.
3. There is list of datastores in values.yaml. StorageClass must be created for every datastore. Datastore mentioned in value `.storageClass.default` must be annotated as default.

*/

package template_tests

import (
	"encoding/base64"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/helm"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "")
}

const globalValues = `
  enabledModules: ["vertical-pod-autoscaler-crd"]
  modules:
    placement: {}
  modulesImages:
    registry: registry.flant.com
    registryDockercfg: cfg
    tags:
      common:
        csiExternalProvisioner116: imagehash
        csiExternalAttacher116: imagehash
        csiExternalProvisioner119: imagehash
        csiExternalAttacher119: imagehash
        csiExternalResizer: imagehash
        csiNodeDriverRegistrar: imagehash
      cloudProviderVsphere:
        vsphereCsi: imagehash
        cloudControllerManager116: imagehash
        cloudControllerManager119: imagehash
  discovery:
    d8SpecificNodeCountByRole:
      worker: 1
      master: 3
    nodeCountByType:
      cloud: 1
    podSubnet: 10.0.1.0/16
    kubernetesVersion: 1.16.4
`

const moduleValuesA = `
    internal:
      storageClasses:
      - name: mydsname1
        path: /my/ds/path/mydsname1
        zones: ["zonea", "zoneb"]
      - name: mydsname2
        path: /my/ds/path/mydsname2
        zones: ["zonea", "zoneb"]
      server: myhost
      username: myuname
      password: myPaSsWd
      insecure: true
      regionTagCategory: myregtagcat
      zoneTagCategory: myzonetagcat
      region: myreg
      sshKey: mysshkey1
      vmFolderPath: dev/test
      datacenter: X1
      zones: ["aaa", "bbb"]
      masterInstanceClass:
        datastore: dev/lun_1
        mainNetwork: k8s-msk/test_187
        memory: 8192
        numCPUs: 4
        template: dev/golden_image
`

const moduleValuesB = `
    internal:
      storageClasses:
      - name: mydsname1
        path: /my/ds/path/mydsname1
        zones: ["zonea", "zoneb"]
      - name: mydsname2
        path: /my/ds/path/mydsname2
        zones: ["zonea", "zoneb"]
      server: myhost
      username: myuname
      password: myPaSsWd
      insecure: true
      regionTagCategory: myregtagcat
      zoneTagCategory: myzonetagcat
      region: myreg
      sshKey: mysshkey1
      vmFolderPath: dev/test
      datacenter: X1
      zones: ["aaa", "bbb"]
      masterInstanceClass: null
`

var _ = Describe("Module :: cloud-provider-vsphere :: helm template ::", func() {
	f := SetupHelmConfig(``)

	Context("Vsphere", func() {
		BeforeEach(func() {
			f.ValuesSetFromYaml("global", globalValues)
			f.ValuesSetFromYaml("cloudProviderVsphere", moduleValuesA)
			f.HelmRender()
		})

		It("Everything must render properly", func() {
			Expect(f.RenderError).ShouldNot(HaveOccurred())

			namespace := f.KubernetesGlobalResource("Namespace", "d8-cloud-provider-vsphere")
			registrySecret := f.KubernetesResource("Secret", "d8-cloud-provider-vsphere", "deckhouse-registry")

			providerRegistrationSecret := f.KubernetesResource("Secret", "kube-system", "d8-node-manager-cloud-provider")

			csiCongrollerPluginSS := f.KubernetesResource("StatefulSet", "d8-cloud-provider-vsphere", "csi-controller")
			csiDriver := f.KubernetesGlobalResource("CSIDriver", "vsphere.csi.vmware.com")
			csiNodePluginDS := f.KubernetesResource("DaemonSet", "d8-cloud-provider-vsphere", "csi-node")
			csiSA := f.KubernetesResource("ServiceAccount", "d8-cloud-provider-vsphere", "csi")
			csiProvisionerCR := f.KubernetesGlobalResource("ClusterRole", "d8:cloud-provider-vsphere:csi:controller:external-provisioner")
			csiProvisionerCRB := f.KubernetesGlobalResource("ClusterRoleBinding", "d8:cloud-provider-vsphere:csi:controller:external-provisioner")
			csiAttacherCR := f.KubernetesGlobalResource("ClusterRole", "d8:cloud-provider-vsphere:csi:controller:external-attacher")
			csiAttacherCRB := f.KubernetesGlobalResource("ClusterRoleBinding", "d8:cloud-provider-vsphere:csi:controller:external-attacher")
			csiResizerCR := f.KubernetesGlobalResource("ClusterRole", "d8:cloud-provider-vsphere:csi:controller:external-resizer")
			csiResizerCRB := f.KubernetesGlobalResource("ClusterRoleBinding", "d8:cloud-provider-vsphere:csi:controller:external-resizer")

			ccmSA := f.KubernetesResource("ServiceAccount", "d8-cloud-provider-vsphere", "cloud-controller-manager")
			ccmCR := f.KubernetesGlobalResource("ClusterRole", "d8:cloud-provider-vsphere:cloud-controller-manager")
			ccmCRB := f.KubernetesGlobalResource("ClusterRoleBinding", "d8:cloud-provider-vsphere:cloud-controller-manager")
			ccmVPA := f.KubernetesResource("VerticalPodAutoscaler", "d8-cloud-provider-vsphere", "cloud-controller-manager")
			ccmDeploy := f.KubernetesResource("Deployment", "d8-cloud-provider-vsphere", "cloud-controller-manager")
			ccmSecret := f.KubernetesResource("Secret", "d8-cloud-provider-vsphere", "cloud-controller-manager")

			userAuthzUser := f.KubernetesGlobalResource("ClusterRole", "d8:user-authz:cloud-provider-vsphere:user")
			userAuthzClusterAdmin := f.KubernetesGlobalResource("ClusterRole", "d8:user-authz:cloud-provider-vsphere:cluster-admin")

			Expect(namespace.Exists()).To(BeTrue())
			Expect(registrySecret.Exists()).To(BeTrue())
			Expect(userAuthzUser.Exists()).To(BeTrue())
			Expect(userAuthzClusterAdmin.Exists()).To(BeTrue())

			// user story #1
			Expect(providerRegistrationSecret.Exists()).To(BeTrue())
			expectedProviderRegistrationJSON := `{
          "server": "myhost",
          "insecure": true,
          "password": "myPaSsWd",
          "region": "myreg",
          "regionTagCategory": "myregtagcat",
          "instanceClassDefaults": {
            "datastore": "dev/lun_1",
            "template": "dev/golden_image",
            "disableTimesync": true
          },
          "sshKey": "mysshkey1",
          "username": "myuname",
          "vmFolderPath": "dev/test",
          "zoneTagCategory": "myzonetagcat"
        }`
			providerRegistrationData, err := base64.StdEncoding.DecodeString(providerRegistrationSecret.Field("data.vsphere").String())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(providerRegistrationData)).To(MatchJSON(expectedProviderRegistrationJSON))

			// user story #2
			Expect(csiDriver.Exists()).To(BeTrue())
			Expect(csiNodePluginDS.Exists()).To(BeTrue())
			Expect(csiSA.Exists()).To(BeTrue())
			Expect(csiCongrollerPluginSS.Exists()).To(BeTrue())
			Expect(csiAttacherCR.Exists()).To(BeTrue())
			Expect(csiAttacherCRB.Exists()).To(BeTrue())
			Expect(csiProvisionerCR.Exists()).To(BeTrue())
			Expect(csiProvisionerCRB.Exists()).To(BeTrue())
			Expect(csiResizerCR.Exists()).To(BeTrue())
			Expect(csiResizerCRB.Exists()).To(BeTrue())
			Expect(csiResizerCR.Exists()).To(BeTrue())
			Expect(csiResizerCRB.Exists()).To(BeTrue())

			Expect(ccmSA.Exists()).To(BeTrue())
			Expect(ccmCR.Exists()).To(BeTrue())
			Expect(ccmCRB.Exists()).To(BeTrue())
			Expect(ccmVPA.Exists()).To(BeTrue())
			Expect(ccmDeploy.Exists()).To(BeTrue())
			Expect(ccmSecret.Exists()).To(BeTrue())

			// user story #3
			scMydsname1 := f.KubernetesGlobalResource("StorageClass", "mydsname1")
			scMydsname2 := f.KubernetesGlobalResource("StorageClass", "mydsname2")

			Expect(scMydsname1.Exists()).To(BeTrue())
			Expect(scMydsname2.Exists()).To(BeTrue())

			Expect(scMydsname1.Field("metadata.annotations").String()).To(MatchYAML(`
storageclass.kubernetes.io/is-default-class: "true"
`))
			Expect(scMydsname2.Field("metadata.annotations").Exists()).To(BeFalse())
		})
	})

	Context("Vsphere", func() {
		BeforeEach(func() {
			f.ValuesSetFromYaml("global", globalValues)
			f.ValuesSetFromYaml("cloudProviderVsphere", moduleValuesB)
			f.HelmRender()
		})

		It("Everything must render properly", func() {
			Expect(f.RenderError).ShouldNot(HaveOccurred())

			providerRegistrationSecret := f.KubernetesResource("Secret", "kube-system", "d8-node-manager-cloud-provider")
			Expect(providerRegistrationSecret.Exists()).To(BeTrue())
			expectedProviderRegistrationJSON := `{
          "server": "myhost",
          "insecure": true,
          "password": "myPaSsWd",
          "region": "myreg",
          "regionTagCategory": "myregtagcat",
          "instanceClassDefaults": {},
          "sshKey": "mysshkey1",
          "username": "myuname",
          "vmFolderPath": "dev/test",
          "zoneTagCategory": "myzonetagcat"
        }`

			providerRegistrationData, err := base64.StdEncoding.DecodeString(providerRegistrationSecret.Field("data.vsphere").String())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(providerRegistrationData)).To(MatchJSON(expectedProviderRegistrationJSON))
		})

		Context("Unsupported Kubernetes version", func() {
			BeforeEach(func() {
				f.ValuesSetFromYaml("global", globalValues)
				f.ValuesSetFromYaml("cloudProviderVsphere", moduleValuesA)
				f.ValuesSet("global.discovery.kubernetesVersion", "1.17.8")
				f.HelmRender()
			})

			It("CCM and CSI controller should not be present on unsupported Kubernetes versions", func() {
				Expect(f.RenderError).ShouldNot(HaveOccurred())
				Expect(f.KubernetesResource("Deployment", "d8-cloud-provider-vsphere", "cloud-controller-manager").Exists()).To(BeFalse())
				Expect(f.KubernetesResource("StatefulSet", "d8-cloud-provider-vsphere", "csi-controller").Exists()).To(BeFalse())

			})
		})
	})

	Context("Vsphere with default StorageClass specified", func() {
		BeforeEach(func() {
			f.ValuesSetFromYaml("global", globalValues)
			f.ValuesSetFromYaml("cloudProviderVsphere", moduleValuesB)
			f.ValuesSetFromYaml("cloudProviderVsphere.internal.defaultStorageClass", `mydsname2`)
			f.HelmRender()
		})

		It("Everything must render properly with proper default StorageClass", func() {
			Expect(f.RenderError).ShouldNot(HaveOccurred())

			scMydsname1 := f.KubernetesGlobalResource("StorageClass", "mydsname1")
			scMydsname2 := f.KubernetesGlobalResource("StorageClass", "mydsname2")

			Expect(scMydsname1.Exists()).To(BeTrue())
			Expect(scMydsname2.Exists()).To(BeTrue())

			Expect(scMydsname1.Field("metadata.annotations").Exists()).To(BeFalse())
			Expect(scMydsname2.Field("metadata.annotations").String()).To(MatchYAML(`
storageclass.kubernetes.io/is-default-class: "true"
`))
		})
	})
})
