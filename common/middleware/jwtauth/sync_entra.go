package jwtauth

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/config"
	msgraph "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

var deltaToken string

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

	cred, _ := azidentity.NewClientSecretCredential(
		config.EntraConfig.TenantId,
		config.EntraConfig.ClientId,
		config.EntraConfig.Mgmt.SecretKey,
		nil,
	)

	graphClient, _ := msgraph.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	delta, err := graphClient.Users().Delta().GetAsDeltaGetResponse(context.TODO(), nil)
	if err != nil {
		return err
	}
	// Use PageIterator to iterate through all users
	pageIterator, err := msgraphcore.NewPageIterator[models.Userable](delta, graphClient.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)

	var listEntra []*SysUser
	err = pageIterator.Iterate(context.Background(), func(user models.Userable) bool {
		var jobTitle string
		if user.GetJobTitle() != nil {
			jobTitle = *user.GetJobTitle()
		}
		roleId, ok := roleIdMap[jobTitle]
		if ok {
			u := toSysUser(user, roleId)
			listEntra = append(listEntra, &u)
		}
		return true
	})

	if pageIterator.GetOdataDeltaLink() != nil {
		deltaToken = *pageIterator.GetOdataDeltaLink()
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

	//requestSkiptoken := "oEBwdSP6uehIAxQOWq_3Ksh_TLol6KIm3stvdc6hGhZRi1hQ7Spe__dpvm3U4zReE4CYXC2zOtaKdi7KHlUtC2CbRiBIUwOxPKLa"
	//requestParameters := &msgraphusers.UsersDeltaWithRequestBuilderGetQueryParameters{
	//	Skiptoken: &requestSkiptoken,
	//}
	//configuration := &msgraphusers.UsersDeltaWithRequestBuilderGetRequestConfiguration{
	//	QueryParameters: requestParameters,
	//}

	// TODO: handle deltaToken
	//fmt.Println(delta.GetOdataDeltaLink(), delta.GetOdataNextLink())

	return nil
}

func toSysUser(user models.Userable, roleId int) SysUser {
	var nickname, phone, email string
	if user.GetDisplayName() != nil {
		nickname = *user.GetDisplayName()
	}
	if user.GetMobilePhone() != nil {
		phone = *user.GetMobilePhone()
	}
	if user.GetMail() != nil {
		email = *user.GetMail()
	}

	return SysUser{
		UserId:   0,
		Username: *user.GetUserPrincipalName(),
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
