// Code generated by "tools/audit_policy.go" DO NOT EDIT.
package hooks

var auditPolicyBasicNamespaces = []string{
	"d8-admission-policy-engine",
	"d8-cdi",
	"d8-ceph-csi",
	"d8-cert-manager",
	"d8-chrony",
	"d8-cloud-instance-manager",
	"d8-cloud-provider-aws",
	"d8-cloud-provider-azure",
	"d8-cloud-provider-gcp",
	"d8-cloud-provider-openstack",
	"d8-cloud-provider-vsphere",
	"d8-cloud-provider-yandex",
	"d8-cni-cilium",
	"d8-cni-flannel",
	"d8-cni-simple-bridge",
	"d8-delivery",
	"d8-descheduler",
	"d8-flant-integration",
	"d8-ingress-nginx",
	"d8-istio",
	"d8-keepalived",
	"d8-linstor",
	"d8-local-path-provisioner",
	"d8-log-shipper",
	"d8-metallb",
	"d8-monitoring",
	"d8-multitenancy-manager",
	"d8-network-gateway",
	"d8-okmeter",
	"d8-openvpn",
	"d8-operator-prometheus",
	"d8-operator-trivy",
	"d8-pod-reloader",
	"d8-runtime-audit-engine",
	"d8-snapshot-controller",
	"d8-system",
	"d8-upmeter",
	"d8-user-authn",
	"d8-user-authz",
	"d8-virtualization",
	"kube-system",
}
var auditPolicyBasicServiceAccounts = []string{
	"agent",
	"alertmanager-internal",
	"alerts-receiver",
	"alliance-ingressgateway",
	"alliance-metadata-exporter",
	"cainjector",
	"cert-exporter",
	"cert-manager",
	"cloud-controller-manager",
	"cloud-data-discoverer",
	"cloud-metrics-exporter",
	"cluster-autoscaler",
	"cni-flannel",
	"cni-simple-bridge",
	"control-plane-proxy",
	"controller",
	"d8-control-plane-manager",
	"d8-kube-dns",
	"d8-kube-proxy",
	"d8-node-local-dns",
	"d8-vertical-pod-autoscaler-admission-controller",
	"d8-vertical-pod-autoscaler-recommender",
	"d8-vertical-pod-autoscaler-updater",
	"dashboard",
	"deckhouse",
	"descheduler",
	"dex",
	"early-oom",
	"ebpf-exporter",
	"events-exporter",
	"extended-monitoring-exporter",
	"grafana",
	"image-availability-exporter",
	"ingress-gateway-controller",
	"ingress-nginx",
	"kiali",
	"kruise",
	"kube-state-metrics",
	"linstor-affinity-controller",
	"linstor-controller",
	"linstor-node",
	"linstor-pools-importer",
	"linstor-scheduler",
	"linstor-scheduler-admission",
	"local-path-provisioner",
	"log-shipper",
	"machine-controller-manager",
	"monitoring-ping",
	"multicluster-api-proxy",
	"network-policy-engine",
	"node-exporter",
	"node-group",
	"node-termination-handler",
	"okmeter",
	"openvpn",
	"operator",
	"operator-prometheus",
	"operator-trivy",
	"piraeus-operator",
	"pod-reloader",
	"pricing",
	"prometheus",
	"relay",
	"report-updater",
	"runtime-audit-engine",
	"smoke-mini",
	"snapshot-controller",
	"speaker",
	"terraform-auto-converger",
	"terraform-state-exporter",
	"trickster",
	"ui",
	"upmeter",
	"upmeter-agent",
	"vmi-router",
	"webhook",
	"webhook-handler",
}
