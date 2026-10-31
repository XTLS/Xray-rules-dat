# Xray-rules-dat

[English](./README.md) | [简体中文](./README_zh-Hans.md) | Русский | [فارسی](./README_fa-IR.md)

---

Данные GeoIP и GeoSite для [Xray](https://github.com/XTLS/Xray-core).

GeoIP основан на [IPinfo Lite](https://ipinfo.io/lite) и дополнен IP-данными с учётом региональных задач маршрутизации и часто используемых сервисов. Поддерживаются IPv4 и IPv6.

GeoSite основан на [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) и дополнен чёрными и белыми списками для обеспечения необходимого доступа в различных регионах.

Данные собираются ежедневно по расписанию в 18:00 UTC (02:00 следующего дня в Пекине, 21:00 в Москве и 21:30 в Тегеране).\
_Из-за очередей и особенностей планирования GitHub Actions сборка может завершиться на несколько минут или часов позже запланированного времени._

## Быстрое скачивание

| Данные           | Файл                 | Содержимое                                                                                                                          | GitHub Releases                                                                               | Зеркала jsDelivr                                                                                                                                                                     |
| ---------------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| GeoIP Lite       | `geoip-lite.dat`     | Чёрные и белые списки для материкового Китая, России и Ирана, а также IP-данные частных сетей, популярных сервисов и др.            | [Скачать](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [Основное](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoIP Full       | `geoip-full.dat`     | Все данные для маршрутизации из версии Lite, а также данные геолокации IP по всему миру и Anycast-адреса                            | [Скачать](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | Слишком большой файл                                                                                                                                                                 |
| GeoSite Standard | `geosite.dat`        | Категории доменов из исходного проекта, дополненные чёрными и белыми списками доменов для материкового Китая, России и Ирана        | [Скачать](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [Основное](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite Expanded | `geosite-expand.dat` | Все категории версии Standard, а также чёрные и белые списки доменов с более широким охватом для материкового Китая, России и Ирана | [Скачать](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | Слишком большой файл                                                                                                                                                                 |

\* Версия Lite подходит для базового разделения трафика на прямые и проксируемые соединения в материковом Китае, России и Иране.\
\* Версия Full позволяет точнее маршрутизировать трафик по стране назначения, а также подходит для разделения трафика на прямые и проксируемые соединения в других странах и регионах.\
\* В поставку Xray по умолчанию входят GeoIP Lite и GeoSite Standard под именами `geoip.dat` и `geosite.dat` соответственно.

### Различия между GeoIP Lite и Full

| Тип                                       | Имена категорий                                                                                                | Lite | Full |
| ----------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ---- | ---- |
| Геолокация IP                             | Двухбуквенные коды стран (см. [таблицу ниже](#все-страны-и-территории))                                        | Нет  | Есть |
| Anycast-адреса                            | `anycast`                                                                                                      | Нет  | Есть |
| Региональные наборы (белые списки)        | `in-cn`, `in-ru`, `in-ir`                                                                                      | Есть | Есть |
| Чёрные списки                             | `blacklist-ru`, `blacklist-ir`                                                                                 | Есть | Есть |
| Частные сети и часто используемые сервисы | `private`, `cloudflare`, `cloudfront`, `facebook`, `fastly`, `google`, `netflix`, `telegram`, `twitter`, `tor` | Есть | Есть |
| Псевдонимы для совместимости              | `cn` → `in-cn`, `ru` → `in-ru`, `ir` → `in-ir`                                                                 | Есть | Нет  |

\* Псевдонимы для совместимости сохраняют привычный для сообщества способ использования и доступны только в версии Lite.\
\* `in-cn` содержит IP-адреса, подходящие для прямого подключения из материкового Китая; это не означает, что сами адреса находятся в материковом Китае. В версии Full категория `cn` определяется строго по геолокации IP; то же относится к `ru` и `ir`.

Принципы построения категорий и источники данных описаны в [документации GeoIP](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_ru-RU.md).

#### Файлы отдельных категорий

Если нужны лишь отдельные категории, можно скачать соответствующие файлы вместо большого `geoip-full.dat`.

Например, можно продолжить использовать версию Lite под именем `geoip.dat` и обращаться к её категории `in-cn` через `geoip:in-cn`. Дополнительно скачайте файл категории IP-адресов США [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) в тот же каталог и обращайтесь к его категории `us` через `ext:us.dat:us`.

DAT-файлы отдельных категорий находятся в каталоге [`/geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) ветки `release`.

#### Все страны и территории

Версия Full включает следующие 250 категорий стран и территорий.

<details>
<summary>Развернуть полный список (по коду)</summary>

| Код  | Страна или территория                    | Код  | Страна или территория            | Код  | Страна или территория                    |
| ---- | ---------------------------------------- | ---- | -------------------------------- | ---- | ---------------------------------------- |
| `ad` | Андорра                                  | `ae` | Объединённые Арабские Эмираты    | `af` | Афганистан                               |
| `ag` | Антигуа и Барбуда                        | `ai` | Ангилья                          | `al` | Албания                                  |
| `am` | Армения                                  | `ao` | Ангола                           | `aq` | Антарктида                               |
| `ar` | Аргентина                                | `as` | Американское Самоа               | `at` | Австрия                                  |
| `au` | Австралия                                | `aw` | Аруба                            | `ax` | Аландские острова                        |
| `az` | Азербайджан                              | `ba` | Босния и Герцеговина             | `bb` | Барбадос                                 |
| `bd` | Бангладеш                                | `be` | Бельгия                          | `bf` | Буркина-Фасо                             |
| `bg` | Болгария                                 | `bh` | Бахрейн                          | `bi` | Бурунди                                  |
| `bj` | Бенин                                    | `bl` | Сен-Бартелеми                    | `bm` | Бермудские острова                       |
| `bn` | Бруней                                   | `bo` | Боливия                          | `bq` | Карибские Нидерланды                     |
| `br` | Бразилия                                 | `bs` | Багамские Острова                | `bt` | Бутан                                    |
| `bv` | Остров Буве                              | `bw` | Ботсвана                         | `by` | Беларусь                                 |
| `bz` | Белиз                                    | `ca` | Канада                           | `cc` | Кокосовые (Килинг) острова               |
| `cd` | Демократическая Республика Конго         | `cf` | Центральноафриканская Республика | `cg` | Республика Конго                         |
| `ch` | Швейцария                                | `ci` | Кот-д’Ивуар                      | `ck` | Острова Кука                             |
| `cl` | Чили                                     | `cm` | Камерун                          | `cn` | Китай                                    |
| `co` | Колумбия                                 | `cr` | Коста-Рика                       | `cu` | Куба                                     |
| `cv` | Кабо-Верде                               | `cw` | Кюрасао                          | `cx` | Остров Рождества                         |
| `cy` | Кипр                                     | `cz` | Чехия                            | `de` | Германия                                 |
| `dj` | Джибути                                  | `dk` | Дания                            | `dm` | Доминика                                 |
| `do` | Доминиканская Республика                 | `dz` | Алжир                            | `ec` | Эквадор                                  |
| `ee` | Эстония                                  | `eg` | Египет                           | `eh` | Западная Сахара                          |
| `er` | Эритрея                                  | `es` | Испания                          | `et` | Эфиопия                                  |
| `fi` | Финляндия                                | `fj` | Фиджи                            | `fk` | Фолклендские острова                     |
| `fm` | Микронезия                               | `fo` | Фарерские острова                | `fr` | Франция                                  |
| `ga` | Габон                                    | `gb` | Великобритания                   | `gd` | Гренада                                  |
| `ge` | Грузия                                   | `gf` | Французская Гвиана               | `gg` | Гернси                                   |
| `gh` | Гана                                     | `gi` | Гибралтар                        | `gl` | Гренландия                               |
| `gm` | Гамбия                                   | `gn` | Гвинея                           | `gp` | Гваделупа                                |
| `gq` | Экваториальная Гвинея                    | `gr` | Греция                           | `gs` | Южная Георгия и Южные Сандвичевы острова |
| `gt` | Гватемала                                | `gu` | Гуам                             | `gw` | Гвинея-Бисау                             |
| `gy` | Гайана                                   | `hk` | Гонконг (САР Китая)              | `hm` | Остров Херд и острова Макдональд         |
| `hn` | Гондурас                                 | `hr` | Хорватия                         | `ht` | Гаити                                    |
| `hu` | Венгрия                                  | `id` | Индонезия                        | `ie` | Ирландия                                 |
| `il` | Израиль                                  | `im` | Остров Мэн                       | `in` | Индия                                    |
| `io` | Британская территория в Индийском океане | `iq` | Ирак                             | `ir` | Иран                                     |
| `is` | Исландия                                 | `it` | Италия                           | `je` | Джерси                                   |
| `jm` | Ямайка                                   | `jo` | Иордания                         | `jp` | Япония                                   |
| `ke` | Кения                                    | `kg` | Кыргызстан                       | `kh` | Камбоджа                                 |
| `ki` | Кирибати                                 | `km` | Коморы                           | `kn` | Сент-Китс и Невис                        |
| `kp` | КНДР                                     | `kr` | Республика Корея                 | `kw` | Кувейт                                   |
| `ky` | Каймановы острова                        | `kz` | Казахстан                        | `la` | Лаос                                     |
| `lb` | Ливан                                    | `lc` | Сент-Люсия                       | `li` | Лихтенштейн                              |
| `lk` | Шри-Ланка                                | `lr` | Либерия                          | `ls` | Лесото                                   |
| `lt` | Литва                                    | `lu` | Люксембург                       | `lv` | Латвия                                   |
| `ly` | Ливия                                    | `ma` | Марокко                          | `mc` | Монако                                   |
| `md` | Молдова                                  | `me` | Черногория                       | `mf` | Сен-Мартен (французская часть)           |
| `mg` | Мадагаскар                               | `mh` | Маршалловы Острова               | `mk` | Северная Македония                       |
| `ml` | Мали                                     | `mm` | Мьянма                           | `mn` | Монголия                                 |
| `mo` | Макао (САР Китая)                        | `mp` | Северные Марианские острова      | `mq` | Мартиника                                |
| `mr` | Мавритания                               | `ms` | Монтсеррат                       | `mt` | Мальта                                   |
| `mu` | Маврикий                                 | `mv` | Мальдивы                         | `mw` | Малави                                   |
| `mx` | Мексика                                  | `my` | Малайзия                         | `mz` | Мозамбик                                 |
| `na` | Намибия                                  | `nc` | Новая Каледония                  | `ne` | Нигер                                    |
| `nf` | Остров Норфолк                           | `ng` | Нигерия                          | `ni` | Никарагуа                                |
| `nl` | Нидерланды                               | `no` | Норвегия                         | `np` | Непал                                    |
| `nr` | Науру                                    | `nu` | Ниуэ                             | `nz` | Новая Зеландия                           |
| `om` | Оман                                     | `pa` | Панама                           | `pe` | Перу                                     |
| `pf` | Французская Полинезия                    | `pg` | Папуа — Новая Гвинея             | `ph` | Филиппины                                |
| `pk` | Пакистан                                 | `pl` | Польша                           | `pm` | Сен-Пьер и Микелон                       |
| `pn` | Острова Питкэрн                          | `pr` | Пуэрто-Рико                      | `ps` | Палестинские территории                  |
| `pt` | Португалия                               | `pw` | Палау                            | `py` | Парагвай                                 |
| `qa` | Катар                                    | `re` | Реюньон                          | `ro` | Румыния                                  |
| `rs` | Сербия                                   | `ru` | Россия                           | `rw` | Руанда                                   |
| `sa` | Саудовская Аравия                        | `sb` | Соломоновы Острова               | `sc` | Сейшельские Острова                      |
| `sd` | Судан                                    | `se` | Швеция                           | `sg` | Сингапур                                 |
| `sh` | Остров Святой Елены                      | `si` | Словения                         | `sj` | Шпицберген и Ян-Майен                    |
| `sk` | Словакия                                 | `sl` | Сьерра-Леоне                     | `sm` | Сан-Марино                               |
| `sn` | Сенегал                                  | `so` | Сомали                           | `sr` | Суринам                                  |
| `ss` | Южный Судан                              | `st` | Сан-Томе и Принсипи              | `sv` | Сальвадор                                |
| `sx` | Синт-Мартен (нидерландская часть)        | `sy` | Сирия                            | `sz` | Эсватини                                 |
| `tc` | Острова Теркс и Кайкос                   | `td` | Чад                              | `tf` | Французские Южные территории             |
| `tg` | Того                                     | `th` | Таиланд                          | `tj` | Таджикистан                              |
| `tk` | Токелау                                  | `tl` | Восточный Тимор                  | `tm` | Туркменистан                             |
| `tn` | Тунис                                    | `to` | Тонга                            | `tr` | Турция                                   |
| `tt` | Тринидад и Тобаго                        | `tv` | Тувалу                           | `tw` | Тайвань                                  |
| `tz` | Танзания                                 | `ua` | Украина                          | `ug` | Уганда                                   |
| `um` | Внешние малые острова США                | `us` | США                              | `uy` | Уругвай                                  |
| `uz` | Узбекистан                               | `va` | Ватикан                          | `vc` | Сент-Винсент и Гренадины                 |
| `ve` | Венесуэла                                | `vg` | Британские Виргинские острова    | `vi` | Американские Виргинские острова          |
| `vn` | Вьетнам                                  | `vu` | Вануату                          | `wf` | Уоллис и Футуна                          |
| `ws` | Самоа                                    | `xk` | Косово                           | `ye` | Йемен                                    |
| `yt` | Майотта                                  | `za` | Южная Африка                     | `zm` | Замбия                                   |
| `zw` | Зимбабве                                 |      |                                  |      |                                          |

</details>

### Различия между GeoSite Standard и Expanded

| Тип                                                                     | Имена категорий                 | Standard | Expanded |
| ----------------------------------------------------------------------- | ------------------------------- | -------- | -------- |
| [domain-list-community](https://github.com/v2fly/domain-list-community) | Соответствуют исходному проекту | Есть     | Есть     |
| Чёрный список для материкового Китая                                    | `blocked-in-cn`                 | Есть     | Есть     |
| Белый список для материкового Китая                                     | `located-in-cn`                 | Есть     | Есть     |
| Расширенный чёрный список для материкового Китая                        | `blocked-in-cn-expand`          | Нет      | Есть     |
| Расширенный белый список для материкового Китая                         | `located-in-cn-expand`          | Нет      | Есть     |
| Чёрный список для России                                                | `blocked-in-ru`                 | Есть     | Есть     |
| Белый список для России                                                 | `located-in-ru`                 | Есть     | Есть     |
| Расширенный чёрный список для России                                    | `blocked-in-ru-expand`          | Нет      | Есть     |
| Белый список для Ирана                                                  | `located-in-ir`                 | Есть     | Есть     |
| Расширенный белый список для Ирана                                      | `located-in-ir-expand`          | Нет      | Есть     |

Сочетания категорий, руководство по миграции и источники данных описаны в [документации GeoSite](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_ru-RU.md).

#### Дельта-файлы версии Expanded

Если расширенные данные нужны лишь для отдельных регионов, можно дополнить версию Standard (`geosite.dat`) соответствующими дельта-файлами, не скачивая `geosite-expand.dat`.

Например, поместите [delta-ru.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geosite/delta/delta-ru.dat) в один каталог с версией Standard (`geosite.dat`). Категории исходного проекта domain-list-community по-прежнему доступны через `geosite:xx`, а расширенный чёрный список для России — через `ext:delta-ru.dat:blocked-in-ru-expand`.

Дельта-файлы DAT находятся в каталоге [`/geosite/delta/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geosite/delta) ветки `release`.

## Инструкции по скачиванию

### GitHub Releases

В GitHub Releases доступны только полные DAT-файлы. Последняя и предыдущие версии находятся на [странице выпусков](https://github.com/XTLS/Xray-rules-dat/releases).

Для автоматического обновления используйте следующий постоянный адрес, заменив `<filename>` нужным именем файла для получения последней версии:

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### Ветка `release`

Ветка `release` также содержит последние полные файлы и дополнительно файлы отдельных категорий и дельта-файлы, но не хранит предыдущие версии.

[Откройте ветку](https://github.com/XTLS/Xray-rules-dat/tree/release), найдите нужный файл и подставьте его путь (например, `geoip/dat/us.dat`) вместо `<path>`. Следующий адрес можно использовать для автоматического обновления:

```text
https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>
```

<details>
<summary>Если GitHub недоступен, используйте зеркало jsDelivr</summary>

- Основное: `https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- Fastly: `https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

\* У jsDelivr есть [ограничение размера отдельного файла](https://github.com/jsdelivr/jsdelivr#restrictions), поэтому даже файлы отдельных категорий или дельта-файлы могут оказаться слишком большими для скачивания. Рекомендуется скачивать напрямую с GitHub.

</details>

Для каждого файла данных в том же расположении доступен соответствующий файл контрольной суммы `.sha256sum`.

Для каждого файла `.dat` в том же расположении также доступны сжатая версия `.gz` и её файл контрольной суммы `.gz.sha256sum`.

## Использование в Xray

В поставку Xray по умолчанию входят GeoIP Lite и GeoSite Standard, уже названные `geoip.dat` и `geosite.dat`. Их можно сразу использовать с синтаксисом `geoip:xx` и `geosite:xx`.

В следующем примере дополнительно используется файл IP-адресов США [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat). Поместите все три файла в один каталог.

Добавьте следующий фрагмент в существующую конфигурацию Xray. Пример использует уже настроенные исходящие подключения `direct`, `xhttp-us` и `xhttp-jp` и ежедневно в 04:00 по местному времени среды выполнения автоматически обновляет все три файла и перезагружает их без перезапуска:

```json
{
  "geodata": {
    "cron": "0 4 * * *",
    "assets": [
      {
        "url": "https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat",
        "file": "geoip.dat"
      },
      {
        "url": "https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat",
        "file": "us.dat"
      },
      {
        "url": "https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat",
        "file": "geosite.dat"
      }
    ]
  },
  "routing": {
    "domainStrategy": "IPOnDemand", // Разрешать доменные имена в IP-адреса по мере необходимости, чтобы доменный трафик также мог соответствовать IP-правилам
    "rules": [
      {
        // Сначала направлять трафик к частным адресам напрямую
        "ip": ["geoip:private"],
        "outboundTag": "direct"
      },
      {
        // Пересечение белых списков доменов и IP: прямое соединение только при совпадении с обоими, чтобы снизить влияние ошибок отдельного списка
        "domain": ["geosite:cn"],
        "ip": ["geoip:in-cn"],
        "outboundTag": "direct"
      },
      {
        // Сопоставить IP-адреса США с категорией us из us.dat и использовать исходящее подключение xhttp-us
        "ip": ["ext:us.dat:us"],
        "outboundTag": "xhttp-us"
      },
      {
        // Использовать исходящее подключение xhttp-jp для доменов из чёрного списка материкового Китая
        "domain": ["geosite:blocked-in-cn"],
        "outboundTag": "xhttp-jp"
      }
    ]
  }
}
```

Обновляемые файлы должны уже существовать. Параметры обновления описаны в [документации по настройке Geodata](https://xtls.github.io/ru/config/geodata.html).

Это лишь пример фрагмента конфигурации; адаптируйте его к существующим настройкам и задачам маршрутизации. Порядок правил и особенности сопоставления IP-адресов и доменов описаны в [документации по маршрутизации](https://xtls.github.io/ru/config/routing.html).

Категория, указанная после последнего двоеточия, должна существовать в соответствующем DAT-файле. См. [Загрузка IP-адресов / доменов из файлов](https://xtls.github.io/ru/config/routing.html#ruleobject).

Если нужны IP-данные других стран, скачайте `geoip-full.dat`, переименуйте его в `geoip.dat`, замените исходный файл и измените `geoip-lite.dat` на `geoip-full.dat` в адресе обновления. После этого можно напрямую использовать ссылки вида `geoip:us`, без отдельного скачивания `us.dat` и без `ext:us.dat:us`.

## Исходный код и лицензии

Исходный код GeoIP, источники данных и инструкции по сборке находятся в [ветке `geoip`](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_ru-RU.md).

Исходный код GeoSite, источники данных и инструкции по сборке находятся в [ветке `geosite`](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_ru-RU.md).

Проект в целом распространяется по лицензии [GPL-3.0](./LICENSE). На отдельные инструменты (GeoSite) распространяется [лицензия MIT](./LICENSE-GeoSite) или [унаследованные условия лицензирования](./LICENSE-GeoSite-inherit).

## Благодарности

- Исходный код GeoIP — форк [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb), который, в свою очередь, является форком [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- Исходный код GeoSite был форкнут и портирован из [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618)
- [IPInfo Lite](https://ipinfo.io/lite)
- [bgptools/anycast-prefixes](https://github.com/bgptools/anycast-prefixes)
- [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community)
- [gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip)
- [misakaio/chnroutes2](https://github.com/misakaio/chnroutes2)
- [zhufengme/block_cn_files](https://github.com/zhufengme/block_cn_files)
- [Loyalsoldier/cn-blocked-domain](https://github.com/Loyalsoldier/cn-blocked-domain)
- [felixonmars/dnsmasq-china-list](https://github.com/felixonmars/dnsmasq-china-list)
- [gfwlist/gfwlist](https://github.com/gfwlist/gfwlist)
- [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip)
- [hxehex/russia-mobile-internet-whitelist](https://github.com/hxehex/russia-mobile-internet-whitelist)
- [Antifilter](https://antifilter.download/)
- [Antifilter.Download community edition](https://community.antifilter.download/)
- [1andrevich/Re-filter-lists](https://github.com/1andrevich/Re-filter-lists)
- [runetfreedom/russia-domains-list](https://github.com/runetfreedom/russia-domains-list)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [bootmortis/iran-hosted-domains](https://github.com/bootmortis/iran-hosted-domains)
