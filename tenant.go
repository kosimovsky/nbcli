package main

import "time"

type T struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Group   struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Depth   int    `json:"_depth"`
	} `json:"group"`
	Description  string   `json:"description"`
	Comments     string   `json:"comments"`
	Tags         []string `json:"tags"`
	CustomFields struct {
	} `json:"custom_fields"`
	Created             time.Time `json:"created"`
	LastUpdated         time.Time `json:"last_updated"`
	CircuitCount        int       `json:"circuit_count"`
	DeviceCount         int       `json:"device_count"`
	IpaddressCount      int       `json:"ipaddress_count"`
	PrefixCount         int       `json:"prefix_count"`
	RackCount           int       `json:"rack_count"`
	SiteCount           int       `json:"site_count"`
	VirtualmachineCount int       `json:"virtualmachine_count"`
	VlanCount           int       `json:"vlan_count"`
	VrfCount            int       `json:"vrf_count"`
	ClusterCount        int       `json:"cluster_count"`
}
