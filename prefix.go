package main

import "time"

type Prefix struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Family  struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"family"`
	Prefix string `json:"prefix"`
	Site   struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"site"`
	Vrf    interface{} `json:"vrf"`
	Tenant interface{} `json:"tenant"`
	Vlan   struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Vid     int    `json:"vid"`
		Name    string `json:"name"`
	} `json:"vlan"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Role         interface{}   `json:"role"`
	IsPool       bool          `json:"is_pool"`
	MarkUtilized bool          `json:"mark_utilized"`
	Description  string        `json:"description"`
	Tags         []interface{} `json:"tags"`
	CustomFields struct {
		Comment interface{} `json:"Comment"`
	} `json:"custom_fields"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Children    int       `json:"children"`
	Depth       int       `json:"_depth"`
}
