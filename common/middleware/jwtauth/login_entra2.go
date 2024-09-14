package jwtauth

import (
	"fmt"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"
	"strings"
)

func wrapError(err any) error {
	return fmt.Errorf("failed to login to Entra, %v", err)
}

func (u *Login) getUserEntraId(tx *gorm.DB) (user SysUser, role SysRole, err error) {

	ac := utils.NewAzureLightClientWithDefaultHttp(config.EntraConfig.TenantId, config.EntraConfig.ClientId)

	token := u.Token
	if len(token) == 0 {
		if !strings.Contains(u.Username, "@") {
			u.Username = u.Username + "@" + config.EntraConfig.Realm
		}
		token, err = ac.AcquireTokenByUsernamePassword(u.Username, u.Password, "User.Read")
		if err != nil {
			err = wrapError(err)
			return
		}
	}

	account, err := ac.Me(token)
	if err != nil {
		err = wrapError(err)
		return
	}

	roleKey := account.JobTitle
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
