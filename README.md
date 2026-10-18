# GeoIP

本仓库基于 IPInfo Lite 构建 GeoIP 数据，提供原始国家和地区分类、面向不同分流环境的地区定制集合及公共分类，支持 Xray、sing-box、Clash 等应用。

## 设计理念

GeoIP 数据的覆盖程度会直接影响分流效果。对本项目的国际分流需求而言，MaxMind 的地址覆盖未能满足要求，因此采用 IPInfo Lite 作为原始国家和地区分类的数据基线。

目标 IP 的国家归属与用户所处环境的分流需求，是两个不同的问题。国家分类回答的是“这个目标 IP 归属于哪个国家或地区”；对于中国大陆、伊朗、俄罗斯等分流环境，用户还需要回答“身处这个国家或地区时，哪些目标地址应当采用相应的路由策略”。仅凭目标地址的国家归属，无法完整表达这类需求。

因此，本项目引入 `in-*` 地区定制集合。`in-cn`、`in-ir`、`in-ru` 分别面向对应地区的分流环境，依据当地的访问条件和分流需求整理。例如，某些海外服务、Telegram、Twitter 或 CDN 地址在当地可以直接、高效地访问，就可以按当地需求纳入相应集合；特殊需求下也可能包含私有地址。集合不受原始国家归属限制，允许彼此重叠，具体采取直连、代理还是封锁，由用户的路由规则决定。

原始国家分类也有独立用途。例如，在本地配置多个国家或地区的出口节点时，用户可能需要按照目标 IP 的国家归属选择节点。此时应使用普通国家码；`in-*` 集合则用于匹配当前所在环境的分流需求。

完整版因此同时保留这两类数据：普通国家码采用 IPInfo Lite 原始分类，`in-*` 承载地区定制集合，定制数据不会改写原始国家分类。精简版中的 `cn` 则作为 `in-cn` 的别名，以兼容社区旧有的使用习惯。

## 数据版本

| 文件             | 包含的分类                                 | `cn`                          | `in-cn`     |
| ---------------- | ------------------------------------------ | ----------------------------- | ----------- |
| `geoip.dat`      | 完整国家和地区分类、地区定制集合、公共分类 | IPInfo Lite 原始 CN 数据      | CN 定制集合 |
| `geoip-lite.dat` | 地区定制集合、公共分类及兼容别名 `cn`      | 与 `in-cn` 完全一致，用于兼容 | CN 定制集合 |

两个版本均包含 IPv4 和 IPv6 数据。精简版不包含其他普通国家和地区分类。

## 数据基线

完整版中的普通国家和地区分类直接使用 IPInfo Lite 原始数据。源库中没有 CIDR 的国家和地区会保留为空类别，避免输出文件缺失。地区定制集合和公共分类独立整理，不修改原始国家分类。

## 中国大陆定制集合

`in-cn` 表示适用于中国大陆分流环境的目标 IP 集合，具体的直连、封锁等行为由用户的路由规则决定。

IPv4 合并以下三个来源，不加入 IPInfo Lite 原始 CN IPv4：

- zhufengme/block_cn_files
- misakaio/chnroutes2
- gaoyifan/china-operator-ip 的 china.txt

IPv6 合并 IPInfo Lite 原始 CN IPv6 与 gaoyifan/china-operator-ip 的 china6.txt。

## 公共分类

完整版和精简版均包含以下公共分类：

- `geoip:private`（`GEOIP,PRIVATE`）
- `geoip:anycast`（`GEOIP,ANYCAST`）
- `geoip:cloudflare`（`GEOIP,CLOUDFLARE`）
- `geoip:cloudfront`（`GEOIP,CLOUDFRONT`）
- `geoip:facebook`（`GEOIP,FACEBOOK`）
- `geoip:fastly`（`GEOIP,FASTLY`）
- `geoip:google`（`GEOIP,GOOGLE`）
- `geoip:netflix`（`GEOIP,NETFLIX`）
- `geoip:telegram`（`GEOIP,TELEGRAM`）
- `geoip:twitter`（`GEOIP,TWITTER`）
- `geoip:tor`（`GEOIP,TOR`）

`private` 提供局域网和私有网络地址。其他公共分类的数据来源包括官方 IP 列表、公开前缀列表，以及基于 IPInfo Lite ASN 数据提取的结果。

## 构建

在仓库根目录准备 `./ipinfo/ipinfo_lite.mmdb`，并确保 `./tmp/geoip-in` 和 `./output` 为干净目录，然后依次执行：

```sh
go run . convert -c config-geoip-in.json
go run . convert -c config-geoip-lite.json
go run . convert -c config-geoip-full.json
```

第一步只生成临时文本 `./tmp/geoip-in/cn.txt`。lite 和 full 均直接读取该文件，第一步完成后可独立运行或交换执行顺序；三个步骤之间不要清理目录。

lite 只输出 `./output/geoip-lite.dat`。full 输出 `./output/geoip.dat`，并生成下列逐分类文件：

| 输出目录                  | 格式                    |
| ------------------------- | ----------------------- |
| `output/dat/`             | Xray GeoIP DAT          |
| `output/srs/`             | sing-box SRS            |
| `output/mrs/`             | Mihomo MRS              |
| `output/text/`            | 纯文本 CIDR             |
| `output/clash/classical/` | Clash classical RuleSet |
| `output/clash/ipcidr/`    | Clash ipcidr RuleSet    |

full 的逐分类文件中，`cn` 使用 IPInfo Lite 原始数据，`in-cn` 使用定制集合。例如，`output/dat/cn.dat` 和 `output/dat/in-cn.dat` 分别保存这两类数据。

构建命令必须显式指定 `-c` 配置文件。更多配置选项见[配置说明](./configuration.md)。

## License

- [GPL-3.0](./LICENSE)

## 致谢

- 本项目从 [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc) 分叉而来
- [IPInfo Lite](https://ipinfo.io/lite)
- [zhufengme/block_cn_files](https://github.com/zhufengme/block_cn_files)
- [misakaio/chnroutes2](https://github.com/misakaio/chnroutes2)
- [gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip)
- [bgptools/anycast-prefixes](https://github.com/bgptools/anycast-prefixes)
