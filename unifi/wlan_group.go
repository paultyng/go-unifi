package unifi

import (
	"context"
	"encoding/json"
)

func (n *WLANGroup) UnmarshalJSON(b []byte) error {
	type Alias WLANGroup
	aux := &struct {
		Maxsta emptyStringInt `json:"maxsta"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	n.Maxsta = int(aux.Maxsta)

	return nil
}

func (c *Client) ListWLANGroup(ctx context.Context, site string) ([]WLANGroup, error) {
	return c.listWLANGroup(ctx, site)
}

func (c *Client) GetWLANGroup(ctx context.Context, site, id string) (*WLANGroup, error) {
	return c.getWLANGroup(ctx, site, id)
}

func (c *Client) DeleteWLANGroup(ctx context.Context, site, id string) error {
	return c.deleteWLANGroup(ctx, site, id)
}

func (c *Client) CreateWLANGroup(ctx context.Context, site string, d *WLANGroup) (*WLANGroup, error) {
	return c.createWLANGroup(ctx, site, d)
}

func (c *Client) UpdateWLANGroup(ctx context.Context, site string, d *WLANGroup) (*WLANGroup, error) {
	return c.updateWLANGroup(ctx, site, d)
}
