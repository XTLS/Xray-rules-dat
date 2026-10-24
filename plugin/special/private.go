package special

import (
	"encoding/json"
	"fmt"

	"github.com/xtls/geoip/lib"
)

const (
	entryNamePrivate = "private"
	TypePrivate      = "private"
	DescPrivate      = "Convert CIDRs unsuitable for use on the public Internet to other formats"
)

// privateCIDRs selects special-purpose and multicast ranges for this project's
// routing policy, including selected globally reachable protocol services.
// Commented entries document deliberate exclusions from the IANA registries.
var privateCIDRs = []string{
	// https://www.iana.org/assignments/iana-ipv4-special-registry
	"0.0.0.0/8",      // "This network"
	"0.0.0.0/32",     // "This host on this network"
	"10.0.0.0/8",     // Private-Use
	"100.64.0.0/10",  // Shared Address Space
	"127.0.0.0/8",    // Loopback
	"169.254.0.0/16", // Link Local
	"172.16.0.0/12",  // Private-Use
	// "192.0.0.0/24",   // IETF Protocol Assignments
	"192.0.0.0/29",   // IPv4 Service Continuity Prefix
	"192.0.0.8/32",   // IPv4 dummy address
	"192.0.0.9/32",   // Port Control Protocol Anycast
	"192.0.0.10/32",  // Traversal Using Relays around NAT Anycast
	"192.0.0.170/32", // NAT64/DNS64 Discovery
	"192.0.0.171/32", // NAT64/DNS64 Discovery
	"192.0.2.0/24",   // Documentation (TEST-NET-1)
	// "192.31.196.0/24",    // AS112-v4
	"192.52.193.0/24", // AMT
	// "192.88.99.0/24",  // Deprecated (6to4 Relay Anycast)
	"192.88.99.2/32", // 6a44-relay anycast address
	"192.168.0.0/16", // Private-Use
	// "192.175.48.0/24",    // Direct Delegation AS112 Service
	"198.18.0.0/15",      // Benchmarking
	"198.51.100.0/24",    // Documentation (TEST-NET-2)
	"203.0.113.0/24",     // Documentation (TEST-NET-3)
	"240.0.0.0/4",        // Reserved
	"255.255.255.255/32", // Limited Broadcast

	// https://www.iana.org/assignments/multicast-addresses
	"224.0.0.0/4",

	// https://www.iana.org/assignments/iana-ipv6-special-registry
	"::1/128", // Loopback Address
	"::/128",  // Unspecified Address
	// "::ffff:0:0/96",     // IPv4-mapped Address
	"64:ff9b::/96",   // IPv4-IPv6 Translat.
	"64:ff9b:1::/48", // IPv4-IPv6 Translat.
	"100::/64",       // Discard-Only Address Block
	"100:0:0:1::/64", // Dummy IPv6 Prefix
	// "2001::/23",      // IETF Protocol Assignments
	"2001::/32",     // TEREDO
	"2001:1::1/128", // Port Control Protocol Anycast
	"2001:1::2/128", // Traversal Using Relays around NAT Anycast
	"2001:1::3/128", // DNS-SD Service Registration Protocol Anycast
	"2001:2::/48",   // Benchmarking
	"2001:3::/32",   // AMT
	// "2001:4:112::/48", // AS112-v6
	// "2001:10::/28",    // Deprecated (previously ORCHID)
	"2001:20::/28",  // ORCHIDv2
	"2001:30::/28",  // Drone Remote ID Protocol Entity Tags (DETs) Prefix
	"2001:db8::/32", // Documentation
	"2002::/16",     // 6to4
	// "2620:4f:8000::/48", // Direct Delegation AS112 Service
	"3fff::/20", // Documentation
	"5f00::/16", // Segment Routing (SRv6) SIDs
	"fc00::/7",  // Unique-Local
	"fe80::/10", // Link-Local Unicast

	// https://www.iana.org/assignments/ipv6-multicast-addresses
	"ff00::/8",
}

func init() {
	lib.RegisterInputConfigCreator(TypePrivate, func(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
		return newPrivate(action, data)
	})
	lib.RegisterInputConverter(TypePrivate, &Private{
		Description: DescPrivate,
	})
}

func newPrivate(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
	var tmp struct {
		OnlyIPType lib.IPType `json:"onlyIPType"`
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, fmt.Errorf("❌ [type %s | action %s] %w", TypePrivate, action, err)
		}
	}

	return &Private{
		Type:        TypePrivate,
		Action:      action,
		Description: DescPrivate,
		OnlyIPType:  tmp.OnlyIPType,
	}, nil
}

type Private struct {
	Type        string
	Action      lib.Action
	Description string
	OnlyIPType  lib.IPType
}

func (p *Private) GetType() string {
	return p.Type
}

func (p *Private) GetAction() lib.Action {
	return p.Action
}

func (p *Private) GetDescription() string {
	return p.Description
}

func (p *Private) Input(container lib.Container) (lib.Container, error) {
	entry := lib.NewEntry(entryNamePrivate)

	for _, cidr := range privateCIDRs {
		if err := entry.AddPrefix(cidr); err != nil {
			return nil, fmt.Errorf("❌ [type %s | action %s] %w", p.Type, p.Action, err)
		}
	}

	ignoreIPType := lib.GetIgnoreIPType(p.OnlyIPType)

	switch p.Action {
	case lib.ActionAdd:
		if err := container.Add(entry, ignoreIPType); err != nil {
			return nil, fmt.Errorf("❌ [type %s | action %s] %w", p.Type, p.Action, err)
		}
	case lib.ActionRemove:
		if err := container.Remove(entry, lib.CaseRemovePrefix, ignoreIPType); err != nil {
			return nil, fmt.Errorf("❌ [type %s | action %s] %w", p.Type, p.Action, err)
		}
	default:
		return nil, fmt.Errorf("❌ [type %s | action %s] %w", p.Type, p.Action, lib.ErrUnknownAction)
	}

	return container, nil
}
