package jwtauth

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kingwel-xie/k2/common/config"
	msgraph "github.com/microsoftgraph/msgraph-sdk-go"
	"gorm.io/gorm"
	"strings"
)

func wrapError(err any) error {
	return fmt.Errorf("failed to login to Entra, %v", err)
}

func (u *Login) getUserEntraId(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	var cred azcore.TokenCredential

	if len(u.Token) > 0 {
		cred = NewSimpleCredentials(u.Token)
	} else {
		if !strings.Contains(u.Username, "@") {
			u.Username = u.Username + "@" + config.EntraConfig.Realm
		}
		cred, err = azidentity.NewUsernamePasswordCredential(
			config.EntraConfig.TenantId,
			config.EntraConfig.ClientId,
			u.Username,
			u.Password,
			nil,
		)
		if err != nil {
			err = wrapError(err)
			return
		}
	}

	graphClient, err := msgraph.NewGraphServiceClientWithCredentials(cred, []string{"User.Read"})
	if err != nil {
		err = wrapError(err)
		return
	}
	account, err := graphClient.Me().Get(context.TODO(), nil)
	if err != nil {
		err = wrapError(err)
		return
	}

	roleKey := *account.GetJobTitle()
	err = tx.Where("role_key = ? ", roleKey).First(&role).Error
	if err != nil {
		err = wrapError("invalid roleKey/JobTitle")
		return
	}
	user = toSysUser(account, role.RoleId)
	err = tx.FirstOrCreate(&user, "username = ?", user.Username).Error
	if err != nil {
		err = wrapError("invalid sysUser: " + user.Username)
		return
	}
	return
}

func NewSimpleCredentials(token string) *SimpleCredentials {
	return &SimpleCredentials{
		TokenValue: token,
	}
}

type SimpleCredentials struct {
	TokenValue string
}

func (m *SimpleCredentials) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token: m.TokenValue,
	}, nil
}
