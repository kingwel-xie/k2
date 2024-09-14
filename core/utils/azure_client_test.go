package utils

import (
	"testing"
)

func TestAcquireToken(t *testing.T) {
	tenantId := "704e3ba9-47c5-4577-93f8-012f7a9c910f"
	clientId := "9c6528fc-392c-435e-9005-ab0576a323d0"

	ac := NewAzureLightClientWithDefaultHttp(tenantId, clientId)
	username := "ss@kingwelxie139.onmicrosoft.com"
	password := "Logu1896"

	token, err := ac.AcquireTokenByUsernamePassword(username, password, "User.Read")
	if err != nil {
		t.Fatal(err)
	}

	u, err := ac.Me(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u.UserPrincipalName)
}
