---
title: "Модуль l2-load-balancer"
---

Модуль реализует улучшенный (относительно стандартного [режима L2 в MetalLB](../380-metallb/#режим-layer-2)) механизм балансировки для сервисов в кластерах bare metal, где нет возможности воспользоваться облачным балансировщиком или [MetalLB](../380-metallb/#режим-bgp) в режиме BGP с настроенным Equal-cost multi-path (ECMP).

## Принцип работы в сравнении с режимом L2 модуля MetalLB

[MetalLB в режиме L2](../380-metallb/#режим-layer-2) позволяет заказать _Service_ с типом `LoadBalancer`, работа которого основана на том, что в пиринговой сети узлы-балансировщики имитируют ARP-ответы от "публичного" IP. Данный режим имеет существенное ограничение — единовременно лишь один балансировочный узел обрабатывает весь входящий трафик данного сервиса. Соответственно:

* Узел, выбранный в качестве лидера для "публичного" IP становится "узким местом" без возможности горизонтального масштабирования.
* В случае выхода узла-балансировщика из строя, все текущие соединения будут оборваны для переключения на новый узел, который будет назначен лидером.

<div data-presentation="../../presentations/381-l2-load-balancer/basics_metallb_ru.pdf"></div>
<!--- Source: https://docs.google.com/presentation/d/1cs1uKeX53DB973EMtLFcc8UQ8BFCW6FY2vmEWua1tu8/ --->

Данный модуль помогает обойти эти ограничения. Он предоставляет новый интерфейс _L2LoadBalancer_, который:

* Позволяет с помощью `nodeSelector` связать группу узлов и пул IP-адресов.

После чего мы можем создать стандартный ресурс _Service_ с типом `LoadBalancer` и в нем с помощью аннотаций указать:

* Какой _L2LoadBalancer_ использовать. Задав этим набор узлов и адресов.
* Указать необходимое количество IP-адресов для L2-анонсирования.

<div data-presentation="../../presentations/381-l2-load-balancer/basics_l2loadbalancer_new_ru.pdf"></div>
<!--- Source: https://docs.google.com/presentation/d/1jDqC4Bhg5NMLZWaFM32bIAzqpkUo0hOkAaRzC0yKRxE/ --->

Таким образом:
* Приложение получит не один, а несколько "публичных" IP. Данные IP потребуется прописать в качестве A-записей для публичного домена приложения. Для последующего горизонтального масштабирования потребуется добавить дополнительные узлы-балансировщики, соответствующие _Service_ будут созданы автоматически, потребуется лишь добавить их в список A-записей прикладного домена.
* При выходе из строя одного из узлов-балансировщиков, лишь часть трафика будет подвержена обрыву для переключения на здоровый узел.
