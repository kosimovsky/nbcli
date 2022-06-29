package proc

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"netbox-client/pkg"
	"strconv"
	"time"
)

const Timeout = 10 * time.Second

func (c *ClientSession) DefaultClient() *ApiClient {
	netboxClient := NetboxClient(Config{
		NetboxHost: viper.GetString("netboxHost"),
		TlsConf:    &tls.Config{InsecureSkipVerify: viper.GetBool("insecure")},
		Token:      viper.GetString("token"),
	})
	return NewApiClient(netboxClient)
}

func NetboxClient(cfg Config) *client.NetBoxAPI {
	//TODO implement me
	httpClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: cfg.TlsConf},
	}
	transport := httptransport.NewWithClient(cfg.NetboxHost, client.DefaultBasePath, []string{"https"}, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+cfg.Token)

	return client.New(transport, nil)
}

func (c *ApiClient) GetCommand(args []string) error {
	err := c.GetAddress(args[1])
	if err != nil {
		return err
	}
	return nil
}

func (c *ApiClient) getDeviceByName(name string) error {

	params := dcim.NewDcimDevicesListParams().WithTimeout(Timeout)
	devicesList, err := c.netBoxAPI.Dcim.DcimDevicesList(params.WithName(&name), nil)
	if err != nil {
		return err
	}
	result := devicesList.GetPayload()

	var device []pkg.Device

	data, _ := json.MarshalIndent(result.Results, "", "")

	err = json.Unmarshal(data, &device)
	if err != nil {
		logrus.Printf("error unmarshalling : %s", err.Error())
	}

	fmt.Println(&name)
	fmt.Println(device[0].Id)

	return nil
}

func (c *ApiClient) GetAddress(name string) error {
	params := ipam.NewIpamIPAddressesListParamsWithTimeout(Timeout)
	id, err := c.deviceIDByName(name)
	if err != nil {
		return err
	}
	strID := strconv.Itoa(id)

	if err != nil {
		return err
	}
	list, err := c.netBoxAPI.Ipam.IpamIPAddressesList(params.WithDeviceID(&strID), nil)
	if err != nil {
		return err
	}
	result := list.GetPayload()
	data, err := json.Marshal(result.Results)
	var addresses []pkg.IpAddresss
	err = json.Unmarshal(data, &addresses)
	if err != nil {
		return err
	}

	fmt.Println(addresses[0].Address)

	return nil
}

func (c *ApiClient) deviceIDByName(name string) (int, error) {
	params := dcim.NewDcimDevicesListParams().WithTimeout(Timeout)
	devicesList, err := c.netBoxAPI.Dcim.DcimDevicesList(params.WithName(&name), nil)
	if err != nil {
		return 0, err
	}
	result := devicesList.GetPayload()

	if *result.Count > 0 {
		var device []pkg.Device
		data, err := json.Marshal(result.Results)
		fmt.Println(*result.Count)
		if err != nil {
			return 0, err
		}
		err = json.Unmarshal(data, &device)
		if err != nil {
			return 0, err
		}
		return device[0].Id, nil
	} else {
		err = errors.New(fmt.Sprintf("there is no object with entry %s in Netbox", name))
		logrus.Errorf("error while getting stat: %s", err)
		return 0, err
	}
}

func (c *ApiClient) Stat(objectName, objectType string) (*OutContent, error) {
	var output []byte
	switch objectType {
	case "device":
		id, err := c.deviceIDByName(objectName)
		if err != nil {
			return nil, err
		}
		params := dcim.NewDcimDevicesReadParams().WithID(int64(id))
		deviceRead, err := c.netBoxAPI.Dcim.DcimDevicesRead(params, nil)
		output, err = deviceRead.Payload.MarshalBinary()
	case "guess":

	}
	content := new(OutContent)
	content.content = output
	content.objectType = objectType
	return content, nil
}
