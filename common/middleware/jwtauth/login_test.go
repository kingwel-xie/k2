package jwtauth

import (
	"github.com/kingwel-xie/k2/common/config"
	"testing"
)

func init() {
	config.EntraConfig.TenantId = "704e3ba9-47c5-4577-93f8-012f7a9c910f"
	config.EntraConfig.ClientId = "9c6528fc-392c-435e-9005-ab0576a323d0"
}

func TestLoginEntraId(t *testing.T) {
	login := Login{
		Username: "ss@kingwelxie139.onmicrosoft.com",
		Password: "*L9^5S-gB('SPMj",
		Code:     "",
		UUID:     "",
		Role:     "",
	}

	login.getUserEntraId(nil)
}
