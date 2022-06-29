package proc

import (
	"crypto/tls"
	"github.com/netbox-community/go-netbox/netbox/client"
)

type Output interface {
	outputJson(interface{}) error
	outputYaml(interface{}) error
	CreateOutput(string) error
}

type ApiClient struct {
	netBoxAPI *client.NetBoxAPI
}

func NewApiClient(c *client.NetBoxAPI) *ApiClient {
	return &ApiClient{netBoxAPI: c}
}

type ClientSession struct {
	client ApiClient
}

func NewClientSession() *ClientSession {
	return &ClientSession{}
}

type Config struct {
	NetboxHost string
	Token      string
	TlsConf    *tls.Config
}

type Client interface {
	//NetboxClient(config Config) *client.NetBoxAPI
	DefaultClient() *ApiClient
}

type Proc struct {
	Output
	Client
	Processor
}

type EntryType map[string]string

func (e EntryType) EntryTypes(cmd string) map[string]string {
	ETMap := make(map[string]string)

	switch cmd {
	case "stat":
		ETMap["server"] = "device"
		ETMap["switch"] = "device"
		ETMap["ip"] = "ipam"
		ETMap["vlan"] = "ipam"
		ETMap["vm"] = "virtualization"
		ETMap["tenant"] = "tenancy"
	default:
		return nil
	}
	return ETMap
}

type Processor interface {
	EntryTypes(cmd string) map[string]string
}
