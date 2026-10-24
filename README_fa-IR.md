# GeoSite

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | فارسی

این مخزن با استفاده از داده‌های اصلی دسته‌بندی دامنه‌ها در پروژهٔ [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community)، بسته‌های دادهٔ GeoSite را تولید می‌کند. برای این کار، نسخه‌ای سازگارشده از کدگذار دادهٔ موجود در `v2fly/domain-list-community` به کار گرفته شده است. داده‌های اصلی با اطلاعات تکمیلی از منابع منتخب ترکیب می‌شوند تا مجموعه‌هایی از دامنه‌ها با معناشناسی استاندارد GeoSite ساخته شوند.

## اصول طراحی

دقت و پوشش داده‌های GeoSite بر کیفیت مسیریابی ترافیک بر اساس دامنه تأثیر می‌گذارند. نحوهٔ تعریف دسته‌بندی‌ها و معناشناسی داده‌های GeoSite نیز در نتیجهٔ مسیریابی نقش دارد. به همین دلیل، این مخزن از کدگذار داده‌ای استفاده می‌کند که از پروژهٔ `v2fly/domain-list-community` به این مخزن منتقل شده است. این روش کمک می‌کند تا ضمن حفظ حداکثری معنای اصلی داده‌ها، تغییرات پروژهٔ بالادستی نیز به‌طور کامل دنبال شوند.

در کنار پیاده‌سازی کامل دسته‌بندی‌های اصلی دامنه‌ها، فهرست‌های تکمیلی و کاربردی دیگری نیز به مخزن اضافه شده‌اند. این فهرست‌ها به مجموعه‌های داده‌ای تبدیل می‌شوند که با معناشناسی GeoSite سازگار هستند. در نتیجه، این داده‌های اضافی لایه‌ای تکمیلی در کنار داده‌های بالادستی ایجاد می‌کنند و پوشش آن‌ها را گسترش می‌دهند.

برای ایجاد تعادل میان کامل‌بودن مجموعهٔ داده‌ها و حجم فایل‌های کدگذاری‌شده، دو نسخه ارائه می‌شود: نسخهٔ استاندارد (`geosite.dat`) و نسخهٔ گسترش‌یافته (`geosite-expand.dat`). همچنین، برای پیگیری تغییرات و اطمینان از هماهنگی رفتار ابزارهای کدگذاری و ساخت با پروژهٔ بالادستی، یک نسخهٔ راستی‌آزمایی (`verify.dat`) نیز در نظر گرفته شده است. این نسخه فقط داده‌های اصلی پروژهٔ بالادستی را شامل می‌شود و برای بررسی صحت فرایند کدگذاری و ساخت، بدون افزودن دسته‌بندی‌های تکمیلی، کاربرد دارد.

## محتوای داده‌ها

نسخهٔ راستی‌آزمایی (`verify.dat`) فقط شامل داده‌های پروژهٔ `v2fly/domain-list-community` است.

### نسخهٔ استاندارد: `geosite.dat`

| دسته‌بندی | توضیحات | شامل‌شده در | شامل | یادداشت |
|---|---|---|---|---|
| شاخص‌های دادهٔ منطبق با `domain-list-community` | مجموعهٔ کامل داده‌های `v2fly/domain-list-community` | | | هیچ تغییری در شاخص‌های دادهٔ مربوطه اعمال نشده است. |
| `gfwlist-cn` | داده‌های برگرفته از GFWList که با ابزار `abpextract` به قالبی سازگار با معناشناسی GeoSite تبدیل شده‌اند | `blocked-in-cn` | | |
| `apple-cn` | دامنه‌های سرورهای Apple در سرزمین اصلی چین، برگرفته از `felixonmars/dnsmasq-china-list` و تبدیل‌شده به فهرست دامنه با ابزار `dnsmasqextract` | `additional-set-cn` | | |
| `google-cn` | دامنه‌های سرورهای Google در سرزمین اصلی چین، برگرفته از `felixonmars/dnsmasq-china-list` و تبدیل‌شده به فهرست دامنه با ابزار `dnsmasqextract` | `additional-set-cn` | | |
| `mainly-ir` | مجموعه‌ای از دامنه‌های بررسی‌شده از `Chocolate4U/Iran-v2ray-rules` که برای دور زدن محدودیت‌های داخل ایران در نظر گرفته شده‌اند | | | با دسته‌بندی `category-ir` در `Chocolate4U/Iran-v2ray-rules` مطابقت دارد. |
| `available-only-in-ru` | دامنه‌هایی که فقط در داخل روسیه در دسترس هستند، برگرفته از `runetfreedom/russia-domains-list` | | | با دسته‌بندی `ru-available-only-inside` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |
| `refilter` | فهرست کامل دامنه‌های پروژهٔ Re:filter | `blocked-in-ru` | | |
| `antifilter-list-community` | فهرست دامنه‌های نسخهٔ انجمنی Antifilter.Download | `blocked-in-ru` | | با دسته‌بندی `antifilter-download-community` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |
| `additional-set-cn` | مجموعهٔ تکمیلی دامنه‌ها برای سرزمین اصلی چین | | `apple-cn`، `google-cn` | |
| `blocked-in-cn` | فهرست دامنه‌های مسدودشده برای سرزمین اصلی چین | | `gfwlist-cn` | |
| `blocked-in-ru` | فهرست دامنه‌های مسدودشده برای روسیه | | `refilter`، `antifilter-list-community` | با دسته‌بندی `ru-blocked` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |

### نسخهٔ گسترش‌یافته: `geosite-expand.dat`

| دسته‌بندی | توضیحات | شامل‌شده در | شامل | یادداشت |
|---|---|---|---|---|
| شاخص‌های دادهٔ منطبق با `domain-list-community` | مجموعهٔ کامل داده‌های `v2fly/domain-list-community` | | | هیچ تغییری در شاخص‌های دادهٔ مربوطه اعمال نشده است. |
| `gfwlist-cn` | داده‌های برگرفته از GFWList که با ابزار `abpextract` به قالبی سازگار با معناشناسی GeoSite تبدیل شده‌اند | `blocked-in-cn`، `blocked-in-cn-expand` | | |
| `apple-cn` | دامنه‌های سرورهای Apple در سرزمین اصلی چین، برگرفته از `felixonmars/dnsmasq-china-list` و تبدیل‌شده به فهرست دامنه با ابزار `dnsmasqextract` | `additional-set-cn` | | |
| `google-cn` | دامنه‌های سرورهای Google در سرزمین اصلی چین، برگرفته از `felixonmars/dnsmasq-china-list` و تبدیل‌شده به فهرست دامنه با ابزار `dnsmasqextract` | `additional-set-cn` | | |
| `expand-cn` | دامنه‌های مربوط به سرزمین اصلی چین از `felixonmars/dnsmasq-china-list` که با ابزار `dnsmasqextract` به فهرست دامنه تبدیل شده‌اند | `additional-set-cn` | | استفادهٔ هم‌زمان با `cn`، `tld-cn` یا `geosite-cn` توصیه نمی‌شود. |
| `greatfire-list` | داده‌های برگرفته از GreatFire.org که توسط پروژهٔ `Loyalsoldier/cn-blocked-domain` تبدیل شده‌اند | `blocked-in-cn-expand` | | |
| `mainly-ir` | مجموعه‌ای از دامنه‌های بررسی‌شده از `Chocolate4U/Iran-v2ray-rules` که برای دور زدن محدودیت‌های داخل ایران در نظر گرفته شده‌اند | | | با دسته‌بندی `category-ir` در `Chocolate4U/Iran-v2ray-rules` مطابقت دارد. |
| `expand-ir` | مجموعه‌ای از دامنه‌ها برای دور زدن محدودیت‌های داخل ایران، برگرفته از Iran Hosted Domains | | | با دسته‌بندی `ir` در `Chocolate4U/Iran-v2ray-rules` مطابقت دارد. |
| `available-only-in-ru` | دامنه‌هایی که فقط در داخل روسیه در دسترس هستند، برگرفته از `runetfreedom/russia-domains-list` | | | با دسته‌بندی `ru-available-only-inside` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |
| `refilter` | فهرست کامل دامنه‌های پروژهٔ Re:filter | `blocked-in-ru`، `blocked-in-ru-expand` | | |
| `antifilter-list-community` | فهرست دامنه‌های نسخهٔ انجمنی Antifilter.Download | `blocked-in-ru`، `blocked-in-ru-expand` | | با دسته‌بندی `antifilter-download-community` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |
| `antifilter-list-full` | فهرست دامنه‌های Antifilter.Download | `blocked-in-ru-expand` | | با دسته‌بندی `antifilter-download` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. تعداد رکوردها بسیار زیاد است؛ با احتیاط استفاده کنید. |
| `additional-set-cn` | مجموعهٔ تکمیلی دامنه‌ها برای سرزمین اصلی چین | | `apple-cn`، `google-cn`، `expand-cn` | |
| `blocked-in-cn` | فهرست دامنه‌های مسدودشده برای سرزمین اصلی چین | `blocked-in-cn-expand` | `gfwlist-cn` | |
| `blocked-in-cn-expand` | فهرست گسترش‌یافتهٔ دامنه‌های مسدودشده برای سرزمین اصلی چین | | `blocked-in-cn`، `greatfire-list` | |
| `blocked-in-ru` | فهرست دامنه‌های مسدودشده برای روسیه | `blocked-in-ru-expand` | `refilter`، `antifilter-list-community` | با دسته‌بندی `ru-blocked` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. |
| `blocked-in-ru-expand` | فهرست گسترش‌یافتهٔ دامنه‌های مسدودشده برای روسیه | | `blocked-in-ru`، `antifilter-list-full` | با دسته‌بندی `ru-blocked-all` در `russia-blocked-geosite` و `russia-v2ray-rules-dat` مطابقت دارد. تعداد رکوردها بسیار زیاد است؛ با احتیاط استفاده کنید. |

## ساخت

ابتدا در ریشهٔ مخزن، پوشهٔ `./data` را ایجاد کنید. این پوشه باید شامل فایل‌های اصلی دسته‌بندی دامنه‌ها باشد که مراحل تبدیل و پاک‌سازی داده‌ها را پشت سر گذاشته‌اند.

پوشه‌های `./predata/std` و `./predata/exp` به‌ترتیب شامل شاخص‌های تکمیلی نسخهٔ استاندارد و نسخهٔ گسترش‌یافته هستند. برای فعال‌کردن این شاخص‌ها، فایل‌های مربوطه را به پوشهٔ `./data` کپی کنید. توجه داشته باشید که برای اعمال هر شاخص تکمیلی، دسته‌بندی دامنهٔ متناظر باید از قبل در پوشهٔ داده‌ها وجود داشته باشد.

سپس دستور زیر را اجرا کنید تا فایل کدگذاری‌شدهٔ `geosite.dat` در پوشهٔ جاری ساخته شود:

```sh
go run ./main.go --datapath ./data --outputname geosite.dat --outputdir ./
```

## مجوزها

- [MIT](./LICENSE-inherited): کدگذار داده و ابزار صادرکنندهٔ داده که از پروژهٔ `v2fly` منتقل شده‌اند.
- [MIT](./LICENSE): ابزارهای تکمیلی ارائه‌شده در این مخزن.
  - `abpextract`: فهرست‌های متنی مبتنی بر قواعد سادهٔ Adblock Plus را به مجموعه‌های داده‌ای تبدیل می‌کند که با قواعد و معناشناسی استاندارد GeoSite سازگار هستند.
  - `dnsmasqextract`: فهرست‌های متنی مبتنی بر قواعد DNSMasq را به فهرست دامنه‌هایی تبدیل می‌کند که با قواعد استاندارد GeoSite سازگار هستند.
  - `sanitizer`: مجموعه‌های دادهٔ دامنه را که شامل ورودی‌های ناسازگار با قواعد استاندارد GeoSite هستند، پاک‌سازی و تبدیل می‌کند تا با این قواعد سازگار شوند.

## سپاسگزاری

- این پروژه، کدگذار و ابزار صادرکنندهٔ داده را از نسخهٔ [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618) دریافت کرده و به این مخزن منتقل کرده است.
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
