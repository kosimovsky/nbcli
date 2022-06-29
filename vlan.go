package main

import "time"

type Vlan struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Site    struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"site"`
	Group  interface{} `json:"group"`
	Vid    int         `json:"vid"`
	Name   string      `json:"name"`
	Tenant struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"tenant"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"status"`
	Role        interface{} `json:"role"`
	Description string      `json:"description"`
	Tags        []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Color   string `json:"color"`
	} `json:"tags"`
	CustomFields struct {
	} `json:"custom_fields"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	PrefixCount int       `json:"prefix_count"`
}
