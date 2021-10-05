package main

import (
	"github.com/kingwel-xie/k2/cmd"
	"github.com/kingwel-xie/k2/cmd/migrate"

	"admin/entry"
	_ "admin/migrate/version"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title admin Web Backend
// @version 1.0.0
// @description admin Web Backend
// @description admin
// @license.name MIT
// @license.url https://github.com/...

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Init("admin", entry.StartCmd, migrate.StartCmd)
	cmd.Execute()
}
