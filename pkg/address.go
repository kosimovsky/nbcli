package pkg

import "time"

type IpAddresss struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Family  struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"family"`
	Address string `json:"address"`
	Vrf     struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Rd      string `json:"rd"`
	} `json:"vrf"`
	Tenant interface{} `json:"tenant"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Role               interface{} `json:"role"`
	AssignedObjectType string      `json:"assigned_object_type"`
	AssignedObjectId   int         `json:"assigned_object_id"`
	AssignedObject     struct {
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
	} `json:"assigned_object"`
	NatInside    interface{}   `json:"nat_inside"`
	NatOutside   interface{}   `json:"nat_outside"`
	DnsName      string        `json:"dns_name"`
	Description  string        `json:"description"`
	Tags         []interface{} `json:"tags"`
	CustomFields struct {
	} `json:"custom_fields"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}
