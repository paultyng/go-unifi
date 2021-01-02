package unifi_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/paultyng/go-unifi/unifi"
)

func TestNetworkUnmarshalJSON(t *testing.T) {
	for n, c := range map[string]struct {
		expected unifi.Network
		json     string
	}{
		"int vlan": {
			expected: unifi.Network{VLAN: 1},
			json:     `{ "vlan": 1 }`,
		},
		"string vlan": {
			expected: unifi.Network{VLAN: 1},
			json:     `{ "vlan": "1" }`,
		},
		"empty string vlan": {
			expected: unifi.Network{VLAN: 0},
			json:     `{ "vlan": "" }`,
		},

		"int dhcpd_leasetime": {
			expected: unifi.Network{DHCPDLeaseTime: 1},
			json:     `{ "dhcpd_leasetime": 1 }`,
		},
		"string dhcpd_leasetime": {
			expected: unifi.Network{DHCPDLeaseTime: 1},
			json:     `{ "dhcpd_leasetime": "1" }`,
		},
		"empty string dhcpd_leasetime": {
			expected: unifi.Network{DHCPDLeaseTime: 0},
			json:     `{ "dhcpd_leasetime": "" }`,
		},

		"int wan_egress_qos": {
			expected: unifi.Network{WANEgressQOS: 1},
			json:     `{ "wan_egress_qos": 1 }`,
		},
		"string wan_egress_qos": {
			expected: unifi.Network{WANEgressQOS: 1},
			json:     `{ "wan_egress_qos": "1" }`,
		},
		"empty string wan_egress_qos": {
			expected: unifi.Network{WANEgressQOS: 0},
			json:     `{ "wan_egress_qos": "" }`,
		},

		"int wan_vlan": {
			expected: unifi.Network{WANVLAN: 1},
			json:     `{ "wan_vlan": 1 }`,
		},
		"string wan_vlan": {
			expected: unifi.Network{WANVLAN: 1},
			json:     `{ "wan_vlan": "1" }`,
		},
		"empty wan_vlan vlan": {
			expected: unifi.Network{WANVLAN: 0},
			json:     `{ "wan_vlan": "" }`,
		},
	} {
		t.Run(n, func(t *testing.T) {
			var actual unifi.Network
			err := json.Unmarshal(([]byte)(c.json), &actual)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("not equal:\nexpected: %#v\nactual: %#v", c.expected, actual)
			}
		})
	}
}
