---
title: "Управление узлами: примеры"
description: Примеры управления узлами кластера Kubernetes. Примеры создания группы узлов. Примеры автоматизации выполнения произвольных настроек на узле.
---

Ниже представлены несколько примеров описания NodeGroup, а также установки плагина cert-manager для `kubectl` и задания параметра `sysctl`.

## Примеры описания NodeGroup

<span id="пример-описания-nodegroup"></span>

### Облачные узлы

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: test
spec:
  nodeType: CloudEphemeral
  cloudInstances:
    zones:
      - eu-west-1a
      - eu-west-1b
    minPerZone: 1
    maxPerZone: 2
    classReference:
      kind: AWSInstanceClass
      name: test
  nodeTemplate:
    labels:
      tier: test
```

### Статические узлы

<span id="пример-описания-статической-nodegroup"></span>

Для виртуальных машин на гипервизорах или физических серверов используйте статические узлы, указав `nodeType: Static` в NodeGroup.

Пример:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: worker
spec:
  nodeType: Static
```

Узлы в такую группу добавляются [вручную](#вручную) с помощью подготовленных скриптов.

Также можно использовать способ [добавления статических узлов с помощью Cluster API Provider Static](#с-помощью-cluster-api-provider-static).

### Системные узлы

<span id="пример-описания-статичной-nodegroup-для-системных-узлов"></span>

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: system
spec:
  nodeTemplate:
    labels:
      node-role.deckhouse.io/system: ""
    taints:
      - effect: NoExecute
        key: dedicated.deckhouse.io
        value: system
  nodeType: Static
```

## Добавление статического узла в кластер

<span id="пример-описания-статичной-nodegroup"></span>

Добавление статического узла можно выполнить вручную или с помощью Cluster API Provider Static.

### Вручную

Чтобы добавить новый статический узел (выделенная ВМ, bare-metal-сервер и т. п.) в кластер вручную, выполните следующие шаги:

1. Для [CloudStatic-узлов](../node-manager/cr.html#nodegroup-v1-spec-nodetype) в облачных провайдерах, перечисленных ниже, выполните описанные в документации шаги:
   - [Для AWS](../cloud-provider-aws/faq.html#добавление-cloudstatic-узлов-в-кластер)
   - [Для GCP](../cloud-provider-gcp/faq.html#добавление-cloudstatic-узлов-в-кластер)
   - [Для YC](../cloud-provider-yandex/faq.html#добавление-cloudstatic-узлов-в-кластер)
1. Используйте существующий или создайте новый ресурс [NodeGroup](cr.html#nodegroup) ([пример](#статические-узлы) NodeGroup с именем `worker`). Параметр [nodeType](cr.html#nodegroup-v1-spec-nodetype) в ресурсе NodeGroup для статических узлов должен быть `Static` или `CloudStatic`.
1. Получите код скрипта в кодировке Base64 для добавления и настройки узла.

   Пример получения кода скрипта в кодировке Base64 для добавления узла в NodeGroup `worker`:

   ```shell
   NODE_GROUP=worker
   kubectl -n d8-cloud-instance-manager get secret manual-bootstrap-for-${NODE_GROUP} -o json | jq '.data."bootstrap.sh"' -r
   ```

1. Выполните предварительную настройку нового узла в соответствии с особенностями вашего окружения. Например:
   - добавьте необходимые точки монтирования в файл `/etc/fstab` (NFS, Ceph и т. д.);
   - установите необходимые пакеты;
   - настройте сетевую связность между новым узлом и остальными узлами кластера.
1. Зайдите на новый узел по SSH и выполните следующую команду, вставив полученную в п. 3 Base64-строку:

   ```shell
   echo <Base64-КОД-СКРИПТА> | base64 -d | bash
   ```

### С помощью Cluster API Provider Static

Простой пример добавления статического узла в кластер с помощью [Cluster API Provider Static (CAPS)](./#cluster-api-provider-static):

1. Подготовьте необходимые ресурсы.

   * Выделите сервер (или виртуальную машину), настройте сетевую связность и т. п., при необходимости установите специфические пакеты ОС и добавьте точки монтирования которые потребуются на узле.

   * Создайте пользователя (в примере — `caps`) с возможностью выполнять `sudo`, выполнив **на сервере** следующую команду:

     ```shell
     useradd -m -s /bin/bash caps 
     usermod -aG sudo caps
     ```

   * Разрешите пользователю выполнять команды через sudo без пароля. Для этого **на сервере** внесите следующую строку в конфигурацию sudo (отредактировав файл `/etc/sudoers`, выполнив команду `sudo visudo` или другим способом):

     ```text
     caps ALL=(ALL) NOPASSWD: ALL
     ```

   * Сгенерируйте **на сервере** пару SSH-ключей с пустой парольной фразой:

     ```shell
     ssh-keygen -t rsa -f caps-id -C "" -N ""
     ```

     Публичный и приватный ключи пользователя `caps` будут сохранены в файлах `caps-id.pub` и `caps-id` в текущей директории на сервере.

   * Добавьте полученный публичный ключ в файл `/home/caps/.ssh/authorized_keys` пользователя `caps`, выполнив в директории с ключами **на сервере** следующие команды:

     ```shell
     mkdir -p /home/caps/.ssh 
     cat caps-id.pub >> /home/caps/.ssh/authorized_keys 
     chmod 700 /home/caps/.ssh 
     chmod 600 /home/caps/.ssh/authorized_keys
     chown -R caps:caps /home/caps/
     ```

   В операционных системах семейства Astra Linux, при использовании модуля мандатного контроля целостности Parsec, сконфигурируйте максимальный уровень целостности для пользователя `caps`:

     ```shell
     pdpl-user -i 63 caps
     ```

1. Создайте в кластере ресурс [SSHCredentials](cr.html#sshcredentials).

   В директории с ключами пользователя **на сервере** выполните следующую команду для получения закрытого ключа в формате Base64:

   ```shell
   base64 -w0 caps-id
   ```

   На любом компьютере с `kubectl`, настроенным на управление кластером, создайте переменную окружения со значением закрытого ключа созданного пользователя в Base64, полученным на предыдущем шаге:

   ```shell
    CAPS_PRIVATE_KEY_BASE64=<ЗАКРЫТЫЙ_КЛЮЧ_В_BASE64>
   ```

   Выполните следующую команду, для создания в кластере ресурса `SSHCredentials` (здесь и далее также используйте `kubectl`, настроенный на управление кластером):

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1alpha1
   kind: SSHCredentials
   metadata:
     name: credentials
   spec:
     user: caps
     privateSSHKey: "${CAPS_PRIVATE_KEY_BASE64}"
   EOF
   ```

1. Создайте в кластере ресурс [StaticInstance](cr.html#staticinstance), указав IP-адрес сервера статического узла:

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1alpha1
   kind: StaticInstance
   metadata:
     name: static-worker-1
     labels:
       role: worker
   spec:
     # Укажите IP-адрес сервера статического узла.
     address: "<SERVER-IP>"
     credentialsRef:
       kind: SSHCredentials
       name: credentials
   EOF
   ```

1. Создайте в кластере ресурс [NodeGroup](cr.html#nodegroup). Параметр `count` обозначает количество `staticInstances`, подпадающих под `labelSelector`, которые будут добавлены в кластер, в данном случае `1`:

   > Поле `labelSelector` в ресурсе `NodeGroup` является неизменным. Чтобы обновить `labelSelector`, нужно создать новую `NodeGroup` и перенести в неё статические узлы, изменив их лейблы (labels).

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1
   kind: NodeGroup
   metadata:
     name: worker
   spec:
     nodeType: Static
     staticInstances:
       count: 1
       labelSelector:
         matchLabels:
           role: worker
   EOF
   ```

   > Если необходимо добавить узлы в уже существующую группу узлов, укажите их желаемое количество в поле `.spec.count` NodeGroup.

### С помощью Cluster API Provider Static для нескольких групп узлов

Пример использования фильтров в [label selector](cr.html#nodegroup-v1-spec-staticinstances-labelselector) StaticInstance, для группировки статических узлов и использования их в разных NodeGroup. В примере используются две группы узлов (`front` и `worker`), предназначенные для разных задач, которые должны содержать разные по характеристикам узлы — два сервера для группы `front` и один для группы `worker`.

1. Подготовьте необходимые ресурсы (3 сервера или виртуальные машины) и создайте ресурс `SSHCredentials`, аналогично п.1 и п.2 [примера](#с-помощью-cluster-api-provider-static).

1. Создайте в кластере два ресурса [NodeGroup](cr.html#nodegroup) (здесь и далее используйте `kubectl`, настроенный на управление кластером):

   > Поле `labelSelector` в ресурсе `NodeGroup` является неизменным. Чтобы обновить labelSelector, нужно создать новую NodeGroup и перенести в неё статические узлы, изменив их лейблы (labels).

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1
   kind: NodeGroup
   metadata:
     name: front
   spec:
     nodeType: Static
     staticInstances:
       count: 2
       labelSelector:
         matchLabels:
           role: front
   ---
   apiVersion: deckhouse.io/v1
   kind: NodeGroup
   metadata:
     name: worker
   spec:
     nodeType: Static
     staticInstances:
       count: 1
       labelSelector:
         matchLabels:
           role: worker
   EOF
   ```

1. Создайте в кластере ресурсы [StaticInstance](cr.html#staticinstance), указав актуальные IP-адреса серверов:

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1alpha1
   kind: StaticInstance
   metadata:
     name: static-front-1
     labels:
       role: front
   spec:
     address: "<SERVER-FRONT-IP1>"
     credentialsRef:
       kind: SSHCredentials
       name: credentials
   ---
   apiVersion: deckhouse.io/v1alpha1
   kind: StaticInstance
   metadata:
     name: static-front-2
     labels:
       role: front
   spec:
     address: "<SERVER-FRONT-IP2>"
     credentialsRef:
       kind: SSHCredentials
       name: credentials
   ---
   apiVersion: deckhouse.io/v1alpha1
   kind: StaticInstance
   metadata:
     name: static-worker-1
     labels:
       role: worker
   spec:
     address: "<SERVER-WORKER-IP>"
     credentialsRef:
       kind: SSHCredentials
       name: credentials
   EOF
   ```

### Cluster API Provider Static: перемещение узлов между NodeGroup

В данном разделе описывается процесс перемещения статических узлов между различными NodeGroup с использованием Cluster API Provider Static (CAPS). Процесс включает изменение конфигурации NodeGroup и обновление лейблов у соответствующих StaticInstance.

#### Исходная конфигурация

Предположим, что в кластере уже существует NodeGroup с именем `worker`, настроенный для управления одним статическим узлом с лейблом `role: worker`.

`NodeGroup` worker:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: worker
spec:
  nodeType: Static
  staticInstances:
    count: 1
    labelSelector:
      matchLabels:
        role: worker
```

`StaticInstance` static-0:

```yaml
apiVersion: deckhouse.io/v1alpha2
kind: StaticInstance
metadata:
  name: static-worker-1
  labels:
    role: worker
spec:
  address: "192.168.1.100"
  credentialsRef:
    kind: SSHCredentials
    name: credentials
```

#### Шаги по перемещению узла между `NodeGroup`

{% alert level="warning" %}
В процессе переноса узлов между NodeGroup будет выполнена очистка и повторный бутстрап узла, объект `Node` будет пересоздан.
{% endalert %}

##### 1. Создание новой `NodeGroup` для целевой группы узлов

Создайте новый ресурс NodeGroup, например, с именем `front`, который будет управлять статическим узлом с лейблом `role: front`.

```shell
kubectl create -f - <<EOF
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: front
spec:
  nodeType: Static
  staticInstances:
    count: 1
    labelSelector:
      matchLabels:
        role: front
EOF
```

##### 2. Обновление лейбла у `StaticInstance`

Измените лейбл `role` у существующего StaticInstance с `worker` на `front`. Это позволит новой NodeGroup `front` начать управлять этим узлом.

```shell
kubectl label staticinstance static-worker-1 role=front --overwrite
```

##### 3. Уменьшение количества статических узлов в исходной `NodeGroup`

Обновите ресурс NodeGroup `worker`, уменьшив значение параметра `count` с `1` до `0`.

```shell
kubectl patch nodegroup worker -p '{"spec": {"staticInstances": {"count": 0}}}' --type=merge
```

## Пример описания `NodeUser`

```yaml
apiVersion: deckhouse.io/v1
kind: NodeUser
metadata:
  name: testuser
spec:
  uid: 1100
  sshPublicKeys:
  - "<SSH_PUBLIC_KEY>"
  passwordHash: <PASSWORD_HASH>
  isSudoer: true
```

## Пример описания `NodeGroupConfiguration`

### Установка плагина cert-manager для kubectl на master-узлах

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: NodeGroupConfiguration
metadata:
  name: add-cert-manager-plugin.sh
spec:
  weight: 100
  bundles:
  - "*"
  nodeGroups:
  - "master"
  content: |
    if [ -x /usr/local/bin/kubectl-cert_manager ]; then
      exit 0
    fi
    curl -L https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/kubectl-cert_manager-linux-amd64.tar.gz -o - | tar -zxvf - kubectl-cert_manager
    mv kubectl-cert_manager /usr/local/bin
```

### Задание параметра sysctl

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: NodeGroupConfiguration
metadata:
  name: sysctl-tune.sh
spec:
  weight: 100
  bundles:
  - "*"
  nodeGroups:
  - "*"
  content: |
    sysctl -w vm.max_map_count=262144
```

### Добавление корневого сертификата в хост

{% alert level="warning" %}
Данный пример приведен для ОС Ubuntu.  
Способ добавления сертификатов в хранилище может отличаться в зависимости от ОС.
  
При адаптации скрипта под другую ОС измените параметры [bundles](cr.html#nodegroupconfiguration-v1alpha1-spec-bundles) и [content](cr.html#nodegroupconfiguration-v1alpha1-spec-content).
{% endalert %}

{% alert level="warning" %}
Для использования сертификата в `containerd` (в т.ч. pull контейнеров из приватного репозитория) после добавления сертификата требуется произвести рестарт сервиса.
{% endalert %}

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: NodeGroupConfiguration
metadata:
  name: add-custom-ca.sh
spec:
  weight: 31
  nodeGroups:
  - '*'  
  bundles:
  - 'ubuntu-lts'
  content: |-
    CERT_FILE_NAME=example_ca
    CERTS_FOLDER="/usr/local/share/ca-certificates"
    CERT_CONTENT=$(cat <<EOF
    -----BEGIN CERTIFICATE-----
    MIIDSjCCAjKgAwIBAgIRAJ4RR/WDuAym7M11JA8W7D0wDQYJKoZIhvcNAQELBQAw
    JTEjMCEGA1UEAxMabmV4dXMuNTEuMjUwLjQxLjIuc3NsaXAuaW8wHhcNMjQwODAx
    MTAzMjA4WhcNMjQxMDMwMTAzMjA4WjAlMSMwIQYDVQQDExpuZXh1cy41MS4yNTAu
    NDEuMi5zc2xpcC5pbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL1p
    WLPr2c4SZX/i4IS59Ly1USPjRE21G4pMYewUjkSXnYv7hUkHvbNL/P9dmGBm2Jsl
    WFlRZbzCv7+5/J+9mPVL2TdTbWuAcTUyaG5GZ/1w64AmAWxqGMFx4eyD1zo9eSmN
    G2jis8VofL9dWDfUYhRzJ90qKxgK6k7tfhL0pv7IHDbqf28fCEnkvxsA98lGkq3H
    fUfvHV6Oi8pcyPZ/c8ayIf4+JOnf7oW/TgWqI7x6R1CkdzwepJ8oU7PGc0ySUWaP
    G5bH3ofBavL0bNEsyScz4TFCJ9b4aO5GFAOmgjFMMUi9qXDH72sBSrgi08Dxmimg
    Hfs198SZr3br5GTJoAkCAwEAAaN1MHMwDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB
    /wQCMAAwUwYDVR0RBEwwSoIPbmV4dXMuc3ZjLmxvY2FsghpuZXh1cy41MS4yNTAu
    NDEuMi5zc2xpcC5pb4IbZG9ja2VyLjUxLjI1MC40MS4yLnNzbGlwLmlvMA0GCSqG
    SIb3DQEBCwUAA4IBAQBvTjTTXWeWtfaUDrcp1YW1pKgZ7lTb27f3QCxukXpbC+wL
    dcb4EP/vDf+UqCogKl6rCEA0i23Dtn85KAE9PQZFfI5hLulptdOgUhO3Udluoy36
    D4WvUoCfgPgx12FrdanQBBja+oDsT1QeOpKwQJuwjpZcGfB2YZqhO0UcJpC8kxtU
    by3uoxJoveHPRlbM2+ACPBPlHu/yH7st24sr1CodJHNt6P8ugIBAZxi3/Hq0wj4K
    aaQzdGXeFckWaxIny7F1M3cIWEXWzhAFnoTgrwlklf7N7VWHPIvlIh1EYASsVYKn
    iATq8C7qhUOGsknDh3QSpOJeJmpcBwln11/9BGRP
    -----END CERTIFICATE-----
    EOF
    )

    # bb-event           - Creating subscription for event function. More information: http://www.bashbooster.net/#event
    ## ca-file-updated   - Event name
    ## update-certs      - The function name that the event will call
    
    bb-event-on "ca-file-updated" "update-certs"
    
    update-certs() {          # Function with commands for adding a certificate to the store
      update-ca-certificates
    }

    # bb-tmp-file - Creating temp file function. More information: http://www.bashbooster.net/#tmp
    CERT_TMP_FILE="$( bb-tmp-file )"
    echo -e "${CERT_CONTENT}" > "${CERT_TMP_FILE}"  
    
    # bb-sync-file                                - File synchronization function. More information: http://www.bashbooster.net/#sync
    ## "${CERTS_FOLDER}/${CERT_FILE_NAME}.crt"    - Destination file
    ##  ${CERT_TMP_FILE}                          - Source file
    ##  ca-file-updated                           - Name of event that will be called if the file changes.

    bb-sync-file \
      "${CERTS_FOLDER}/${CERT_FILE_NAME}.crt" \
      ${CERT_TMP_FILE} \
      ca-file-updated   
```

### Добавление сертификата в ОС и containerd

{% alert level="warning" %}
Данный пример приведен для ОС Ubuntu.  
Способ добавления сертификатов в хранилище может отличаться в зависимости от ОС.
  
При адаптации скрипта под другую ОС измените параметры [bundles](cr.html#nodegroupconfiguration-v1alpha1-spec-bundles) и [content](cr.html#nodegroupconfiguration-v1alpha1-spec-content).
{% endalert %}

{% alert level="info" %}
Пример NodeGroupConfiguration основан на функциях, заложенных в скрипте [032_configure_containerd.sh](./#особенности-написания-скриптов).
{% endalert %}

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: NodeGroupConfiguration
metadata:
  name: add-custom-ca-containerd..sh
spec:
  weight: 31
  nodeGroups:
  - '*'  
  bundles:
  - 'ubuntu-lts'
  content: |-
    REGISTRY_URL=private.registry.example
    CERT_FILE_NAME=${REGISTRY_URL}
    CERTS_FOLDER="/usr/local/share/ca-certificates"
    CERT_CONTENT=$(cat <<EOF
    -----BEGIN CERTIFICATE-----
    MIIDSjCCAjKgAwIBAgIRAJ4RR/WDuAym7M11JA8W7D0wDQYJKoZIhvcNAQELBQAw
    JTEjMCEGA1UEAxMabmV4dXMuNTEuMjUwLjQxLjIuc3NsaXAuaW8wHhcNMjQwODAx
    MTAzMjA4WhcNMjQxMDMwMTAzMjA4WjAlMSMwIQYDVQQDExpuZXh1cy41MS4yNTAu
    NDEuMi5zc2xpcC5pbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL1p
    WLPr2c4SZX/i4IS59Ly1USPjRE21G4pMYewUjkSXnYv7hUkHvbNL/P9dmGBm2Jsl
    WFlRZbzCv7+5/J+9mPVL2TdTbWuAcTUyaG5GZ/1w64AmAWxqGMFx4eyD1zo9eSmN
    G2jis8VofL9dWDfUYhRzJ90qKxgK6k7tfhL0pv7IHDbqf28fCEnkvxsA98lGkq3H
    fUfvHV6Oi8pcyPZ/c8ayIf4+JOnf7oW/TgWqI7x6R1CkdzwepJ8oU7PGc0ySUWaP
    G5bH3ofBavL0bNEsyScz4TFCJ9b4aO5GFAOmgjFMMUi9qXDH72sBSrgi08Dxmimg
    Hfs198SZr3br5GTJoAkCAwEAAaN1MHMwDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB
    /wQCMAAwUwYDVR0RBEwwSoIPbmV4dXMuc3ZjLmxvY2FsghpuZXh1cy41MS4yNTAu
    NDEuMi5zc2xpcC5pb4IbZG9ja2VyLjUxLjI1MC40MS4yLnNzbGlwLmlvMA0GCSqG
    SIb3DQEBCwUAA4IBAQBvTjTTXWeWtfaUDrcp1YW1pKgZ7lTb27f3QCxukXpbC+wL
    dcb4EP/vDf+UqCogKl6rCEA0i23Dtn85KAE9PQZFfI5hLulptdOgUhO3Udluoy36
    D4WvUoCfgPgx12FrdanQBBja+oDsT1QeOpKwQJuwjpZcGfB2YZqhO0UcJpC8kxtU
    by3uoxJoveHPRlbM2+ACPBPlHu/yH7st24sr1CodJHNt6P8ugIBAZxi3/Hq0wj4K
    aaQzdGXeFckWaxIny7F1M3cIWEXWzhAFnoTgrwlklf7N7VWHPIvlIh1EYASsVYKn
    iATq8C7qhUOGsknDh3QSpOJeJmpcBwln11/9BGRP
    -----END CERTIFICATE-----
    EOF
    )
    CONFIG_CONTENT=$(cat <<EOF
    [plugins]
      [plugins."io.containerd.grpc.v1.cri".registry.configs."${REGISTRY_URL}".tls]
        ca_file = "${CERTS_FOLDER}/${CERT_FILE_NAME}.crt"
    EOF
    )
    
    mkdir -p /etc/containerd/conf.d

    # bb-tmp-file - Create temp file function. More information: http://www.bashbooster.net/#tmp

    CERT_TMP_FILE="$( bb-tmp-file )"
    echo -e "${CERT_CONTENT}" > "${CERT_TMP_FILE}"  
    
    CONFIG_TMP_FILE="$( bb-tmp-file )"
    echo -e "${CONFIG_CONTENT}" > "${CONFIG_TMP_FILE}"  

    # bb-event           - Creating subscription for event function. More information: http://www.bashbooster.net/#event
    ## ca-file-updated   - Event name
    ## update-certs      - The function name that the event will call
    
    bb-event-on "ca-file-updated" "update-certs"
    
    update-certs() {          # Function with commands for adding a certificate to the store
      update-ca-certificates  # Restarting the containerd service is not required as this is done automatically in the script 032_configure_containerd.sh
    }

    # bb-sync-file                                - File synchronization function. More information: http://www.bashbooster.net/#sync
    ## "${CERTS_FOLDER}/${CERT_FILE_NAME}.crt"    - Destination file
    ##  ${CERT_TMP_FILE}                          - Source file
    ##  ca-file-updated                           - Name of event that will be called if the file changes.

    bb-sync-file \
      "${CERTS_FOLDER}/${CERT_FILE_NAME}.crt" \
      ${CERT_TMP_FILE} \
      ca-file-updated   
      
    bb-sync-file \
      "/etc/containerd/conf.d/${REGISTRY_URL}.toml" \
      ${CONFIG_TMP_FILE} 
```

### Добавление в containerd возможности скачивать образы из insecure container registry

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: NodeGroupConfiguration
metadata:
  name: containerd-additional-registry.sh
spec:
  bundles:
    - '*'
  content: |
    REGISTRY_URL=private.registry.example
    mkdir -p /etc/containerd/conf.d
    bb-sync-file /etc/containerd/conf.d/additional_registry.toml - << EOF
    [plugins]
      [plugins."io.containerd.grpc.v1.cri"]
        [plugins."io.containerd.grpc.v1.cri".registry]
          [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
            [plugins."io.containerd.grpc.v1.cri".registry.mirrors."${REGISTRY_URL}"]
              endpoint = ["http://${REGISTRY_URL}"]
          [plugins."io.containerd.grpc.v1.cri".registry.configs]
            [plugins."io.containerd.grpc.v1.cri".registry.configs."${REGISTRY_URL}".auth]
              auth = "AAAABBBCCCDDD=="
            [plugins."io.containerd.grpc.v1.cri".registry.configs."${REGISTRY_URL}".tls]
              insecure_skip_verify = true
    EOF
  nodeGroups:
    - "*"
  weight: 31
```

Установите следующие параметры в ресурсе `NodeGroupConfiguration`:

* `REGISTRY_URL: <ADDITIONAL_REGISTRY_URL>` — адрес insecure container registry. Пример: `REGISTRY_URL=private.registry.example`;
* `auth: <BASE64>` — права доступа к стороннему registry, зашифрованные в Base64.

Если разрешен анонимный доступ к образам в стороннем registry, значение параметра `auth` должно выглядеть следующим образом:

```json
{"auths": { "<ADDITIONAL_REGISTRY>": {}}}
```

Приведенное значение должно быть закодировано в Base64.

Если для доступа к образам в стороннем registry необходима аутентификация, значение параметра `auth` должно выглядеть следующим образом:

```json
{"auths": { "<ADDITIONAL_REGISTRY>": {"username":"<ADDITIONAL_USERNAME>","password":"<ADDITIONAL_PASSWORD>","auth":"<AUTH_BASE64>"}}}
```

где:

* `<ADDITIONAL_USERNAME>` — имя пользователя для аутентификации на `<ADDITIONAL_REGISTRY>`;
* `<ADDITIONAL_PASSWORD>` — пароль пользователя для аутентификации на `<ADDITIONAL_REGISTRY>`;
* `<ADDITIONAL_REGISTRY>` — адрес стороннего registry в виде `<HOSTNAME>[:PORT]`;
* `<AUTH_BASE64>` — строка вида `<ADDITIONAL_USERNAME>:<ADDITIONAL_PASSWORD>`, закодированная в Base64.

Итоговое значение для `auth` должно быть также закодировано в Base64.

Вы можете использовать следующий скрипт для генерации `auth`:

```shell
declare MYUSER='<ADDITIONAL_USERNAME>'
declare MYPASSWORD='<ADDITIONAL_PASWORD>'
declare MYREGISTRY='<ADDITIONAL_REGISTRY>'

MYAUTH=$(echo -n "$MYUSER:$MYPASSWORD" | base64 -w0)
MYRESULTSTRING=$(echo -n "{\"auths\":{\"$MYREGISTRY\":{\"username\":\"$MYUSER\",\"password\":\"$MYPASSWORD\",\"auth\":\"$MYAUTH\"}}}" | base64 -w0)

echo "$MYRESULTSTRING"
```
