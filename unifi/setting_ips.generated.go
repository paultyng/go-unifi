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

type SettingIps struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DNSFiltering bool `json:"dns_filtering"`
	DNSFilters   []struct {
		Filter    string `json:"filter,omitempty"` // security|adult|family
		NetworkID string `json:"network_id"`
		Version   string `json:"version,omitempty"` // v4|v6

	} `json:"dns_filters,omitempty"`
	EnabledCategories []string `json:"enabled_categories,omitempty"` // emerging-activex|emerging-attackresponse|botcc|emerging-chat|ciarmy|compromised|emerging-dns|emerging-dos|dshield|emerging-exploit|emerging-ftp|emerging-games|emerging-icmp|emerging-icmpinfo|emerging-imap|emerging-inappropriate|emerging-info|emerging-malware|emerging-misc|emerging-mobile|emerging-netbios|emerging-p2p|emerging-policy|emerging-pop3|emerging-rpc|emerging-scada|emerging-scan|emerging-shellcode|emerging-smtp|emerging-snmp|spamhaus|emerging-sql|emerging-telnet|emerging-tftp|tor|emerging-trojan|emerging-useragent|emerging-voip|emerging-webapps|emerging-webclient|emerging-webserver|emerging-worm
	EndpointScanning  bool     `json:"endpoint_scanning"`
	Honeypot          []struct {
		IPAddress string `json:"ip_address,omitempty"`
		NetworkID string `json:"network_id"`
		Version   string `json:"version,omitempty"` // v4|v6

	} `json:"honeypot,omitempty"`
	HoneypotEnabled     bool   `json:"honeypot_enabled"`
	IPsMode             string `json:"ips_mode,omitempty"` // ids|ips|ipsInline|disabled
	RestrictIPAddresses bool   `json:"restrict_ip_addresses"`
	RestrictTor         bool   `json:"restrict_tor"`
	RestrictTorrents    bool   `json:"restrict_torrents"`
	Suppression         struct {
		Alerts []struct {
			Category  string `json:"category,omitempty"`
			Gid       int    `json:"gid,omitempty"`
			ID        int    `json:"id,omitempty"`
			Signature string `json:"signature,omitempty"`
			Tracking  []struct {
				Direction string `json:"direction,omitempty"` // both|src|dest
				Mode      string `json:"mode,omitempty"`      // ip|subnet|network
				Value     string `json:"value,omitempty"`
			} `json:"tracking,omitempty"`
			Type string `json:"type,omitempty"` // all|track

		} `json:"alerts,omitempty"`
		Whitelist []struct {
			Direction string `json:"direction,omitempty"` // both|src|dest
			Mode      string `json:"mode,omitempty"`      // ip|subnet|network
			Value     string `json:"value,omitempty"`
		} `json:"whitelist,omitempty"`
	} `json:"suppression"`
}
