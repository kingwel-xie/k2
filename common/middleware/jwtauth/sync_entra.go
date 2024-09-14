package jwtauth

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/core/utils"
)

func SyncUpEntra() error {
	// start a cron job to sync up EntraId, every day at 0:01am
	_ = common.Runtime.GetCrontab().AddJob("syncEntraId", "0 15 21 * * ?", func() {
		_ = syncUpEntra()
	})
	return syncUpEntra()
}

func syncUpEntra() error {
	// load SysUser
	db := common.Runtime.GetDb()

	var listLocal []*SysUser
	err := db.Find(&listLocal).Error
	if err != nil {
		return err
	}
	// construct rolsIdMap: roleKey => roleId
	var roles []*SysRole
	var roleIdMap = make(map[string]int)
	err = db.Find(&roles).Error
	if err != nil {
		return err
	}
	for _, r := range roles {
		roleIdMap[r.RoleKey] = r.RoleId
	}

	ac := utils.NewAzureLightClientWithDefaultHttp(config.EntraConfig.TenantId, config.EntraConfig.ClientId)

	var listEntra []*SysUser
	err = ac.EnumUsersDelta(config.EntraConfig.Mgmt.SecretKey, func(user *utils.AzureUser) bool {
		var jobTitle string
		if len(user.JobTitle) > 0 {
			jobTitle = user.JobTitle
		}
		roleId, ok := roleIdMap[jobTitle]
		if ok {
			u := toSysUser(user, roleId)
			listEntra = append(listEntra, &u)
		}
		return true
	})
	if err != nil {
		return err
	}

	// cross-reference between listLocal & listEntra
	// suppose listEntra is always addedOrUpdated
	for _, u := range listEntra {
		// lookup local list to see if the username is already there
		// if so, copy the primary key - meaning updated, otherwise, leave it as 0 - newly added
		for _, u2 := range listLocal {
			if u.Username == u2.Username {
				u.UserId = u2.UserId
				break
			}
		}
	}
	// we now have the processed listEntra, save it
	err = db.Save(&listEntra).Error
	if err != nil {
		return err
	}

	return nil
}

func toSysUser(user *utils.AzureUser, roleId int) SysUser {
	var nickname, phone, email string
	if len(user.DisplayName) > 0 {
		nickname = user.DisplayName
	}
	if len(user.MobilePhone) > 0 {
		phone = user.MobilePhone
	}
	if len(user.Mail) > 0 {
		email = user.Mail
	}

	return SysUser{
		UserId:   0,
		Username: user.UserPrincipalName,
		NickName: nickname,
		RoleId:   roleId,
		Phone:    phone,
		Avatar:   "",
		Sex:      "",
		Email:    email,
		DeptId:   0,
		PostId:   0,
		Remark:   "",
		Status:   "2",
	}
}
