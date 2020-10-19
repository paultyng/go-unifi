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

type SettingMgmt struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AdvancedFeatureEnabled  bool     `json:"advanced_feature_enabled"`
	AlertEnabled            bool     `json:"alert_enabled"`
	AutoUpgrade             bool     `json:"auto_upgrade"`
	BootSound               bool     `json:"boot_sound"`
	LedEnabled              bool     `json:"led_enabled"`
	OutdoorModeEnabled      bool     `json:"outdoor_mode_enabled"`
	UnifiIDpEnabled         bool     `json:"unifi_idp_enabled"`
	WifimanEnabled          bool     `json:"wifiman_enabled"`
	XMgmtKey                string   `json:"x_mgmt_key,omitempty"` // [0-9a-f]{32}
	XSshAuthPasswordEnabled bool     `json:"x_ssh_auth_password_enabled"`
	XSshBindWildcard        bool     `json:"x_ssh_bind_wildcard"`
	XSshEnabled             bool     `json:"x_ssh_enabled"`
	XSshKeys                []string `json:"x_ssh_keys,omitempty"`
	XSshMd5Passwd           string   `json:"x_ssh_md5passwd,omitempty"`
	XSshPassword            string   `json:"x_ssh_password,omitempty"` // .{1,128}
	XSshSha512Passwd        string   `json:"x_ssh_sha512passwd,omitempty"`
	XSshUsername            string   `json:"x_ssh_username,omitempty"` // ^[_A-Za-z0-9][-_.A-Za-z0-9]{0,29}$
}

func (c *Client) getSettingMgmt(ctx context.Context, site string) (*SettingMgmt, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []SettingMgmt `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/mgmt", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingMgmt(ctx context.Context, site string, d *SettingMgmt) (*SettingMgmt, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []SettingMgmt `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/mgmt", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
