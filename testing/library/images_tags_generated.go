// Code generated by "tools/images_tags.go" DO NOT EDIT.
// To generate run 'make generate'
package library

var DefaultImagesDigests = map[string]interface{}{
	"admissionPolicyEngine": map[string]interface{}{
		"constraintExporter": "imageHash-admissionPolicyEngine-constraintExporter",
		"gatekeeper":         "imageHash-admissionPolicyEngine-gatekeeper",
		"trivyProvider":      "imageHash-admissionPolicyEngine-trivyProvider",
	},
	"basicAuth": map[string]interface{}{
		"nginx": "imageHash-basicAuth-nginx",
	},
	"cephCsi": map[string]interface{}{
		"cephcsi": "imageHash-cephCsi-cephcsi",
	},
	"certManager": map[string]interface{}{
		"certManagerAcmeSolver": "imageHash-certManager-certManagerAcmeSolver",
		"certManagerCainjector": "imageHash-certManager-certManagerCainjector",
		"certManagerController": "imageHash-certManager-certManagerController",
		"certManagerWebhook":    "imageHash-certManager-certManagerWebhook",
	},
	"chrony": map[string]interface{}{
		"chrony": "imageHash-chrony-chrony",
	},
	"ciliumHubble": map[string]interface{}{
		"relay":      "imageHash-ciliumHubble-relay",
		"uiBackend":  "imageHash-ciliumHubble-uiBackend",
		"uiFrontend": "imageHash-ciliumHubble-uiFrontend",
	},
	"cloudProviderAws": map[string]interface{}{
		"cloudControllerManager126": "imageHash-cloudProviderAws-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderAws-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderAws-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderAws-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderAws-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderAws-cloudDataDiscoverer",
		"ebsCsiPlugin":              "imageHash-cloudProviderAws-ebsCsiPlugin",
		"nodeTerminationHandler":    "imageHash-cloudProviderAws-nodeTerminationHandler",
	},
	"cloudProviderAzure": map[string]interface{}{
		"azurediskCsi":              "imageHash-cloudProviderAzure-azurediskCsi",
		"cloudControllerManager126": "imageHash-cloudProviderAzure-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderAzure-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderAzure-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderAzure-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderAzure-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderAzure-cloudDataDiscoverer",
	},
	"cloudProviderGcp": map[string]interface{}{
		"cloudControllerManager126": "imageHash-cloudProviderGcp-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderGcp-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderGcp-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderGcp-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderGcp-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderGcp-cloudDataDiscoverer",
		"pdCsiPlugin":               "imageHash-cloudProviderGcp-pdCsiPlugin",
	},
	"cloudProviderOpenstack": map[string]interface{}{
		"cinderCsiPlugin126":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin126",
		"cinderCsiPlugin127":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin127",
		"cinderCsiPlugin128":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin128",
		"cinderCsiPlugin129":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin129",
		"cinderCsiPlugin130":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin130",
		"cloudControllerManager126": "imageHash-cloudProviderOpenstack-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderOpenstack-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderOpenstack-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderOpenstack-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderOpenstack-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderOpenstack-cloudDataDiscoverer",
	},
	"cloudProviderVcd": map[string]interface{}{
		"capcdControllerManager126": "imageHash-cloudProviderVcd-capcdControllerManager126",
		"capcdControllerManager127": "imageHash-cloudProviderVcd-capcdControllerManager127",
		"capcdControllerManager128": "imageHash-cloudProviderVcd-capcdControllerManager128",
		"capcdControllerManager129": "imageHash-cloudProviderVcd-capcdControllerManager129",
		"capcdControllerManager130": "imageHash-cloudProviderVcd-capcdControllerManager130",
		"cloudControllerManager126": "imageHash-cloudProviderVcd-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderVcd-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderVcd-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderVcd-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderVcd-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderVcd-cloudDataDiscoverer",
		"vcdCsiPlugin":              "imageHash-cloudProviderVcd-vcdCsiPlugin",
		"vcdCsiPluginLegacy":        "imageHash-cloudProviderVcd-vcdCsiPluginLegacy",
	},
	"cloudProviderVsphere": map[string]interface{}{
		"cloudControllerManager126": "imageHash-cloudProviderVsphere-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderVsphere-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderVsphere-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderVsphere-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderVsphere-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderVsphere-cloudDataDiscoverer",
		"vsphereCsiPlugin126":       "imageHash-cloudProviderVsphere-vsphereCsiPlugin126",
		"vsphereCsiPlugin127":       "imageHash-cloudProviderVsphere-vsphereCsiPlugin127",
		"vsphereCsiPlugin128":       "imageHash-cloudProviderVsphere-vsphereCsiPlugin128",
		"vsphereCsiPlugin129":       "imageHash-cloudProviderVsphere-vsphereCsiPlugin129",
		"vsphereCsiPlugin130":       "imageHash-cloudProviderVsphere-vsphereCsiPlugin130",
		"vsphereCsiPluginLegacy":    "imageHash-cloudProviderVsphere-vsphereCsiPluginLegacy",
	},
	"cloudProviderYandex": map[string]interface{}{
		"cloudControllerManager126": "imageHash-cloudProviderYandex-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderYandex-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderYandex-cloudControllerManager128",
		"cloudControllerManager129": "imageHash-cloudProviderYandex-cloudControllerManager129",
		"cloudControllerManager130": "imageHash-cloudProviderYandex-cloudControllerManager130",
		"cloudDataDiscoverer":       "imageHash-cloudProviderYandex-cloudDataDiscoverer",
		"cloudMetricsExporter":      "imageHash-cloudProviderYandex-cloudMetricsExporter",
		"cloudMigrator":             "imageHash-cloudProviderYandex-cloudMigrator",
		"yandexCsiPlugin":           "imageHash-cloudProviderYandex-yandexCsiPlugin",
	},
	"cloudProviderZvirt": map[string]interface{}{
		"capzControllerManager":  "imageHash-cloudProviderZvirt-capzControllerManager",
		"cloudControllerManager": "imageHash-cloudProviderZvirt-cloudControllerManager",
		"cloudDataDiscoverer":    "imageHash-cloudProviderZvirt-cloudDataDiscoverer",
		"zvirtCsiDriver":         "imageHash-cloudProviderZvirt-zvirtCsiDriver",
	},
	"cniCilium": map[string]interface{}{
		"agentDistroless":    "imageHash-cniCilium-agentDistroless",
		"checkKernelVersion": "imageHash-cniCilium-checkKernelVersion",
		"egressGatewayAgent": "imageHash-cniCilium-egressGatewayAgent",
		"kubeRbacProxy":      "imageHash-cniCilium-kubeRbacProxy",
		"operator":           "imageHash-cniCilium-operator",
		"safeAgentUpdater":   "imageHash-cniCilium-safeAgentUpdater",
	},
	"cniFlannel": map[string]interface{}{
		"flanneld": "imageHash-cniFlannel-flanneld",
	},
	"cniSimpleBridge": map[string]interface{}{
		"simpleBridge": "imageHash-cniSimpleBridge-simpleBridge",
	},
	"common": map[string]interface{}{
		"alpine":                    "imageHash-common-alpine",
		"checkKernelVersion":        "imageHash-common-checkKernelVersion",
		"csiExternalAttacher126":    "imageHash-common-csiExternalAttacher126",
		"csiExternalAttacher127":    "imageHash-common-csiExternalAttacher127",
		"csiExternalAttacher128":    "imageHash-common-csiExternalAttacher128",
		"csiExternalAttacher129":    "imageHash-common-csiExternalAttacher129",
		"csiExternalAttacher130":    "imageHash-common-csiExternalAttacher130",
		"csiExternalProvisioner126": "imageHash-common-csiExternalProvisioner126",
		"csiExternalProvisioner127": "imageHash-common-csiExternalProvisioner127",
		"csiExternalProvisioner128": "imageHash-common-csiExternalProvisioner128",
		"csiExternalProvisioner129": "imageHash-common-csiExternalProvisioner129",
		"csiExternalProvisioner130": "imageHash-common-csiExternalProvisioner130",
		"csiExternalResizer126":     "imageHash-common-csiExternalResizer126",
		"csiExternalResizer127":     "imageHash-common-csiExternalResizer127",
		"csiExternalResizer128":     "imageHash-common-csiExternalResizer128",
		"csiExternalResizer129":     "imageHash-common-csiExternalResizer129",
		"csiExternalResizer130":     "imageHash-common-csiExternalResizer130",
		"csiExternalSnapshotter126": "imageHash-common-csiExternalSnapshotter126",
		"csiExternalSnapshotter127": "imageHash-common-csiExternalSnapshotter127",
		"csiExternalSnapshotter128": "imageHash-common-csiExternalSnapshotter128",
		"csiExternalSnapshotter129": "imageHash-common-csiExternalSnapshotter129",
		"csiExternalSnapshotter130": "imageHash-common-csiExternalSnapshotter130",
		"csiLivenessprobe126":       "imageHash-common-csiLivenessprobe126",
		"csiLivenessprobe127":       "imageHash-common-csiLivenessprobe127",
		"csiLivenessprobe128":       "imageHash-common-csiLivenessprobe128",
		"csiLivenessprobe129":       "imageHash-common-csiLivenessprobe129",
		"csiLivenessprobe130":       "imageHash-common-csiLivenessprobe130",
		"csiNodeDriverRegistrar126": "imageHash-common-csiNodeDriverRegistrar126",
		"csiNodeDriverRegistrar127": "imageHash-common-csiNodeDriverRegistrar127",
		"csiNodeDriverRegistrar128": "imageHash-common-csiNodeDriverRegistrar128",
		"csiNodeDriverRegistrar129": "imageHash-common-csiNodeDriverRegistrar129",
		"csiNodeDriverRegistrar130": "imageHash-common-csiNodeDriverRegistrar130",
		"csiVsphereSyncer126":       "imageHash-common-csiVsphereSyncer126",
		"csiVsphereSyncer127":       "imageHash-common-csiVsphereSyncer127",
		"csiVsphereSyncer128":       "imageHash-common-csiVsphereSyncer128",
		"csiVsphereSyncer129":       "imageHash-common-csiVsphereSyncer129",
		"csiVsphereSyncer130":       "imageHash-common-csiVsphereSyncer130",
		"distroless":                "imageHash-common-distroless",
		"init":                      "imageHash-common-init",
		"iptablesWrapper":           "imageHash-common-iptablesWrapper",
		"kubeRbacProxy":             "imageHash-common-kubeRbacProxy",
		"nginxStatic":               "imageHash-common-nginxStatic",
		"pause":                     "imageHash-common-pause",
		"pythonStatic":              "imageHash-common-pythonStatic",
		"redisStatic":               "imageHash-common-redisStatic",
		"shellOperator":             "imageHash-common-shellOperator",
	},
	"controlPlaneManager": map[string]interface{}{
		"controlPlaneManager126":   "imageHash-controlPlaneManager-controlPlaneManager126",
		"controlPlaneManager127":   "imageHash-controlPlaneManager-controlPlaneManager127",
		"controlPlaneManager128":   "imageHash-controlPlaneManager-controlPlaneManager128",
		"controlPlaneManager129":   "imageHash-controlPlaneManager-controlPlaneManager129",
		"controlPlaneManager130":   "imageHash-controlPlaneManager-controlPlaneManager130",
		"etcd":                     "imageHash-controlPlaneManager-etcd",
		"etcdBackup":               "imageHash-controlPlaneManager-etcdBackup",
		"kubeApiserver126":         "imageHash-controlPlaneManager-kubeApiserver126",
		"kubeApiserver127":         "imageHash-controlPlaneManager-kubeApiserver127",
		"kubeApiserver128":         "imageHash-controlPlaneManager-kubeApiserver128",
		"kubeApiserver129":         "imageHash-controlPlaneManager-kubeApiserver129",
		"kubeApiserver130":         "imageHash-controlPlaneManager-kubeApiserver130",
		"kubeApiserverHealthcheck": "imageHash-controlPlaneManager-kubeApiserverHealthcheck",
		"kubeControllerManager126": "imageHash-controlPlaneManager-kubeControllerManager126",
		"kubeControllerManager127": "imageHash-controlPlaneManager-kubeControllerManager127",
		"kubeControllerManager128": "imageHash-controlPlaneManager-kubeControllerManager128",
		"kubeControllerManager129": "imageHash-controlPlaneManager-kubeControllerManager129",
		"kubeControllerManager130": "imageHash-controlPlaneManager-kubeControllerManager130",
		"kubeScheduler126":         "imageHash-controlPlaneManager-kubeScheduler126",
		"kubeScheduler127":         "imageHash-controlPlaneManager-kubeScheduler127",
		"kubeScheduler128":         "imageHash-controlPlaneManager-kubeScheduler128",
		"kubeScheduler129":         "imageHash-controlPlaneManager-kubeScheduler129",
		"kubeScheduler130":         "imageHash-controlPlaneManager-kubeScheduler130",
		"kubernetesApiProxy":       "imageHash-controlPlaneManager-kubernetesApiProxy",
	},
	"dashboard": map[string]interface{}{
		"dashboard":      "imageHash-dashboard-dashboard",
		"metricsScraper": "imageHash-dashboard-metricsScraper",
	},
	"deckhouse": map[string]interface{}{
		"webhookHandler": "imageHash-deckhouse-webhookHandler",
	},
	"deckhouseTools": map[string]interface{}{
		"web": "imageHash-deckhouseTools-web",
	},
	"descheduler": map[string]interface{}{
		"descheduler": "imageHash-descheduler-descheduler",
	},
	"documentation": map[string]interface{}{
		"docsBuilder": "imageHash-documentation-docsBuilder",
		"web":         "imageHash-documentation-web",
	},
	"extendedMonitoring": map[string]interface{}{
		"certExporter":               "imageHash-extendedMonitoring-certExporter",
		"eventsExporter":             "imageHash-extendedMonitoring-eventsExporter",
		"extendedMonitoringExporter": "imageHash-extendedMonitoring-extendedMonitoringExporter",
		"imageAvailabilityExporter":  "imageHash-extendedMonitoring-imageAvailabilityExporter",
	},
	"ingressNginx": map[string]interface{}{
		"controller110":         "imageHash-ingressNginx-controller110",
		"controller19":          "imageHash-ingressNginx-controller19",
		"kruise":                "imageHash-ingressNginx-kruise",
		"kruiseStateMetrics":    "imageHash-ingressNginx-kruiseStateMetrics",
		"kubeRbacProxy":         "imageHash-ingressNginx-kubeRbacProxy",
		"nginxExporter":         "imageHash-ingressNginx-nginxExporter",
		"protobufExporter":      "imageHash-ingressNginx-protobufExporter",
		"proxyFailover":         "imageHash-ingressNginx-proxyFailover",
		"proxyFailoverIptables": "imageHash-ingressNginx-proxyFailoverIptables",
	},
	"istio": map[string]interface{}{
		"apiProxy":          "imageHash-istio-apiProxy",
		"cniV1x16x2":        "imageHash-istio-cniV1x16x2",
		"cniV1x19x7":        "imageHash-istio-cniV1x19x7",
		"kiali":             "imageHash-istio-kiali",
		"metadataDiscovery": "imageHash-istio-metadataDiscovery",
		"metadataExporter":  "imageHash-istio-metadataExporter",
		"operatorV1x16x2":   "imageHash-istio-operatorV1x16x2",
		"operatorV1x19x7":   "imageHash-istio-operatorV1x19x7",
		"pilotV1x16x2":      "imageHash-istio-pilotV1x16x2",
		"pilotV1x19x7":      "imageHash-istio-pilotV1x19x7",
		"proxyv2V1x16x2":    "imageHash-istio-proxyv2V1x16x2",
		"proxyv2V1x19x7":    "imageHash-istio-proxyv2V1x19x7",
	},
	"keepalived": map[string]interface{}{
		"keepalived": "imageHash-keepalived-keepalived",
	},
	"kubeDns": map[string]interface{}{
		"coredns":                           "imageHash-kubeDns-coredns",
		"resolvWatcher":                     "imageHash-kubeDns-resolvWatcher",
		"stsPodsHostsAppenderInitContainer": "imageHash-kubeDns-stsPodsHostsAppenderInitContainer",
		"stsPodsHostsAppenderWebhook":       "imageHash-kubeDns-stsPodsHostsAppenderWebhook",
	},
	"kubeProxy": map[string]interface{}{
		"initContainer": "imageHash-kubeProxy-initContainer",
		"kubeProxy126":  "imageHash-kubeProxy-kubeProxy126",
		"kubeProxy127":  "imageHash-kubeProxy-kubeProxy127",
		"kubeProxy128":  "imageHash-kubeProxy-kubeProxy128",
		"kubeProxy129":  "imageHash-kubeProxy-kubeProxy129",
		"kubeProxy130":  "imageHash-kubeProxy-kubeProxy130",
	},
	"l2LoadBalancer": map[string]interface{}{
		"controller": "imageHash-l2LoadBalancer-controller",
		"speaker":    "imageHash-l2LoadBalancer-speaker",
	},
	"localPathProvisioner": map[string]interface{}{
		"helper":               "imageHash-localPathProvisioner-helper",
		"localPathProvisioner": "imageHash-localPathProvisioner-localPathProvisioner",
	},
	"logShipper": map[string]interface{}{
		"vector": "imageHash-logShipper-vector",
	},
	"loki": map[string]interface{}{
		"loki": "imageHash-loki-loki",
	},
	"metallb": map[string]interface{}{
		"controller": "imageHash-metallb-controller",
		"speaker":    "imageHash-metallb-speaker",
	},
	"monitoringKubernetes": map[string]interface{}{
		"ebpfExporter":                      "imageHash-monitoringKubernetes-ebpfExporter",
		"kubeStateMetrics":                  "imageHash-monitoringKubernetes-kubeStateMetrics",
		"kubeletEvictionThresholdsExporter": "imageHash-monitoringKubernetes-kubeletEvictionThresholdsExporter",
		"nodeExporter":                      "imageHash-monitoringKubernetes-nodeExporter",
	},
	"monitoringPing": map[string]interface{}{
		"monitoringPing": "imageHash-monitoringPing-monitoringPing",
	},
	"multitenancyManager": map[string]interface{}{
		"multitenancyManager": "imageHash-multitenancyManager-multitenancyManager",
	},
	"networkGateway": map[string]interface{}{
		"dnsmasq": "imageHash-networkGateway-dnsmasq",
		"snat":    "imageHash-networkGateway-snat",
	},
	"networkPolicyEngine": map[string]interface{}{
		"kubeRouter": "imageHash-networkPolicyEngine-kubeRouter",
	},
	"nodeLocalDns": map[string]interface{}{
		"coredns":                    "imageHash-nodeLocalDns-coredns",
		"iptablesLoop":               "imageHash-nodeLocalDns-iptablesLoop",
		"staleDnsConnectionsCleaner": "imageHash-nodeLocalDns-staleDnsConnectionsCleaner",
	},
	"nodeManager": map[string]interface{}{
		"bashibleApiserver":        "imageHash-nodeManager-bashibleApiserver",
		"capiControllerManager":    "imageHash-nodeManager-capiControllerManager",
		"capsControllerManager":    "imageHash-nodeManager-capsControllerManager",
		"clusterAutoscaler126":     "imageHash-nodeManager-clusterAutoscaler126",
		"clusterAutoscaler127":     "imageHash-nodeManager-clusterAutoscaler127",
		"clusterAutoscaler128":     "imageHash-nodeManager-clusterAutoscaler128",
		"clusterAutoscaler129":     "imageHash-nodeManager-clusterAutoscaler129",
		"clusterAutoscaler130":     "imageHash-nodeManager-clusterAutoscaler130",
		"earlyOom":                 "imageHash-nodeManager-earlyOom",
		"fencingAgent":             "imageHash-nodeManager-fencingAgent",
		"machineControllerManager": "imageHash-nodeManager-machineControllerManager",
	},
	"openvpn": map[string]interface{}{
		"easyrsaMigrator": "imageHash-openvpn-easyrsaMigrator",
		"openvpn":         "imageHash-openvpn-openvpn",
		"ovpnAdmin":       "imageHash-openvpn-ovpnAdmin",
		"pmacct":          "imageHash-openvpn-pmacct",
	},
	"operatorPrometheus": map[string]interface{}{
		"prometheusConfigReloader": "imageHash-operatorPrometheus-prometheusConfigReloader",
		"prometheusOperator":       "imageHash-operatorPrometheus-prometheusOperator",
	},
	"operatorTrivy": map[string]interface{}{
		"nodeCollector": "imageHash-operatorTrivy-nodeCollector",
		"operator":      "imageHash-operatorTrivy-operator",
		"reportUpdater": "imageHash-operatorTrivy-reportUpdater",
		"trivy":         "imageHash-operatorTrivy-trivy",
	},
	"podReloader": map[string]interface{}{
		"podReloader": "imageHash-podReloader-podReloader",
	},
	"prometheus": map[string]interface{}{
		"alertmanager":                "imageHash-prometheus-alertmanager",
		"alertsReceiver":              "imageHash-prometheus-alertsReceiver",
		"grafana":                     "imageHash-prometheus-grafana",
		"grafanaDashboardProvisioner": "imageHash-prometheus-grafanaDashboardProvisioner",
		"grafanaV10":                  "imageHash-prometheus-grafanaV10",
		"memcached":                   "imageHash-prometheus-memcached",
		"memcachedExporter":           "imageHash-prometheus-memcachedExporter",
		"mimir":                       "imageHash-prometheus-mimir",
		"prometheus":                  "imageHash-prometheus-prometheus",
		"promxy":                      "imageHash-prometheus-promxy",
		"trickster":                   "imageHash-prometheus-trickster",
	},
	"prometheusMetricsAdapter": map[string]interface{}{
		"k8sPrometheusAdapter":   "imageHash-prometheusMetricsAdapter-k8sPrometheusAdapter",
		"prometheusReverseProxy": "imageHash-prometheusMetricsAdapter-prometheusReverseProxy",
	},
	"prometheusPushgateway": map[string]interface{}{
		"pushgateway": "imageHash-prometheusPushgateway-pushgateway",
	},
	"registryPackagesProxy": map[string]interface{}{
		"registryPackagesProxy": "imageHash-registryPackagesProxy-registryPackagesProxy",
	},
	"registrypackages": map[string]interface{}{
		"amazonEc2Utils220":         "imageHash-registrypackages-amazonEc2Utils220",
		"containerd1720":            "imageHash-registrypackages-containerd1720",
		"crictl126":                 "imageHash-registrypackages-crictl126",
		"crictl127":                 "imageHash-registrypackages-crictl127",
		"crictl128":                 "imageHash-registrypackages-crictl128",
		"crictl129":                 "imageHash-registrypackages-crictl129",
		"crictl130":                 "imageHash-registrypackages-crictl130",
		"d8":                        "imageHash-registrypackages-d8",
		"d8CaUpdater060824":         "imageHash-registrypackages-d8CaUpdater060824",
		"d8Curl821":                 "imageHash-registrypackages-d8Curl821",
		"dockerRegistry283":         "imageHash-registrypackages-dockerRegistry283",
		"drbd":                      "imageHash-registrypackages-drbd",
		"e2fsprogs1470":             "imageHash-registrypackages-e2fsprogs1470",
		"ec2DescribeTagsV001Flant2": "imageHash-registrypackages-ec2DescribeTagsV001Flant2",
		"ecrCredentialProvider126":  "imageHash-registrypackages-ecrCredentialProvider126",
		"ecrCredentialProvider127":  "imageHash-registrypackages-ecrCredentialProvider127",
		"ecrCredentialProvider128":  "imageHash-registrypackages-ecrCredentialProvider128",
		"ecrCredentialProvider129":  "imageHash-registrypackages-ecrCredentialProvider129",
		"ecrCredentialProvider130":  "imageHash-registrypackages-ecrCredentialProvider130",
		"growpart033":               "imageHash-registrypackages-growpart033",
		"iptables189":               "imageHash-registrypackages-iptables189",
		"jq16":                      "imageHash-registrypackages-jq16",
		"kubeadm12615":              "imageHash-registrypackages-kubeadm12615",
		"kubeadm12716":              "imageHash-registrypackages-kubeadm12716",
		"kubeadm12815":              "imageHash-registrypackages-kubeadm12815",
		"kubeadm12910":              "imageHash-registrypackages-kubeadm12910",
		"kubeadm1306":               "imageHash-registrypackages-kubeadm1306",
		"kubectl12615":              "imageHash-registrypackages-kubectl12615",
		"kubectl12716":              "imageHash-registrypackages-kubectl12716",
		"kubectl12815":              "imageHash-registrypackages-kubectl12815",
		"kubectl12910":              "imageHash-registrypackages-kubectl12910",
		"kubectl1306":               "imageHash-registrypackages-kubectl1306",
		"kubelet12615":              "imageHash-registrypackages-kubelet12615",
		"kubelet12716":              "imageHash-registrypackages-kubelet12716",
		"kubelet12815":              "imageHash-registrypackages-kubelet12815",
		"kubelet12910":              "imageHash-registrypackages-kubelet12910",
		"kubelet1306":               "imageHash-registrypackages-kubelet1306",
		"kubernetesCni140":          "imageHash-registrypackages-kubernetesCni140",
		"lsblk2402":                 "imageHash-registrypackages-lsblk2402",
		"netcat110481":              "imageHash-registrypackages-netcat110481",
		"socat1734":                 "imageHash-registrypackages-socat1734",
		"tomlMerge01":               "imageHash-registrypackages-tomlMerge01",
		"virtWhat125":               "imageHash-registrypackages-virtWhat125",
		"xfsprogs670":               "imageHash-registrypackages-xfsprogs670",
	},
	"runtimeAuditEngine": map[string]interface{}{
		"falco":            "imageHash-runtimeAuditEngine-falco",
		"falcosidekick":    "imageHash-runtimeAuditEngine-falcosidekick",
		"k8sMetacollector": "imageHash-runtimeAuditEngine-k8sMetacollector",
		"rulesLoader":      "imageHash-runtimeAuditEngine-rulesLoader",
	},
	"snapshotController": map[string]interface{}{
		"snapshotController":        "imageHash-snapshotController-snapshotController",
		"snapshotValidationWebhook": "imageHash-snapshotController-snapshotValidationWebhook",
	},
	"staticRoutingManager": map[string]interface{}{
		"agent": "imageHash-staticRoutingManager-agent",
	},
	"terraformManager": map[string]interface{}{
		"baseTerraformManager":      "imageHash-terraformManager-baseTerraformManager",
		"terraformManagerAws":       "imageHash-terraformManager-terraformManagerAws",
		"terraformManagerAzure":     "imageHash-terraformManager-terraformManagerAzure",
		"terraformManagerGcp":       "imageHash-terraformManager-terraformManagerGcp",
		"terraformManagerOpenstack": "imageHash-terraformManager-terraformManagerOpenstack",
		"terraformManagerVcd":       "imageHash-terraformManager-terraformManagerVcd",
		"terraformManagerVsphere":   "imageHash-terraformManager-terraformManagerVsphere",
		"terraformManagerYandex":    "imageHash-terraformManager-terraformManagerYandex",
		"terraformManagerZvirt":     "imageHash-terraformManager-terraformManagerZvirt",
	},
	"upmeter": map[string]interface{}{
		"smokeMini": "imageHash-upmeter-smokeMini",
		"status":    "imageHash-upmeter-status",
		"upmeter":   "imageHash-upmeter-upmeter",
		"webui":     "imageHash-upmeter-webui",
	},
	"userAuthn": map[string]interface{}{
		"basicAuthProxy":      "imageHash-userAuthn-basicAuthProxy",
		"dex":                 "imageHash-userAuthn-dex",
		"dexAuthenticator":    "imageHash-userAuthn-dexAuthenticator",
		"kubeconfigGenerator": "imageHash-userAuthn-kubeconfigGenerator",
		"selfSignedGenerator": "imageHash-userAuthn-selfSignedGenerator",
	},
	"userAuthz": map[string]interface{}{
		"webhook": "imageHash-userAuthz-webhook",
	},
	"verticalPodAutoscaler": map[string]interface{}{
		"admissionController": "imageHash-verticalPodAutoscaler-admissionController",
		"recommender":         "imageHash-verticalPodAutoscaler-recommender",
		"updater":             "imageHash-verticalPodAutoscaler-updater",
	},
}
