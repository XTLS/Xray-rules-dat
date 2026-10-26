# GeoSite

[English](./README.md) | 简体中文 | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

本仓库基于 [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) 仓库中的原始域名分类数据，使用基于 v2fly/domain-list-community 包含的标准数据编码器移植的编码器，在原有域名分类数据的基础上通过增补特定数据源的方式，构建符合标准语义的 GeoSite 数据包。

## 设计理念

GeoSite 数据的精准程度与覆盖程度，会对基于域名的分流效果带来影响，而 GeoSite 的语义设计也会对此存在影响。因此，本仓库采用由 v2fly/domain-list-community 移植的数据编码器对相关数据分类进行编码构建，以最大程度上在保证数据语义完整的同时，能完整跟进上游更改。

同时，在完整实现原有上游域名分类数据的基础上，引入附加的实用域名分类表，并且将引入的额外域名表转换成合乎 GeoSite 语义的数据。这些新引入的域名分类表可以视为在上游数据之外的另一层补充和扩充。

考虑到数据集在完备程度和编码后体积之间的平衡，本仓库将编码后的数据包分为两个版本：标准版（geosite.dat）和扩展版（geosite-expand.dat）。另外为了便于跟踪以保证编码构建工具与上游行为保持一致，另有可以视为仅对上游的数据进行镜像构建的验证版本（validation.dat）。

## 数据内容

验证版本（validation.dat）仅包含来自 v2fly/domain-list-community 的数据。

### 标准版 geosite.dat

| 内容 | 描述 | 包含于 | 包含 | 备注 |
|----|----|----|----|----|
| 与 domain-list-community 一致的数据索引 | 来自 v2fly/domain-list-community 的完整数据 | | | 未对相关数据索引进行改动或调整 |
| `gfwlist-cn` | 来自 GFWList 的数据，经过 abpextract 转换为 GeoSite 语义 | `blocked-in-cn` | | |
| `mainly-ir` | 来自 Chocolate4U/Iran-v2ray-rules 中经过有效性检查的用于绕过伊朗境内的域名集合 | | | 与 Chocolate4U/Iran-v2ray-rules 中 `category-ir` 一致 |
| `available-only-in-ru` | 来自 runetfreedom/russia-domains-list 中仅能在俄罗斯境内使用的域名 | | | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `ru-available-only-inside` 一致 |
| `refilter` | 来自 Re:filter 的完整域名列表 | `blocked-in-ru` | | |
| `antifilter-list-community` | 来自 Antifilter.Download community edition 的域名列表 | `blocked-in-ru` | | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `antifilter-download-community` 一致 |
| `blocked-in-cn` | 中国大陆黑名单域名 | |  `gfwlist-cn` | |
| `blocked-in-ru` | 俄罗斯黑名单域名 | | `refilter` `antifilter-list-community` | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `ru-blocked` 一致 |

### 扩展版 geosite-expand.dat

| 内容 | 描述 | 包含于 | 包含 | 备注 |
|----|----|----|----|----|
| 与 domain-list-community 一致的数据索引 | 来自 v2fly/domain-list-community 的完整数据 | | | 未对相关数据索引进行改动或调整 |
| `gfwlist-cn` | 来自 GFWList 的数据，经过 abpextract 转换为 GeoSite 语义 | `blocked-in-cn` `blocked-in-cn-expand` | | |
| `expand-cn` | 来自 felixonmars/dnsmasq-china-list 中的中国大陆域名数据，经过 dnsmasqextract 转换为域名数据列表 | `additional-set-cn` | | 不建议与 `cn` `tld-cn` `geosite-cn` 同时使用 |
| `greatfire-list` | 来自 GreatFire.org 的数据，由 Loyalsoldier/cn-blocked-domain 转换 | `blocked-in-cn-expand` | | |
| `mainly-ir` | 来自 Chocolate4U/Iran-v2ray-rules 中经过有效性检查的用于绕过伊朗境内的域名集合 | | | 与 Chocolate4U/Iran-v2ray-rules 中 `category-ir` 一致 |
| `expand-ir` | 来自 Iran Hosted Domains 中用于绕过伊朗境内的域名集合 | | | 与 Chocolate4U/Iran-v2ray-rules 中 `ir` 一致 |
| `available-only-in-ru` | 来自 runetfreedom/russia-domains-list 中仅能在俄罗斯境内使用的域名 | | | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `ru-available-only-inside` 一致 |
| `refilter` | 来自 Re:filter 的完整域名列表 | `blocked-in-ru` `blocked-in-ru-expand` | | |
| `antifilter-list-community` | 来自 Antifilter.Download community edition 的域名列表 | `blocked-in-ru` `blocked-in-ru-expand` | | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `antifilter-download-community` 一致 |
| `antifilter-list-full` | 来自 Antifilter.Download 的域名列表 | `blocked-in-ru-expand` | | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `antifilter-download` 一致；记录数过多，慎用 |
| `blocked-in-cn` | 中国大陆黑名单域名 | `blocked-in-cn-expand` | `gfwlist-cn` | |
| `blocked-in-cn-expand` | 扩展中国大陆黑名单域名 | | `blocked-in-cn` `greatfire-list` | |
| `blocked-in-ru` | 俄罗斯黑名单域名 | `blocked-in-ru-expand` | `refilter` `antifilter-list-community` | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `ru-blocked` 一致 |
| `blocked-in-ru-expand` | 扩展俄罗斯黑名单域名 | | `blocked-in-ru` `antifilter-list-full` | 与 russia-blocked-geosite 和 russia-v2ray-rules-dat 中 `ru-blocked-all` 一致；记录数过多，慎用 |

## 构建

在仓库根目录准备 `./data` 目录，目录中应包含经过转换与清洗的原始分类数据文件。其中，`./predata/std` 与 `./predata/exp` 中包含的内容分别为标准版附加索引以及扩展版附加索引，须在数据目录中存在相关分类的情况下方可复制到 `./data` 下启用这些附加索引。

然后执行以下命令，在当前目录中生成编码构建后的 geosite.dat：

``` sh
go run ./main.go --datapath ./data --outputname geosite.dat --outputdir ./
```

``` sh
go run ./main.go --datapath ./data --datprofile ./tasks/standard.json --outputdir ./
go run ./main.go --datapath ./data --datprofile ./tasks/expand.json --outputdir ./
```

## License

- [MIT](./LICENSE-inherited)：由 v2fly 移植的数据编码器以及数据导出器
- [MIT](./LICENSE)：本仓库提供的附加工具
  - abpextract：此工具将使用简单 AdBlock Plus 规则编码的文本列表转换成符合 GeoSite 标准规则及语义的数据列表
  - dnsmasqextract：此工具将使用 DNSMasq 规则编码的文本列表转换成符合 GeoSite 标准规则的数据列表
  - sanitizer：此工具将包含不符合 GeoSite 标准规则的数据列表进行数据清洗及转换使其符合标准规则

## 致谢

- 本项目从 [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618) 获得数据编码器、数据导出器并进行移植而来
- [Domain list community](https://github.com/v2fly/domain-list-community)
- [Antifilter](https://antifilter.download/)
- [Antifilter.Download community edition](https://community.antifilter.download/)
- [Loyalsoldier/cn-blocked-domain](https://github.com/Loyalsoldier/cn-blocked-domain)
- [felixonmars/dnsmasq-china-list](https://github.com/felixonmars/dnsmasq-china-list)
- [GFWList](https://github.com/gfwlist/gfwlist)
- [Iran Hosted Domains](https://github.com/bootmortis/iran-hosted-domains)
- [Chocolate4U/Iran-v2ray-rules](https://github.com/Chocolate4U/Iran-v2ray-rules)
- [`Re:filter`](https://github.com/1andrevich/Re-filter-lists)
- [runetfreedom/russia-domains-list](https://github.com/runetfreedom/russia-domains-list)
