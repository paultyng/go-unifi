package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

func (dst *Account) UnmarshalJSON(b []byte) error {
	type Alias Account
	aux := &struct {
		TunnelType       emptyStringInt `json:"tunnel_type"`
		TunnelMediumType emptyStringInt `json:"tunnel_medium_type"`
		VLAN             emptyStringInt `json:"vlan"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	dst.TunnelType = aux.TunnelType.val
	dst.TunnelMediumType = aux.TunnelMediumType.val
	dst.VLAN = aux.VLAN.val

	return nil
}

func (dst *Account) MarshalJSON() ([]byte, error) {
	type Alias Account
	aux := &struct {
		*Alias

		TunnelType       emptyStringInt `json:"tunnel_type"`
		TunnelMediumType emptyStringInt `json:"tunnel_medium_type"`
		VLAN             emptyStringInt `json:"vlan"`
	}{
		Alias: (*Alias)(dst),
	}

	aux.TunnelType = EmptyStringIntZero(dst.TunnelType)
	aux.TunnelMediumType = EmptyStringIntZero(dst.TunnelMediumType)
	aux.VLAN = EmptyStringIntZero(dst.VLAN)

	b, err := json.Marshal(aux)
	return b, err
}

func (c *Client) ListAccount(ctx context.Context, site string) ([]Account, error) {
	return c.listAccount(ctx, site)
}

func (c *Client) GetAccount(ctx context.Context, site, id string) (*Account, error) {
	return c.getAccount(ctx, site, id)
}

func (c *Client) DeleteAccount(ctx context.Context, site, id string) error {
	return c.deleteAccount(ctx, site, id)
}

func (c *Client) CreateAccount(ctx context.Context, site string, d *Account) (*Account, error) {
	return c.createAccount(ctx, site, d)
}

func (c *Client) UpdateAccount(ctx context.Context, site string, d *Account) (*Account, error) {
	return c.updateAccount(ctx, site, d)
}
