package app

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/kingwel-xie/k2/core/utils"
	"github.com/spf13/cobra"
	"text/template"
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
	path := "app/"
	appPath := path + appName
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

	t1, err := template.ParseFiles("template/cmd_api.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["appName"] = appName
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	utils.FileCreate(b1, "./cmd/"+appName+"server.go")
	t2, err := template.ParseFiles("template/router.template")
	var b2 bytes.Buffer
	err = t2.Execute(&b2, nil)
	utils.FileCreate(b2, appPath+"/router/router.go")
	return nil
}