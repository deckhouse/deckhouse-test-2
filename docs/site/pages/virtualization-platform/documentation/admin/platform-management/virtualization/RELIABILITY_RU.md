---
title: "Механизмы обеспечения надежности"
permalink: ru/virtualization-platform/documentation/admin/platform-management/virtualization/reliability.html
lang: ru
---

## Миграция / Режим обслуживания

Миграция виртуальных машин — ключевая функция управления виртуализованной инфраструктурой, которая дает возможность переносить работающие виртуальные машины с одного физического узла на другой без необходимости их выключения. Этот процесс критичен для выполнения различных задач и сценариев:

- Балансировка нагрузки: перемещение виртуальных машин между узлами помогает равномерно распределить нагрузку, обеспечивая эффективное использование ресурсов серверов.
- Перевод узла в режим обслуживания: виртуальные машины можно перемещать с узлов, которые требуется вывести из эксплуатации для планового обслуживания или обновлений.
- Обновление «прошивки» виртуальных машин: миграция позволяет обновить «прошивку» виртуальных машин без остановки их работы.

### Запуск миграции произвольной машины

Далее рассмотрен пример миграции выбранной виртуальной машины.

Перед началом миграции проверьте текущий статус виртуальной машины:

```bash
d8 k get vm
# NAME                                   PHASE     NODE           IPADDRESS     AGE
# linux-vm                              Running   virtlab-pt-1   10.66.10.14   79m
```

Как видно, виртуальная машина в данный момент работает на узле `virtlab-pt-1`.

Для выполнения миграции виртуальной машины с одного узла на другой, с учетом требований по размещению, используется ресурс [VirtualMachineOperations](../../../reference/cr/virtualmachineoperations.html) (`vmop`) с типом `Evict`.

```yaml
d8 k apply -f - <<EOF
apiVersion: virtualization.deckhouse.io/v1alpha2
kind: VirtualMachineOperation
metadata:
  name: migrate-linux-vm-$(date +%s)
spec:
  # имя виртуальной машины
  virtualMachineName: linux-vm
  # операция для миграции
  type: Evict
EOF
```

После создания ресурса `vmop`, выполните команду:

```bash
d8 k get vm -w
# NAME                                   PHASE       NODE           IPADDRESS     AGE
# linux-vm                              Running     virtlab-pt-1   10.66.10.14   79m
# linux-vm                              Migrating   virtlab-pt-1   10.66.10.14   79m
# linux-vm                              Migrating   virtlab-pt-1   10.66.10.14   79m
# linux-vm                              Running     virtlab-pt-2   10.66.10.14   79m
```

Эта команда отображает статус виртуальной машины в процессе миграции. Она позволяет наблюдать за перемещением виртуальной машины с одного узла на другой.

#### Режим обслуживания

При проведении работ на узлах, на которых работают виртуальные машины, существует риск нарушения их работоспособности. Чтобы избежать этого, узел можно перевести в режим обслуживания, предварительно переместив все виртуальные машины на другие доступные узлы.

Для перевода узла в режим обслуживания и перемещения виртуальных машин выполните следующую команду:

```bash
d8 k drain <nodename> --ignore-daemonsets --delete-emptydir-data
```

где `<nodename>` — это имя узла, на котором будут проводиться работы, и который нужно очистить от всех ресурсов, включая системные.

Эта команда выполняет несколько задач:

- Она эвакуирует все поды с указанного узла.
- Игнорирует DaemonSet'ы, чтобы не остановить критически важные сервисы.
- Удаляет временные данные, хранящиеся в emptyDir, чтобы освободить ресурсы узла.

Если необходимо вытеснить с узла только виртуальные машины, можно использовать более точную команду с фильтрацией по метке, которая соответствует виртуальным машинам. Для этого выполните следующую команду:

```bash
d8 k drain <nodename> --pod-selector vm.kubevirt.internal.virtualization.deckhouse.io/name --delete-emptydir-data
```

После выполнения команды узел перейдет в режим обслуживания, и виртуальные машины на нем запускаться больше не будут. Чтобы вывести узел из режима обслуживания и вернуть его в рабочее состояние, выполните команду:

```bash
d8 k uncordon <nodename>
```

![Режим обслуживания, схема](/images/virtualization-platform/drain.ru.png)

### Восстановление после сбоя

ColdStandby обеспечивает механизм восстановления работы виртуальной машины в случае сбоя узла, на котором она была запущена.

Для работы данного механизма необходимо выполнить следующие требования:

- Политика запуска виртуальной машины (`.spec.runPolicy`) должна быть установлена в одно из следующих значений: `AlwaysOnUnlessStoppedManually`, `AlwaysOn`.
- На узлах, где запускаются виртуальные машины, должен быть активирован механизм [fencing](../../../../reference/cr/nodegroup.html#nodegroup-v1-spec-fencing-mode).

Рассмотрим, как работает механизм ColdStandby на примере:

1. Кластер состоит из трех узлов **master**, **workerA** и **workerB**. На worker-узлах включен механизм Fencing.
1. Виртуальная машина `linux-vm` изначально запущена на узле  **workerA**.
1. На узле **workerA** происходит сбой (например, отключение питания, потеря сети и т.д.).
1. Контроллер Kubernetes проверяет доступность узлов и обнаруживает, что **workerA** недоступен.
1. Контроллер удаляет недоступный узел **workerA** из кластера.
1. Виртуальная машина `linux-vm` автоматически запускается на другом доступном узле — **workerB**.

![Восстановление после сбоя, схема](/images/virtualization-platform/coldstandby.ru.png)
