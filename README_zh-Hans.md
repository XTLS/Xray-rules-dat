# GeoIP

[English](./README.md) | 简体中文 | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

本仓库基于 IPInfo Lite 构建 GeoIP 数据，提供原始国家和地区分类、面向不同分流环境的地区定制集合及公共分类，支持 Xray、sing-box、Clash 等应用。

## 设计理念

GeoIP 数据的覆盖程度会直接影响分流效果。在国际多出口分流场景下，MaxMind 的地址覆盖未能满足要求，因此采用 IPInfo Lite 作为原始国家和地区分类的数据基线，并单列 anycast，避免按国家分流造成绕路。

目标 IP 的国家归属与用户所处环境的分流需求，是两个不同的问题。国家分类回答的是“这个目标 IP 归属于哪个国家或地区”；对于中国大陆、伊朗、俄罗斯等分流环境，用户还需要回答“身处这个国家或地区时，哪些目标地址应当采用相应的路由策略”。仅凭目标地址的国家归属，无法完整表达这类需求。

因此，本项目引入 `in-*` 地区定制集合，作为相应地区的直连白名单。`in-cn`、`in-ir`、`in-ru` 等分别面向对应地区的分流环境，依据当地的访问条件和分流需求整理。例如，某些海外服务、Telegram、Twitter 或 CDN 地址在当地可以直接、高效地访问，就可以按当地需求纳入相应集合；特殊需求下也可能包含私有地址。

原始国家分类也有独立用途。例如，配置多个国家或地区的出口节点时，用户可能需要按照目标 IP 的国家归属选择节点。此时应使用普通国家码；`in-*` 集合则用于匹配当前所在环境的分流需求。

完整版 `geoip-full.dat` 因此同时保留这两类数据：普通国家码采用 IPInfo Lite 原始分类，`in-*` 承载地区定制集合，定制数据不会改写原始国家分类。

精简版 `geoip-lite.dat` 不含完整国家分类，不适用于按目标 IP 的国家归属选择出口节点。它保留的普通国家码仅是地区定制集合的别名，用于兼容社区旧有的使用习惯。

两个版本均包含 IPv4 和 IPv6 数据。

## 地区定制集合

### 中国大陆：`in-cn`

`in-cn` 表示适用于中国大陆分流环境的直连白名单。

IPv4 合并以下两个来源：

- misakaio/chnroutes2
- gaoyifan/china-operator-ip 的 china.txt

IPv6 仅使用 gaoyifan/china-operator-ip 的 china6.txt。

中国大陆因 ISP 劣化国际网络，境外地址约等于黑名单，因此只维护直连白名单，不单列黑名单。

### 伊朗：`in-ir`

以 IPInfo Lite 原始 IR IPv4、IPv6 为基础，使用本地 ASN 数据补充 Eitaa（AS202798）和 GAP（AS210877），并合并 ArvanCloud、DerakCloud、IranServer、ParsPack 的 CDN 地址列表，以及伊朗特殊的六个本地服务私有地址。ParsPack 使用 Chocolate4U/Iran-v2ray-rules 提供的纯文本镜像。另提供 `blacklist-ir`，收录封锁注入地址与其他注入地址。

### 俄罗斯：`in-ru`

以 IPInfo Lite 原始 RU IPv4、IPv6 为基础，合并 hxehex/russia-mobile-internet-whitelist 提供的移动网络白名单。另提供 `blacklist-ru`，合并 Antifilter、1andrevich/Re-filter-lists 的 IP 列表及参考配置补充的 Telegram 网段。

## 公共分类

完整版和精简版均包含以下公共分类：

- `geoip:private`
- `geoip:cloudflare`
- `geoip:cloudfront`
- `geoip:facebook`
- `geoip:fastly`
- `geoip:google`
- `geoip:netflix`
- `geoip:telegram`
- `geoip:twitter`
- `geoip:tor`

`geoip:anycast` 仅包含在完整版中。

## 构建

在仓库根目录准备 `./ipinfo/ipinfo_lite.mmdb`，并确保 `./tmp/geoip-in` 和 `./output` 为干净目录，然后依次执行：

```sh
go run . convert -c config-geoip-in.json
go run . convert -c config-geoip-lite.json
go run . convert -c config-geoip-full.json
```

第一步在 `./tmp/geoip-in/` 生成地区定制集合与黑名单的临时文本。lite 和 full 均直接读取这些文件，第一步完成后可独立运行或交换执行顺序；三个步骤之间不要清理目录。

lite 只输出 `./output/geoip-lite.dat`。full 输出 `./output/geoip-full.dat`，并生成下列逐分类文件：

| 输出目录                  | 格式                    |
| ------------------------- | ----------------------- |
| `output/dat/`             | Xray GeoIP DAT          |
| `output/srs/`             | sing-box SRS            |
| `output/mrs/`             | Mihomo MRS              |
| `output/text/`            | 纯文本 CIDR             |
| `output/clash/classical/` | Clash classical RuleSet |
| `output/clash/ipcidr/`    | Clash ipcidr RuleSet    |

构建命令必须显式指定 `-c` 配置文件。更多配置选项见[配置说明](./configuration.md)。

## License

- [GPL-3.0](./LICENSE)

## 致谢

- 本项目从 [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb) 分叉而来，后者分叉自 [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- [IPInfo Lite](https://ipinfo.io/lite)
- [misakaio/chnroutes2](https://github.com/misakaio/chnroutes2)
- [gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip)
- [bgptools/anycast-prefixes](https://github.com/bgptools/anycast-prefixes)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip)
- [hxehex/russia-mobile-internet-whitelist](https://github.com/hxehex/russia-mobile-internet-whitelist)
- [Antifilter](https://antifilter.download/)
- [1andrevich/Re-filter-lists](https://github.com/1andrevich/Re-filter-lists)
