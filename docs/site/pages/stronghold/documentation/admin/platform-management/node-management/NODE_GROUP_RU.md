---
title: "Группы узлов"
permalink: ru/stronghold/documentation/admin/platform-management/node-management/node-group.html
lang: ru
---

## Управление узлами кластера

Управление узлами осуществляется с помощью модуля `node-manager`, основные функции которого:

1. Управление несколькими узлами как связанной группой (NodeGroup):
   * Возможность задать метаданные, которые будут применяться ко всем узлам в группе.
   * Мониторинг группы узлов как целого объекта (сегментация узлов на графиках, агрегация алертов о недоступности узлов, а также уведомления при недоступности определенного числа узлов или их процента в группе).

1. Установка и обновление и настройка ПО узла (containerd, kubelet и др.), подключение узла в кластер:
   * Установка операционной системы (см. [список поддерживаемых ОС](../../../about/requirements.html#поддерживаемые-ос)) вне зависимости от типа используемой инфраструктуры (в любом облаке или на физическом оборудовании).
   * Базовая настройка операционной системы (отключение автообновления, установка необходимых пакетов, настройка параметров журналирования и т.д.).
   * Конфигурация Nginx для балансировки запросов от узлов (kubelet) между API-серверами, включая настройку автоматического обновления списка upstream-серверов.
   * Установка и настройка CRI containerd и Kubernetes, включение узла в кластер.
   * Управление обновлениями узлов и их простоем (disruptions):
     * Автоматическое определение допустимой минорной версии Kubernetes для группы узлов, исходя из её конфигурации (`kubernetesVersion`), версии по умолчанию для всего кластера и текущей версии control plane. Обновление узлов не допускается, если оно опережает обновление control plane.
     * Из группы одновременно производится обновление только одного узла и только если все узлы группы доступны.
     * Два варианта обновлений узлов:
       * обычные — всегда происходят автоматически;
       * требующие прерывания работы (disruption), например: обновление ядра, смена версии containerd, значительная смена версии kubelet и пр.При разрешении автоматических disruptive-обновлений перед обновлением выполняется процесс drain узла (можно отключить).
   * Мониторинг состояния и прогресса обновления.

1. Масштабирование кластера.
   * В рамках платформы виртуализации доступно поддержание желаемого количества узлов в группе при использовании [Cluster API Provider Static](#работа-со-статическими-узлами).

1. Управление Linux-пользователями на узлах.

## Типы узлов

Платформа виртуализации предполагает запуск на bare-metal серверах, поэтому далее будет освещено управление `Static` узлами.

Ознакомиться с другими типами узлов и возможностями работы с облачными провайдерами можно в документации платформы Deckhouse.

## Группа узлов

Для управления узлами в системе используются группы узлов, которые описываются с помощью ресурсов [NodeGroup](../../../../reference/cr/nodegroup.html). Каждая группа узлов выполняет свои специфические задачи, например:

- группа для control plane компонентов Kubernetes;
- группа для компонентов мониторинга;
- группа для control plane компонентов платформы виртуализации;
- группа узлов с виртуальными машинами (vm-worker-узлы);
- группа узлов с контейнерными приложениями (worker-узлы) и т.п.

Разбиение узлов по группам и распределение компонентов по группам узлов зависит от задач кластера. Примеры конфигураций кластера платформы виртуализации можно посмотреть в разделе [Установка платформы](../../install/steps/install.html).

Узлы в группе обладают общими метаданными и параметрами, что позволяет настроить их автоматически в соответствии с конфигурацией группы. Deckhouse также отслеживает количество узлов в группе и выполняет обновления программного обеспечения на них.

Для таких групп узлов доступны следующие функции мониторинга:

- с группировкой параметров узлов на графиках группы;
- с группировкой алертов о недоступности узлов;
- с алертами о недоступности определённого числа или процента узлов в группе.

## Развертывание, настройка и обновление узлов Kubernetes

### Развертывание узлов Kubernetes

Deckhouse автоматически выполняет следующие неизменяемые операции для развертывания узлов кластера:

1. Настройка и оптимизация операционной системы для работы с containerd и Kubernetes:
   - устанавливаются необходимые пакеты из репозиториев соответствующего дистрибутива.
   - настраиваются параметры работы ядра, параметры журналирования, ротация журналов и другие параметры системы.
1. Установка требуемых версий containerd и kubelet, включение узла в кластер Kubernetes.
1. Настройка Nginx и обновление списка upstream для балансировки запросов от узла к Kubernetes API.

### Поддержка актуального состояния узлов

Для поддержания узлов кластера в актуальном состоянии могут применяться два типа обновлений:

- **Обычные** — такие обновления всегда применяются автоматически, и не приводят к остановке или перезагрузке узла.
- **Требующие прерывания** (disruption) — например, обновление версии ядра или containerd, значительная смена версии kubelet и т.д. Для этого типа обновлений можно выбрать ручной или автоматический режим (секция параметров [disruptions](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-disruptions)). В автоматическом режиме перед обновлением выполняется корректная приостановка работы узла (drain) и только после этого производится обновление.

В любой момент времени обновляется только один узел из группы, и это возможно только в случае, когда все узлы группы находятся в доступном состоянии.

Модуль `node-manager` имеет набор встроенных метрик мониторинга, которые позволяют контролировать прогресс обновления, получать уведомления о возникающих во время обновления проблемах или о необходимости получения разрешения на обновление (ручное подтверждение обновления).

## Работа со статическими узлами

### Ограничения

При работе со статическими узлами функции модуля `node-manager` выполняются со следующими ограничениями:

- **Отсутствует заказ узлов** — непосредственное выделение ресурсов (серверов bare-metal, виртуальных машин, связанных ресурсов) выполняется вручную. Дальнейшая настройка ресурсов (подключение узла к кластеру, настройка мониторинга и т.п.) выполняется  автоматически или частично.
- **Отсутствует автоматическое масштабирование узлов** — доступно поддержание в группе указанного количества узлов при использовании [Cluster API Provider Static](#работа-со-статическими-узлами) (параметр [staticInstances.count](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-count)). Deckhouse будет пытаться поддерживать указанное количество узлов в группе, очищая лишние узлы и настраивая новые при необходимости (выбирая их из ресурсов [StaticInstance](../../../../reference/cr/staticinstance.html), находящихся в состоянии *Pending*).

### Ручное управление статическим узлом

Настройка или очистка узла, его подключение к кластеру и отключение могут выполняться с помощью подготовленных скриптов.

Для настройки сервера (ВМ) и ввода узла в кластер нужно загрузить и выполнить специальный бутстрап скрипт. Такой скрипт генерируется для каждой группы статических узлов (каждого ресурса `NodeGroup`). Он находится в секрете `d8-cloud-instance-manager/manual-bootstrap-for-<ИМЯ-NODEGROUP>`. Пример добавления статического узла в кластер можно найти в [FAQ](examples.html#вручную).

Для отключения узла кластера и очистки сервера (виртуальной машины) нужно выполнить скрипт `/var/lib/bashible/cleanup_static_node.sh`, который уже находится на каждом статическом узле. Пример отключения узла кластера и очистки сервера можно найти в [FAQ](faq.html#как-вручную-очистить-статический-узел).

### Автоматическое управление узлом

Автоматическое управление статическим узлом происходит с помощью [Cluster API Provider Static](#работа-со-статическими-узлами).

Cluster API Provider Static (CAPS) подключается к серверу (ВМ) используя ресурсы [StaticInstance](../../../../reference/cr/staticinstance.html) и [SSHCredentials](../../../../reference/cr/sshcredentials.html), выполняет настройку, и вводит узел в кластер.

При необходимости (например, если удален соответствующий серверу ресурс [StaticInstance](../../../../reference/cr/staticinstance.html) или уменьшено [количество узлов группы](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-count)), Cluster API Provider Static подключается к узлу кластера, очищает его и отключает от кластера.

### Автоматическое управление существующим узлом

> Поддерживается в версиях Deckhouse 1.63 и выше.

Для передачи существующего узла кластера под управление CAPS, необходимо подготовить для этого узла ресурсы [StaticInstance](../../../../reference/cr/staticinstance.html) и [SSHCredentials](../../../../reference/cr/sshcredentials.html), как при автоматическом управлении в пункте выше, однако ресурс [StaticInstance](../../../../reference/cr/staticinstance.html) должен дополнительно быть помечен аннотацией `static.node.deckhouse.io/skip-bootstrap-phase: ""`.

### Настройка узла через CAPS

Cluster API Provider Static (CAPS), это реализация провайдера декларативного управления статическими узлами (серверами bare-metal или виртуальными машинами) для проекта [Cluster API](https://cluster-api.sigs.k8s.io/) Kubernetes. По сути, CAPS это дополнительный слой абстракции к уже существующему функционалу Deckhouse по автоматической настройке и очистке статических узлов с помощью скриптов, генерируемых для каждой группы узлов (см. раздел [работа со статическими узлами](#работа-со-статическими-узлами)).

CAPS выполняет следующие функции:

- настройка сервера bare-metal (или виртуальной машины) для подключения к кластеру Kubernetes;
- подключение узла в кластер Kubernetes;
- отключение узла от кластера Kubernetes;
- очистка сервера bare-metal (или виртуальной машины) после отключения узла из кластера Kubernetes.

CAPS использует следующие ресурсы (CustomResource) при работе:

- **[StaticInstance](../../../../reference/cr/staticinstance.html).** Каждый ресурс `StaticInstance` описывает конкретный хост (сервер, ВМ), который управляется с помощью CAPS.
- **[SSHCredentials](../../../reference/cr/sshcredentials.html)**. Содержит данные SSH, необходимые для подключения к хосту (`SSHCredentials` указывается в параметре [credentialsRef](../../../../cr/staticinstance.html#staticinstance-v1alpha1-spec-credentialsref) ресурса `StaticInstance`).
- **[NodeGroup](../../../reference/cr/nodegroup.html)**. Секция параметров [staticInstances](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances) определяет необходимое количество узлов в группе и фильтр множества ресурсов `StaticInstance` которые могут использоваться в группе.

CAPS включается автоматически, если в NodeGroup заполнена секция параметров [staticInstances](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances). Если в `NodeGroup` секция параметров `staticInstances` не заполнена, то настройка и очистка узлов для работы в этой группе выполняется вручную (см. примеры [добавления статического узла в кластер](examples.html#вручную) и [очистки узла](faq.html#как-вручную-очистить-статический-узел)), а не с помощью CAPS.

Схема работы со статическими узлами при использовании CAPS:

1. **Подготовка ресурсов.**

Перед тем, как передать сервер bare-metal или виртуальную машину под управление CAPS, может понадобиться предварительная подготовка, например:

- Подготовка системы хранения, добавление точек монтирования и т.п.;
- Установка специфических пакетов ОС. Например, установка пакета `ceph-common`, если на сервере используется тома CEPH;
- Настройка необходимой сетевой связанности. Например, между сервером и узлами кластера;
- Настройка доступа по SSH на сервер, создание пользователя для управления с root-доступом через `sudo`. Хорошей практикой является создание отдельного пользователя и уникальных ключей для каждого сервера.

**Создание ресурса [SSHCredentials](../../../../reference/cr/sshcredentials.html).**

В ресурсе `SSHCredentials` указываются параметры, необходимые CAPS для подключения к серверу по SSH. Один ресурс `SSHCredentials` может использоваться для подключения к нескольким серверам, но хорошей практикой является создание уникальных пользователей и ключей доступа для подключения к каждому серверу. В этом случае ресурс `SSHCredentials` будет отдельным на каждый сервер.

**Создание ресурса [StaticInstance](../../../../reference/cr/staticinstance.html).**

На каждый сервер (ВМ) в кластере создается отдельный ресурс `StaticInstance`. В нем указан IP-адрес для подключения и ссылка на ресурс `SSHCredentials`, данные которого нужно использовать при подключении.

Возможные состояния `StaticInstances` и связанных с ним серверов (ВМ) и узлов кластера:

- `Pending`. Сервер не настроен, и в кластере нет соответствующего узла.
- `Bootstraping`. Выполняется процедура настройки сервера (ВМ) и подключения узла в кластер.
- `Running`. Сервер настроен, и в кластер добавлен соответствующий узел.
- `Cleaning`. Выполняется процедура очистки сервера и отключение узла из кластера.

> Можно передать существующий узел кластера, заранее введенный в кластер вручную, под управление CAPS, пометив его StaticInstance аннотацией `static.node.deckhouse.io/skip-bootstrap-phase: ""`.

**Создание ресурса [NodeGroup](../../../../reference/cr/nodegroup.html).**

В контексте CAPS в ресурсе `NodeGroup` нужно обратить внимание на параметр [nodeType](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-nodetype) (должен быть `Static`) и секцию параметров [staticInstances](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances).

Секция параметров [staticInstances.labelSelector](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-labelselector) определяет фильтр, по которому CAPS выбирает ресурсы `StaticInstance`, которые нужно использовать в группе. Фильтр позволяет использовать для разных групп узлов только определенные `StaticInstance`, а также позволяет использовать один `StaticInstance` в разных группах узлов. Фильтр можно не определять, чтобы использовать в группе узлов любой доступный `StaticInstance`.

Параметр [staticInstances.count](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-count) определяет желаемое количество узлов в группе.  При изменении параметра, CAPS начинает добавлять или удалять необходимое количество узлов, запуская этот процесс параллельно.

В соответствии с данными секции параметров [staticInstances](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances), CAPS будет пытаться поддерживать указанное (параметр [count](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-count))) количество узлов в группе. При необходимости добавить узел в группу, CAPS выбирает соответствующий [фильтру](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-staticinstances-labelselector) ресурс StaticInstance находящийся в статусе `Pending`, настраивает сервер (ВМ) и добавляет узел в кластер. При необходимости удалить узел из группы, CAPS выбирает StaticInstance находящийся в статусе `Running`, очищает сервер и удаляет узел из кластера (после чего, соответствующий StaticInstance переходит в состояние `Pending` и снова может быть использован).

[Пример добавления узла](examples.html#с-помощью-cluster-api-provider-static).

## Как интерпретировать состояние группы узлов?

**Ready** — группа узлов содержит минимально необходимое число запланированных узлов с состоянием `Ready` для всех зон.

Пример 1. Группа узлов в состоянии `Ready`:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: ng1
spec:
  nodeType: CloudEphemeral
  cloudInstances:
    maxPerZone: 5
    minPerZone: 1
status:
  conditions:
  - status: "True"
    type: Ready
---
apiVersion: v1
kind: Node
metadata:
  name: node1
  labels:
    node.deckhouse.io/group: ng1
status:
  conditions:
  - status: "True"
    type: Ready
```

Пример 2. Группа узлов в состоянии `Not Ready`:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: ng1
spec:
  nodeType: CloudEphemeral
  cloudInstances:
    maxPerZone: 5
    minPerZone: 2
status:
  conditions:
  - status: "False"
    type: Ready
---
apiVersion: v1
kind: Node
metadata:
  name: node1
  labels:
    node.deckhouse.io/group: ng1
status:
  conditions:
  - status: "True"
    type: Ready
```

**Updating** — группа узлов содержит как минимум один узел, в котором присутствует аннотация с префиксом `update.node.deckhouse.io` (например, `update.node.deckhouse.io/waiting-for-approval`).

**WaitingForDisruptiveApproval** — группа узлов содержит как минимум один узел, в котором присутствует аннотация `update.node.deckhouse.io/disruption-required` и
отсутствует аннотация `update.node.deckhouse.io/disruption-approved`.

**Scaling** — рассчитывается только для групп узлов с типом `CloudEphemeral`. Состояние `True` может быть в двух случаях:

1. Когда число узлов меньше желаемого числа узлов в группе, то есть когда нужно увеличить число узлов в группе.
1. Когда какой-то узел помечается к удалению или число узлов больше желаемого числа узлов, то есть когда нужно уменьшить число узлов в группе.

Желаемое число узлов — это сумма всех реплик, входящих в группу узлов.

Пример. Желаемое число узлов равно 2:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: ng1
spec:
  nodeType: CloudEphemeral
  cloudInstances:
    maxPerZone: 5
    minPerZone: 2
status:
...
  desired: 2
...
```

**Error** — содержит последнюю ошибку, возникшую при создании узла в группе узлов.

## Как влияют параметры NodeGroup

| Параметр NG                           | Disruption update          | Перезаказ узлов   | Рестарт kubelet |
|---------------------------------------|----------------------------|-------------------|-----------------|
| chaos                                 | -                          | -                 | -               |
| cloudInstances.classReference         | -                          | +                 | -               |
| cloudInstances.maxSurgePerZone        | -                          | -                 | -               |
| cri.containerd.maxConcurrentDownloads | -                          | -                 | +               |
| cri.type                              | - (NotManaged) / + (other) | -                 | -               |
| disruptions                           | -                          | -                 | -               |
| kubelet.maxPods                       | -                          | -                 | +               |
| kubelet.rootDir                       | -                          | -                 | +               |
| kubernetesVersion                     | -                          | -                 | +               |
| nodeTemplate                          | -                          | -                 | -               |
| static                                | -                          | -                 | +               |
| update.maxConcurrent                  | -                          | -                 | -               |

Подробно о всех параметрах можно прочитать в описании Custom Resource [NodeGroup](../../../../reference/cr/nodegroup).

В случае изменения параметров `instanceClass` или `instancePrefix` в конфигурации Deckhouse не будет происходить `RollingUpdate`. Deckhouse создаст новые `MachineDeployment`, а старые удалит. Количество заказываемых одновременно `MachineDeployment` определяется параметром `cloudInstances.maxSurgePerZone`.

При обновлении, которое требует прерывания работы узла (disruption update), выполняется вытеснение подов с данного узла. Если некоторые поды не удалось вытеснить, попытки повторяются каждые 20 секунд в течение максимум 5 минут. По истечении этого времени поды, которые не удалось вытеснить, удаляются принудительно.

## Как выделить узлы под специфические нагрузки?

{% alert level="warning" %}
Запрещено использование домена `deckhouse.io` в ключах labels и taints у NodeGroup. Он зарезервирован для компонентов **Deckhouse**. Следует отдавать предпочтение в пользу ключей dedicated или dedicated.client.com.
{% endalert %}

Для решений данной задачи существуют два механизма:

1. Установка меток в NodeGroup `spec.nodeTemplate.labels` для последующего использования их в [spec.nodeSelector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/) или [spec.affinity.nodeAffinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity). Указывает, какие именно узлы будут выбраны планировщиком для запуска целевого приложения.
1. Установка ограничений в NodeGroup `spec.nodeTemplate.taints` с дальнейшим снятием их в [spec.tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). Запрещает исполнение не разрешенных явно приложений на этих узлах.

> Deckhouse по умолчанию допускает использование taint с ключом `dedicated`. Поэтому рекомендуется применять именно этот ключ с любым значением для taint на выделенных узлах.
> Если необходимо использовать нестандартные ключи taint (например, dedicated.client.com), нужно добавить их в параметр [modules.placement.customTolerationKeys](../../../../reference/mc.html#global-parameters-modules-placement-customtolerationkeys). Это позволит системным компонентам, таким как `cni-flannel`, работать на этих узлах.

Подробнее [в статье на Habr](https://habr.com/ru/company/flant/blog/432748/).

### Системные

Компоненты Deckhouse используют лейблы и тэйнты для назначения узлов. Системные компоненты могут быть назначены на отдельные узлы с помощью такой `NodeGroup`:

```yaml
nodeTemplate:
  labels:
    node-role.deckhouse.io/system: ""
  taints:
    - effect: NoExecute
      key: dedicated.deckhouse.io
      value: system
```

<!-- TODO ссылка про подробности куда-то в dkp? Или в этой доке сделать раздел? -->

<!-- ### Компоненты controler plane виртуализации

TODO Надо какую-то группу придумать для control plane компонентов виртуализации. Или они просто на system выезжают? -->

## Как выделить узлы под виртуальные машины?

Чтобы виртуальные машины запускались на узлах определённой группы, помимо создания самой группы, нужен ресурс VirtualMachineClass с nodeSelector.

Например, для группы vm-workers это может выглядеть так:

```yaml
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: vm-worker
spec:
  nodeType: Static
```

VirtualMachineClass с nodeSelector для группы vm-worker:

```yaml
apiVersion: virtualization.deckhouse.io/v1alpha2
kind: VirtualMachineClass
metadata:
  name: vm-worker
spec:
  nodeSelector:
    matchExpressions:
    - key: node.deckhouse.io/group
      operator: In
      values:
        - vm-worker
```

Фрагмент манифеста виртуальной машины, которая будет запускаться на узлах группы vm-worker:

```yaml
apiVersion: virtualization.deckhouse.io/v1alpha2
kind: VirtualMachine
metadata:
  name: vm-name
spec:
  virtualMachineClassName: vm-workers
  # more VM fields ...
```
