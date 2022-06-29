// Package cmd
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile: configuration file that stores information about the Netbox server
var cfgFile string

// outputFormat: output format (json, yaml)
var outputFormat string

const Logfile = "netbox-client.log"

//TODO

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nbcli",
	Short: "A brief description for client",
	Long:  `A longer description for client`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is configs/config.yaml)")
	RootCmd.PersistentFlags().StringVar(&outputFormat, "out", "yaml", `output format [yaml, json]`)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
