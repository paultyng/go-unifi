package unifi

import "encoding/json"

func (n *WLANGroup) UnmarshalJSON(b []byte) error {
	type Alias WLANGroup
	aux := &struct {
		Maxsta json.Number `json:"maxsta"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	n.Maxsta = 0
	if aux.Maxsta.String() != "" {
		maxsta, err := aux.Maxsta.Int64()
		if err != nil {
			return err
		}
		n.Maxsta = int(maxsta)
	}
	return nil
}

func (c *Client) ListWLANGroup(site string) ([]WLANGroup, error) {
	return c.listWLANGroup(site)
}

func (c *Client) GetWLANGroup(site, id string) (*WLANGroup, error) {
	return c.getWLANGroup(site, id)
}

func (c *Client) DeleteWLANGroup(site, id string) error {
	return c.deleteWLANGroup(site, id)
}

func (c *Client) CreateWLANGroup(site string, d *WLANGroup) (*WLANGroup, error) {
	return c.createWLANGroup(site, d)
}

func (c *Client) UpdateWLANGroup(site string, d *WLANGroup) (*WLANGroup, error) {
	return c.updateWLANGroup(site, d)
}
