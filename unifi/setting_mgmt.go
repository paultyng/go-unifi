package unifi

import (
	"context"
)

func (c *Client) GetSettingMgmt(ctx context.Context, site string) (*SettingMgmt, error) {
	return c.getSettingMgmt(ctx, site)
}

func (c *Client) UpdateSettingMgmt(ctx context.Context, site string, d *SettingMgmt) (*SettingMgmt, error) {
	return c.updateSettingMgmt(ctx, site, d)
}
