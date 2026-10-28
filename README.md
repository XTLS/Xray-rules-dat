# Xray-rules-dat

English | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

---

GeoIP and GeoSite data for [Xray](https://github.com/XTLS/Xray-core).

GeoIP is based on [IPinfo Lite](https://ipinfo.io/lite), supplemented with IP data for regional routing needs and commonly used services. It supports both IPv4 and IPv6.

GeoSite is based on [v2fly/domain-list-community](https://github.com/v2fly/domain-list-community), supplemented with blacklist and whitelist datasets for essential access needs in various regions.

Data is built daily at a scheduled time of 18:00 UTC (02:00 the following day in Beijing, 21:00 in Moscow, and 21:30 in Tehran).\
_Due to GitHub Actions workflow queues and scheduling, builds may finish a few minutes to several hours later than scheduled._

## Quick downloads

| Data             | File                 | Contents                                                                                                                                   | GitHub Releases                                                                                | jsDelivr mirrors                                                                                                                                                                    |
| ---------------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GeoIP Lite       | `geoip-lite.dat`     | Blacklists and whitelists for mainland China, Russia, and Iran, plus IP data for private networks, commonly used services, and more        | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-lite.dat)     | [Default](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geoip/geoip-lite.dat) |
| GeoIP Full       | `geoip-full.dat`     | All routing data from the Lite edition, plus global IP geolocation data and Anycast IP addresses                                           | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geoip-full.dat)     | File too large                                                                                                                                                                      |
| GeoSite Standard | `geosite.dat`        | Upstream domain categories, plus domain blacklists and whitelists for mainland China, Russia, and Iran                                     | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite.dat)        | [Default](https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat) / [Fastly](https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/geosite/geosite.dat)   |
| GeoSite Expanded | `geosite-expand.dat` | All categories from the Standard edition, plus domain blacklists and whitelists with broader coverage for mainland China, Russia, and Iran | [Download](https://github.com/XTLS/Xray-rules-dat/releases/latest/download/geosite-expand.dat) | File too large                                                                                                                                                                      |

\* The Lite edition is suitable for basic direct/proxy routing in mainland China, Russia, and Iran.\
\* The Full edition supports more granular routing by destination country, as well as direct/proxy routing needs in other countries and regions.\
\* Xray releases include GeoIP Lite and GeoSite Standard by default, named `geoip.dat` and `geosite.dat`, respectively.

### GeoIP Lite vs. Full

| Type                                        | Category names                                                                                                 | Lite         | Full          |
| ------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------ | ------------- |
| IP geolocation                              | Two-letter country codes (see the [table below](#all-countries-and-territories))                               | Not included | Included      |
| Anycast addresses                           | `anycast`                                                                                                      | Not included | Included      |
| Region-specific sets (whitelists)           | `in-cn`, `in-ru`, `in-ir`                                                                                      | Included     | Included      |
| Blacklists                                  | `blacklist-ru`, `blacklist-ir`                                                                                 | Included     | Included      |
| Private networks and commonly used services | `private`, `cloudflare`, `cloudfront`, `facebook`, `fastly`, `google`, `netflix`, `telegram`, `twitter`, `tor` | Included     | Included      |
| Compatibility aliases                       | `cn` → `in-cn`, `ru` → `in-ru`, `ir` → `in-ir`                                                                 | Available    | Not available |

\* Compatibility aliases preserve established community usage and are available only in the Lite edition.\
\* `in-cn` contains IP addresses suitable for direct connections from mainland China; it does not imply that these IP addresses are located in mainland China. In the Full edition, `cn` is strictly based on IP geolocation; the same applies to `ru` and `ir`.

For the category design and data sources, see the [GeoIP documentation](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README.md).

#### Per-category files

If you need only a few categories, download their individual files instead of the larger `geoip-full.dat`.

For example, you can keep using the Lite edition as `geoip.dat` and continue to access its `in-cn` category through `geoip:in-cn`. Download the U.S. IP category file [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat) into the same directory, then access its `us` category through `ext:us.dat:us`.

Per-category DAT files are available in the [`/geoip/dat/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geoip/dat) directory on the `release` branch.

#### All countries and territories

The Full edition includes the following 250 country and territory categories.

<details>
<summary>Expand the full list (sorted by code)</summary>

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
| `sx` | Sint Maarten (Dutch part)            | `sy` | Syria                    | `sz` | Eswatini                                     |
| `tc` | Turks and Caicos Islands             | `td` | Chad                     | `tf` | French Southern Territories                  |
| `tg` | Togo                                 | `th` | Thailand                 | `tj` | Tajikistan                                   |
| `tk` | Tokelau                              | `tl` | Timor-Leste              | `tm` | Turkmenistan                                 |
| `tn` | Tunisia                              | `to` | Tonga                    | `tr` | Türkiye                                      |
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

### GeoSite Standard vs. Expanded

| Type                                                                    | Category names         | Standard     | Expanded |
| ----------------------------------------------------------------------- | ---------------------- | ------------ | -------- |
| [domain-list-community](https://github.com/v2fly/domain-list-community) | Same as upstream       | Included     | Included |
| Mainland China blacklist                                                | `blocked-in-cn`        | Included     | Included |
| Mainland China whitelist                                                | `located-in-cn`        | Included     | Included |
| Expanded mainland China blacklist                                       | `blocked-in-cn-expand` | Not included | Included |
| Expanded mainland China whitelist                                       | `located-in-cn-expand` | Not included | Included |
| Russia blacklist                                                        | `blocked-in-ru`        | Included     | Included |
| Russia whitelist                                                        | `located-in-ru`        | Included     | Included |
| Expanded Russia blacklist                                               | `blocked-in-ru-expand` | Not included | Included |
| Iran whitelist                                                          | `located-in-ir`        | Included     | Included |
| Expanded Iran whitelist                                                 | `located-in-ir-expand` | Not included | Included |

For category combinations, migration guidance, and data sources, see the [GeoSite documentation](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README.md).

#### Expanded delta files

If you need expanded data for only a few regions, download the corresponding delta files alongside GeoSite Standard (`geosite.dat`) instead of downloading `geosite-expand.dat`.

For example, place [delta-ru.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geosite/delta/delta-ru.dat) in the same directory as GeoSite Standard (`geosite.dat`). Upstream domain-list-community categories remain available through `geosite:xx`, and the expanded Russia blacklist through `ext:delta-ru.dat:blocked-in-ru-expand`.

Delta DAT files are available in the [`/geosite/delta/`](https://github.com/XTLS/Xray-rules-dat/tree/release/geosite/delta) directory on the `release` branch.

## Download instructions

### GitHub Releases

GitHub Releases provides only complete DAT packages. Both current and historical versions are available on the [releases page](https://github.com/XTLS/Xray-rules-dat/releases).

Use the following fixed URL for automatic updates. Replace `<filename>` with the desired filename to get the latest version:

```text
https://github.com/XTLS/Xray-rules-dat/releases/latest/download/<filename>
```

### The `release` branch

The `release` branch also provides the latest complete packages, along with per-category and delta files, but does not retain historical versions.

[Browse the branch](https://github.com/XTLS/Xray-rules-dat/tree/release) to find the file you need, then replace `<path>` with its path (for example, `geoip/dat/us.dat`). The following URL can be used for automatic updates:

```text
https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/<path>
```

<details>
<summary>If GitHub is inaccessible, use a jsDelivr mirror</summary>

- Default: `https://cdn.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`
- Fastly: `https://fastly.jsdelivr.net/gh/XTLS/Xray-rules-dat@release/<path>`

\* jsDelivr has a [per-file size limit](https://github.com/jsdelivr/jsdelivr#restrictions), so even per-category or delta files may be too large to download. Downloading directly from GitHub is recommended.

</details>

Each data file has a corresponding `.sha256sum` checksum file in the same location.

## Xray usage

Xray releases include GeoIP Lite and GeoSite Standard by default, already named `geoip.dat` and `geosite.dat`. You can use them directly with the `geoip:xx` and `geosite:xx` syntax.

The following example also uses the U.S. IP file [us.dat](https://raw.githubusercontent.com/XTLS/Xray-rules-dat/release/geoip/dat/us.dat). Place all three files in the same directory.

Merge the following snippet into your existing Xray configuration. The example uses existing `direct`, `xhttp-us`, and `xhttp-jp` outbounds, and automatically updates and hot-reloads all three files daily at 04:00 in the runtime environment's local time:

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
    "domainStrategy": "IPOnDemand", // Resolve domains on demand so domain traffic can also match IP rules
    "rules": [
      {
        // Route private addresses directly first
        "ip": ["geoip:private"],
        "outboundTag": "direct"
      },
      {
        // Intersect the domain and IP whitelists; route directly only if both match, reducing misclassification from either list alone
        "domain": ["geosite:cn"],
        "ip": ["geoip:in-cn"],
        "outboundTag": "direct"
      },
      {
        // Match U.S. IP addresses using the us category in us.dat and use the xhttp-us outbound
        "ip": ["ext:us.dat:us"],
        "outboundTag": "xhttp-us"
      },
      {
        // Use the xhttp-jp outbound for domains on the mainland China blacklist
        "domain": ["geosite:blocked-in-cn"],
        "outboundTag": "xhttp-jp"
      }
    ]
  }
}
```

Files must already exist before they can be updated. For update options, see the [Geodata configuration documentation](https://xtls.github.io/en/config/geodata.html).

This is only a configuration snippet; adapt it to your existing configuration and routing needs. For rule order and IP/domain matching behavior, see the [routing documentation](https://xtls.github.io/en/config/routing.html).

The category name after the last colon must exist in the corresponding DAT file. See [Loading IP addresses / domains from files](https://xtls.github.io/en/config/routing.html#ruleobject).

If you need IP data for more countries, download `geoip-full.dat`, rename it to `geoip.dat`, replace the original file, and change `geoip-lite.dat` to `geoip-full.dat` in the update URL. You can then use references such as `geoip:us` directly, without downloading `us.dat` separately or using `ext:us.dat:us`.

## Source code and licenses

GeoIP source code, data sources, and build instructions are maintained on the [`geoip` branch](https://github.com/XTLS/Xray-rules-dat/blob/geoip/README.md).

GeoSite source code, data sources, and build instructions are maintained on the [`geosite` branch](https://github.com/XTLS/Xray-rules-dat/blob/geosite/README.md).

The project as a whole is licensed under [GPL-3.0](./LICENSE). Some tools (GeoSite) are separately covered by the [MIT license](./LICENSE-GeoSite) or [inherited license terms](./LICENSE-GeoSite-inherit).

## Acknowledgments

- GeoIP source code was forked from [`Meo597/geoip@56a84321`](https://github.com/Meo597/geoip/commit/56a8432119d6f03d8689b36ac3ba9a8ed880f9cb), which was forked from [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
- GeoSite source code was forked and ported from [`v2fly/domain-list-community@e449533`](https://github.com/v2fly/domain-list-community/commit/e449533a4ab2c00716e351f13b592153b39a5618)
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
