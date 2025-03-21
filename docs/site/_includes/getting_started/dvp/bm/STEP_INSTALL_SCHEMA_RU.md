## Что необходимо для установки

1. **Персональный компьютер.** Компьютер, с которого будет производиться установка. Он нужен только для запуска инсталлятора и не будет частью кластера.

   Требования...

   - ОС: Windows 10+, macOS 10.15+, Linux (Ubuntu 20.04+, Fedora 35+);
   - установленный docker для запуска инсталлятора Deckhouse (инструкции для [Ubuntu](https://docs.docker.com/engine/install/ubuntu/), [macOS](https://docs.docker.com/desktop/mac/install/), [Windows](https://docs.docker.com/desktop/windows/install/));
   - HTTPS-доступ до хранилища образов контейнеров `registry.deckhouse.ru`;
   - SSH-доступ по ключу до узла, который будет **master-узлом** будущего кластера;
   - SSH-доступ по ключу до узла, который будет **worker-узлом** будущего кластера.

1. **Физический сервер или виртуальная машина для master-узла.**

   Требования...

   - не менее 4 ядер CPU;
   - не менее 8 ГБ RAM;
   - не менее 60 ГБ дискового пространства на быстром диске (400+ IOPS);
   - [поддерживаемая ОС](/products/virtualization-platform/documentation/admin/install/requirements.html#поддерживаемые-ос-для-узлов-платформы);
   - ядро Linux версии `5.7` или новее;
   - ЦП с архитектурой x86_64 с поддержкой инструкций Intel-VT (VMX) или AMD-V (SVM);
   - **Уникальный hostname** в пределах серверов (виртуальных машин) кластера;
   - HTTPS-доступ до хранилища образов контейнеров `registry.deckhouse.ru`;
   - доступ к стандартным для используемой ОС репозиториям пакетов;
   - SSH-доступ от **персонального компьютера** (см. п.1) по ключу;
   - сетевой доступ от **персонального компьютера** (см. п.1) по порту `22322/TCP`;
   - на узле не должно быть установлено пакетов container runtime, например containerd или Docker;
   - на узле должны быть установлены пакеты `cloud-utils` и `cloud-init`.

1. **Физический сервер или виртуальная машина для worker-узла.**

   Требования аналогичны требованиям к master-узлу, но также зависят от характера запускаемой на узлах нагрузки.
   Также на worker-узлах требуется дополнительные диски для развертывания программно определяемого хранилища.
