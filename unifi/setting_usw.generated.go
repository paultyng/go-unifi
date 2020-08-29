// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ fmt.Formatter
	_ context.Context
)

type SettingUsw struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DHCPSnoop bool `json:"dhcp_snoop"`
}

func (c *Client) getSettingUsw(ctx context.Context, site string) (*SettingUsw, error) {
	var respBody struct {
		Meta meta         `json:"meta"`
		Data []SettingUsw `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/usw", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingUsw(ctx context.Context, site string, d *SettingUsw) (*SettingUsw, error) {
	var respBody struct {
		Meta meta         `json:"meta"`
		Data []SettingUsw `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/usw", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
