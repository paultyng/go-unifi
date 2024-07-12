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
			`{"vlan":"","tunnel_type":"","tunnel_medium_type":""}`,
			unifi.Account{},
		},
		"response": {
			`{"vlan":10,"tunnel_type":1,"tunnel_medium_type":1}`,
			unifi.Account{
				VLAN:             intPtr(10),
				TunnelType:       intPtr(1),
				TunnelMediumType: intPtr(1),
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

func intPtr(i int) *int {
	return &i
}
