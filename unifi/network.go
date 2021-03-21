package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

func (n *Network) MarshalJSON() ([]byte, error) {
	type Alias Network
	var removeFields []string

	switch n.Purpose {
	case "wan":
		removeFields = []string{
			"auto_scale_enabled",
			"dhcpd_boot_enabled",
			"dhcpd_boot_filename",
			"dhcpd_boot_server",
			"dhcpd_dns_1",
			"dhcpd_dns_2",
			"dhcpd_dns_3",
			"dhcpd_dns_4",
			"dhcpd_dns_enabled",
			"dhcpd_enabled",
			"dhcpd_gateway",
			"dhcpd_gateway_enabled",
			"dhcpd_ip_1",
			"dhcpd_ip_2",
			"dhcpd_ip_3",
			"dhcpd_leasetime",
			"dhcpd_mac_1",
			"dhcpd_mac_2",
			"dhcpd_mac_3",
			"dhcpd_ntp_1",
			"dhcpd_ntp_2",
			"dhcpd_ntp_enabled",
			"dhcpd_start",
			"dhcpd_stop",
			"dhcpd_tftp_server",
			"dhcpd_time_offset",
			"dhcpd_time_offset_enabled",
			"dhcpd_unifi_controller",
			"dhcpdv6_dns_1",
			"dhcpdv6_dns_2",
			"dhcpdv6_dns_3",
			"dhcpdv6_dns_4",
			"dhcpdv6_dns_auto",
			"dhcpdv6_enabled",
			"dhcpdv6_leasetime",
			"dhcpdv6_start",
			"dhcpdv6_stop",
			"dhcpd_wpad_url",
			"dhcpd_wins_1",
			"dhcpd_wins_2",
			"dhcpd_wins_enabled",
			"dhcp_relay_enabled",
			"dhcpguard_enabled",
			"dpi_enabled",
			"dpigroup_id",
			"domain_name",
			"exposed_to_site_vpn",
			"gateway_device",
			"gateway_type",
			"igmp_fastleave",
			"igmp_groupmembership",
			"igmp_maxresponse",
			"igmp_mcrtrexpiretime",
			"igmp_querier",
			"igmp_snooping",
			"igmp_supression",
			"ipsec_dh_group",
			"ipsec_dynamic_routing",
			"ipsec_encryption",
			"ipsec_esp_dh_group",
			"ipsec_hash",
			"ipsec_ike_dh_group",
			"ipsec_interface",
			"ipsec_key_exchange",
			"ipsec_local_ip",
			"ipsec_peer_ip",
			"ipsec_pfs",
			"ipsec_profile",
			"ip_subnet",
			"ipv6_interface_type",
			"ipv6_pd_interface",
			"ipv6_pd_prefixid",
			"ipv6_pd_start",
			"ipv6_pd_stop",
			"ipv6_ra_enabled",
			"ipv6_ra_preferred_lifetime",
			"ipv6_ra_priority",
			"ipv6_ra_valid_lifetime",
			"ipv6_subnet",
			"is_nat",
			"l2tp_interface",
			"lte_lan_enabled",
			"nat_outbound_ip_addresses",
			"networkgroup",
			"openvpn_local_address",
			"openvpn_local_port",
			"openvpn_mode",
			"openvpn_remote_address",
			"openvpn_remote_host",
			"openvpn_remote_port",
			"pptpc_require_mppe",
			"pptpc_route_distance",
			"pptpc_server_ip",
			"pptpc_username",
			"priority",
			"purpose",
			"radiusprofile_id",
			"remote_site_id",
			"remote_site_subnets",
			"remote_vpn_subnets",
			"report_wan_event",
			"require_mschapv2",
			"route_distance",
			"upnp_lan_enabled",
			"usergroup_id",
			"vlan",
			"vlan_enabled",
			"vpn_client_default_route",
			"vpn_client_pull_dns",
			"vpn_type",
			"x_ipsec_pre_shared_key",
			"x_openvpn_shared_secret_key",
			"x_pptpc_password",
		}

	default:
		removeFields = []string{
			"wan_dhcp_options",
			"wan_dhcpv6_pd_size",
			"wan_dns1",
			"wan_dns2",
			"wan_dns3",
			"wan_dns4",
			"wan_egress_qos",
			"wan_gateway",
			"wan_gateway_v6",
			"wan_ip",
			"wan_ip_aliases",
			"wan_ipv6",
			"wan_load_balance_type",
			"wan_load_balance_weight",
			"wan_netmask",
			"wan_networkgroup",
			"wan_prefixlen",
			"wan_smartq_down_rate",
			"wan_smartq_enabled",
			"wan_smartq_up_rate",
			"wan_type",
			"wan_type_v6",
			"wan_username",
			"wan_vlan",
			"wan_vlan_enabled",
			"x_wan_password",
		}
	}

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(n),
	}
	b, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(b, &jsonMap)
	if err != nil {
		return nil, err
	}

	for _, key := range removeFields {
		delete(jsonMap, key)
	}

	return json.Marshal(jsonMap)
}

func (c *Client) DeleteNetwork(ctx context.Context, site, id, name string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), struct {
		Name string `json:"name"`
	}{
		Name: name,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ListNetwork(ctx context.Context, site string) ([]Network, error) {
	return c.listNetwork(ctx, site)
}

func (c *Client) GetNetwork(ctx context.Context, site, id string) (*Network, error) {
	return c.getNetwork(ctx, site, id)
}

func (c *Client) CreateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.createNetwork(ctx, site, d)
}

func (c *Client) UpdateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	return c.updateNetwork(ctx, site, d)
}
