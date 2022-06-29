package proc

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"netbox-client/pkg"
)

type OutContent struct {
	content    []byte
	objectType string
}

func (o *OutContent) outputJson(device interface{}) error {
	ret, _ := json.MarshalIndent(device, "", " ")
	fmt.Println(string(ret))
	return nil
}

func (o *OutContent) outputYaml(device interface{}) error {
	ret, _ := yaml.Marshal(device)
	fmt.Println(string(ret))
	return nil
}

func (o *OutContent) CreateOutput(outputFormat string) error {
	switch o.objectType {
	case "device":
		var device pkg.Device
		err := json.Unmarshal(o.content, &device)
		if err != nil {
			return err
		}
		switch outputFormat {
		case "json":
			err = o.outputJson(device)
			if err != nil {
				return err
			}
		default:
			err = o.outputYaml(device)
			if err != nil {
				return err
			}
		}
	}

	//switch outputFormat {
	//case "json":
	//	err := o.outputJson()
	//	if err != nil {
	//		return err
	//	}
	//default:
	//	err := o.outputYaml()
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
	//var device pkg.Device
	//err := json.Unmarshal(o.content, &device)
	//if err != nil {
	//	logrus.Fatalf("error while unmarshalling content: %s", err.Error())
	//}
	//if err != nil {
	//	return
	//}
	//ret, err := json.MarshalIndent(device.PrimaryIp.Address, "", " ")
	//if err != nil {
	//	logrus.Fatalf("error while marshalling with indent : %s", err.Error())
	//}
	//fmt.Println(string(ret))
}
