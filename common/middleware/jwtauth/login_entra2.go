package jwtauth

import (
	"context"
	"fmt"
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
	if !strings.Contains(u.Username, "@") {
		u.Username = u.Username + "@" + config.EntraConfig.Realm
	}
	cred, _ := azidentity.NewUsernamePasswordCredential(
		config.EntraConfig.TenantId,
		config.EntraConfig.ClientId,
		u.Username,
		u.Password,
		nil,
	)

	graphClient, _ := msgraph.NewGraphServiceClientWithCredentials(cred, []string{"User.Read"})
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
	err = tx.FirstOrCreate(&user, "username = ?", u.Username).Error
	if err != nil {
		err = wrapError("invalid sysUser: " + u.Username)
		return
	}
	return
}
