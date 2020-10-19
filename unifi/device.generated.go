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

type Device struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	MAC string `json:"mac,omitempty"`

	AtfEnabled                 bool                      `json:"atf_enabled,omitempty"`
	BandsteeringMode           string                    `json:"bandsteering_mode,omitempty"` // off|equal|prefer_5g
	BaresipAuthUser            string                    `json:"baresip_auth_user,omitempty"` // ^\+?[a-zA-Z0-9_.\-!~*'()]*
	BaresipEnabled             bool                      `json:"baresip_enabled,omitempty"`
	BaresipExtension           string                    `json:"baresip_extension,omitempty"` // ^\+?[a-zA-Z0-9_.\-!~*'()]*
	ConfigNetwork              DeviceConfigNetwork       `json:"config_network,omitempty"`
	DPIEnabled                 bool                      `json:"dpi_enabled,omitempty"`
	Disabled                   bool                      `json:"disabled,omitempty"`
	Dot1XFallbackNetworkID     string                    `json:"dot1x_fallback_networkconf_id,omitempty"` // [\d\w]+|
	Dot1XPortctrlEnabled       bool                      `json:"dot1x_portctrl_enabled,omitempty"`
	EthernetOverrides          []DeviceEthernetOverrides `json:"ethernet_overrides,omitempty"`
	FlowctrlEnabled            bool                      `json:"flowctrl_enabled,omitempty"`
	HeightInMeters             float64                   `json:"heightInMeters,omitempty"`
	JumboframeEnabled          bool                      `json:"jumboframe_enabled,omitempty"`
	LcmBrightness              int                       `json:"lcm_brightness,omitempty"` // [1-9]|[1-9][0-9]|100
	LcmBrightnessOverride      bool                      `json:"lcm_brightness_override,omitempty"`
	LcmIDleTimeout             int                       `json:"lcm_idle_timeout,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|3600
	LcmIDleTimeoutOverride     bool                      `json:"lcm_idle_timeout_override,omitempty"`
	LcmTrackerEnabled          bool                      `json:"lcm_tracker_enabled,omitempty"`
	LcmTrackerSeed             string                    `json:"lcm_tracker_seed,omitempty"`              // .{0,50}
	LedOverride                string                    `json:"led_override,omitempty"`                  // default|on|off
	LedOverrideColor           string                    `json:"led_override_color,omitempty"`            // ^#(?:[0-9a-fA-F]{3}){1,2}$
	LedOverrideColorBrightness int                       `json:"led_override_color_brightness,omitempty"` // ^[0-9][0-9]?$|^100$
	Locked                     bool                      `json:"locked,omitempty"`
	LteApn                     string                    `json:"lte_apn,omitempty"` // .{1,128}
	LteExtAnt                  bool                      `json:"lte_ext_ant,omitempty"`
	LtePoe                     bool                      `json:"lte_poe,omitempty"`
	LteSimPin                  int                       `json:"lte_sim_pin,omitempty"`
	LteSoftLimit               int                       `json:"lte_soft_limit,omitempty"`
	MapID                      string                    `json:"map_id,omitempty"`
	MeshStaVapEnabled          bool                      `json:"mesh_sta_vap_enabled,omitempty"`
	MgmtNetworkID              string                    `json:"mgmt_network_id,omitempty"`       // [\d\w]+
	Name                       string                    `json:"name,omitempty"`                  // .{1,128}
	OutdoorModeOverride        string                    `json:"outdoor_mode_override,omitempty"` // default|on|off
	OutletEnabled              bool                      `json:"outlet_enabled,omitempty"`
	OutletOverrides            []DeviceOutletOverrides   `json:"outlet_overrides,omitempty"`
	PortOverrides              []DevicePortOverrides     `json:"port_overrides,omitempty"`
	PowerSourceCtrl            string                    `json:"power_source_ctrl,omitempty"` // auto|8023af|8023at|8023bt-type3|8023bt-type4|pasv24|poe-injector|ac|adapter|dc|rps
	PowerSourceCtrlEnabled     bool                      `json:"power_source_ctrl_enabled,omitempty"`
	RADIUSProfileID            string                    `json:"radiusprofile_id,omitempty"`
	RadioTable                 []DeviceRadioTable        `json:"radio_table,omitempty"`
	ResetbtnEnabled            string                    `json:"resetbtn_enabled,omitempty"` // on|off
	RpsOverride                DeviceRpsOverride         `json:"rps_override,omitempty"`
	SnmpContact                string                    `json:"snmp_contact,omitempty"`  // .{0,255}
	SnmpLocation               string                    `json:"snmp_location,omitempty"` // .{0,255}
	StpPriority                string                    `json:"stp_priority,omitempty"`  // 0|4096|8192|12288|16384|20480|24576|28672|32768|36864|40960|45056|49152|53248|57344|61440
	StpVersion                 string                    `json:"stp_version,omitempty"`   // stp|rstp|disabled
	SwitchVLANEnabled          bool                      `json:"switch_vlan_enabled,omitempty"`
	UbbPairName                string                    `json:"ubb_pair_name,omitempty"` // .{1,128}
	Volume                     int                       `json:"volume,omitempty"`        // [0-9]|[1-9][0-9]|100
	WLANOverrides              []DeviceWLANOverrides     `json:"wlan_overrides,omitempty"`
	X                          float64                   `json:"x,omitempty"`
	XBaresipPassword           string                    `json:"x_baresip_password,omitempty"` // ^[a-zA-Z0-9_.\-!~*'()]*
	Y                          float64                   `json:"y,omitempty"`
}

type DeviceConfigNetwork struct {
	BondingEnabled bool   `json:"bonding_enabled,omitempty"`
	DNS1           string `json:"dns1,omitempty"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$|^$
	DNS2           string `json:"dns2,omitempty"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$|^$
	DNSsuffix      string `json:"dnssuffix,omitempty"`
	Gateway        string `json:"gateway,omitempty"` // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	IP             string `json:"ip,omitempty"`      // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
	Netmask        string `json:"netmask,omitempty"` // ^((128|192|224|240|248|252|254)\.0\.0\.0)|(255\.(((0|128|192|224|240|248|252|254)\.0\.0)|(255\.(((0|128|192|224|240|248|252|254)\.0)|255\.(0|128|192|224|240|248|252|254)))))$
	Type           string `json:"type,omitempty"`    // dhcp|static
}

type DeviceEthernetOverrides struct {
	Ifname       string `json:"ifname,omitempty"`       // eth[0-9]{1,2}
	NetworkGroup string `json:"networkgroup,omitempty"` // LAN[2-8]?|WAN[2]?
}

type DeviceOutletOverrides struct {
	CycleEnabled bool   `json:"cycle_enabled,omitempty"`
	Index        int    `json:"index,omitempty"`
	Name         string `json:"name,omitempty"` // .{0,128}
	RelayState   bool   `json:"relay_state,omitempty"`
}

type DevicePortOverrides struct {
	AggregateNumPorts            int      `json:"aggregate_num_ports,omitempty"` // [2-6]
	Autoneg                      bool     `json:"autoneg,omitempty"`
	Dot1XCtrl                    string   `json:"dot1x_ctrl,omitempty"`             // auto|force_authorized|force_unauthorized|mac_based|multi_host
	Dot1XIDleTimeout             int      `json:"dot1x_idle_timeout,omitempty"`     // [0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]
	EgressRateLimitKbps          int      `json:"egress_rate_limit_kbps,omitempty"` // 6[4-9]|[7-9][0-9]|[1-9][0-9]{2,6}
	EgressRateLimitKbpsEnabled   bool     `json:"egress_rate_limit_kbps_enabled,omitempty"`
	FullDuplex                   bool     `json:"full_duplex,omitempty"`
	Isolation                    bool     `json:"isolation,omitempty"`
	LldpmedEnabled               bool     `json:"lldpmed_enabled,omitempty"`
	LldpmedNotifyEnabled         bool     `json:"lldpmed_notify_enabled,omitempty"`
	MirrorPortIDX                int      `json:"mirror_port_idx,omitempty"` // [1-9]|[1-4][0-9]|5[0-2]
	Name                         string   `json:"name,omitempty"`            // .{0,128}
	OpMode                       string   `json:"op_mode,omitempty"`         // switch|mirror|aggregate
	PoeMode                      string   `json:"poe_mode,omitempty"`        // auto|pasv24|passthrough|off
	PortIDX                      int      `json:"port_idx,omitempty"`        // [1-9]|[1-4][0-9]|5[0-2]
	PortProfileID                string   `json:"portconf_id,omitempty"`     // [\d\w]+
	PortSecurityEnabled          bool     `json:"port_security_enabled,omitempty"`
	PortSecurityMACAddress       []string `json:"port_security_mac_address,omitempty"` // ^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$
	PriorityQueue1Level          int      `json:"priority_queue1_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue2Level          int      `json:"priority_queue2_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue3Level          int      `json:"priority_queue3_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	PriorityQueue4Level          int      `json:"priority_queue4_level,omitempty"`     // [0-9]|[1-9][0-9]|100
	Speed                        int      `json:"speed,omitempty"`                     // 10|100|1000|2500|5000|10000|20000|25000|40000|50000|100000
	StormctrlBroadcastastEnabled bool     `json:"stormctrl_bcast_enabled,omitempty"`
	StormctrlBroadcastastLevel   int      `json:"stormctrl_bcast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlBroadcastastRate    int      `json:"stormctrl_bcast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StormctrlMcastEnabled        bool     `json:"stormctrl_mcast_enabled,omitempty"`
	StormctrlMcastLevel          int      `json:"stormctrl_mcast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlMcastRate           int      `json:"stormctrl_mcast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StormctrlType                string   `json:"stormctrl_type,omitempty"`        // level|rate
	StormctrlUcastEnabled        bool     `json:"stormctrl_ucast_enabled,omitempty"`
	StormctrlUcastLevel          int      `json:"stormctrl_ucast_level,omitempty"` // [0-9]|[1-9][0-9]|100
	StormctrlUcastRate           int      `json:"stormctrl_ucast_rate,omitempty"`  // [0-9]|[1-9][0-9]{1,6}|1[0-3][0-9]{6}|14[0-7][0-9]{5}|148[0-7][0-9]{4}|14880000
	StpPortMode                  bool     `json:"stp_port_mode,omitempty"`
}

type DeviceRadioTable struct {
	AntennaGain           int    `json:"antenna_gain,omitempty"`   // ^-?([0-9]|[1-9][0-9])
	AntennaID             int    `json:"antenna_id,omitempty"`     // -1|[0-9]
	BackupChannel         int    `json:"backup_channel,omitempty"` // [0-9]|[1][0-4]|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	Channel               int    `json:"channel,omitempty"`        // [0-9]|[1][0-4]|4.5|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	HardNoiseFloorEnabled bool   `json:"hard_noise_floor_enabled,omitempty"`
	Ht                    string `json:"ht,omitempty"` // 20|40|80|160|1080|2160
	LoadbalanceEnabled    bool   `json:"loadbalance_enabled,omitempty"`
	Maxsta                int    `json:"maxsta,omitempty"`   // [1-9]|[1-9][0-9]|1[0-9]{2}|200|^$
	MinRssi               int    `json:"min_rssi,omitempty"` // ^-([1-9]|[1-8][0-9]|9[0-4])$
	MinRssiEnabled        bool   `json:"min_rssi_enabled,omitempty"`
	Name                  string `json:"name,omitempty"`
	Radio                 string `json:"radio,omitempty"`      // ng|na|ad
	SensLevel             int    `json:"sens_level,omitempty"` // ^-([5-8][0-9]|90)$
	SensLevelEnabled      bool   `json:"sens_level_enabled,omitempty"`
	TxPower               int    `json:"tx_power,omitempty"`      // [\d]+|auto
	TxPowerMode           string `json:"tx_power_mode,omitempty"` // auto|medium|high|low|custom
	VwireEnabled          bool   `json:"vwire_enabled,omitempty"`
}

type DeviceRpsOverride struct {
	PowerManagementMode string               `json:"power_management_mode,omitempty"` // dynamic|static
	RpsPortTable        []DeviceRpsPortTable `json:"rps_port_table,omitempty"`
}

type DeviceRpsPortTable struct {
	Name     string `json:"name,omitempty"`      // .{0,32}
	PortIDX  int    `json:"port_idx,omitempty"`  // [1-6]
	PortMode string `json:"port_mode,omitempty"` // auto|force_active|manual|disabled
}

type DeviceWLANOverrides struct {
	Enabled            bool   `json:"enabled,omitempty"`
	Name               string `json:"name,omitempty"` // .{1,32}
	NameCombineEnabled bool   `json:"name_combine_enabled,omitempty"`
	NameCombineSuffix  string `json:"name_combine_suffix,omitempty"` // .{0,8}
	Radio              string `json:"radio,omitempty"`               // ng|na
	RadioName          string `json:"radio_name,omitempty"`
	VLAN               int    `json:"vlan,omitempty"` // [2-9]|[1-9][0-9]{1,2}|[1-3][0-9]{3}|40[0-8][0-9]|409[0-5]|^$
	VLANEnabled        bool   `json:"vlan_enabled,omitempty"`
	WLANID             string `json:"wlan_id,omitempty"`      // [\d\w]+
	XPassphrase        string `json:"x_passphrase,omitempty"` // [\x20-\x7E]{8,63}|[0-9a-fA-F]{64}
}

func (c *Client) listDevice(ctx context.Context, site string) ([]Device, error) {
	var respBody struct {
		Meta meta     `json:"meta"`
		Data []Device `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/stat/device", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getDevice(ctx context.Context, site, id string) (*Device, error) {
	var respBody struct {
		Meta meta     `json:"meta"`
		Data []Device `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/stat/device/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteDevice(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/device/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	var respBody struct {
		Meta meta     `json:"meta"`
		Data []Device `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/device", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	var respBody struct {
		Meta meta     `json:"meta"`
		Data []Device `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/device/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
