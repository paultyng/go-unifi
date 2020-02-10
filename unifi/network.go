package unifi

import (
	"encoding/json"
	"fmt"
)

func (dst *Network) UnmarshalJSON(b []byte) error {
	type Alias Network
	aux := &struct {
		VLAN           json.Number `json:"vlan"`
		DHCPDLeaseTime json.Number `json:"dhcpd_leasetime"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	dst.VLAN = 0
	if aux.VLAN.String() != "" {
		n, err := aux.VLAN.Int64()
		if err != nil {
			return err
		}
		dst.VLAN = int(n)
	}

	dst.DHCPDLeaseTime = 0
	if aux.DHCPDLeaseTime.String() != "" {
		n, err := aux.DHCPDLeaseTime.Int64()
		if err != nil {
			return err
		}
		dst.DHCPDLeaseTime = int(n)
	}

	return nil
}

func (c *Client) DeleteNetwork(site, id, name string) error {
	err := c.do("DELETE", fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), struct {
		Name string `json:"name"`
	}{
		Name: name,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ListNetwork(site string) ([]Network, error) {
	return c.listNetwork(site)
}

func (c *Client) GetNetwork(site, id string) (*Network, error) {
	return c.getNetwork(site, id)
}

func (c *Client) CreateNetwork(site string, d *Network) (*Network, error) {
	return c.createNetwork(site, d)
}

func (c *Client) UpdateNetwork(site string, d *Network) (*Network, error) {
	return c.updateNetwork(site, d)
}
