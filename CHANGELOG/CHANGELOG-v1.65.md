# Changelog v1.65

## [MALFORMED]


 - #9390 unknown section "autoscaling"
 - #9703 unknown section "moniotring"
 - #9852 missing section, missing summary, missing type, unknown section ""
 - #9944 unknown section "docs, documentation"
 - #9966 unknown section "cloud-provider-openstack,cloud-provider-vcd,cloud-provider-yandex,cloud-provider-vsphere,cloud-provider-zvirt,istio"
 - #9973 unknown section "ci, docs"

## Know before update


 - Restart containerd.

## Features


 - **[candi]** Install CA certificates on nodes using d8-ca-updater, which is installed from the registrypackages. [#9246](https://github.com/deckhouse/deckhouse/pull/9246)
 - **[candi]** Update containerd to 1.7.20. [#9246](https://github.com/deckhouse/deckhouse/pull/9246)
    Restart containerd.
 - **[cni-cilium]** Adding support for configuring each node individually using CiliumNodeConfig resources. [#9754](https://github.com/deckhouse/deckhouse/pull/9754)
 - **[deckhouse]** Get rid of the rbacgen tool. [#9622](https://github.com/deckhouse/deckhouse/pull/9622)
 - **[deckhouse]** Extend Deckhouse update settings. [#9314](https://github.com/deckhouse/deckhouse/pull/9314)
 - **[deckhouse-controller]** Add discovered GVKs from modules' CRDs to global values. [#9963](https://github.com/deckhouse/deckhouse/pull/9963)
 - **[deckhouse-controller]** adding an alert that manual confirmation is required to install mr [#9943](https://github.com/deckhouse/deckhouse/pull/9943)
 - **[deckhouse-controller]** Get rid of crd modules. [#9593](https://github.com/deckhouse/deckhouse/pull/9593)
 - **[dhctl]** Upon editing configuration secrets, create them if they are missing from cluster [#9689](https://github.com/deckhouse/deckhouse/pull/9689)
 - **[dhctl]** Reduces code duplication in the gRPC server message handler and log sender, refactors the graceful shutdown mechanism, and adds support for proper log output for multiple parallel instances of the dhctl server. [#9096](https://github.com/deckhouse/deckhouse/pull/9096)
 - **[dhctl]** Reduce manual operations when converging control plane nodes. [#8380](https://github.com/deckhouse/deckhouse/pull/8380)
 - **[multitenancy-manager]** Add projects render validation. [#9607](https://github.com/deckhouse/deckhouse/pull/9607)
 - **[user-authn]** dex support base64 encoded and PEM encoded certs [#9894](https://github.com/deckhouse/deckhouse/pull/9894)
 - **[user-authn]** refresh groups on updating tokens [#9598](https://github.com/deckhouse/deckhouse/pull/9598)

## Fixes


 - **[candi]** Step "check_hostname_uniqueness" works without temporary files creation [#9756](https://github.com/deckhouse/deckhouse/pull/9756)
 - **[cloud-provider-vcd]** Fix vcd catalogs sharing. [#9802](https://github.com/deckhouse/deckhouse/pull/9802)
 - **[dhctl]** Added repo check to validateRegistryDockerCfg [#9688](https://github.com/deckhouse/deckhouse/pull/9688)
 - **[dhctl]** Break circle and output error in log on check dependencies if get first error [#9679](https://github.com/deckhouse/deckhouse/pull/9679)
 - **[go_lib]** cloud-data-discoverer continues its operation despite temporary issues within the cluster. [#9570](https://github.com/deckhouse/deckhouse/pull/9570)
 - **[kube-dns]** Graceful rollout of the `kube-dns` deployment without disrupting connections. [#9565](https://github.com/deckhouse/deckhouse/pull/9565)
 - **[monitoring-kubernetes]** add tag main for dashboard [#9677](https://github.com/deckhouse/deckhouse/pull/9677)
    dashbord can be seen on the home page
 - **[multitenancy-manager]** Change logs format to json format. [#9955](https://github.com/deckhouse/deckhouse/pull/9955)

## Chore


 - **[cni-cilium]** Updating `cilium` and its components to version 1.14.14 [#9650](https://github.com/deckhouse/deckhouse/pull/9650)
    All cilium pods will be restarted.
 - **[common]** Bump shell-operator to optimize conversion hooks in the webhook-handler. [#9983](https://github.com/deckhouse/deckhouse/pull/9983)
 - **[dhctl]** Remove support for deprecated 'InitConfiguration.configOverrides' parameter. [#9920](https://github.com/deckhouse/deckhouse/pull/9920)
 - **[ingress-nginx]** Remove v1.6 IngressNginxController. [#9935](https://github.com/deckhouse/deckhouse/pull/9935)
 - **[ingress-nginx]** Update kruise controller to v1.7.2. [#9898](https://github.com/deckhouse/deckhouse/pull/9898)
    kriuse controller will be restarted, pods of an ingress nginx controller of v1.10 will be recreated.
 - **[node-manager]** Fix the module's snapshots debugging. [#9995](https://github.com/deckhouse/deckhouse/pull/9995)

