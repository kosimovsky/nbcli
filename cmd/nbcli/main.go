// Package main
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
	"log"
	"netbox-client/cmd"
	"os"
)

func main() {
	logfile, err := os.OpenFile(cmd.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error occured opening file : %v, %s", err, cmd.Logfile)
	}
	defer logfile.Close()
	logrus.New()
	logrus.SetOutput(logfile)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err = doc.GenMarkdownTree(cmd.RootCmd, "docs")
	if err != nil {
		logrus.Errorf("Error generating docs : %s", err.Error())
	}

	cmd.Execute()
}
