# Xray-rules-dat

English | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

---

GeoIP and GeoSite data for [Xray](https://github.com/XTLS/Xray-core).

GeoIP is based on [IPinfo Lite](https://ipinfo.io/lite), supplemented with IP data for regional routing needs and commonly used services. Both IPv4 and IPv6 are supported.

GeoSite is based on [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community), with additional blocklists and allowlists for regions where such lists are required for practical use.

Builds are scheduled daily at 18:00 UTC (02:00 the following day in China, 21:00 in Moscow, and 21:30 in Iran). Due to GitHub Actions workflow queuing and scheduling, actual build completion may be delayed by several minutes to several hours.

## Quick download

| Data             | File                 | Contents                                                                                                                 | GitHub Releases                                                                                | jsDelivr mirrors                                                                                                                                                                    |
| ---------------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GeoIP full       | `geoip-full.dat`     | Original categories for all countries and territories, regional custom sets, blocklists, and common categories           | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | File too large                                                                                                                                                                      |
| GeoIP lite       | `geoip-lite.dat`     | Regional custom sets, blocklists, and common categories, with compatibility aliases                                      | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [Default](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoSite Standard | `geosite.dat`        | A baseline domain classification dataset that balances data coverage against encoded file size, suitable for general use | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [Default](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite Extended | `geosite-expand.dat` | A domain classification dataset with broader list coverage, intended for use cases that require more comprehensive data  | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | File too large                                                                                                                                                                      |

\* Xray release packages include the GeoSite standard dataset and the GeoIP lite dataset by default, named `geosite.dat` and `geoip.dat`, respectively.

### Differences between GeoIP lite and full

| Category                                  | Name or scope                                                                                                  | Full         | Lite         |
| ----------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------ | ------------ |
| Original country and territory categories | Original IPinfo data for all countries and territories                                                         | Included     | Not included |
| Anycast addresses                         | `anycast`                                                                                                      | Included     | Not included |
| Regional custom sets                      | `in-cn`, `in-ru`, `in-ir`                                                                                      | Included     | Included     |
| Blocklists                                | `blacklist-ru`, `blacklist-ir`                                                                                 | Included     | Included     |
| Common categories                         | `private`, `cloudflare`, `cloudfront`, `facebook`, `fastly`, `google`, `netflix`, `telegram`, `twitter`, `tor` | Included     | Included     |
| Compatibility aliases                     | `cn` → `in-cn`, `ru` → `in-ru`, `ir` → `in-ir`                                                                 | Not provided | Provided     |

See the [GeoIP documentation](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README.md) for the category design and data sources.

### Individual category files

If you only need specific categories, download their files individually without downloading the larger `geoip-full.dat`.

For example, keep using the lite version as `geoip.dat`; the existing `in-cn` category remains available through `geoip:in-cn`. Download [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) into the same directory to use the `us` category through `ext:us.dat:us`.

Individual category DAT files are available in the [`geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) directory on the `release` branch.

### All countries and territories

The full version includes the following 250 country and territory categories.

<details>
<summary>Show the complete list (sorted by code)</summary>

| Code | Country or territory                 | Code | Country or territory     | Code | Country or territory                         |
| ---- | ------------------------------------ | ---- | ------------------------ | ---- | -------------------------------------------- |
| `ad` | Andorra                              | `ae` | United Arab Emirates     | `af` | Afghanistan                                  |
| `ag` | Antigua and Barbuda                  | `ai` | Anguilla                 | `al` | Albania                                      |
| `am` | Armenia                              | `ao` | Angola                   | `aq` | Antarctica                                   |
| `ar` | Argentina                            | `as` | American Samoa           | `at` | Austria                                      |
| `au` | Australia                            | `aw` | Aruba                    | `ax` | Åland Islands                                |
| `az` | Azerbaijan                           | `ba` | Bosnia and Herzegovina   | `bb` | Barbados                                     |
| `bd` | Bangladesh                           | `be` | Belgium                  | `bf` | Burkina Faso                                 |
| `bg` | Bulgaria                             | `bh` | Bahrain                  | `bi` | Burundi                                      |
| `bj` | Benin                                | `bl` | Saint Barthélemy         | `bm` | Bermuda                                      |
| `bn` | Brunei                               | `bo` | Bolivia                  | `bq` | Caribbean Netherlands                        |
| `br` | Brazil                               | `bs` | Bahamas                  | `bt` | Bhutan                                       |
| `bv` | Bouvet Island                        | `bw` | Botswana                 | `by` | Belarus                                      |
| `bz` | Belize                               | `ca` | Canada                   | `cc` | Cocos (Keeling) Islands                      |
| `cd` | Democratic Republic of the Congo     | `cf` | Central African Republic | `cg` | Republic of the Congo                        |
| `ch` | Switzerland                          | `ci` | Côte d’Ivoire            | `ck` | Cook Islands                                 |
| `cl` | Chile                                | `cm` | Cameroon                 | `cn` | China                                        |
| `co` | Colombia                             | `cr` | Costa Rica               | `cu` | Cuba                                         |
| `cv` | Cape Verde                           | `cw` | Curaçao                  | `cx` | Christmas Island                             |
| `cy` | Cyprus                               | `cz` | Czechia                  | `de` | Germany                                      |
| `dj` | Djibouti                             | `dk` | Denmark                  | `dm` | Dominica                                     |
| `do` | Dominican Republic                   | `dz` | Algeria                  | `ec` | Ecuador                                      |
| `ee` | Estonia                              | `eg` | Egypt                    | `eh` | Western Sahara                               |
| `er` | Eritrea                              | `es` | Spain                    | `et` | Ethiopia                                     |
| `fi` | Finland                              | `fj` | Fiji                     | `fk` | Falkland Islands                             |
| `fm` | Micronesia                           | `fo` | Faroe Islands            | `fr` | France                                       |
| `ga` | Gabon                                | `gb` | United Kingdom           | `gd` | Grenada                                      |
| `ge` | Georgia                              | `gf` | French Guiana            | `gg` | Guernsey                                     |
| `gh` | Ghana                                | `gi` | Gibraltar                | `gl` | Greenland                                    |
| `gm` | Gambia                               | `gn` | Guinea                   | `gp` | Guadeloupe                                   |
| `gq` | Equatorial Guinea                    | `gr` | Greece                   | `gs` | South Georgia and the South Sandwich Islands |
| `gt` | Guatemala                            | `gu` | Guam                     | `gw` | Guinea-Bissau                                |
| `gy` | Guyana                               | `hk` | Hong Kong SAR, China     | `hm` | Heard Island and McDonald Islands            |
| `hn` | Honduras                             | `hr` | Croatia                  | `ht` | Haiti                                        |
| `hu` | Hungary                              | `id` | Indonesia                | `ie` | Ireland                                      |
| `il` | Israel                               | `im` | Isle of Man              | `in` | India                                        |
| `io` | British Indian Ocean Territory       | `iq` | Iraq                     | `ir` | Iran                                         |
| `is` | Iceland                              | `it` | Italy                    | `je` | Jersey                                       |
| `jm` | Jamaica                              | `jo` | Jordan                   | `jp` | Japan                                        |
| `ke` | Kenya                                | `kg` | Kyrgyzstan               | `kh` | Cambodia                                     |
| `ki` | Kiribati                             | `km` | Comoros                  | `kn` | Saint Kitts and Nevis                        |
| `kp` | North Korea                          | `kr` | South Korea              | `kw` | Kuwait                                       |
| `ky` | Cayman Islands                       | `kz` | Kazakhstan               | `la` | Laos                                         |
| `lb` | Lebanon                              | `lc` | Saint Lucia              | `li` | Liechtenstein                                |
| `lk` | Sri Lanka                            | `lr` | Liberia                  | `ls` | Lesotho                                      |
| `lt` | Lithuania                            | `lu` | Luxembourg               | `lv` | Latvia                                       |
| `ly` | Libya                                | `ma` | Morocco                  | `mc` | Monaco                                       |
| `md` | Moldova                              | `me` | Montenegro               | `mf` | Saint Martin (French part)                   |
| `mg` | Madagascar                           | `mh` | Marshall Islands         | `mk` | North Macedonia                              |
| `ml` | Mali                                 | `mm` | Myanmar                  | `mn` | Mongolia                                     |
| `mo` | Macao SAR, China                     | `mp` | Northern Mariana Islands | `mq` | Martinique                                   |
| `mr` | Mauritania                           | `ms` | Montserrat               | `mt` | Malta                                        |
| `mu` | Mauritius                            | `mv` | Maldives                 | `mw` | Malawi                                       |
| `mx` | Mexico                               | `my` | Malaysia                 | `mz` | Mozambique                                   |
| `na` | Namibia                              | `nc` | New Caledonia            | `ne` | Niger                                        |
| `nf` | Norfolk Island                       | `ng` | Nigeria                  | `ni` | Nicaragua                                    |
| `nl` | Netherlands                          | `no` | Norway                   | `np` | Nepal                                        |
| `nr` | Nauru                                | `nu` | Niue                     | `nz` | New Zealand                                  |
| `om` | Oman                                 | `pa` | Panama                   | `pe` | Peru                                         |
| `pf` | French Polynesia                     | `pg` | Papua New Guinea         | `ph` | Philippines                                  |
| `pk` | Pakistan                             | `pl` | Poland                   | `pm` | Saint Pierre and Miquelon                    |
| `pn` | Pitcairn Islands                     | `pr` | Puerto Rico              | `ps` | Palestinian Territories                      |
| `pt` | Portugal                             | `pw` | Palau                    | `py` | Paraguay                                     |
| `qa` | Qatar                                | `re` | Réunion                  | `ro` | Romania                                      |
| `rs` | Serbia                               | `ru` | Russia                   | `rw` | Rwanda                                       |
| `sa` | Saudi Arabia                         | `sb` | Solomon Islands          | `sc` | Seychelles                                   |
| `sd` | Sudan                                | `se` | Sweden                   | `sg` | Singapore                                    |
| `sh` | Saint Helena                         | `si` | Slovenia                 | `sj` | Svalbard and Jan Mayen                       |
| `sk` | Slovakia                             | `sl` | Sierra Leone             | `sm` | San Marino                                   |
| `sn` | Senegal                              | `so` | Somalia                  | `sr` | Suriname                                     |
| `ss` | South Sudan                          | `st` | São Tomé and Príncipe    | `sv` | El Salvador                                  |
| `sx` | Sint Maarten (Dutch part)            | `sy` | Syria                    | `sz` | Swaziland                                    |
| `tc` | Turks and Caicos Islands             | `td` | Chad                     | `tf` | French Southern Territories                  |
| `tg` | Togo                                 | `th` | Thailand                 | `tj` | Tajikistan                                   |
| `tk` | Tokelau                              | `tl` | East Timor               | `tm` | Turkmenistan                                 |
| `tn` | Tunisia                              | `to` | Tonga                    | `tr` | Turkey                                       |
| `tt` | Trinidad and Tobago                  | `tv` | Tuvalu                   | `tw` | Taiwan                                       |
| `tz` | Tanzania                             | `ua` | Ukraine                  | `ug` | Uganda                                       |
| `um` | United States Minor Outlying Islands | `us` | United States            | `uy` | Uruguay                                      |
| `uz` | Uzbekistan                           | `va` | Vatican City             | `vc` | Saint Vincent and the Grenadines             |
| `ve` | Venezuela                            | `vg` | British Virgin Islands   | `vi` | U.S. Virgin Islands                          |
| `vn` | Vietnam                              | `vu` | Vanuatu                  | `wf` | Wallis and Futuna                            |
| `ws` | Samoa                                | `xk` | Kosovo                   | `ye` | Yemen                                        |
| `yt` | Mayotte                              | `za` | South Africa             | `zm` | Zambia                                       |
| `zw` | Zimbabwe                             |      |                          |      |                                              |

</details>

### GeoSite Data Collections

For details on category combinations, migration guidance, and data sources, see the [GeoSite documentation](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README.md).

#### GeoSite Standard

The standard edition (`geosite.dat`) consists of the following tags and their corresponding datasets:

- Original data from `v2fly/domain-list-community`, imported without modification.
- `available-only-in-ru`: Domains available exclusively within Russia, sourced from `runetfreedom/russia-domains-list`
- `blocked-in-cn`: Blocklist for mainland China
  - `gfwlist-cn`: Data sourced from `gfwlist/gfwlist`
- `blocked-in-ru`: Blocklist for Russia
  - `antifilter-list-community`: Domain list from the community edition of Antifilter.Download
  - `refilter`: Complete domain list from `1andrevich/Re-filter-lists`
- `mainly-ir`: A validated set of domains inside Iran, sourced from `Chocolate4U/Iran-v2ray-rules`

#### GeoSite Extended

The extended edition (`geosite-expand.dat`) consists of the following tags and their corresponding datasets:

- Original data from `v2fly/domain-list-community`, imported without modification.
- `available-only-in-ru`: Domains available exclusively within Russia, sourced from `runetfreedom/russia-domains-list`
- `blocked-in-cn-expand`: Expanded blocklist for mainland China
  - `blocked-in-cn`: Blocklist for mainland China
    - `gfwlist-cn`: Data sourced from `gfwlist/gfwlist`
  - `greatfire-list`: Data sourced from GreatFire.org
- `blocked-in-ru-expand`: Expanded blocklist for Russia
  - `antifilter-list-full`: Domain list from Antifilter.Download
  - `blocked-in-ru`: Blocklist for Russia
    - `antifilter-list-community`: Domain list from the community edition of Antifilter.Download
    - `refilter`: Complete domain list from `1andrevich/Re-filter-lists`
- `expand-cn`: Mainland China domain data from `felixonmars/dnsmasq-china-list`
- `expand-ir`: A set of domains inside Iran, sourced from `bootmortis/iran-hosted-domains`
- `mainly-ir`: A validated set of domains inside Iran, sourced from `Chocolate4U/Iran-v2ray-rules`

## Download instructions

### GitHub Releases

GitHub Releases only provides single-file DAT databases. For manual downloads, open the [latest release](https://github.com/XTLS/Xray-rules-dat/releases/latest). Older versions are available on [GitHub Releases](https://github.com/XTLS/Xray-rules-dat/releases).

Automatic updaters can download the latest version from the following fixed URL. Replace `<filename>` with the appropriate file name:

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### The `release` branch

The latest versions of all files are published on the [`release` branch](https://github.com/XTLS/Xray-rules-dat/tree/release). When you need a specific category, download its file directly from this branch.

Replace `<path>` with the file path within the branch, such as `geoip/dat/us.dat`:

- GitHub Raw: `https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>`
- jsDelivr mirror (default): `https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- jsDelivr mirror (Fastly): `https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

jsDelivr mirrors have a [per-file size limit](https://github.com/jsdelivr/jsdelivr#restrictions), so even individual category files may be too large to download. Prefer GitHub Releases or GitHub Raw, and use mirrors when GitHub is inaccessible.

Each data file has a corresponding `.sha256sum` checksum file in the same location.

## Xray usage

Xray releases include GeoIP lite as `geoip.dat`. Its regional custom sets and common categories can be used directly.

The following example keeps the lite version as `geoip.dat` and downloads [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) into the same directory.

Merge the following snippet into your Xray configuration to automatically update and hot-reload both files daily at 04:00 in the runtime environment’s local time. Private addresses and the `in-cn` category use the existing `direct` outbound, while the `us` category uses the existing `wireguard` outbound:

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

The files to be updated must already exist. See the [Geodata configuration documentation](https://xtls.github.io/en/config/geodata.html) for more options.

Replace the category names in the example as needed. See Xray’s [routing documentation](https://xtls.github.io/en/config/routing.html) for rule order and IP matching behavior.

The existing `in-cn` category remains available through `geoip:in-cn`, and the additional `us` category is loaded through `ext:us.dat:us`. For other categories, replace `us` in both the file name and category name.

The category name after the last colon must exist in that DAT file. See [loading IPs from a file](https://xtls.github.io/en/config/routing.html#ruleobject) in the Xray documentation.

To use original categories for more countries and territories, you can also download `geoip-full.dat`, rename it to `geoip.dat`, and replace the existing file. Also change `geoip-lite.dat` to `geoip-full.dat` in the update URL above. After replacing the file, you can use original categories directly, such as `geoip:us`, without downloading `us.dat` separately or using `ext:us.dat:us`.

## Source code and license

GeoIP source code, data sources, and build steps are maintained on the [`geoip` branch](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README.md).

GeoSite source code, data sources, and build tools are maintained in the [`geosite` branch](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README.md).

This project is licensed under [GPL-3.0](./LICENSE). Certain tools (GeoSite) are licensed under the [MIT](./LICENSE-GeoSite) License or are subject to [inherited license terms](./LICENSE-GeoSite-inherit).

## Acknowledgments

- The GeoIP source code was forked from [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb), which was forked from [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- GeoSite source code was forked from and ported from [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618)
- [IPInfo Lite](https://ipinfo.io/lite)
- [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community): upstream GeoSite data
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
