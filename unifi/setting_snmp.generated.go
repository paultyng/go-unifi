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

type SettingSnmp struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Community string `json:"community,omitempty"` // .{1,256}
	Enabled   bool   `json:"enabled"`
	EnabledV3 bool   `json:"enabledV3"`
	Username  string `json:"username,omitempty"`   // [a-zA-Z0-9_-]{1,30}
	XPassword string `json:"x_password,omitempty"` // [^'"]{8,32}
}

func (c *Client) getSettingSnmp(ctx context.Context, site string) (*SettingSnmp, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []SettingSnmp `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/snmp", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingSnmp(ctx context.Context, site string, d *SettingSnmp) (*SettingSnmp, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []SettingSnmp `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/snmp", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
