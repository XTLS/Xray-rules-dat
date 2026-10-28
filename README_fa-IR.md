# Xray-rules-dat

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | فارسی

---

داده‌های GeoIP و GeoSite برای [Xray](https://github.com/XTLS/Xray-core).

GeoIP بر پایهٔ [IPinfo Lite](https://ipinfo.io/lite) تهیه شده و با توجه به نیازهای مسیریابی در مناطق مختلف و داده‌های IP سرویس‌های پرکاربرد تکمیل می‌شود. این مجموعه از IPv4 و IPv6 پشتیبانی می‌کند.

GeoSite بر پایهٔ [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) تهیه شده و با مجموعه‌های فهرست سیاه و سفید برای پاسخ‌گویی به نیازهای ضروری دسترسی در مناطق مختلف تکمیل می‌شود.

ساخت داده‌ها هر روز برای ساعت 18:00 به وقت UTC برنامه‌ریزی شده است (02:00 روز بعد در پکن، 21:00 در مسکو و 21:30 در تهران).\
_به دلیل صف گردش‌کارها و سازوکار زمان‌بندی GitHub Actions، پایان ساخت ممکن است از چند دقیقه تا چند ساعت پس از زمان برنامه‌ریزی‌شده به تأخیر بیفتد._

## دانلود سریع

| داده                | فایل                 | محتوا                                                                                                                             | GitHub Releases                                                                              | آینه‌های jsDelivr                                                                                                                                                                   |
| ------------------- | -------------------- | --------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GeoIP سبک           | `geoip-lite.dat`     | فهرست‌های سیاه و سفید برای سرزمین اصلی چین، روسیه و ایران، به‌همراه داده‌های IP شبکه‌های خصوصی، سرویس‌های پرکاربرد و موارد دیگر   | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [پیش‌فرض](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoIP کامل          | `geoip-full.dat`     | تمام داده‌های مسیریابی نسخهٔ سبک، به‌همراه داده‌های موقعیت جغرافیایی IP در سراسر جهان و نشانی‌های IP از نوع Anycast               | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | حجم فایل بیش از حد مجاز است                                                                                                                                                         |
| GeoSite استاندارد   | `geosite.dat`        | دسته‌بندی دامنه‌های پروژهٔ بالادستی، به‌همراه فهرست‌های سیاه و سفید دامنه‌ها برای سرزمین اصلی چین، روسیه و ایران                  | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [پیش‌فرض](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite گسترش‌یافته | `geosite-expand.dat` | تمام دسته‌بندی‌های نسخهٔ استاندارد، به‌همراه فهرست‌های سیاه و سفید دامنه‌ها با پوشش گسترده‌تر برای سرزمین اصلی چین، روسیه و ایران | [دانلود](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | حجم فایل بیش از حد مجاز است                                                                                                                                                         |

\* نسخهٔ سبک برای تفکیک پایهٔ ترافیک مستقیم و پراکسی در سرزمین اصلی چین، روسیه و ایران مناسب است.\
\* نسخهٔ کامل برای مسیریابی دقیق‌تر بر اساس کشور مقصد ترافیک و همچنین نیازهای تفکیک ترافیک مستقیم و پراکسی در سایر کشورها و مناطق مناسب است.\
\* بسته‌های انتشار Xray به‌طور پیش‌فرض شامل داده‌های GeoIP سبک و GeoSite استاندارد هستند که به‌ترتیب با نام‌های `geoip.dat` و `geosite.dat` ارائه می‌شوند.

### تفاوت GeoIP سبک و کامل

| نوع                                    | نام دسته‌ها                                                                                                    | سبک          | کامل          |
| -------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------ | ------------- |
| موقعیت جغرافیایی IP                    | کدهای دوحرفی کشورها (نگاه کنید به [جدول زیر](#همه-کشورها-و-مناطق))                                             | ندارد        | دارد          |
| نشانی‌های Anycast                      | `anycast`                                                                                                      | ندارد        | دارد          |
| مجموعه‌های ویژهٔ هر منطقه (فهرست سفید) | `in-cn`، `in-ru`، `in-ir`                                                                                      | دارد         | دارد          |
| فهرست‌های سیاه                         | `blacklist-ru`، `blacklist-ir`                                                                                 | دارد         | دارد          |
| شبکه‌های خصوصی و سرویس‌های پرکاربرد    | `private`، `cloudflare`، `cloudfront`، `facebook`، `fastly`، `google`، `netflix`، `telegram`، `twitter`، `tor` | دارد         | دارد          |
| نام‌های مستعار برای سازگاری            | `cn` → `in-cn`، `ru` → `in-ru`، `ir` → `in-ir`                                                                 | ارائه می‌شود | ارائه نمی‌شود |

\* نام‌های مستعار برای سازگاری، شیوهٔ استفادهٔ رایج در جامعهٔ کاربران را حفظ می‌کنند و فقط در نسخهٔ سبک ارائه می‌شوند.\
\* دستهٔ `in-cn` شامل نشانی‌های IP مناسب برای اتصال مستقیم از سرزمین اصلی چین است؛ این به معنی قرار داشتن آن نشانی‌ها در سرزمین اصلی چین نیست. در نسخهٔ کامل، دستهٔ `cn` صرفاً بر اساس موقعیت جغرافیایی IP تعریف می‌شود؛ همین موضوع دربارهٔ `ru` و `ir` نیز صدق می‌کند.

برای آشنایی با منطق طراحی دسته‌ها و منابع داده، [مستندات GeoIP](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_fa-IR.md) را ببینید.

#### فایل‌های جداگانهٔ هر دسته

اگر فقط به چند دسته نیاز دارید، می‌توانید فایل‌های مربوط به همان دسته‌ها را جداگانه دانلود کنید و نیازی به دانلود فایل حجیم `geoip-full.dat` ندارید.

برای مثال، می‌توانید همچنان از نسخهٔ سبک با نام `geoip.dat` استفاده کنید و به دستهٔ `in-cn` آن از طریق `geoip:in-cn` دسترسی داشته باشید. سپس فایل دستهٔ IP آمریکا، یعنی [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat)، را در همان پوشه قرار دهید و به دستهٔ `us` آن از طریق `ext:us.dat:us` دسترسی پیدا کنید.

فایل‌های DAT هر دسته در پوشهٔ [`/geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) از شاخهٔ `release` قرار دارند.

#### همه کشورها و مناطق

نسخهٔ کامل شامل 250 دستهٔ کشور و منطقهٔ زیر است.

<details>
<summary>نمایش فهرست کامل (مرتب‌شده بر اساس کد)</summary>

| کد   | کشور یا منطقه                    | کد   | کشور یا منطقه                   | کد   | کشور یا منطقه                       |
| ---- | -------------------------------- | ---- | ------------------------------- | ---- | ----------------------------------- |
| `ad` | آندورا                           | `ae` | امارات متحدهٔ عربی              | `af` | افغانستان                           |
| `ag` | آنتیگوا و باربودا                | `ai` | آنگویلا                         | `al` | آلبانی                              |
| `am` | ارمنستان                         | `ao` | آنگولا                          | `aq` | جنوبگان                             |
| `ar` | آرژانتین                         | `as` | ساموآی آمریکا                   | `at` | اتریش                               |
| `au` | استرالیا                         | `aw` | آروبا                           | `ax` | جزایر آلاند                         |
| `az` | جمهوری آذربایجان                 | `ba` | بوسنی و هرزگوین                 | `bb` | باربادوس                            |
| `bd` | بنگلادش                          | `be` | بلژیک                           | `bf` | بورکینافاسو                         |
| `bg` | بلغارستان                        | `bh` | بحرین                           | `bi` | بوروندی                             |
| `bj` | بنین                             | `bl` | سن بارتلمی                      | `bm` | برمودا                              |
| `bn` | برونئی                           | `bo` | بولیوی                          | `bq` | جزایر کارائیب هلند                  |
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
| `gq` | گینهٔ استوایی                    | `gr` | یونان                           | `gs` | جورجیای جنوبی و جزایر ساندویچ جنوبی |
| `gt` | گواتمالا                         | `gu` | گوام                            | `gw` | گینهٔ بیسائو                        |
| `gy` | گویان                            | `hk` | هنگ‌کنگ، منطقهٔ ویژهٔ اداری چین | `hm` | جزیرهٔ هرد و جزایر مک‌دونالد        |
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
| `mo` | ماکائو، منطقهٔ ویژهٔ اداری چین   | `mp` | جزایر ماریانای شمالی            | `mq` | مارتینیک                            |
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
| `sx` | سینت مارتن (بخش هلندی)           | `sy` | سوریه                           | `sz` | اسواتینی                            |
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

### تفاوت GeoSite استاندارد و گسترش‌یافته

| نوع                                                                     | نام دسته‌ها              | استاندارد | گسترش‌یافته |
| ----------------------------------------------------------------------- | ------------------------ | --------- | ----------- |
| [domain-list-community](https://github.com/v2fly/domain-list-community) | مطابق با پروژهٔ بالادستی | دارد      | دارد        |
| فهرست سیاه سرزمین اصلی چین                                              | `blocked-in-cn`          | دارد      | دارد        |
| فهرست سفید سرزمین اصلی چین                                              | `cn`                     | دارد      | دارد        |
| فهرست سیاه گسترش‌یافتهٔ سرزمین اصلی چین                                 | `blocked-in-cn-expand`   | ندارد     | دارد        |
| فهرست سفید گسترش‌یافتهٔ سرزمین اصلی چین                                 | `expand-cn`              | ندارد     | دارد        |
| فهرست سیاه روسیه                                                        | `blocked-in-ru`          | دارد      | دارد        |
| فهرست سفید روسیه                                                        | `available-only-in-ru`   | دارد      | دارد        |
| فهرست سیاه گسترش‌یافتهٔ روسیه                                           | `blocked-in-ru-expand`   | ندارد     | دارد        |
| فهرست سفید ایران                                                        | `mainly-ir`              | دارد      | دارد        |
| فهرست سفید گسترش‌یافتهٔ ایران                                           | `expand-ir`              | ندارد     | دارد        |

برای ترکیب دسته‌ها، راهنمای مهاجرت و منابع داده، [مستندات GeoSite](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_fa-IR.md) را ببینید.

#### فایل‌های افزایشی نسخهٔ گسترش‌یافته

اگر فقط به داده‌های گسترش‌یافتهٔ برخی مناطق نیاز دارید، فایل‌های افزایشی مربوط را در کنار نسخهٔ استاندارد `geosite.dat` دانلود کنید؛ نیازی به دانلود `geosite-expand.dat` نیست.

برای مثال، [delta-ru.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geosite/delta/delta-ru.dat) را در همان پوشهٔ نسخهٔ استاندارد `geosite.dat` قرار دهید. دسته‌های پروژهٔ بالادستی domain-list-community همچنان با `geosite:xx` و فهرست سیاه گسترش‌یافتهٔ روسیه با `ext:delta-ru.dat:blocked-in-ru-expand` قابل استفاده‌اند.

فایل‌های افزایشی DAT در پوشهٔ [`/geosite/delta/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geosite/delta) از شاخهٔ `release` قرار دارند.

## راهنمای دانلود

### GitHub Releases

در GitHub Releases فقط فایل‌های کامل DAT ارائه می‌شوند. آخرین نسخه و نسخه‌های پیشین در [صفحهٔ انتشارها](https://github.com/XTLS/Xray-rules-dat/releases) در دسترس هستند.

برای به‌روزرسانی خودکار می‌توانید از نشانی ثابت زیر استفاده کنید. کافی است `<filename>` را با نام فایل مورد نظر جایگزین کنید تا آخرین نسخه را دریافت کنید:

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### شاخهٔ `release`

شاخهٔ `release` نیز آخرین فایل‌های کامل را به‌همراه فایل‌های جداگانهٔ هر دسته و فایل‌های افزایشی ارائه می‌کند، اما نسخه‌های پیشین را نگه نمی‌دارد.

با [مرور این شاخه](https://github.com/XTLS/Xray-rules-dat/tree/release)، فایل مورد نظر را پیدا کنید و مسیر آن (مانند `geoip/dat/us.dat`) را به‌جای `<path>` قرار دهید. نشانی زیر برای به‌روزرسانی خودکار قابل استفاده است:

```text
https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>
```

<details>
<summary>اگر به GitHub دسترسی ندارید، از آینهٔ jsDelivr استفاده کنید</summary>

- پیش‌فرض: `https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- Fastly: `https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

\* jsDelivr دارای [محدودیت حجم برای هر فایل](https://github.com/jsdelivr/jsdelivr#restrictions) است؛ بنابراین حتی فایل‌های جداگانهٔ دسته‌ها یا فایل‌های افزایشی نیز ممکن است به دلیل حجم زیاد قابل دانلود نباشند. توصیه می‌شود ابتدا دانلود مستقیم از GitHub را امتحان کنید.

</details>

برای هر فایل داده، یک فایل متناظر با پسوند `.sha256sum` برای بررسی صحت در همان مسیر وجود دارد.

## استفاده در Xray

بسته‌های انتشار Xray به‌طور پیش‌فرض شامل GeoIP سبک و GeoSite استاندارد با نام‌های `geoip.dat` و `geosite.dat` هستند. می‌توانید مستقیماً با ساختار `geoip:xx` و `geosite:xx` از آن‌ها استفاده کنید.

در مثال زیر، فایل IP آمریکا یعنی [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) نیز استفاده می‌شود. هر سه فایل را در یک پوشه قرار دهید.

قطعهٔ زیر را در پیکربندی فعلی Xray خود ادغام کنید. این مثال از اتصال‌های خروجی موجود `direct`، `xhttp-us` و `xhttp-jp` استفاده می‌کند و هر روز ساعت 04:00 به وقت محلی محیط اجرا، هر سه فایل را به‌طور خودکار به‌روزرسانی و بدون راه‌اندازی دوباره بارگذاری می‌کند:

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
    "domainStrategy": "IPOnDemand", // نام دامنه‌ها را در صورت نیاز به IP تبدیل کنید تا ترافیک دامنه نیز بتواند با قوانین IP تطبیق پیدا کند
    "rules": [
      {
        // ابتدا ترافیک نشانی‌های خصوصی را مستقیماً ارسال کنید
        "ip": ["geoip:private"],
        "outboundTag": "direct"
      },
      {
        // اشتراک فهرست‌های سفید دامنه و IP: اتصال مستقیم فقط در صورت تطبیق با هر دو، برای کاهش اثر تشخیص نادرست در یک فهرست
        "domain": ["geosite:cn"],
        "ip": ["geoip:in-cn"],
        "outboundTag": "direct"
      },
      {
        // نشانی‌های IP آمریکا را با دستهٔ us در us.dat تطبیق دهید و از اتصال خروجی xhttp-us استفاده کنید
        "ip": ["ext:us.dat:us"],
        "outboundTag": "xhttp-us"
      },
      {
        // برای دامنه‌های فهرست سیاه سرزمین اصلی چین از اتصال خروجی xhttp-jp استفاده کنید
        "domain": ["geosite:blocked-in-cn"],
        "outboundTag": "xhttp-jp"
      }
    ]
  }
}
```

فایل‌هایی که قرار است به‌روزرسانی شوند باید از قبل موجود باشند. گزینه‌های به‌روزرسانی در [مستندات پیکربندی Geodata](https://xtls.github.io/en/config/geodata.html) توضیح داده شده‌اند.

این فقط نمونه‌ای از یک قطعهٔ پیکربندی است؛ برای استفادهٔ واقعی، آن را با پیکربندی فعلی و نیازهای مسیریابی خود سازگار کنید. ترتیب قوانین و شیوهٔ تطبیق IP و دامنه در [مستندات مسیریابی](https://xtls.github.io/en/config/routing.html) آمده است.

نام دسته پس از آخرین دونقطه باید در فایل DAT مربوطه وجود داشته باشد. برای جزئیات، [بارگذاری IP / دامنه از فایل](https://xtls.github.io/en/config/routing.html#ruleobject) را ببینید.

اگر به داده‌های IP کشورهای بیشتری نیاز دارید، `geoip-full.dat` را دانلود کنید، نام آن را به `geoip.dat` تغییر دهید، فایل قبلی را جایگزین کنید و در نشانی به‌روزرسانی، `geoip-lite.dat` را به `geoip-full.dat` تغییر دهید. پس از آن می‌توانید مستقیماً از ارجاع‌هایی مانند `geoip:us` استفاده کنید و دیگر نیازی به دانلود جداگانهٔ `us.dat` یا استفاده از `ext:us.dat:us` ندارید.

## کد منبع و مجوزها

کد منبع GeoIP، منابع داده و راهنمای ساخت در [شاخهٔ `geoip`](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_fa-IR.md) نگهداری می‌شوند.

کد منبع GeoSite، منابع داده و راهنمای ساخت در [شاخهٔ `geosite`](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_fa-IR.md) نگهداری می‌شوند.

این پروژه در مجموع تحت مجوز [GPL-3.0](./LICENSE) منتشر می‌شود. برخی ابزارها (GeoSite) به‌صورت جداگانه تحت [مجوز MIT](./LICENSE-GeoSite) یا [شرایط مجوز به‌ارث‌رسیده](./LICENSE-GeoSite-inherit) قرار دارند.

## قدردانی

- کد منبع GeoIP از [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb) منشعب شده است؛ این پروژه نیز از [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc) منشعب شده بود
- کد منبع GeoSite از [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618) منشعب و منتقل شده است
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
