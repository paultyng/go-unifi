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

type SettingRsyslogd struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	Debug                       bool   `json:"debug"`
	Enabled                     bool   `json:"enabled"`
	IP                          string `json:"ip,omitempty"`
	NetconsoleEnabled           bool   `json:"netconsole_enabled"`
	NetconsoleHost              string `json:"netconsole_host,omitempty"`
	NetconsolePort              int    `json:"netconsole_port,omitempty"` // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	Port                        int    `json:"port,omitempty"`            // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]
	ThisController              bool   `json:"this_controller"`
	ThisControllerEncryptedOnly bool   `json:"this_controller_encrypted_only"`
}

func (c *Client) getSettingRsyslogd(ctx context.Context, site string) (*SettingRsyslogd, error) {
	var respBody struct {
		Meta meta              `json:"meta"`
		Data []SettingRsyslogd `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/rsyslogd", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingRsyslogd(ctx context.Context, site string, d *SettingRsyslogd) (*SettingRsyslogd, error) {
	var respBody struct {
		Meta meta              `json:"meta"`
		Data []SettingRsyslogd `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/rsyslogd", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
