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

type SettingSuperSdn struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	AuthToken         string   `json:"auth_token,omitempty"`
	DeviceID          string   `json:"device_id"`
	Enabled           bool     `json:"enabled"`
	Migrated          bool     `json:"migrated"`
	OauthAppID        string   `json:"oauth_app_id"`
	OauthEnabled      bool     `json:"oauth_enabled"`
	OauthRedirectUris []string `json:"oauth_redirect_uris,omitempty"`
	SsoLoginEnabled   string   `json:"sso_login_enabled,omitempty"`
	UbicUuid          string   `json:"ubic_uuid,omitempty"`
	XOauthAppSecret   string   `json:"x_oauth_app_secret,omitempty"`
}

func (c *Client) getSettingSuperSdn(ctx context.Context, site string) (*SettingSuperSdn, error) {
	var respBody struct {
		Meta meta              `json:"meta"`
		Data []SettingSuperSdn `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/get/setting/super_sdn", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) updateSettingSuperSdn(ctx context.Context, site string, d *SettingSuperSdn) (*SettingSuperSdn, error) {
	var respBody struct {
		Meta meta              `json:"meta"`
		Data []SettingSuperSdn `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/set/setting/super_sdn", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
