# Xray-rules-dat

[English](./README.md) | 简体中文 | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

---

供 [Xray](https://github.com/XTLS/Xray-core) 使用的 GeoIP 和 GeoSite 数据。

GeoIP 以 [IPinfo Lite](https://ipinfo.io/lite) 为基础，结合各地分流需求和常用服务的 IP 数据整理，支持 IPv4 和 IPv6。

GeoSite 以 [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) 为基础，补充各地区存在刚性使用需求的黑名单与白名单数据集合。

数据每日定时构建，计划执行时间为 UTC 18:00（北京次日 02:00、莫斯科 21:00、德黑兰 21:30）。\
_受 GitHub Actions 工作流排队及调度机制影响，实际构建完成时间可能较计划时间延后数分钟至数小时。_

## 快速下载

| 数据           | 文件                 | 内容                                                                         | GitHub Releases                                                                            | jsDelivr 镜像                                                                                                                                                                    |
| -------------- | -------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GeoIP 精简版   | `geoip-lite.dat`     | 面向中国大陆、俄罗斯和伊朗的黑白名单，以及私有网络、常用服务等 IP 数据       | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [默认](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoIP 完整版   | `geoip-full.dat`     | 包含精简版分流数据，另提供全球 IP 地理位置数据及任播（Anycast）IP            | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | 体积过大                                                                                                                                                                         |
| GeoSite 标准版 | `geosite.dat`        | 上游域名分类，以及针对中国大陆、俄罗斯和伊朗补充的域名黑白名单               | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [默认](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite 扩展版 | `geosite-expand.dat` | 包含标准版全部分类，并为中国大陆、俄罗斯和伊朗提供覆盖范围更广的域名黑白名单 | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | 体积过大                                                                                                                                                                         |

\* 精简版适合中国大陆、俄罗斯和伊朗的基本直连／代理分流。\
\* 完整版适合按流量目的地国家进一步精细分流，也适用于其他国家和地区的直连／代理分流需求。\
\* Xray 发布包默认随附精简版 GeoIP 与标准版 GeoSite 数据文件，分别以 `geoip.dat` 和 `geosite.dat` 为文件名。

### GeoIP 精简版与完整版的区别

| 类型                   | 分类名称                                                                                                       | 精简版 | 完整版 |
| ---------------------- | -------------------------------------------------------------------------------------------------------------- | ------ | ------ |
| IP 地理位置            | 两字母国家代码（见[下表](#全部国家和地区)）                                                                    | 不包含 | 包含   |
| 任播地址               | `anycast`                                                                                                      | 不包含 | 包含   |
| 地区定制集合（白名单） | `in-cn`、`in-ru`、`in-ir`                                                                                      | 包含   | 包含   |
| 黑名单                 | `blacklist-ru`、`blacklist-ir`                                                                                 | 包含   | 包含   |
| 私有网络与常用服务     | `private`、`cloudflare`、`cloudfront`、`facebook`、`fastly`、`google`、`netflix`、`telegram`、`twitter`、`tor` | 包含   | 包含   |
| 兼容别名               | `cn` → `in-cn`、`ru` → `in-ru`、`ir` → `in-ir`                                                                 | 提供   | 不提供 |

\* 兼容别名用于延续社区旧有的使用习惯，仅在精简版中提供。\
\* `in-cn` 收录适合从中国大陆直连的 IP，不代表这些 IP 位于中国大陆；完整版中的 `cn` 严格按 IP 地理位置划分，`ru`、`ir` 同理。

分类设计理念和数据来源详见 [GeoIP 说明文档](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_zh-Hans.md)。

#### 逐分类文件

只需个别分类时，可单独下载对应文件，无需下载体积较大的 `geoip-full.dat`。

例如，继续使用精简版的 `geoip.dat`，原有 `in-cn` 分类仍通过 `geoip:in-cn` 使用。额外下载美国 IP 分类文件 [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) 放在同一目录后，其中的 `us` 分类通过 `ext:us.dat:us` 使用。

逐分类 DAT 文件见 “release” 分支的 [`/geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) 目录。

#### 全部国家和地区

完整版包含以下 250 个国家和地区分类。

<details>
<summary>展开完整列表（按代码排序）</summary>

| 代码 | 国家或地区         | 代码 | 国家或地区           | 代码 | 国家或地区             |
| ---- | ------------------ | ---- | -------------------- | ---- | ---------------------- |
| `ad` | 安道尔             | `ae` | 阿拉伯联合酋长国     | `af` | 阿富汗                 |
| `ag` | 安提瓜和巴布达     | `ai` | 安圭拉               | `al` | 阿尔巴尼亚             |
| `am` | 亚美尼亚           | `ao` | 安哥拉               | `aq` | 南极洲                 |
| `ar` | 阿根廷             | `as` | 美属萨摩亚           | `at` | 奥地利                 |
| `au` | 澳大利亚           | `aw` | 阿鲁巴               | `ax` | 奥兰群岛               |
| `az` | 阿塞拜疆           | `ba` | 波斯尼亚和黑塞哥维那 | `bb` | 巴巴多斯               |
| `bd` | 孟加拉国           | `be` | 比利时               | `bf` | 布基纳法索             |
| `bg` | 保加利亚           | `bh` | 巴林                 | `bi` | 布隆迪                 |
| `bj` | 贝宁               | `bl` | 圣巴泰勒米           | `bm` | 百慕大                 |
| `bn` | 文莱               | `bo` | 玻利维亚             | `bq` | 荷属加勒比区           |
| `br` | 巴西               | `bs` | 巴哈马               | `bt` | 不丹                   |
| `bv` | 布韦岛             | `bw` | 博茨瓦纳             | `by` | 白俄罗斯               |
| `bz` | 伯利兹             | `ca` | 加拿大               | `cc` | 科科斯（基林）群岛     |
| `cd` | 刚果（金）         | `cf` | 中非共和国           | `cg` | 刚果（布）             |
| `ch` | 瑞士               | `ci` | 科特迪瓦             | `ck` | 库克群岛               |
| `cl` | 智利               | `cm` | 喀麦隆               | `cn` | 中国                   |
| `co` | 哥伦比亚           | `cr` | 哥斯达黎加           | `cu` | 古巴                   |
| `cv` | 佛得角             | `cw` | 库拉索               | `cx` | 圣诞岛                 |
| `cy` | 塞浦路斯           | `cz` | 捷克                 | `de` | 德国                   |
| `dj` | 吉布提             | `dk` | 丹麦                 | `dm` | 多米尼克               |
| `do` | 多米尼加共和国     | `dz` | 阿尔及利亚           | `ec` | 厄瓜多尔               |
| `ee` | 爱沙尼亚           | `eg` | 埃及                 | `eh` | 西撒哈拉               |
| `er` | 厄立特里亚         | `es` | 西班牙               | `et` | 埃塞俄比亚             |
| `fi` | 芬兰               | `fj` | 斐济                 | `fk` | 福克兰群岛             |
| `fm` | 密克罗尼西亚       | `fo` | 法罗群岛             | `fr` | 法国                   |
| `ga` | 加蓬               | `gb` | 英国                 | `gd` | 格林纳达               |
| `ge` | 格鲁吉亚           | `gf` | 法属圭亚那           | `gg` | 根西岛                 |
| `gh` | 加纳               | `gi` | 直布罗陀             | `gl` | 格陵兰                 |
| `gm` | 冈比亚             | `gn` | 几内亚               | `gp` | 瓜德罗普               |
| `gq` | 赤道几内亚         | `gr` | 希腊                 | `gs` | 南乔治亚和南桑威奇群岛 |
| `gt` | 危地马拉           | `gu` | 关岛                 | `gw` | 几内亚比绍             |
| `gy` | 圭亚那             | `hk` | 中国香港特别行政区   | `hm` | 赫德岛和麦克唐纳群岛   |
| `hn` | 洪都拉斯           | `hr` | 克罗地亚             | `ht` | 海地                   |
| `hu` | 匈牙利             | `id` | 印度尼西亚           | `ie` | 爱尔兰                 |
| `il` | 以色列             | `im` | 马恩岛               | `in` | 印度                   |
| `io` | 英属印度洋领地     | `iq` | 伊拉克               | `ir` | 伊朗                   |
| `is` | 冰岛               | `it` | 意大利               | `je` | 泽西岛                 |
| `jm` | 牙买加             | `jo` | 约旦                 | `jp` | 日本                   |
| `ke` | 肯尼亚             | `kg` | 吉尔吉斯斯坦         | `kh` | 柬埔寨                 |
| `ki` | 基里巴斯           | `km` | 科摩罗               | `kn` | 圣基茨和尼维斯         |
| `kp` | 朝鲜               | `kr` | 韩国                 | `kw` | 科威特                 |
| `ky` | 开曼群岛           | `kz` | 哈萨克斯坦           | `la` | 老挝                   |
| `lb` | 黎巴嫩             | `lc` | 圣卢西亚             | `li` | 列支敦士登             |
| `lk` | 斯里兰卡           | `lr` | 利比里亚             | `ls` | 莱索托                 |
| `lt` | 立陶宛             | `lu` | 卢森堡               | `lv` | 拉脱维亚               |
| `ly` | 利比亚             | `ma` | 摩洛哥               | `mc` | 摩纳哥                 |
| `md` | 摩尔多瓦           | `me` | 黑山                 | `mf` | 法属圣马丁             |
| `mg` | 马达加斯加         | `mh` | 马绍尔群岛           | `mk` | 北马其顿               |
| `ml` | 马里               | `mm` | 缅甸                 | `mn` | 蒙古                   |
| `mo` | 中国澳门特别行政区 | `mp` | 北马里亚纳群岛       | `mq` | 马提尼克               |
| `mr` | 毛里塔尼亚         | `ms` | 蒙特塞拉特           | `mt` | 马耳他                 |
| `mu` | 毛里求斯           | `mv` | 马尔代夫             | `mw` | 马拉维                 |
| `mx` | 墨西哥             | `my` | 马来西亚             | `mz` | 莫桑比克               |
| `na` | 纳米比亚           | `nc` | 新喀里多尼亚         | `ne` | 尼日尔                 |
| `nf` | 诺福克岛           | `ng` | 尼日利亚             | `ni` | 尼加拉瓜               |
| `nl` | 荷兰               | `no` | 挪威                 | `np` | 尼泊尔                 |
| `nr` | 瑙鲁               | `nu` | 纽埃                 | `nz` | 新西兰                 |
| `om` | 阿曼               | `pa` | 巴拿马               | `pe` | 秘鲁                   |
| `pf` | 法属波利尼西亚     | `pg` | 巴布亚新几内亚       | `ph` | 菲律宾                 |
| `pk` | 巴基斯坦           | `pl` | 波兰                 | `pm` | 圣皮埃尔和密克隆群岛   |
| `pn` | 皮特凯恩群岛       | `pr` | 波多黎各             | `ps` | 巴勒斯坦领土           |
| `pt` | 葡萄牙             | `pw` | 帕劳                 | `py` | 巴拉圭                 |
| `qa` | 卡塔尔             | `re` | 留尼汪               | `ro` | 罗马尼亚               |
| `rs` | 塞尔维亚           | `ru` | 俄罗斯               | `rw` | 卢旺达                 |
| `sa` | 沙特阿拉伯         | `sb` | 所罗门群岛           | `sc` | 塞舌尔                 |
| `sd` | 苏丹               | `se` | 瑞典                 | `sg` | 新加坡                 |
| `sh` | 圣赫勒拿           | `si` | 斯洛文尼亚           | `sj` | 斯瓦尔巴和扬马延       |
| `sk` | 斯洛伐克           | `sl` | 塞拉利昂             | `sm` | 圣马力诺               |
| `sn` | 塞内加尔           | `so` | 索马里               | `sr` | 苏里南                 |
| `ss` | 南苏丹             | `st` | 圣多美和普林西比     | `sv` | 萨尔瓦多               |
| `sx` | 荷属圣马丁         | `sy` | 叙利亚               | `sz` | 斯威士兰               |
| `tc` | 特克斯和凯科斯群岛 | `td` | 乍得                 | `tf` | 法属南部领地           |
| `tg` | 多哥               | `th` | 泰国                 | `tj` | 塔吉克斯坦             |
| `tk` | 托克劳             | `tl` | 东帝汶               | `tm` | 土库曼斯坦             |
| `tn` | 突尼斯             | `to` | 汤加                 | `tr` | 土耳其                 |
| `tt` | 特立尼达和多巴哥   | `tv` | 图瓦卢               | `tw` | 台湾                   |
| `tz` | 坦桑尼亚           | `ua` | 乌克兰               | `ug` | 乌干达                 |
| `um` | 美国本土外小岛屿   | `us` | 美国                 | `uy` | 乌拉圭                 |
| `uz` | 乌兹别克斯坦       | `va` | 梵蒂冈               | `vc` | 圣文森特和格林纳丁斯   |
| `ve` | 委内瑞拉           | `vg` | 英属维尔京群岛       | `vi` | 美属维尔京群岛         |
| `vn` | 越南               | `vu` | 瓦努阿图             | `wf` | 瓦利斯和富图纳         |
| `ws` | 萨摩亚             | `xk` | 科索沃               | `ye` | 也门                   |
| `yt` | 马约特             | `za` | 南非                 | `zm` | 赞比亚                 |
| `zw` | 津巴布韦           |      |                      |      |                        |

</details>

### GeoSite 标准版与扩展版的区别

| 类型                                                                    | 分类名称               | 标准版 | 扩展版 |
| ----------------------------------------------------------------------- | ---------------------- | ------ | ------ |
| [domain-list-community](https://github.com/v2fly/domain-list-community) | 与上游保持一致         | 包含   | 包含   |
| 中国大陆黑名单                                                          | `blocked-in-cn`        | 包含   | 包含   |
| 中国大陆白名单                                                          | `located-in-cn`        | 包含   | 包含   |
| 中国大陆扩展黑名单                                                      | `blocked-in-cn-expand` | 不包含 | 包含   |
| 中国大陆扩展白名单                                                      | `located-in-cn-expand` | 不包含 | 包含   |
| 俄罗斯黑名单                                                            | `blocked-in-ru`        | 包含   | 包含   |
| 俄罗斯白名单                                                            | `located-in-ru`        | 包含   | 包含   |
| 俄罗斯扩展黑名单                                                        | `blocked-in-ru-expand` | 不包含 | 包含   |
| 伊朗白名单                                                              | `located-in-ir`        | 包含   | 包含   |
| 伊朗扩展白名单                                                          | `located-in-ir-expand` | 不包含 | 包含   |

分类组合、迁移指南及数据来源详见 [GeoSite 说明文档](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_zh-Hans.md)。

#### 扩展版增量文件

只需个别地区的扩展数据时，可在标准版 `geosite.dat` 基础上下载对应增量，无需下载 `geosite-expand.dat`。

例如，将 [delta-ru.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geosite/delta/delta-ru.dat) 与标准版 `geosite.dat` 放在同一目录，上游 domain-list-community 的分类仍可通过 `geosite:xx` 使用，俄罗斯扩展黑名单则通过 `ext:delta-ru.dat:blocked-in-ru-expand` 使用。

增量 DAT 文件见 “release” 分支的 [`/geosite/delta/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geosite/delta) 目录。

## 下载说明

### GitHub Releases

GitHub Releases 仅提供整包 DAT 文件。最新及历史版本均见[发布页面](https://github.com/XTLS/Xray-rules-dat/releases)。

以下固定地址可用于自动更新，将 `<filename>` 替换为所需文件名即可获取最新版本：

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### “release” 分支

“release” 分支同样提供最新的整包文件，并额外提供逐分类和增量文件，但不保留历史版本。

[浏览该分支](https://github.com/XTLS/Xray-rules-dat/tree/release)查找所需文件，将其路径（如 `geoip/dat/us.dat`）填入 `<path>`，以下地址即可用于自动更新：

```text
https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>
```

<details>
<summary>无法访问 GitHub 时，可改用 jsDelivr 镜像</summary>

- 默认：`https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- Fastly：`https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

\* jsDelivr 受[单文件体积限制](https://github.com/jsdelivr/jsdelivr#restrictions)，即便是逐分类或增量文件也可能因体积过大而无法下载。建议优先从 GitHub 下载。

</details>

每个数据文件在同一位置都有对应的 `.sha256sum` 校验文件。

## Xray 用法

Xray 发布包默认随附精简版 GeoIP 和标准版 GeoSite，已分别命名为 `geoip.dat` 和 `geosite.dat`，可直接通过 `geoip:xx` 和 `geosite:xx` 语法使用。

以下示例额外使用美国 IP 文件 [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat)，三个文件放在同一目录。

将以下片段合并到现有 Xray 配置中。示例使用已有的 `direct`、`xhttp-us` 和 `xhttp-jp` 出站，每天按运行环境本地时间 04:00 自动更新并热重载三个文件：

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
    "domainStrategy": "IPOnDemand", // 按需解析域名，使域名流量也能匹配 IP 规则
    "rules": [
      {
        // 私有地址优先直连
        "ip": ["geoip:private"],
        "outboundTag": "direct"
      },
      {
        // 域名与 IP 白名单取交集，两项均命中才直连，减少单一名单误判的影响
        "domain": ["geosite:cn"],
        "ip": ["geoip:in-cn"],
        "outboundTag": "direct"
      },
      {
        // 从 us.dat 的 us 分类匹配美国 IP，使用 xhttp-us 出站
        "ip": ["ext:us.dat:us"],
        "outboundTag": "xhttp-us"
      },
      {
        // 中国大陆黑名单域名使用 xhttp-jp 出站
        "domain": ["geosite:blocked-in-cn"],
        "outboundTag": "xhttp-jp"
      }
    ]
  }
}
```

待更新文件须已存在，更新选项见 [Geodata 配置文档](https://xtls.github.io/config/geodata.html)。

以上仅为配置片段示例，实际使用需结合现有配置和分流需求调整。规则顺序及 IP、域名匹配方式见[路由文档](https://xtls.github.io/config/routing.html)。

最后一个冒号后的分类名称须存在于对应的 DAT 文件中，详见[从文件中加载 IP / 域名](https://xtls.github.io/config/routing.html#ruleobject)。

如需更多国家的 IP 数据，可下载 `geoip-full.dat`，重命名为 `geoip.dat` 后替换原文件，并将更新地址中的 `geoip-lite.dat` 改为 `geoip-full.dat`。此时可直接使用 `geoip:us` 等引用，无需额外下载 `us.dat` 或使用 `ext:us.dat:us`。

## 源码与许可证

GeoIP 源码、数据来源和构建说明维护在 [`geoip` 分支](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_zh-Hans.md)。

GeoSite 源码、数据来源和构建说明维护在 [`geosite` 分支](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_zh-Hans.md)。

本项目整体适用 [GPL-3.0](./LICENSE) 许可证。部分工具（GeoSite）分别适用 [MIT](./LICENSE-GeoSite) 许可证或 [继承许可条款](./LICENSE-GeoSite-inherit)。

## 致谢

- GeoIP 源码从 [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb) 分叉而来，后者分叉自 [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- GeoSite 源码从 [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618) 分叉并移植而来
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
