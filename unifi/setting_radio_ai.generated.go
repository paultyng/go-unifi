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

type SettingRadioAi struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ChannelsNa     []int    `json:"channels_na,omitempty"` // 36|40|44|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165
	ChannelsNg     []int    `json:"channels_ng,omitempty"` // 1|2|3|4|5|6|7|8|9|10|11|12|13|14
	CronExpr       string   `json:"cron_expr,omitempty"`
	Default        bool     `json:"default"`
	Enabled        bool     `json:"enabled"`
	ExcludeDevices []string `json:"exclude_devices,omitempty"` // ([0-9a-z]{2}:){5}[0-9a-z]{2}
	HtModesNa      []int    `json:"ht_modes_na,omitempty"`     // ^(20|40|80|160)$
	HtModesNg      []int    `json:"ht_modes_ng,omitempty"`     // ^(20|40)$
	Optimize       []string `json:"optimize,omitempty"`        // channel|power
	Radios         []string `json:"radios,omitempty"`          // na|ng
	UseXY          bool     `json:"useXY"`
}

func (c *Client) getSettingRadioAi(ctx context.Context, site string) (*SettingRadioAi, error) {
	var respBody struct {
		Meta meta             `json:"meta"`
		Data []SettingRadioAi `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/radio_ai", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingRadioAi(ctx context.Context, site string, d *SettingRadioAi) (*SettingRadioAi, error) {
	var respBody struct {
		Meta meta             `json:"meta"`
		Data []SettingRadioAi `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/radio_ai", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
