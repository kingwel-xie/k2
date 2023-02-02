package apis

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kingwel-xie/k2/common/api"
	cerr "github.com/kingwel-xie/k2/common/error"
	"io/ioutil"
	"net/http"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)


var (
	passwordModificationFailedErr = cerr.New(403, "密码修改失败", "password modification failed")
	loginFailedErr = cerr.New(403, "登录失败", "login failed")
)

type SysUser struct {
	api.Api
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [get]
// @Security Bearer
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.SysUser, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int false "用户编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [get]
// @Security Bearer
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	var object models.SysUser
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserInsertReq true "用户数据"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [post]
// @Security Bearer
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.Insert(&req)
	if err != nil {
		e.Error(err)
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param userId path int true "用户编码"
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [put]
// @Security Bearer
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.Update(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [delete]
// @Security Bearer
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// CheckExistence
// @Summary 检查用户存在
// @Description 检查用户存在
// @Tags 用户
// @Param username path string true "username"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/check-existence/{username} [get]
// @Security Bearer
func (e SysUser) CheckExistence(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserCheckExistenceReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.CheckExistence(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "成功")
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/status [put]
// @Security Bearer
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.UpdateStatus(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/pwd/reset [put]
// @Security Bearer
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.ResetPwd(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// UpdatePwd
// @Summary 重置密码
// @Description 获取JSON
// @Tags 个人中心
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUser) UpdatePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.PassWord{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.UpdatePwd(s.Identity.UserId, req.OldPassword, req.NewPassword)
	if err != nil {
		e.Error(passwordModificationFailedErr.Wrap(err))
		return
	}
	e.OK(nil, "密码修改成功")
}

// UpdateProfile
// @Summary 修改个人中心
// @Description 获取JSON
// @Tags 个人中心
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [post]
// @Security Bearer
func (e SysUser) UpdateProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateProfileReq{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.UpdateProfile(s.Identity.UserId, &req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(nil, "修改成功")
}

// InsetAvatar
// @Summary 修改头像
// @Description 获取JSON
// @Tags 个人中心
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/avatar [post]
// @Security Bearer
func (e SysUser) InsetAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	file, err := c.FormFile("upload[]")
	if err != nil {
		e.Error(err)
		return
	}
	f, err := file.Open()
	if err != nil {
		e.Error(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		e.Error(err)
	}
	contentType := http.DetectContentType(buf)

	req.UserId = s.Identity.UserId
	req.Avatar = fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(buf))

	err = s.UpdateAvatar(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.Avatar, "修改成功")
}

// GetProfile
// @Summary 获取个人中心
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security Bearer
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	req.Id = s.Identity.UserId

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = s.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	}, "查询成功")
}

// GetInfo
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/getinfo [get]
// @Security Bearer
func (e SysUser) GetInfo(c *gin.Context) {
	req := dto.SysUserById{}
	s := service.SysUser{}
	r := service.SysRole{}
	err := e.MakeContext(c).
		MakeService(&r.Service).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	var identity = s.Identity
	var roles = make([]string, 1)
	roles[0] = identity.RoleKey
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if identity.RoleKey == "admin" {
		mp["permissions"] = permissions
	} else {
		list, _ := r.GetById(identity.RoleId)
		mp["permissions"] = list
	}
	sysUser, err := s.GetCurrentUser()
	if err != nil {
		e.Error(loginFailedErr.Wrap(err))
		return
	}
	mp["introduction"] = "If I die before I wake"
	mp["avatar"] = sysUser.Avatar
	mp["userName"] = sysUser.Username
	mp["userId"] = sysUser.UserId
	mp["deptId"] = sysUser.DeptId
	mp["name"] = sysUser.NickName
	e.OK(mp, "")
}
