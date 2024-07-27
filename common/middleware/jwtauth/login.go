package jwtauth

import (
	"errors"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"
)

var ErrMismatchRoleKey = errors.New("mismatch role key")

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
	Role     string `form:"Role" json:"role"`
}

func (u *Login) getUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {

	userTableName := "sys_user"
	roleTableName := "sys_role"
	err = tx.Table(userTableName).Where("username = ?  and status = 2", u.Username).First(&user).Error
	if err != nil {
		return
	}
	_, err = utils.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		return
	}
	err = tx.Table(roleTableName).Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		return
	}
	// check roleKey when Role exists
	if len(u.Role) > 0 && u.Role != role.RoleKey {
		err = ErrMismatchRoleKey
		return
	}
	return
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	// when Entra is enabled
	if config.EntraConfig.Enable {
		return u.getUserEntraId(tx)
	} else {
		return u.getUser(tx)
	}
}
