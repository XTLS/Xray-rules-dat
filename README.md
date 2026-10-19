# GeoIP

English | [简体中文](./README_zh-Hans.md) | [Русский](./README_ru-RU.md) | [فارسی](./README_fa-IR.md)

This repository builds GeoIP data from IPInfo Lite. It provides original country and territory classifications, regional sets tailored to different routing environments, and common categories for applications such as Xray, sing-box, and Clash.

## Design

The coverage of GeoIP data directly affects routing decisions. For international routing across multiple egress nodes, MaxMind's address coverage did not meet the project's requirements. IPInfo Lite is therefore used as the baseline for original country and territory classifications, with anycast addresses listed separately to avoid detours caused by country-based routing.

The country associated with a destination IP and the routing needs of the user's environment are two different questions. Country classifications answer “Which country or territory does this destination IP belong to?” In routing environments such as mainland China, Iran, and Russia, users also need to answer “While in this country or territory, which destination addresses should follow a particular routing policy?” The destination's country classification alone cannot fully express these requirements.

This project therefore introduces regional `in-*` sets as direct-connection allowlists for their respective regions. Sets such as `in-cn`, `in-ir`, and `in-ru` are tailored to local access conditions and routing needs. For example, addresses for overseas services, Telegram, Twitter, or CDNs may be included when they can be accessed directly and efficiently from that region. Private addresses may also be included for specific requirements.

Original country classifications have their own use cases. For example, when configuring egress nodes in multiple countries or territories, users may want to select a node based on the destination IP's country. Use ordinary country codes for this purpose; use `in-*` sets for the routing needs of your current environment.

The full `geoip-full.dat` therefore keeps both types of data: ordinary country codes use IPInfo Lite's original classifications, while `in-*` entries contain the regional sets. Regional customizations do not overwrite the original country classifications.

The lite `geoip-lite.dat` does not include complete country classifications and is not suitable for choosing egress nodes by destination IP country. The ordinary country codes it retains are only aliases for the regional sets, provided for compatibility with established community usage.

Both versions include IPv4 and IPv6 data.

## Regional sets

### Mainland China: `in-cn`

`in-cn` is a direct-connection allowlist for routing from mainland China.

IPv4 data combines the following three sources, without adding IPInfo Lite's original CN IPv4 data:

- zhufengme/block_cn_files
- misakaio/chnroutes2
- china.txt from gaoyifan/china-operator-ip

IPv6 data combines IPInfo Lite's original CN IPv6 data with china6.txt from gaoyifan/china-operator-ip.

In theory, announced prefixes are sufficient for routing needs. More specific prefixes that have not yet been announced are also included so that routing rules can match them as soon as they are announced, without waiting for a data update.

Because ISPs in mainland China degrade international connectivity, overseas addresses effectively constitute a blocklist. This project therefore maintains only a direct-connection allowlist for mainland China, without a separate blocklist.

### Iran: `in-ir`

Starting with IPInfo Lite's original IR IPv4 and IPv6 data, this set adds Eitaa (AS202798) and GAP (AS210877) from the local ASN database, CDN address lists from ArvanCloud, DerakCloud, IranServer, and ParsPack, and six private addresses used by specific local services in Iran. ParsPack data comes from the plain-text mirror provided by Chocolate4U/Iran-v2ray-rules. A separate `blacklist-ir` contains addresses injected for blocking and other injected addresses.

### Russia: `in-ru`

Starting with IPInfo Lite's original RU IPv4 and IPv6 data, this set adds the mobile internet allowlist from hxehex/russia-mobile-internet-whitelist. A separate `blacklist-ru` combines IP lists from Antifilter and 1andrevich/Re-filter-lists with Telegram prefixes taken from the reference configuration.

## Common categories

Both the full and lite versions include the following common categories:

- `geoip:private`
- `geoip:anycast`
- `geoip:cloudflare`
- `geoip:cloudfront`
- `geoip:facebook`
- `geoip:fastly`
- `geoip:google`
- `geoip:netflix`
- `geoip:telegram`
- `geoip:twitter`
- `geoip:tor`

## Build

Place `./ipinfo/ipinfo_lite.mmdb` under the repository root and ensure that `./tmp/geoip-in` and `./output` are clean directories. Then run:

```sh
go run . convert -c config-geoip-in.json
go run . convert -c config-geoip-lite.json
go run . convert -c config-geoip-full.json
```

The first step generates temporary text files for the regional sets and blocklists in `./tmp/geoip-in/`. Both lite and full read these files directly. After the first step, lite and full can run independently or in either order. Do not clean the directories between these three steps.

The lite build outputs only `./output/geoip-lite.dat`. The full build outputs `./output/geoip-full.dat` and generates the following files for each category:

| Output directory         | Format                  |
| ------------------------ | ----------------------- |
| `output/dat/`            | Xray GeoIP DAT          |
| `output/srs/`            | sing-box SRS            |
| `output/mrs/`            | Mihomo MRS              |
| `output/text/`           | Plain-text CIDR         |
| `output/clash/classical/` | Clash classical RuleSet |
| `output/clash/ipcidr/`    | Clash ipcidr RuleSet    |

Build commands must explicitly specify a configuration file with `-c`. See the [configuration guide (Chinese)](./configuration.md) for more options.

## License

- [GPL-3.0](./LICENSE)

## Acknowledgements

- This project was forked from [`Loyalsoldier/geoip@b84c0db4`](https://github.com/Loyalsoldier/geoip/commit/b84c0db4ad7286eb5cec870cabf6642a156290bc)
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
