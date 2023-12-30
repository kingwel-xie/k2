package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/core/utils"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)

type SysDept struct {
	api.Api
}

// GetPage
// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param deptName query string false "deptName"
// @Param deptId query string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept [get]
// @Security Bearer
func (e SysDept) GetPage(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	list := make([]models.SysDept, 0)
	list, err = s.SetDeptPage(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(list, "查询成功")
}

// Get
// @Summary 获取部门数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func (e SysDept) Get(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	var object models.SysDept
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(err)
		return
	}

	e.OK(object, "查询成功")
}

// Insert 添加部门
// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDeptInsertReq true "data"
// @Success 200 {string} string	"{"code": 200, "message": "创建成功"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func (e SysDept) Insert(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptInsertReq{}
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
// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param deptId path int true "deptId"
// @Param data body dto.SysDeptUpdateReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/dept/{deptId} [put]
// @Security Bearer
func (e SysDept) Update(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
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
// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param data body dto.SysDeptDeleteReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/dept [delete]
// @Security Bearer
func (e SysDept) Delete(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
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

// Get2Tree 查询部门树
// @Summary 查询部门树
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param deptName query string false "deptName"
// @Param deptId query string false "deptId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/deptTree [get]
// @Security Bearer
func (e SysDept) Get2Tree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}
	list := make([]dto.DeptLabel, 0)
	list, err = s.SetDeptTree(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(list, "")
}

// GetDeptTreeRoleSelect 根据角色ID查询部门结构 TODO: 此接口需要调整不应该将list和选中放在一起
// @Summary 根据角色ID查询部门结构
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param roleId path int true "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/roleDeptTreeSelect/{roleId} [get]
// @Security Bearer
func (e SysDept) GetDeptTreeRoleSelect(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	id, err := utils.StringToInt(c.Param("roleId"))
	result, err := s.SetDeptLabel()
	if err != nil {
		e.Error(err)
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = s.GetWithRoleId(id)
		if err != nil {
			e.Error(err)
			return
		}
	}
	e.OK(gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	}, "")
}

// ListNoCheck
// @Summary 部门列表数据
// @Description Get JSON
// @Tags 部门
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-list [get]
// @Security Bearer
func (e SysDept) ListNoCheck(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	list := make([]models.SysDept, 0)

	err = s.ListNoCheck(&req, &list)
	if err != nil {
		e.Error(err)
		return
	}

	e.PageOK(list, len(list), 1, len(list), "查询成功")
}