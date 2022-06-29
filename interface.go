package main

import "time"

type Interface struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Device  struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
	} `json:"device"`
	Module interface{} `json:"module"`
	Name   string      `json:"name"`
	Label  string      `json:"label"`
	Type   struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"type"`
	Enabled bool        `json:"enabled"`
	Parent  interface{} `json:"parent"`
	Bridge  interface{} `json:"bridge"`
	Lag     struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Device  struct {
			Id      int    `json:"id"`
			Url     string `json:"url"`
			Display string `json:"display"`
			Name    string `json:"name"`
		} `json:"device"`
		Name     string      `json:"name"`
		Cable    interface{} `json:"cable"`
		Occupied bool        `json:"_occupied"`
	} `json:"lag"`
	Mtu                        interface{}   `json:"mtu"`
	MacAddress                 interface{}   `json:"mac_address"`
	Speed                      interface{}   `json:"speed"`
	Duplex                     interface{}   `json:"duplex"`
	Wwn                        interface{}   `json:"wwn"`
	MgmtOnly                   bool          `json:"mgmt_only"`
	Description                string        `json:"description"`
	Mode                       interface{}   `json:"mode"`
	RfRole                     interface{}   `json:"rf_role"`
	RfChannel                  interface{}   `json:"rf_channel"`
	RfChannelFrequency         interface{}   `json:"rf_channel_frequency"`
	RfChannelWidth             interface{}   `json:"rf_channel_width"`
	TxPower                    interface{}   `json:"tx_power"`
	UntaggedVlan               interface{}   `json:"untagged_vlan"`
	TaggedVlans                []interface{} `json:"tagged_vlans"`
	MarkConnected              bool          `json:"mark_connected"`
	Cable                      interface{}   `json:"cable"`
	WirelessLink               interface{}   `json:"wireless_link"`
	LinkPeer                   interface{}   `json:"link_peer"`
	LinkPeerType               interface{}   `json:"link_peer_type"`
	WirelessLans               []interface{} `json:"wireless_lans"`
	Vrf                        interface{}   `json:"vrf"`
	ConnectedEndpoint          interface{}   `json:"connected_endpoint"`
	ConnectedEndpointType      interface{}   `json:"connected_endpoint_type"`
	ConnectedEndpointReachable interface{}   `json:"connected_endpoint_reachable"`
	Tags                       []interface{} `json:"tags"`
	CustomFields               struct {
	} `json:"custom_fields"`
	Created          time.Time `json:"created"`
	LastUpdated      time.Time `json:"last_updated"`
	CountIpaddresses int       `json:"count_ipaddresses"`
	CountFhrpGroups  int       `json:"count_fhrp_groups"`
	Occupied         bool      `json:"_occupied"`
}
