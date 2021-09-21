package unifi_test

import (
	"encoding/json"
	"testing"

	"github.com/paultyng/go-unifi/unifi"
	"github.com/tj/assert"
)

func TestAccountMarshalJSON(t *testing.T) {
	for n, c := range map[string]struct {
		expectedJSON string
		acc          unifi.Account
	}{
		"empty strings": {
			`{"vlan":"","tunnel_type":"","tunnel_medium_type":"","networkconf_id":""}`,
			unifi.Account{},
		},
		"response": {
			`{"vlan":10,"tunnel_type":1,"tunnel_medium_type":1,"networkconf_id":""}`,
			unifi.Account{
				VLAN:             10,
				TunnelType:       1,
				TunnelMediumType: 1,
			},
		},
	} {
		t.Run(n, func(t *testing.T) {
			actual, err := json.Marshal(&c.acc)
			if err != nil {
				t.Fatal(err)
			}
			assert.JSONEq(t, c.expectedJSON, string(actual))
		})
	}
}
