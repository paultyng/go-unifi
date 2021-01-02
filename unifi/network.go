package unifi

import (
	"context"
	"fmt"
)

func (c *Client) DeleteNetwork(ctx context.Context, site, id, name string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), struct {
		Name string `json:"name"`
	}{
		Name: name,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ListNetwork(ctx context.Context, site string) ([]Network, error) {
	return c.listNetwork(ctx, site)
}

func (c *Client) GetNetwork(ctx context.Context, site, id string) (*Network, error) {
	return c.getNetwork(ctx, site, id)
}

func (c *Client) CreateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.createNetwork(ctx, site, d)
}

func (c *Client) UpdateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.updateNetwork(ctx, site, d)
}
