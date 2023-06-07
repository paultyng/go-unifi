package unifi_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/paultyng/go-unifi/unifi"
)

func TestNetworkUnmarshalJSON(t *testing.T) {
	for n, c := range map[string]struct {
		expected func(n *unifi.Network)
		json     string
	}{
		"int vlan": {
			expected: func(n *unifi.Network) { n.VLAN = 1 },
			json:     `{ "vlan": 1 }`,
		},
		"string vlan": {
			expected: func(n *unifi.Network) { n.VLAN = 1 },
			json:     `{ "vlan": "1" }`,
		},
		"empty string vlan": {
			expected: func(n *unifi.Network) { n.VLAN = 0 },
			json:     `{ "vlan": "" }`,
		},

		"int dhcpd_leasetime": {
			expected: func(n *unifi.Network) { n.DHCPDLeaseTime = 1 },
			json:     `{ "dhcpd_leasetime": 1 }`,
		},
		"string dhcpd_leasetime": {
			expected: func(n *unifi.Network) { n.DHCPDLeaseTime = 1 },
			json:     `{ "dhcpd_leasetime": "1" }`,
		},
		"empty string dhcpd_leasetime": {
			expected: func(n *unifi.Network) { n.DHCPDLeaseTime = 0 },
			json:     `{ "dhcpd_leasetime": "" }`,
		},

		"int wan_egress_qos": {
			expected: func(n *unifi.Network) { n.WANEgressQOS = 1 },
			json:     `{ "wan_egress_qos": 1 }`,
		},
		"string wan_egress_qos": {
			expected: func(n *unifi.Network) { n.WANEgressQOS = 1 },
			json:     `{ "wan_egress_qos": "1" }`,
		},
		"empty string wan_egress_qos": {
			expected: func(n *unifi.Network) { n.WANEgressQOS = 0 },
			json:     `{ "wan_egress_qos": "" }`,
		},

		"int wan_vlan": {
			expected: func(n *unifi.Network) { n.WANVLAN = 1 },
			json:     `{ "wan_vlan": 1 }`,
		},
		"string wan_vlan": {
			expected: func(n *unifi.Network) { n.WANVLAN = 1 },
			json:     `{ "wan_vlan": "1" }`,
		},
		"empty wan_vlan vlan": {
			expected: func(n *unifi.Network) { n.WANVLAN = 0 },
			json:     `{ "wan_vlan": "" }`,
		},
	} {
		t.Run(n, func(t *testing.T) {
			// set some non-zero value defaults
			expected := unifi.Network{
				InternetAccessEnabled: true,
			}
			c.expected(&expected)
			var actual unifi.Network
			err := json.Unmarshal(([]byte)(c.json), &actual)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("not equal:\nexpected: %#v\nactual: %#v", expected, actual)
			}
		})
	}
}
