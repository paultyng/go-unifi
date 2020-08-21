package unifi

import (
	"context"
)

func (c *Client) ListPortConf(ctx context.Context, site string) ([]PortConf, error) {
	return c.listPortConf(ctx, site)
}

func (c *Client) GetPortConf(ctx context.Context, site, id string) (*PortConf, error) {
	return c.getPortConf(ctx, site, id)
}

func (c *Client) DeletePortConf(ctx context.Context, site, id string) error {
	return c.deletePortConf(ctx, site, id)
}

func (c *Client) CreatePortConf(ctx context.Context, site string, d *PortConf) (*PortConf, error) {
	return c.createPortConf(ctx, site, d)
}

func (c *Client) UpdatePortConf(ctx context.Context, site string, d *PortConf) (*PortConf, error) {
	return c.updatePortConf(ctx, site, d)
}
