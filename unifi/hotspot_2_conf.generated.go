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

type Hotspot2Conf struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	AnqpDomainID            int                                 `json:"anqp_domain_id,omitempty"` // ^0|[1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]|$
	Capab                   []Hotspot2ConfCapab                 `json:"capab,omitempty"`
	CellularNetworkList     []Hotspot2ConfCellularNetworkList   `json:"cellular_network_list,omitempty"`
	DeauthReqTimeout        int                                 `json:"deauth_req_timeout,omitempty"` // [1-9][0-9]|[1-9][0-9][0-9]|[1-2][0-9][0-9][0-9]|3[0-5][0-9][0-9]|3600
	DisableDgaf             bool                                `json:"disable_dgaf"`
	DomainNameList          []string                            `json:"domain_name_list,omitempty"` // .{1,128}
	FriendlyName            []Hotspot2ConfFriendlyName          `json:"friendly_name,omitempty"`
	GasAdvanced             bool                                `json:"gas_advanced"`
	GasComebackDelay        int                                 `json:"gas_comeback_delay,omitempty"`
	GasFragLimit            int                                 `json:"gas_frag_limit,omitempty"`
	Hessid                  string                              `json:"hessid"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$|^$
	HessidUsed              bool                                `json:"hessid_used"`
	IPaddrTypeAvailV4       int                                 `json:"ipaddr_type_avail_v4,omitempty"` // 0|1|2|3|4|5|6|7
	IPaddrTypeAvailV6       int                                 `json:"ipaddr_type_avail_v6,omitempty"` // 0|1|2
	Icons                   []Hotspot2ConfIcons                 `json:"icons,omitempty"`
	MetricsDownlinkLoad     int                                 `json:"metrics_downlink_load,omitempty"`
	MetricsDownlinkLoadSet  bool                                `json:"metrics_downlink_load_set"`
	MetricsDownlinkSpeed    int                                 `json:"metrics_downlink_speed,omitempty"`
	MetricsDownlinkSpeedSet bool                                `json:"metrics_downlink_speed_set"`
	MetricsInfoAtCapacity   bool                                `json:"metrics_info_at_capacity"`
	MetricsInfoLinkStatus   string                              `json:"metrics_info_link_status,omitempty"` // up|down|test
	MetricsInfoSymmetric    bool                                `json:"metrics_info_symmetric"`
	MetricsMeasurement      int                                 `json:"metrics_measurement,omitempty"`
	MetricsMeasurementSet   bool                                `json:"metrics_measurement_set"`
	MetricsStatus           bool                                `json:"metrics_status"`
	MetricsUplinkLoad       int                                 `json:"metrics_uplink_load,omitempty"`
	MetricsUplinkLoadSet    bool                                `json:"metrics_uplink_load_set"`
	MetricsUplinkSpeed      int                                 `json:"metrics_uplink_speed,omitempty"`
	MetricsUplinkSpeedSet   bool                                `json:"metrics_uplink_speed_set"`
	NaiRealmList            []Hotspot2ConfNaiRealmList          `json:"nai_realm_list,omitempty"`
	Name                    string                              `json:"name,omitempty"` // .{1,128}
	NetworkAccessAsra       bool                                `json:"network_access_asra"`
	NetworkAccessEsr        bool                                `json:"network_access_esr"`
	NetworkAccessInternet   bool                                `json:"network_access_internet"`
	NetworkAccessUesa       bool                                `json:"network_access_uesa"`
	NetworkAuthType         int                                 `json:"network_auth_type,omitempty"` // -1|0|1|2|3
	NetworkAuthUrl          string                              `json:"network_auth_url,omitempty"`
	NetworkType             int                                 `json:"network_type,omitempty"` // 0|1|2|3|4|5|14|15
	Osu                     []Hotspot2ConfOsu                   `json:"osu,omitempty"`
	OsuSSID                 string                              `json:"osu_ssid"`
	QOSMapDcsp              []Hotspot2ConfQOSMapDcsp            `json:"qos_map_dcsp,omitempty"`
	QOSMapExceptions        []Hotspot2ConfQOSMapExceptions      `json:"qos_map_exceptions,omitempty"`
	QOSMapStatus            bool                                `json:"qos_map_status"`
	RoamingConsortiumList   []Hotspot2ConfRoamingConsortiumList `json:"roaming_consortium_list,omitempty"`
	SaveTimestamp           string                              `json:"save_timestamp,omitempty"`
	TCFilename              string                              `json:"t_c_filename,omitempty"` // .{1,256}
	TCTimestamp             int                                 `json:"t_c_timestamp,omitempty"`
	VenueGroup              int                                 `json:"venue_group,omitempty"` // 0|1|2|3|4|5|6|7|8|9|10|11
	VenueName               []Hotspot2ConfVenueName             `json:"venue_name,omitempty"`
	VenueType               int                                 `json:"venue_type,omitempty"` // 0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15
}

type Hotspot2ConfCapab struct {
	Port     int    `json:"port,omitempty"`     // ^(0|[1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])|$
	Protocol string `json:"protocol,omitempty"` // icmp|tcp_udp|tcp|udp|esp
	Status   string `json:"status,omitempty"`   // closed|open|unknown
}

type Hotspot2ConfCellularNetworkList struct {
	Mcc  int    `json:"mcc,omitempty"`
	Mnc  int    `json:"mnc,omitempty"`
	Name string `json:"name,omitempty"` // .{1,128}
}

type Hotspot2ConfDescription struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Text     string `json:"text,omitempty"`     // .{1,128}
}

type Hotspot2ConfFriendlyName struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Text     string `json:"text,omitempty"`     // .{1,128}
}

type Hotspot2ConfIcon struct {
	Name string `json:"name,omitempty"` // .{1,128}
}

type Hotspot2ConfIcons struct {
	Data     string `json:"data,omitempty"`
	Filename string `json:"filename,omitempty"` // .{1,256}
	Height   int    `json:"height,omitempty"`
	Language string `json:"language,omitempty"` // [a-z]{3}
	Media    string `json:"media,omitempty"`    // .{1,256}
	Name     string `json:"name,omitempty"`     // .{1,256}
	Size     int    `json:"size,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type Hotspot2ConfNaiRealmList struct {
	AuthIDs   string `json:"auth_ids,omitempty"`
	AuthVals  string `json:"auth_vals,omitempty"`
	EapMethod int    `json:"eap_method,omitempty"` // 13|21|18|23|50
	Encoding  int    `json:"encoding,omitempty"`   // 0|1
	Name      string `json:"name,omitempty"`       // .{1,128}
	Status    bool   `json:"status"`
}

type Hotspot2ConfOsu struct {
	Description      []Hotspot2ConfDescription  `json:"description,omitempty"`
	FriendlyName     []Hotspot2ConfFriendlyName `json:"friendly_name,omitempty"`
	Icon             []Hotspot2ConfIcon         `json:"icon,omitempty"`
	MethodOmaDm      bool                       `json:"method_oma_dm"`
	MethodSoapXmlSpp bool                       `json:"method_soap_xml_spp"`
	Nai              string                     `json:"nai,omitempty"`
	Nai2             string                     `json:"nai2,omitempty"`
	OperatingClass   string                     `json:"operating_class,omitempty"` // [0-9A-Fa-f]{12}
	ServerUri        string                     `json:"server_uri,omitempty"`
}

type Hotspot2ConfQOSMapDcsp struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

type Hotspot2ConfQOSMapExceptions struct {
	Dcsp int `json:"dcsp,omitempty"`
	Up   int `json:"up,omitempty"` // [0-7]
}

type Hotspot2ConfRoamingConsortiumList struct {
	Name string `json:"name,omitempty"` // .{1,128}
	Oid  string `json:"oid,omitempty"`  // .{1,128}
}

type Hotspot2ConfVenueName struct {
	Language string `json:"language,omitempty"` // [a-z]{3}
	Name     string `json:"name,omitempty"`
	Url      string `json:"url,omitempty"`
}

func (c *Client) listHotspot2Conf(ctx context.Context, site string) ([]Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/hotspot2conf", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getHotspot2Conf(ctx context.Context, site, id string) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteHotspot2Conf(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createHotspot2Conf(ctx context.Context, site string, d *Hotspot2Conf) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/hotspot2conf", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateHotspot2Conf(ctx context.Context, site string, d *Hotspot2Conf) (*Hotspot2Conf, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []Hotspot2Conf `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/hotspot2conf/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
