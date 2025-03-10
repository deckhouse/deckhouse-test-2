---
title: "Список алертов"
permalink: ru/alerts.html
lang: ru
---

На странице приведен список всех алертов мониторинга в Deckhouse Kubernetes Platform.

Алерты сгруппированы по модулям. Справа от названия алерта указаны иконки минимальной редакции DKP в которой доступен алерт и уровня критичности алерта (severity).

Для каждого алерта приведено краткое описание (summary), раскрыв которое можно увидеть подробное описание алерта (description), при их наличии.

## Критичность алерта

В описании алертов присутствует параметр Severity (S), означающий уровень критичности. Его значение варьируется от `S1` до `S9` и может расцениваться следующим образом:

* `S1` — максимальный уровень, критический сбой/авария (требуются незамедлительные действия);
* `S2` — высокий уровень, близкий к максимальному, возможная авария (необходимо быстрое реагирование);
* `S3` — средний уровень, потенциально серьёзная проблема (необходима проверка);
* `S4`-`S9` — низкий уровень. Есть проблема, но в целом работоспособность не нарушена.

{% include deckhouse-alerts.liquid %}
