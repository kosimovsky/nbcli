package pkg

import "time"

type Device struct {
	Id int `json:"id"`
	//Url        string `json:"url"`
	//Display    string `json:"display"`
	Name       string `json:"name"`
	DeviceType struct {
		//Id           int    `json:"id"`
		//Url          string `json:"url"`
		//Display      string `json:"display"`
		Manufacturer struct {
			//Id      int    `json:"id"`
			//Url     string `json:"url"`
			//Display string `json:"display"`
			Name string `json:"name"`
			//Slug    string `json:"slug"`
		} `json:"manufacturer"`
		Model string `json:"model"`
		//Slug  string `json:"slug"`
	} `json:"device_type"`
	//DeviceRole struct {
	//	Id      int    `json:"id"`
	//	Url     string `json:"url"`
	//	Display string `json:"display"`
	//	Name    string `json:"name"`
	//	Slug    string `json:"slug"`
	//} `json:"device_role"`
	//Tenant struct {
	//	Id      int    `json:"id"`
	//	Url     string `json:"url"`
	//	Display string `json:"display"`
	//	Name    string `json:"name"`
	//	Slug    string `json:"slug"`
	//} `json:"tenant"`
	//Platform interface{} `json:"platform"`
	Serial   string      `json:"serial"`
	AssetTag interface{} `json:"asset_tag"`
	Site     struct {
		//Id      int    `json:"id"`
		//Url     string `json:"url"`
		//Display string `json:"display"`
		Name string `json:"name"`
		//Slug    string `json:"slug"`
	} `json:"site"`
	Location struct {
		//Id      int    `json:"id"`
		//Url     string `json:"url"`
		//Display string `json:"display"`
		Name string `json:"name"`
		//Slug    string `json:"slug"`
		//Depth   int    `json:"_depth"`
	} `json:"location"`
	Rack struct {
		//Id      int    `json:"id"`
		//Url     string `json:"url"`
		//Display string `json:"display"`
		Name string `json:"name"`
	} `json:"rack"`
	//Position int `json:"position"`
	//Face     struct {
	//	Value string `json:"value"`
	//	Label string `json:"label"`
	//} `json:"face"`
	//ParentDevice interface{} `json:"parent_device"`
	//Status       struct {
	//	Value string `json:"value"`
	//	Label string `json:"label"`
	//} `json:"status"`
	//Airflow   interface{} `json:"airflow"`
	PrimaryIp struct {
		//Id      int    `json:"id"`
		//Url     string `json:"url"`
		//Display string `json:"display"`
		//Family  int    `json:"family"`
		Address string `json:"address"`
	} `json:"primary_ip"`
	//PrimaryIp4 struct {
	//	Id      int    `json:"id"`
	//	Url     string `json:"url"`
	//	Display string `json:"display"`
	//	Family  int    `json:"family"`
	//	Address string `json:"address"`
	//} `json:"primary_ip4"`
	//PrimaryIp6     interface{} `json:"primary_ip6"`
	//Cluster        interface{} `json:"cluster"`
	//VirtualChassis interface{} `json:"virtual_chassis"`
	//VcPosition     interface{} `json:"vc_position"`
	//VcPriority     interface{} `json:"vc_priority"`
	//Comments       string      `json:"comments"`
	//LocalContextData interface{} `json:"local_context_data"`
	Tags []struct {
		Id      int    `json:"id"`
		Url     string `json:"url"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Color   string `json:"color"`
	} `json:"tags"`
	CustomFields struct {
		Cpu                interface{} `json:"cpu"`
		Memory             interface{} `json:"memory"`
		NumberHdd          interface{} `json:"number_hdd"`
		Hdd                interface{} `json:"hdd"`
		NumberSsd          interface{} `json:"number_ssd"`
		Ssd                interface{} `json:"ssd"`
		NumberNvme         interface{} `json:"number_nvme"`
		Nvme               interface{} `json:"nvme"`
		InventoryNumber    string      `json:"Inventory_number"`
		Barcode            string      `json:"Barcode"`
		TalmerSsn          string      `json:"Talmer_ssn"`
		ServiceEnvironment string      `json:"Service Environment"`
		ServiceContract    interface{} `json:"service_contract"`
		PropertyOf         string      `json:"Property of"`
		Change             interface{} `json:"Change"`
		Sobol              interface{} `json:"Sobol"`
		SNSobol            interface{} `json:"SN Sobol"`
	} `json:"custom_fields"`
	ConfigContext struct {
	} `json:"config_context"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}
