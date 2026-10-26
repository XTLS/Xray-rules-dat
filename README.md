# GeoSite

English | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

This repository builds GeoSite data packages that conform to the standard GeoSite semantics, using the original domain classification data from [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community) as its upstream baseline. The build process uses a ported version of the standard data encoder included in `v2fly/domain-list-community`, augmenting the original domain classification data with additional datasets from selected sources.

## Design Philosophy

The accuracy and coverage of GeoSite data affect the effectiveness of domain-based traffic routing. The semantic design of GeoSite data also plays a role in determining routing outcomes. Accordingly, this repository uses the data encoder ported from `v2fly/domain-list-community` to encode and build the relevant domain categories, preserving semantic integrity to the greatest extent possible while maintaining full compatibility with upstream changes.

On top of implementing the complete upstream domain classification dataset, this repository incorporates additional practical domain lists and transforms them into datasets that conform to GeoSite semantics. These additional lists serve as a supplementary layer of coverage beyond the upstream dataset.

To balance dataset completeness against the size of the encoded artifacts, this repository provides two builds: the standard edition (`geosite.dat`) and the expanded edition (`geosite-expand.dat`). For tracking and verification purposes, a separate verification build (`validation.dat`) is also provided. This build can be regarded as a mirror build of upstream data only, intended to verify that the encoding and build tooling remains consistent with upstream behavior.

## Data Contents

The verification build (`validation.dat`) contains only data sourced from `v2fly/domain-list-community`.

### Standard Edition: `geosite.dat`

| Entry | Description | Included In | Includes | Notes |
|---|---|---|---|---|
| Data indices matching `domain-list-community` | Complete dataset from `v2fly/domain-list-community` | | | No modifications or adjustments have been made to the corresponding data indices. |
| `gfwlist-cn` | Data sourced from GFWList, converted into GeoSite-compatible semantics by `abpextract` | `blocked-in-cn` | | |
| `mainly-ir` | A validated set of domains inside Iran, sourced from `Chocolate4U/Iran-v2ray-rules` | | | Matches `category-ir` in `Chocolate4U/Iran-v2ray-rules`. |
| `available-only-in-ru` | Domains available exclusively within Russia, sourced from `runetfreedom/russia-domains-list` | | | Matches `ru-available-only-inside` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |
| `refilter` | Complete domain list from Re:filter | `blocked-in-ru` | | |
| `antifilter-list-community` | Domain list from the community edition of Antifilter.Download | `blocked-in-ru` | | Matches `antifilter-download-community` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |
| `blocked-in-cn` | Blocklist for mainland China | | `gfwlist-cn` | |
| `blocked-in-ru` | Blocklist for Russia | | `refilter`, `antifilter-list-community` | Matches `ru-blocked` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |

### Expanded Edition: `geosite-expand.dat`

| Entry | Description | Included In | Includes | Notes |
|---|---|---|---|---|
| Data indices matching `domain-list-community` | Complete dataset from `v2fly/domain-list-community` | | | No modifications or adjustments have been made to the corresponding data indices. |
| `gfwlist-cn` | Data sourced from GFWList, converted into GeoSite-compatible semantics by `abpextract` | `blocked-in-cn`, `blocked-in-cn-expand` | | |
| `expand-cn` | Mainland China domain data from `felixonmars/dnsmasq-china-list`, converted into a domain list by `dnsmasqextract` | `additional-set-cn` | | Not recommended for simultaneous use with `cn`, `tld-cn`, or `geosite-cn`. |
| `greatfire-list` | Data sourced from GreatFire.org and converted by `Loyalsoldier/cn-blocked-domain` | `blocked-in-cn-expand` | | |
| `mainly-ir` | A validated set of domains inside Iran, sourced from `Chocolate4U/Iran-v2ray-rules` | | | Matches `category-ir` in `Chocolate4U/Iran-v2ray-rules`. |
| `expand-ir` | A set of domains inside Iran, sourced from Iran Hosted Domains | | | Matches `ir` in `Chocolate4U/Iran-v2ray-rules`. |
| `available-only-in-ru` | Domains available exclusively within Russia, sourced from `runetfreedom/russia-domains-list` | | | Matches `ru-available-only-inside` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |
| `refilter` | Complete domain list from Re:filter | `blocked-in-ru`, `blocked-in-ru-expand` | | |
| `antifilter-list-community` | Domain list from the community edition of Antifilter.Download | `blocked-in-ru`, `blocked-in-ru-expand` | | Matches `antifilter-download-community` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |
| `antifilter-list-full` | Domain list from Antifilter.Download | `blocked-in-ru-expand` | | Matches `antifilter-download` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. Excessive record count; use with caution. |
| `blocked-in-cn` | Blocklist for mainland China | `blocked-in-cn-expand` | `gfwlist-cn` | |
| `blocked-in-cn-expand` | Expanded blocklist for mainland China | | `blocked-in-cn`, `greatfire-list` | |
| `blocked-in-ru` | Blocklist for Russia | `blocked-in-ru-expand` | `refilter`, `antifilter-list-community` | Matches `ru-blocked` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. |
| `blocked-in-ru-expand` | Expanded blocklist for Russia | | `blocked-in-ru`, `antifilter-list-full` | Matches `ru-blocked-all` in `russia-blocked-geosite` and `russia-v2ray-rules-dat`. Excessive record count; use with caution. |

## Build

Prepare a `./data` directory at the repository root. This directory must contain the source classification files after conversion and sanitization.

The `./predata/std` and `./predata/exp` directories contain the supplemental indices for the standard and expanded editions, respectively. To enable these additional indices, copy the relevant files into `./data`. A corresponding domain category must already exist in the data directory for its supplemental index to be applicable.

Then execute the following command to generate the encoded `geosite.dat` artifact in the current directory:

```sh
go run ./main.go --datapath ./data --outputname geosite.dat --outputdir ./
```

``` sh
go run ./main.go --datapath ./data --datprofile ./tasks/standard.json --outputdir ./
go run ./main.go --datapath ./data --datprofile ./tasks/expand.json --outputdir ./
```

## License

- [MIT](./LICENSE-inherited): The data encoder and data exporter ported from `v2fly`.
- [MIT](./LICENSE): The additional utilities provided by this repository.
  - `abpextract`: Converts text lists encoded using basic Adblock Plus rules into datasets conforming to GeoSite's standard rules and semantics.
  - `dnsmasqextract`: Converts text lists encoded using DNSMasq rules into domain datasets conforming to GeoSite's standard rules.
  - `sanitizer`: Sanitizes and transforms domain datasets containing entries that do not conform to GeoSite's standard rules, bringing them into compliance.

## Acknowledgments

- This project ports the data encoder and data exporter from [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618).
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
