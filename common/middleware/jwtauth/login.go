package jwtauth

import (
	"gorm.io/gorm"
	"github.com/kingwel-xie/k2/core/utils"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

func (u *Login) GetUser(tx *gorm.DB, prefix string) (user SysUser, role SysRole, err error) {
	userTableName := prefix + "sys_user"
	roleTableName := prefix + "sys_role"
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
	return
}
