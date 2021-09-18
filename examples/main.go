package main

import (
	"github.com/kingwel-xie/k2/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title Kobh Web Framework
// @version 1.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description KOBH
// @license.name MIT
// @license.url https://github.com

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Init("myApp")
	cmd.Execute()
}
