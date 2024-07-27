package jwtauth

import (
	"context"
	"encoding/json"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"gorm.io/gorm"
	"io"
	"net/http"
)

func (u *Login) getRoleByKey(tx *gorm.DB, roleKey string) (role SysRole, err error) {
	return
}

func (u *Login) loginViaEntraId(tx *gorm.DB, cfg any) (user SysUser, role SysRole, err error) {
	// 创建 MSAL 公开客户端应用程序
	pca, err := public.New(
		"73f691ce-3596-437e-99ad-6f8de687603f",
		public.WithAuthority("https://login.microsoftonline.com/704e3ba9-47c5-4577-93f8-012f7a9c910f"),
	)
	if err != nil {
		err = wrapError(err)
		return
	}

	// 使用用户输入的用户名和密码获取访问令牌
	account, err := pca.AcquireTokenByUsernamePassword(context.Background(), []string{"user.read"}, u.Username, u.Password)
	if err != nil {
		err = wrapError(err)
		return
	}

	url := "https://graph.microsoft.com/v1.0/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = wrapError(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+account.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = wrapError(err)
		return
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)

	// 解析响应并输出用户个人资料信息
	var profile MsUserProfile
	err = json.Unmarshal(result, &profile)
	if err != nil {
		err = wrapError(err)
		return
	}

	// get SysRole
	role, err = u.getRoleByKey(tx, profile.JobTitle)
	if err != nil {
		err = wrapError(err)
		return
	}
	user = SysUser{
		UserId:   0,
		Username: profile.DisplayName,
		Password: "",
		NickName: profile.GivenName,
		Phone:    profile.MobilePhone,
		RoleId:   role.RoleId,
		Salt:     "",
		Avatar:   "",
		Sex:      "",
		Email:    profile.Mail,
		DeptId:   0,
		PostId:   0,
		Remark:   "",
		Status:   "2",
	}
	return

	// below commented code, using msgraph sdk
	//cred := NewSimpleCredentials(account.AccessToken)
	//client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})
	//if err != nil {
	//	fmt.Printf("Error creating client: %v\n", err)
	//	return
	//}
	//
	//result, err := client.Me().Drive().Get(context.Background(), nil)
	//if err != nil {
	//	fmt.Printf("Error getting the drive: %v\n", err)
	//	return
	//}
	//fmt.Printf("Found Drive : %v\n", *result.GetId())
}

type MsUserProfile struct {
	OdataContext      string        `json:"@odata.context"`
	BusinessPhones    []interface{} `json:"businessPhones"`
	DisplayName       string        `json:"displayName"`
	GivenName         string        `json:"givenName"`
	JobTitle          string        `json:"jobTitle"`
	Mail              string        `json:"mail"`
	MobilePhone       string        `json:"mobilePhone"`
	OfficeLocation    interface{}   `json:"officeLocation"`
	PreferredLanguage interface{}   `json:"preferredLanguage"`
	Surname           string        `json:"surname"`
	UserPrincipalName string        `json:"userPrincipalName"`
	Id                string        `json:"id"`
}

type T struct {
	OdataContext      string        `json:"@odata.context"`
	BusinessPhones    []interface{} `json:"businessPhones"`
	DisplayName       string        `json:"displayName"`
	GivenName         string        `json:"givenName"`
	JobTitle          string        `json:"jobTitle"`
	Mail              string        `json:"mail"`
	MobilePhone       string        `json:"mobilePhone"`
	OfficeLocation    string        `json:"officeLocation"`
	PreferredLanguage interface{}   `json:"preferredLanguage"`
	Surname           string        `json:"surname"`
	UserPrincipalName string        `json:"userPrincipalName"`
	Id                string        `json:"id"`
}
