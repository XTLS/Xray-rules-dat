# Xray-rules-dat

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | فارسی

---

داده‌های GeoIP و GeoSite برای استفاده در [Xray](https://github.com/XTLS/Xray-core).

GeoIP بر پایهٔ [IPinfo Lite](https://ipinfo.io/lite) تهیه شده و با داده‌های IP متناسب با نیازهای مسیریابی منطقه‌ای و سرویس‌های پرکاربرد تکمیل می‌شود. از IPv4 و IPv6 پشتیبانی می‌شود.

GeoSite در حال حاضر از داده‌های بالادستی [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) بدون تغییر استفاده می‌کند.

ساخت داده‌ها هر روز ساعت ۱۸:۰۰ به وقت UTC انجام می‌شود (۰۲:۰۰ روز بعد در چین، ۲۱:۰۰ در مسکو و ۲۱:۳۰ در ایران).

## دانلود سریع

| داده             | فایل             | محتوا                                                                                                           | GitHub Releases                                                                          | آینه‌های jsDelivr                                                                                                                                                                   |
| ---------------- | ---------------- | --------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| نسخهٔ کامل GeoIP | `geoip-full.dat` | دسته‌بندی‌های اصلی همهٔ کشورها و قلمروها، مجموعه‌های سفارشی منطقه‌ای، فهرست‌های مسدودسازی و دسته‌بندی‌های عمومی | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat) | حجم فایل بیش از حد مجاز است                                                                                                                                                         |
| نسخهٔ سبک GeoIP  | `geoip-lite.dat` | مجموعه‌های سفارشی منطقه‌ای، فهرست‌های مسدودسازی و دسته‌بندی‌های عمومی، به همراه نام‌های مستعار برای سازگاری     | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat) | [پیش‌فرض](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoSite          | `geosite.dat`    | داده‌های دسته‌بندی دامنه‌ها                                                                                     | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)    | [پیش‌فرض](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |

\* بسته‌های انتشار Xray شامل GeoSite و نسخهٔ سبک GeoIP با نام‌های `geosite.dat` و `geoip.dat` هستند.

### تفاوت نسخهٔ سبک و کامل GeoIP

| دسته‌بندی                           | نام یا محدوده                                                                                                  | نسخهٔ کامل    | نسخهٔ سبک    |
| ----------------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------- | ------------ |
| دسته‌بندی‌های اصلی کشورها و قلمروها | داده‌های اصلی IPinfo برای همهٔ کشورها و قلمروها                                                                | دارد          | ندارد        |
| آدرس‌های Anycast                    | `anycast`                                                                                                      | دارد          | ندارد        |
| مجموعه‌های سفارشی منطقه‌ای          | `in-cn`، `in-ru`، `in-ir`                                                                                      | دارد          | دارد         |
| فهرست‌های مسدودسازی                 | `blacklist-ru`، `blacklist-ir`                                                                                 | دارد          | دارد         |
| دسته‌بندی‌های عمومی                 | `private`، `cloudflare`، `cloudfront`، `facebook`، `fastly`، `google`، `netflix`، `telegram`، `twitter`، `tor` | دارد          | دارد         |
| نام‌های مستعار برای سازگاری         | `cn` → `in-cn`، `ru` → `in-ru`، `ir` → `in-ir`                                                                 | ارائه نمی‌شود | ارائه می‌شود |

برای آشنایی با اصول طراحی دسته‌بندی‌ها و منابع داده، [مستندات GeoIP](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_fa-IR.md) را ببینید.

### فایل‌های جداگانهٔ دسته‌بندی‌ها

اگر فقط به چند دسته‌بندی نیاز دارید، می‌توانید فایل‌های مربوط به آن‌ها را جداگانه دانلود کنید و نیازی به دانلود فایل بزرگ‌تر `geoip-full.dat` نیست.

برای نمونه، می‌توانید همچنان از نسخهٔ سبک با نام `geoip.dat` استفاده کنید؛ دسته‌بندی موجود `in-cn` همچنان از طریق `geoip:in-cn` در دسترس است. فایل [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) را نیز در همان پوشه قرار دهید تا دسته‌بندی `us` از طریق `ext:us.dat:us` قابل استفاده باشد.

فایل‌های DAT هر دسته‌بندی در پوشهٔ [`geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) در شاخهٔ `release` قرار دارند.

### همهٔ کشورها و قلمروها

نسخهٔ کامل شامل ۲۵۰ دسته‌بندی کشور و قلمرو زیر است.

<details>
<summary>نمایش فهرست کامل (مرتب‌شده بر اساس کد)</summary>

| کد   | کشور یا قلمرو                    | کد   | کشور یا قلمرو                   | کد   | کشور یا قلمرو                       |
| ---- | -------------------------------- | ---- | ------------------------------- | ---- | ----------------------------------- |
| `ad` | آندورا                           | `ae` | امارات متحدهٔ عربی              | `af` | افغانستان                           |
| `ag` | آنتیگوا و باربودا                | `ai` | آنگویلا                         | `al` | آلبانی                              |
| `am` | ارمنستان                         | `ao` | آنگولا                          | `aq` | جنوبگان                             |
| `ar` | آرژانتین                         | `as` | ساموآی آمریکا                   | `at` | اتریش                               |
| `au` | استرالیا                         | `aw` | آروبا                           | `ax` | جزایر آلند                          |
| `az` | جمهوری آذربایجان                 | `ba` | بوسنی و هرزگوین                 | `bb` | باربادوس                            |
| `bd` | بنگلادش                          | `be` | بلژیک                           | `bf` | بورکینافاسو                         |
| `bg` | بلغارستان                        | `bh` | بحرین                           | `bi` | بوروندی                             |
| `bj` | بنین                             | `bl` | سن بارتلمی                      | `bm` | برمودا                              |
| `bn` | برونئی                           | `bo` | بولیوی                          | `bq` | هلند کارائیب                        |
| `br` | برزیل                            | `bs` | باهاما                          | `bt` | بوتان                               |
| `bv` | جزیرهٔ بووه                      | `bw` | بوتسوانا                        | `by` | بلاروس                              |
| `bz` | بلیز                             | `ca` | کانادا                          | `cc` | جزایر کوکوس (کیلینگ)                |
| `cd` | جمهوری دموکراتیک کنگو            | `cf` | جمهوری آفریقای مرکزی            | `cg` | جمهوری کنگو                         |
| `ch` | سوئیس                            | `ci` | ساحل عاج                        | `ck` | جزایر کوک                           |
| `cl` | شیلی                             | `cm` | کامرون                          | `cn` | چین                                 |
| `co` | کلمبیا                           | `cr` | کاستاریکا                       | `cu` | کوبا                                |
| `cv` | کیپ ورد                          | `cw` | کوراسائو                        | `cx` | جزیرهٔ کریسمس                       |
| `cy` | قبرس                             | `cz` | چک                              | `de` | آلمان                               |
| `dj` | جیبوتی                           | `dk` | دانمارک                         | `dm` | دومینیکا                            |
| `do` | جمهوری دومینیکن                  | `dz` | الجزایر                         | `ec` | اکوادور                             |
| `ee` | استونی                           | `eg` | مصر                             | `eh` | صحرای غربی                          |
| `er` | اریتره                           | `es` | اسپانیا                         | `et` | اتیوپی                              |
| `fi` | فنلاند                           | `fj` | فیجی                            | `fk` | جزایر فالکلند                       |
| `fm` | میکرونزی                         | `fo` | جزایر فارو                      | `fr` | فرانسه                              |
| `ga` | گابن                             | `gb` | بریتانیا                        | `gd` | گرنادا                              |
| `ge` | گرجستان                          | `gf` | گویان فرانسه                    | `gg` | گرنزی                               |
| `gh` | غنا                              | `gi` | جبل‌الطارق                      | `gl` | گرینلند                             |
| `gm` | گامبیا                           | `gn` | گینه                            | `gp` | گوادلوپ                             |
| `gq` | گینهٔ استوایی                    | `gr` | یونان                           | `gs` | جزایر جورجیای جنوبی و ساندویچ جنوبی |
| `gt` | گواتمالا                         | `gu` | گوام                            | `gw` | گینهٔ بیسائو                        |
| `gy` | گویان                            | `hk` | هنگ‌کنگ، منطقهٔ اداری ویژهٔ چین | `hm` | جزیرهٔ هرد و جزایر مک‌دونالد        |
| `hn` | هندوراس                          | `hr` | کرواسی                          | `ht` | هائیتی                              |
| `hu` | مجارستان                         | `id` | اندونزی                         | `ie` | ایرلند                              |
| `il` | اسرائیل                          | `im` | جزیرهٔ من                       | `in` | هند                                 |
| `io` | قلمرو بریتانیا در اقیانوس هند    | `iq` | عراق                            | `ir` | ایران                               |
| `is` | ایسلند                           | `it` | ایتالیا                         | `je` | جرزی                                |
| `jm` | جامائیکا                         | `jo` | اردن                            | `jp` | ژاپن                                |
| `ke` | کنیا                             | `kg` | قرقیزستان                       | `kh` | کامبوج                              |
| `ki` | کیریباتی                         | `km` | کومور                           | `kn` | سنت کیتس و نویس                     |
| `kp` | کرهٔ شمالی                       | `kr` | کرهٔ جنوبی                      | `kw` | کویت                                |
| `ky` | جزایر کیمن                       | `kz` | قزاقستان                        | `la` | لائوس                               |
| `lb` | لبنان                            | `lc` | سنت لوسیا                       | `li` | لیختن‌اشتاین                        |
| `lk` | سری‌لانکا                        | `lr` | لیبریا                          | `ls` | لسوتو                               |
| `lt` | لیتوانی                          | `lu` | لوکزامبورگ                      | `lv` | لتونی                               |
| `ly` | لیبی                             | `ma` | مراکش                           | `mc` | موناکو                              |
| `md` | مولداوی                          | `me` | مونته‌نگرو                      | `mf` | سن مارتن (بخش فرانسوی)              |
| `mg` | ماداگاسکار                       | `mh` | جزایر مارشال                    | `mk` | مقدونیهٔ شمالی                      |
| `ml` | مالی                             | `mm` | میانمار                         | `mn` | مغولستان                            |
| `mo` | ماکائو، منطقهٔ اداری ویژهٔ چین   | `mp` | جزایر ماریانای شمالی            | `mq` | مارتینیک                            |
| `mr` | موریتانی                         | `ms` | مونتسرات                        | `mt` | مالت                                |
| `mu` | موریس                            | `mv` | مالدیو                          | `mw` | مالاوی                              |
| `mx` | مکزیک                            | `my` | مالزی                           | `mz` | موزامبیک                            |
| `na` | نامیبیا                          | `nc` | کالدونیای جدید                  | `ne` | نیجر                                |
| `nf` | جزیرهٔ نورفک                     | `ng` | نیجریه                          | `ni` | نیکاراگوئه                          |
| `nl` | هلند                             | `no` | نروژ                            | `np` | نپال                                |
| `nr` | نائورو                           | `nu` | نیوئه                           | `nz` | نیوزیلند                            |
| `om` | عمان                             | `pa` | پاناما                          | `pe` | پرو                                 |
| `pf` | پلی‌نزی فرانسه                   | `pg` | پاپوآ گینهٔ نو                  | `ph` | فیلیپین                             |
| `pk` | پاکستان                          | `pl` | لهستان                          | `pm` | سن پیر و میکلون                     |
| `pn` | جزایر پیت‌کرن                    | `pr` | پورتوریکو                       | `ps` | سرزمین‌های فلسطینی                  |
| `pt` | پرتغال                           | `pw` | پالائو                          | `py` | پاراگوئه                            |
| `qa` | قطر                              | `re` | رئونیون                         | `ro` | رومانی                              |
| `rs` | صربستان                          | `ru` | روسیه                           | `rw` | رواندا                              |
| `sa` | عربستان سعودی                    | `sb` | جزایر سلیمان                    | `sc` | سیشل                                |
| `sd` | سودان                            | `se` | سوئد                            | `sg` | سنگاپور                             |
| `sh` | سنت هلنا                         | `si` | اسلوونی                         | `sj` | سوالبارد و یان ماین                 |
| `sk` | اسلواکی                          | `sl` | سیرالئون                        | `sm` | سان مارینو                          |
| `sn` | سنگال                            | `so` | سومالی                          | `sr` | سورینام                             |
| `ss` | سودان جنوبی                      | `st` | سائوتومه و پرنسیپ               | `sv` | السالوادور                          |
| `sx` | سینت مارتن (بخش هلندی)           | `sy` | سوریه                           | `sz` | سوازیلند                            |
| `tc` | جزایر تورکس و کایکوس             | `td` | چاد                             | `tf` | سرزمین‌های جنوبی فرانسه             |
| `tg` | توگو                             | `th` | تایلند                          | `tj` | تاجیکستان                           |
| `tk` | توکلائو                          | `tl` | تیمور شرقی                      | `tm` | ترکمنستان                           |
| `tn` | تونس                             | `to` | تونگا                           | `tr` | ترکیه                               |
| `tt` | ترینیداد و توباگو                | `tv` | تووالو                          | `tw` | تایوان                              |
| `tz` | تانزانیا                         | `ua` | اوکراین                         | `ug` | اوگاندا                             |
| `um` | جزایر کوچک حاشیه‌ای ایالات متحده | `us` | ایالات متحدهٔ آمریکا            | `uy` | اروگوئه                             |
| `uz` | ازبکستان                         | `va` | واتیکان                         | `vc` | سنت وینسنت و گرنادین‌ها             |
| `ve` | ونزوئلا                          | `vg` | جزایر ویرجین بریتانیا           | `vi` | جزایر ویرجین آمریکا                 |
| `vn` | ویتنام                           | `vu` | وانواتو                         | `wf` | والیس و فوتونا                      |
| `ws` | ساموآ                            | `xk` | کوزوو                           | `ye` | یمن                                 |
| `yt` | مایوت                            | `za` | آفریقای جنوبی                   | `zm` | زامبیا                              |
| `zw` | زیمبابوه                         |      |                                 |      |                                     |

</details>

## راهنمای دانلود

### GitHub Releases

در GitHub Releases فقط پایگاه‌های دادهٔ DAT به‌صورت تک‌فایل ارائه می‌شوند. برای دانلود دستی، [آخرین نسخه](https://github.com/XTLS/Xray-rules-dat/releases/latest) را باز کنید. نسخه‌های پیشین در [GitHub Releases](https://github.com/XTLS/Xray-rules-dat/releases) در دسترس هستند.

برنامه‌های به‌روزرسانی خودکار می‌توانند آخرین نسخه را از نشانی ثابت زیر دانلود کنند. به‌جای `<filename>` نام فایل موردنظر را قرار دهید:

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### شاخهٔ `release`

آخرین نسخهٔ همهٔ فایل‌ها در [شاخهٔ `release`](https://github.com/XTLS/Xray-rules-dat/tree/release) منتشر می‌شود. اگر به دسته‌بندی مشخصی نیاز دارید، فایل مربوط به آن را مستقیماً از این شاخه دانلود کنید.

به‌جای `<path>` مسیر فایل در شاخه، مانند `geoip/dat/us.dat`، را قرار دهید:

- GitHub Raw: `https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>`
- آینهٔ jsDelivr (پیش‌فرض): `https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- آینهٔ jsDelivr (Fastly): `https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

آینه‌های jsDelivr دارای [محدودیت حجم هر فایل](https://github.com/jsdelivr/jsdelivr#restrictions) هستند؛ بنابراین حتی فایل‌های جداگانهٔ دسته‌بندی‌ها نیز ممکن است به‌دلیل حجم زیاد قابل دانلود نباشند. توصیه می‌شود ابتدا از GitHub Releases یا GitHub Raw استفاده کنید و فقط هنگامی که GitHub در دسترس نیست، از آینه‌ها کمک بگیرید.

برای هر فایل داده، یک فایل متناظر با پسوند `.sha256sum` برای بررسی صحت فایل در همان محل قرار دارد.

## استفاده در Xray

بسته‌های انتشار Xray شامل نسخهٔ سبک GeoIP با نام `geoip.dat` هستند. می‌توانید مستقیماً از مجموعه‌های سفارشی منطقه‌ای و دسته‌بندی‌های عمومی آن استفاده کنید.

در مثال زیر، نسخهٔ سبک با نام `geoip.dat` حفظ می‌شود و فایل [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) نیز در همان پوشه دانلود می‌شود.

قطعهٔ زیر را به پیکربندی Xray اضافه کنید تا هر دو فایل هر روز ساعت ۰۴:۰۰ به وقت محلی محیط اجرا، به‌صورت خودکار به‌روزرسانی و بدون راه‌اندازی مجدد بارگذاری شوند. آدرس‌های خصوصی و دسته‌بندی `in-cn` از خروجی موجود `direct` و دسته‌بندی `us` از خروجی موجود `wireguard` استفاده می‌کنند:

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
      }
    ]
  },
  "routing": {
    "rules": [
      {
        "ip": ["geoip:private", "geoip:in-cn"],
        "outboundTag": "direct"
      },
      {
        "ip": ["ext:us.dat:us"],
        "outboundTag": "wireguard"
      }
    ]
  }
}
```

فایل‌هایی که قرار است به‌روزرسانی شوند باید از قبل وجود داشته باشند. برای گزینه‌های بیشتر، [مستندات پیکربندی Geodata](https://xtls.github.io/en/config/geodata.html) را ببینید.

می‌توانید نام دسته‌بندی‌ها را در مثال بر اساس نیاز تغییر دهید. برای ترتیب قوانین و روش تطبیق IP، [مستندات مسیریابی Xray](https://xtls.github.io/en/config/routing.html) را ببینید.

دسته‌بندی موجود `in-cn` همچنان از طریق `geoip:in-cn` استفاده می‌شود و دسته‌بندی اضافه‌شدهٔ `us` از طریق `ext:us.dat:us` بارگذاری می‌شود. برای دسته‌بندی‌های دیگر، `us` را در نام فایل و نام دسته‌بندی جایگزین کنید.

نام دسته‌بندی پس از آخرین دونقطه باید در همان فایل DAT وجود داشته باشد. برای جزئیات، بخش [بارگذاری IP از فایل](https://xtls.github.io/en/config/routing.html#ruleobject) در مستندات Xray را ببینید.

اگر به دسته‌بندی‌های اصلی برای کشورها و قلمروهای بیشتری نیاز دارید، می‌توانید `geoip-full.dat` را دانلود کرده، نام آن را به `geoip.dat` تغییر دهید و جایگزین فایل قبلی کنید. همچنین در نشانی به‌روزرسانی بالا، `geoip-lite.dat` را به `geoip-full.dat` تغییر دهید. پس از جایگزینی فایل، می‌توانید دسته‌بندی‌های اصلی را مستقیماً، برای نمونه از طریق `geoip:us`، استفاده کنید و نیازی به دانلود جداگانهٔ `us.dat` یا استفاده از `ext:us.dat:us` نیست.

## کد منبع و مجوز

کد منبع GeoIP، منابع داده و مراحل ساخت در [شاخهٔ `geoip`](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_fa-IR.md) نگهداری می‌شوند.

این پروژه تحت مجوز [GPL-3.0](./LICENSE) منتشر می‌شود.

## قدردانی

- کد منبع GeoIP از [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb) منشعب شده است که خود انشعابی از [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc) است
- [IPInfo Lite](https://ipinfo.io/lite)
- [zhufengme/block_cn_files](https://github.com/zhufengme/block_cn_files)
- [misakaio/chnroutes2](https://github.com/misakaio/chnroutes2)
- [gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip)
- [bgptools/anycast-prefixes](https://github.com/bgptools/anycast-prefixes)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip)
- [hxehex/russia-mobile-internet-whitelist](https://github.com/hxehex/russia-mobile-internet-whitelist)
- [Antifilter](https://antifilter.download/)
- [1andrevich/Re-filter-lists](https://github.com/1andrevich/Re-filter-lists)

---

- [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community): منبع بالادستی داده‌های GeoSite
