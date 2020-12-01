package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

func (dst *FirewallRule) UnmarshalJSON(b []byte) error {
	type Alias FirewallRule
	aux := &struct {
		RuleIndex emptyStringInt `json:"rule_index"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	dst.RuleIndex = int(aux.RuleIndex)

	return nil
}

func (c *Client) ListFirewallRule(ctx context.Context, site string) ([]FirewallRule, error) {
	return c.listFirewallRule(ctx, site)
}

func (c *Client) GetFirewallRule(ctx context.Context, site, id string) (*FirewallRule, error) {
	return c.getFirewallRule(ctx, site, id)
}

func (c *Client) DeleteFirewallRule(ctx context.Context, site, id string) error {
	return c.deleteFirewallRule(ctx, site, id)
}

func (c *Client) CreateFirewallRule(ctx context.Context, site string, d *FirewallRule) (*FirewallRule, error) {
	return c.createFirewallRule(ctx, site, d)
}

func (c *Client) UpdateFirewallRule(ctx context.Context, site string, d *FirewallRule) (*FirewallRule, error) {
	return c.updateFirewallRule(ctx, site, d)
}
