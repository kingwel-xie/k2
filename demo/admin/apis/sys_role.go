package apis

import (
	"fmt"
	"github.com/kingwel-xie/k2/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"admin/models"
	"admin/service"
	"admin/service/dto"
	"github.com/kingwel-xie/k2/common/api"
)

type SysRole struct {
	api.Api
}

// GetPage
// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色/Role
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [get]
// @Security Bearer
func (e SysRole) GetPage(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.SysRole, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [get]
// @Security Bearer
func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	var object models.SysRole
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(err)
		return
	}

	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [post]
// @Security Bearer
func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	if req.Status == "" {
		req.Status = "2"
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(err)
		return
	}
	// reload casbin policies
	err = common.Runtime.GetCasbin().LoadPolicy()
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户角色
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body dto.SysRoleUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [put]
// @Security Bearer
func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, nil, binding.JSON).
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
	// reload casbin policies
	err = common.Runtime.GetCasbin().LoadPolicy()
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色/Role
// @Param data body dto.SysRoleDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [delete]
// @Security Bearer
func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
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
	// reload casbin policies
	err = common.Runtime.GetCasbin().LoadPolicy()
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("删除角色 %v 状态成功！", req.GetId()))
}

// Update2Status 修改用户角色状态
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-status [put]
// @Security Bearer
func (e SysRole) Update2Status(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
	err := e.MakeContext(c).
		Bind(&req).
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
	e.OK(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

// Update2DataScope 更新角色数据权限
// @Summary 更新角色数据权限
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.RoleDataScopeReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-datascope [put]
// @Security Bearer
func (e SysRole) Update2DataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.UpdateDataScope(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(nil, "操作成功")
}

// GetAll
// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色/Role
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-list [get]
// @Security Bearer
func (e SysRole) GetAll(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.SysRole, 0)
	var count int64

	err = s.ListNoCheck(&req, &list, &count)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}