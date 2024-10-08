---
title: "Cloud provider — VMware Cloud Director: FAQ"
---

## Как поднять гибридный кластер?

Гибридный кластер представляет собой объединенные в один кластер bare-metal-узлы и узлы VMware Cloud Director. Для создания такого кластера
необходимо наличие L2-сети между всеми узлами кластера.

Чтобы поднять гибридный кластер, необходимо:

1. Включить DHCP-сервер для внутренней сети.

1. Подготовить файл с конфигурацией провайдера (используйте корректные значения для вашего облака):

   ```yaml
   apiVersion: deckhouse.io/v1
   internalNetworkCIDR: <NETWORK_CIRD>
   kind: VCDClusterConfiguration
   layout: Standard
   mainNetwork: <NETWORK_NAME>
   masterNodeGroup:
     instanceClass:
       etcdDiskSizeGb: 10
       mainNetworkIPAddresses:
       - 192.168.199.2
       rootDiskSizeGb: 20
       sizingPolicy: not_exists
       storageProfile: not_exists
       template: not_exists
     replicas: 1
   organization: <ORGANIZATION>
   provider:
     insecure: true
     password: <PASSWORD>
     server: <API_URL>
     username: <USER_NAME>
   sshPublicKey: <SSH_PUBLIC_KEY>
   virtualApplicationName: <VAPP_NAME>
   virtualDataCenter: <VDC_NAME>
   ```

   Обратите внимание, что `masterNodeGroup` является обязательным, но его можно оставить как есть.
1. Закодируйте полученный файл в Base64.
1. Создайте секрет со следующим содержимым:

   ```yaml
   
   apiVersion: v1
   data:
     cloud-provider-cluster-configuration.yaml: <BASE64_СТРОКА_ПОЛУЧЕННАЯ_НА_ПРЕДЫДУЩЕМ_ЭТАПЕ> 
     cloud-provider-discovery-data.json: eyJhcGlWZXJzaW9uIjoiZGVja2hvdXNlLmlvL3YxIiwia2luZCI6IlZDRENsb3VkUHJvdmlkZXJEaXNjb3ZlcnlEYXRhIiwiem9uZXMiOlsiZGVmYXVsdCJdfQo=
   kind: Secret
   metadata:
     labels:
       heritage: deckhouse
       name: d8-provider-cluster-configuration
     name: d8-provider-cluster-configuration
     namespace: kube-system
   type: Opaque
   ```

1. Включите модуль `cloud-provider-vcd`:

   ```shell
   kubectl -n d8-system exec -it svc/deckhouse-leader -c deckhouse -- deckhouse-controller module enable cloud-provider-vcd
   ```
