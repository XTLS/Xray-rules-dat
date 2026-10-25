# Xray-rules-dat

[English](./README.md) | 简体中文 | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

---

供 [Xray](https://github.com/XTLS/Xray-core) 使用的 GeoIP 和 GeoSite 数据。

GeoIP 以 [IPinfo Lite](https://ipinfo.io/lite) 为基础，结合各地分流需求和常用服务的 IP 数据整理，支持 IPv4 和 IPv6。

GeoSite 基于 [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) 构建，并在此基础上补充各地区存在刚性使用需求的黑名单与白名单数据集合。

数据按日定时构建，计划执行时间为 UTC 18:00（中国标准时间次日 02:00、俄罗斯莫斯科时间 21:00、伊朗时间 21:30）。受 GitHub Actions 工作流排队及调度机制影响，实际构建完成时间可能较计划时间延后数分钟至数小时。

## 快速下载

| 数据           | 文件                 | 内容                                                                               | GitHub Releases                                                                            | jsDelivr 镜像                                                                                                                                                                    |
| -------------- | -------------------- | ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GeoIP 完整版   | `geoip-full.dat`     | 涵盖全部国家和地区的原始分类、地区定制集合、黑名单及公共分类                       | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | 体积过大                                                                                                                                                                         |
| GeoIP 精简版   | `geoip-lite.dat`     | 涵盖地区定制集合、黑名单及公共分类，并提供兼容别名                                 | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [默认](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoSite 标准版 | `geosite.dat`        | 在数据覆盖完整性与编码后文件体积之间取得平衡的基础域名分类数据，适用于一般使用场景 | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [默认](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite 扩展版 | `geosite-expand.dat` | 以扩大列表覆盖范围为目标的域名分类数据，适用于需要更高数据覆盖程度的使用场景       | [下载](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | 体积过大                                                                                                                                                                         |

\* Xray 发布包默认随附标准版 GeoSite 与精简版 GeoIP 数据文件，分别以 `geosite.dat` 和 `geoip.dat` 为文件名。

### GeoIP 精简版与完整版的区别

| 分类               | 名称或范围                                                                                                     | 完整版 | 精简版 |
| ------------------ | -------------------------------------------------------------------------------------------------------------- | ------ | ------ |
| 原始国家和地区分类 | IPinfo 全部国家和地区的原始数据                                                                                | 包含   | 不包含 |
| 任播地址           | `anycast`                                                                                                      | 包含   | 不包含 |
| 地区定制集合       | `in-cn`、`in-ru`、`in-ir`                                                                                      | 包含   | 包含   |
| 黑名单             | `blacklist-ru`、`blacklist-ir`                                                                                 | 包含   | 包含   |
| 公共分类           | `private`、`cloudflare`、`cloudfront`、`facebook`、`fastly`、`google`、`netflix`、`telegram`、`twitter`、`tor` | 包含   | 包含   |
| 兼容别名           | `cn` → `in-cn`、`ru` → `in-ru`、`ir` → `in-ir`                                                                 | 不提供 | 提供   |

分类设计理念和数据来源详见 [GeoIP 说明文档](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_zh-Hans.md)。

### 逐分类文件

只需个别分类时，可单独下载对应文件，无需下载体积较大的 `geoip-full.dat`。

例如，继续使用精简版的 `geoip.dat`，原有 `in-cn` 分类仍通过 `geoip:in-cn` 使用。额外下载 [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) 放在同一目录后，`us` 分类通过 `ext:us.dat:us` 使用。

逐分类 DAT 文件见 “release” 分支的 [`geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) 目录。

### 全部国家和地区

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

### GeoSite 数据集合

分类组合、迁移指南及数据来源的具体说明，请参阅 [GeoSite 说明文档](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_zh-Hans.md)。

#### GeoSite 标准版

标准版（`geosite.dat`）由以下标签及其对应的数据集合构成：

- `v2fly/domain-list-community` 提供的原始数据，保持原样引入
- `available-only-in-ru`：来自 `runetfreedom/russia-domains-list` 中仅能在俄罗斯境内使用的域名
- `blocked-in-cn`：中国大陆黑名单域名
  - `gfwlist-cn`：来自 `gfwlist/gfwlist` 的数据
- `blocked-in-ru`：俄罗斯黑名单域名
  - `antifilter-list-community`：来自 Antifilter.Download community edition 的域名列表
  - `refilter`：来自 `1andrevich/Re-filter-lists` 的完整域名列表
- `mainly-ir`：来自 `Chocolate4U/Iran-v2ray-rules` 中经过有效性检查的用于绕过伊朗境内的域名集合

#### GeoSite 扩展版

扩展版（`geosite-expand.dat`）由以下标签及其对应的数据集合构成：

- `v2fly/domain-list-community` 提供的原始数据，保持原样引入
- `available-only-in-ru`：来自 `runetfreedom/russia-domains-list` 中仅能在俄罗斯境内使用的域名
- `blocked-in-cn-expand`：扩展中国大陆黑名单域名
  - `blocked-in-cn`：中国大陆黑名单域名
    - `gfwlist-cn`：来自 `gfwlist/gfwlist` 的数据
  - `greatfire-list`：来自 GreatFire.org 的数据
- `blocked-in-ru-expand`：扩展俄罗斯黑名单域名
  - `antifilter-list-full`：来自 Antifilter.Download 的域名列表
  - `blocked-in-ru`：俄罗斯黑名单域名
    - `antifilter-list-community`：来自 Antifilter.Download community edition 的域名列表
    - `refilter`：来自 `1andrevich/Re-filter-lists` 的完整域名列表
- `expand-cn`：来自 `felixonmars/dnsmasq-china-list` 中的中国大陆域名数据
- `expand-ir`：来自 `bootmortis/iran-hosted-domains` 中用于绕过伊朗境内的域名集合
- `mainly-ir`：来自 `Chocolate4U/Iran-v2ray-rules` 中经过有效性检查的用于绕过伊朗境内的域名集合

## 下载说明

### GitHub Releases

GitHub Releases 仅提供单文件 DAT 数据库。手动下载请打开[最新版本](https://github.com/XTLS/Xray-rules-dat/releases/latest)，历史版本见 [GitHub Releases](https://github.com/XTLS/Xray-rules-dat/releases)。

自动更新程序可通过以下固定地址下载最新版本，将 `<filename>` 替换为对应文件名：

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### “release” 分支

所有文件的最新版本统一发布在 [“release” 分支](https://github.com/XTLS/Xray-rules-dat/tree/release)。需要特定分类时，可直接从该分支下载对应文件。

将 `<path>` 替换为分支内的文件路径，如 `geoip/dat/us.dat`：

- GitHub Raw：`https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>`
- jsDelivr 镜像（默认）：`https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- jsDelivr 镜像（Fastly）：`https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

jsDelivr 镜像受[单文件体积限制](https://github.com/jsdelivr/jsdelivr#restrictions)，逐分类文件也可能因体积过大而无法下载。建议优先通过 GitHub Releases 或 GitHub Raw 下载，仅在无法访问 GitHub 时使用镜像。

每个数据文件在同一位置都有对应的 `.sha256sum` 校验文件。

## Xray 用法

Xray 发布包随附的是精简版 GeoIP，文件名为 `geoip.dat`。可直接使用其中的地区定制集合和公共分类。

以下示例继续使用精简版的 `geoip.dat`，并额外下载 [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) 到同一目录。

将以下片段合并到 Xray 配置中，每天按运行环境本地时间 04:00 自动更新并热重载这两个文件。私有地址和 `in-cn` 分类通过已有的 `direct` 出站，`us` 分类通过已有的 `wireguard` 出站：

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

待更新文件须已存在，更多选项见 [Geodata 配置文档](https://xtls.github.io/config/geodata.html)。

可按需替换示例中的分类名称。规则顺序和 IP 匹配方式见 Xray 的[路由文档](https://xtls.github.io/config/routing.html)。

原有 `in-cn` 分类仍通过 `geoip:in-cn` 使用，额外下载的 `us` 分类通过 `ext:us.dat:us` 使用。如需其他分类，替换文件名和分类名中的 `us`。

最后一个冒号后的分类名称必须存在于该 DAT 文件中。详见 Xray 文档中的[从文件加载 IP](https://xtls.github.io/config/routing.html#ruleobject)。

如需使用更多国家和地区的原始分类，也可以下载 `geoip-full.dat`，重命名为 `geoip.dat` 后替换原文件，并将上述更新地址中的 `geoip-lite.dat` 改为 `geoip-full.dat`。替换后，可直接通过 `geoip:us` 等方式使用原始分类，无需额外下载 `us.dat` 或使用 `ext:us.dat:us`。

## 源码与许可证

GeoIP 源码、数据来源和构建步骤维护在 [`geoip` 分支](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README_zh-Hans.md)。

GeoSite 源码、数据来源及构建工具统一维护于 [`geosite` 分支](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README_zh-Hans.md)。

本项目整体适用 [GPL-3.0](./LICENSE) 许可证。部分工具（GeoSite）分别适用 [MIT](./LICENSE-GeoSite) 许可证或 [继承许可条款](./LICENSE-GeoSite-inherit)。

## 致谢

- GeoIP 源码从 [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb) 分叉而来，后者分叉自 [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- GeoSite 源码由 [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618) 分叉并移植而来
- [IPInfo Lite](https://ipinfo.io/lite)
- [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community)：GeoSite 上游数据
- [zhufengme/block_cn_files](https://github.com/zhufengme/block_cn_files)
- [misakaio/chnroutes2](https://github.com/misakaio/chnroutes2)
- [gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip)
- [Loyalsoldier/cn-blocked-domain](https://github.com/Loyalsoldier/cn-blocked-domain)
- [felixonmars/dnsmasq-china-list](https://github.com/felixonmars/dnsmasq-china-list)
- [gfwlist/gfwlist](https://github.com/gfwlist/gfwlist)
- [bgptools/anycast-prefixes](https://github.com/bgptools/anycast-prefixes)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [bootmortis/iran-hosted-domains](https://github.com/bootmortis/iran-hosted-domains)
- [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip)
- [hxehex/russia-mobile-internet-whitelist](https://github.com/hxehex/russia-mobile-internet-whitelist)
- [Antifilter](https://antifilter.download/)
- [Antifilter.Download community edition](https://community.antifilter.download/)
- [1andrevich/Re-filter-lists](https://github.com/1andrevich/Re-filter-lists)
- [runetfreedom/russia-domains-list](https://github.com/runetfreedom/russia-domains-list)
