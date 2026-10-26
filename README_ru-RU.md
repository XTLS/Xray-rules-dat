# GeoSite

[English](./README.md) | [简体中文](./README_zh-Hans.md) | Русский | [فارسی](./README_fa-IR.md)

Этот репозиторий предназначен для сборки пакетов данных GeoSite на основе исходных классификаций доменов из [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community). Для этого используется адаптированный кодировщик данных из `v2fly/domain-list-community`. Исходные классификации дополняются данными из дополнительных источников, после чего формируются наборы данных, соответствующие семантике GeoSite.

## Принципы проектирования

Точность и полнота данных GeoSite влияют на эффективность маршрутизации трафика по доменам. На результат маршрутизации также влияет семантическая модель классификации GeoSite. Поэтому в этом репозитории используется портированный кодировщик данных из `v2fly/domain-list-community`. Он позволяет формировать соответствующие категории доменов, максимально сохраняя исходную семантику данных и обеспечивая совместимость с изменениями в upstream-проекте.

Помимо полного набора исходных классификаций доменов, в репозиторий добавлены дополнительные списки доменов, которые могут быть полезны на практике. Эти списки преобразуются в наборы данных, соответствующие семантике GeoSite. Таким образом, они дополняют исходные данные и расширяют их охват.

Чтобы найти баланс между полнотой наборов данных и размером закодированных файлов, предусмотрены две сборки: стандартная (`geosite.dat`) и расширенная (`geosite-expand.dat`). Кроме того, для отслеживания изменений и проверки соответствия поведения инструментов сборки upstream-проекту предусмотрена отдельная проверочная сборка (`validation.dat`). Она содержит только исходные данные upstream-проекта и предназначена для проверки процесса кодирования и сборки без дополнительных классификаций.

## Состав данных

Проверочная сборка (`validation.dat`) содержит только данные из `v2fly/domain-list-community`.

### Стандартная сборка: `geosite.dat`

| Категория | Описание | Входит в | Содержит | Примечания |
|---|---|---|---|---|
| Индексы данных, соответствующие `domain-list-community` | Полный набор данных из `v2fly/domain-list-community` | | | Соответствующие индексы данных не изменялись и не корректировались. |
| `gfwlist-cn` | Данные из GFWList, преобразованные утилитой `abpextract` в соответствии с семантикой GeoSite | `blocked-in-cn` | | |
| `mainly-ir` | Проверенный набор доменов из `Chocolate4U/Iran-v2ray-rules`, предназначенный для обхода ограничений внутри Ирана | | | Соответствует категории `category-ir` в `Chocolate4U/Iran-v2ray-rules`. |
| `available-only-in-ru` | Домены, доступные только на территории России, из `runetfreedom/russia-domains-list` | | | Соответствует категории `ru-available-only-inside` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |
| `refilter` | Полный список доменов из Re:filter | `blocked-in-ru` | | |
| `antifilter-list-community` | Список доменов из сообщества Antifilter.Download | `blocked-in-ru` | | Соответствует категории `antifilter-download-community` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |
| `blocked-in-cn` | Список заблокированных доменов для материкового Китая | | `gfwlist-cn` | |
| `blocked-in-ru` | Список заблокированных доменов для России | | `refilter`, `antifilter-list-community` | Соответствует категории `ru-blocked` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |

### Расширенная сборка: `geosite-expand.dat`

| Категория | Описание | Входит в | Содержит | Примечания |
|---|---|---|---|---|
| Индексы данных, соответствующие `domain-list-community` | Полный набор данных из `v2fly/domain-list-community` | | | Соответствующие индексы данных не изменялись и не корректировались. |
| `gfwlist-cn` | Данные из GFWList, преобразованные утилитой `abpextract` в соответствии с семантикой GeoSite | `blocked-in-cn`, `blocked-in-cn-expand` | | |
| `expand-cn` | Домены материкового Китая из `felixonmars/dnsmasq-china-list`, преобразованные утилитой `dnsmasqextract` в список доменов | `additional-set-cn` | | Не рекомендуется использовать одновременно с `cn`, `tld-cn` или `geosite-cn`. |
| `greatfire-list` | Данные из GreatFire.org, преобразованные проектом `Loyalsoldier/cn-blocked-domain` | `blocked-in-cn-expand` | | |
| `mainly-ir` | Проверенный набор доменов из `Chocolate4U/Iran-v2ray-rules`, предназначенный для обхода ограничений внутри Ирана | | | Соответствует категории `category-ir` в `Chocolate4U/Iran-v2ray-rules`. |
| `expand-ir` | Набор доменов для обхода ограничений внутри Ирана из Iran Hosted Domains | | | Соответствует категории `ir` в `Chocolate4U/Iran-v2ray-rules`. |
| `available-only-in-ru` | Домены, доступные только на территории России, из `runetfreedom/russia-domains-list` | | | Соответствует категории `ru-available-only-inside` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |
| `refilter` | Полный список доменов из Re:filter | `blocked-in-ru`, `blocked-in-ru-expand` | | |
| `antifilter-list-community` | Список доменов из сообщества Antifilter.Download | `blocked-in-ru`, `blocked-in-ru-expand` | | Соответствует категории `antifilter-download-community` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |
| `antifilter-list-full` | Список доменов из Antifilter.Download | `blocked-in-ru-expand` | | Соответствует категории `antifilter-download` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. Содержит очень много записей; используйте с осторожностью. |
| `blocked-in-cn` | Список заблокированных доменов для материкового Китая | `blocked-in-cn-expand` | `gfwlist-cn` | |
| `blocked-in-cn-expand` | Расширенный список заблокированных доменов для материкового Китая | | `blocked-in-cn`, `greatfire-list` | |
| `blocked-in-ru` | Список заблокированных доменов для России | `blocked-in-ru-expand` | `refilter`, `antifilter-list-community` | Соответствует категории `ru-blocked` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. |
| `blocked-in-ru-expand` | Расширенный список заблокированных доменов для России | | `blocked-in-ru`, `antifilter-list-full` | Соответствует категории `ru-blocked-all` в `russia-blocked-geosite` и `russia-v2ray-rules-dat`. Содержит очень много записей; используйте с осторожностью. |

## Сборка

В корне репозитория необходимо подготовить каталог `./data`. В нём должны находиться исходные файлы классификаций доменов после преобразования и очистки.

Каталоги `./predata/std` и `./predata/exp` содержат дополнительные индексы для стандартной и расширенной сборок соответственно. Чтобы включить эти индексы, скопируйте соответствующие файлы в `./data`. При этом в каталоге данных уже должна существовать соответствующая категория доменов, иначе дополнительный индекс не будет применён.

Затем выполните следующую команду, чтобы создать закодированный файл `geosite.dat` в текущем каталоге:

```sh
go run ./main.go --datapath ./data --outputname geosite.dat --outputdir ./
```

``` sh
go run ./main.go --datapath ./data --datprofile ./tasks/standard.json --outputdir ./
go run ./main.go --datapath ./data --datprofile ./tasks/expand.json --outputdir ./
```

## Лицензии

- [MIT](./LICENSE-inherited): кодировщик данных и экспортёр данных, портированные из `v2fly`.
- [MIT](./LICENSE): дополнительные утилиты, предоставляемые этим репозиторием.
  - `abpextract`: преобразует текстовые списки, использующие базовые правила Adblock Plus, в наборы данных, соответствующие стандартным правилам и семантике GeoSite.
  - `dnsmasqextract`: преобразует текстовые списки, использующие правила DNSMasq, в списки доменов, соответствующие стандартным правилам GeoSite.
  - `sanitizer`: очищает и преобразует наборы данных доменов, содержащие записи, не соответствующие стандартным правилам GeoSite, приводя их к требуемому формату.

## Благодарности

- Проект использует портированные кодировщик и экспортёр данных из [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618).
- [Domain List Community](https://github.com/v2fly/domain-list-community)
- [Antifilter](https://antifilter.download/)
- [Antifilter.Download Community Edition](https://community.antifilter.download/)
- [Loyalsoldier/cn-blocked-domain](https://github.com/Loyalsoldier/cn-blocked-domain)
- [felixonmars/dnsmasq-china-list](https://github.com/felixonmars/dnsmasq-china-list)
- [GFWList](https://github.com/gfwlist/gfwlist)
- [Iran Hosted Domains](https://github.com/bootmortis/iran-hosted-domains)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [`Re:filter`](https://github.com/1andrevich/Re-filter-lists)
- [runetfreedom/russia-domains-list](https://github.com/runetfreedom/russia-domains-list)
