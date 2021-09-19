package app

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/kingwel-xie/k2/core/utils"
	k2template "github.com/kingwel-xie/k2/template"
	"github.com/spf13/cobra"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:     "create-app",
		Short:   "Create a new app",
		Long:    "Use when you need to create a new app",
		Example: "create-app -n myapp",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "", "Start server with provided configuration file")
}

func run() {
	fmt.Println(`Generating skeletons...`)
	if err := genFile(); err != nil {
		fmt.Println(`Generating failed`, err)
	}
}

func genFile() error {
	if appName == "" {
		return errors.New("arg `name` invalid ：name is empty")
	}
	appPath := appName
	err := utils.IsNotExistMkDir(appPath)
	if err != nil {
		return err
	}
	apiPath := appPath + "/apis/"
	err = utils.IsNotExistMkDir(apiPath)
	if err != nil {
		return err
	}
	modelsPath := appPath + "/models/"
	err = utils.IsNotExistMkDir(modelsPath)
	if err != nil {
		return err
	}
	routerPath := appPath + "/router/"
	err = utils.IsNotExistMkDir(routerPath)
	if err != nil {
		return err
	}
	servicePath := appPath + "/service/"
	err = utils.IsNotExistMkDir(servicePath)
	if err != nil {
		return err
	}
	dtoPath := appPath + "/service/dto/"
	err = utils.IsNotExistMkDir(dtoPath)
	if err != nil {
		return err
	}
	entryPath := appPath + "/entry/"
	err = utils.IsNotExistMkDir(entryPath)
	if err != nil {
		return err
	}
	docsPath := appPath + "/docs/"
	err = utils.IsNotExistMkDir(docsPath)
	if err != nil {
		return err
	}
	configPath := appPath + "/config/"
	err = utils.IsNotExistMkDir(configPath)
	if err != nil {
		return err
	}

	c1, err := k2template.Asset("cmd_api.template")
	if err != nil {
		return err
	}
	t1, err := template.New("cmd_api.template").Parse(string(c1))
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["appName"] = appName
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	utils.FileCreate(b1, entryPath+"server.go")

	c2, err := k2template.Asset("router.template")
	if err != nil {
		return err
	}
	t2, err := template.New("router.template").Parse(string(c2))
	var b2 bytes.Buffer
	err = t2.Execute(&b2, nil)
	utils.FileCreate(b2, routerPath+"router.go")

	c3, err := k2template.Asset("config.template")
	if err != nil {
		return err
	}
	t3, err := template.New("config.template").Parse(string(c3))
	var b3 bytes.Buffer
	err = t3.Execute(&b3, nil)
	utils.FileCreate(b3, configPath+"extend.go")
	return nil
}
