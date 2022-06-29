// Package cmd
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"netbox-client/pkg/proc"
)

// objectType is variable to which you need to pass the type of the netbox object.
var objectType string

// objectName is variable to which you need to pass the name of the netbox object.
var objectName string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:              "get",
	TraverseChildren: true,
	Short:            "A brief description",
	Long:             `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("get called")
		return nil
	},
}

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "A brief description",
	Long:  `A longer description`,
	Args:  cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := initClient()
		err := client.GetCommand(args)
		if err != nil {
			logrus.Errorf("Error occured: %s", err.Error())
			return err
		}
		return nil
	},
}

var addrCmd = &cobra.Command{
	Use:   "addr",
	Short: "Get ip address",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := initClient()
		err := client.GetAddress(args[1])
		if err != nil {
			logrus.Errorf("Error occured: %s", err.Error())
			return err
		}
		return nil
	},
}

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get netbox object (server, virtual machine, ipaddress, vlan etc.) statistics",
	Example: `get stat --type server --entry <hostname>
get stat --type vm --entry <hostname> --out json
get stat --type ip4 --entry 192.168.1.1/24`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if objectType != "" {
			t := new(proc.EntryType)
			et := t.EntryTypes(cmd.Use)
			if ok, err := contains(et, objectType); !ok {
				logrus.Errorf("%s: only valid entry types are %v", err.Error(), et)
				return err
			} else {
				client := initClient()
				output, err := client.Stat(objectName, et[objectType])
				if err != nil {
					return err
				}
				if err := output.CreateOutput(outputFormat); err != nil {
					return err
				}
			}
		} else {
			err := errors.New("objectType is undefined, pass it via --type flag")
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	// get command initiation
	getCmd.AddCommand(deviceCmd, addrCmd, statCmd)

	//deviceCmd.Flags().StringP("name", "n", "", "device by name")

	addrCmd.PersistentFlags().StringP("name", "n", "", "device by name")

	// stat command flags
	statCmd.PersistentFlags().StringVarP(&objectName, "entry", "e", "", "entry (hostname, address etc)")
	err := statCmd.MarkPersistentFlagRequired("entry")
	if err != nil {
		logrus.Printf("flag name required: %s", err.Error())
	}
	statCmd.PersistentFlags().StringVarP(&objectType, "type", "t", "", validTypes(statCmd.Use))
	err = statCmd.MarkPersistentFlagRequired("type")
	if err != nil {
		logrus.Printf("flag type required: %s", err.Error())
	}
}

// initClient returns ready to use netbox API client
func initClient() *proc.ApiClient {
	session := proc.NewClientSession()
	return session.DefaultClient()
}

func contains(m map[string]string, t string) (bool, error) {
	for r := range m {
		if r == t {
			return true, nil
		}
	}
	err := errors.New("wrong entry type")
	return false, err
}

func validTypes(cmd string) string {
	var e proc.EntryType
	vt := e.EntryTypes(cmd)
	var ret []string
	for r := range vt {
		ret = append(ret, r)
	}
	return fmt.Sprintf("type of entry (valid types: %s)", ret)
}
