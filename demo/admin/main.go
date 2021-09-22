package main

import (
	"github.com/kingwel-xie/k2/cmd"
	"github.com/kingwel-xie/k2/cmd/migrate"

	"kx/entry"
	_ "kx/migrate/version"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title KX Web Backend
// @version 1.0.0
// @description KX Web Backend
// @description KX
// @license.name MIT
// @license.url https://github.com/...

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Init("kx", entry.StartCmd, migrate.StartCmd)
	cmd.Execute()
}
