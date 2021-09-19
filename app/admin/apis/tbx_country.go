package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/kingwel-xie/k2/app/admin/models"
	"github.com/kingwel-xie/k2/app/admin/service"
	"github.com/kingwel-xie/k2/app/admin/service/dto"
	"github.com/kingwel-xie/k2/common/api"
)

type TbxCountry struct {
	api.Api
}

// GetPage 获取国家编码列表
// @Summary 获取国家编码列表
// @Description 获取国家编码列表
// @Tags 国家编码
// @Param code query string false "编码"
// @Param cName query string false "中文名称"
// @Param eName query string false "English Name"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbxCountry}} "{"code": 200, "data": [...]}"
// @Router /api/v1/country [get]
// @Security Bearer
func (e TbxCountry) GetPage(c *gin.Context) {
	s := service.TbxCountry{}
	req := dto.TbxCountryGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	list := make([]models.TbxCountry, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取国家编码
// @Summary 获取国家编码
// @Description 获取国家编码
// @Tags 国家编码
// @Param code path string false "编码"
// @Success 200 {object} response.Response{data=models.TbxCountry} "{"code": 200, "data": [...]}"
// @Router /api/v1/country/{code} [get]
// @Security Bearer
func (e TbxCountry) Get(c *gin.Context) {
	s := service.TbxCountry{}
	req := dto.TbxCountryGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	var object models.TbxCountry
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建国家编码
// @Summary 创建国家编码
// @Description 创建国家编码
// @Tags 国家编码
// @Accept application/json
// @Product application/json
// @Param data body dto.TbxCountryInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "创建成功", "data": code}"
// @Router /api/v1/country [post]
// @Security Bearer
func (e TbxCountry) Insert(c *gin.Context) {
	s := service.TbxCountry{}
	req := dto.TbxCountryInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改国家编码
// @Summary 修改国家编码
// @Description 修改国家编码
// @Tags 国家编码
// @Accept application/json
// @Product application/json
// @Param code path string true "编码"
// @Param data body dto.TbxCountryUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功", "data": code}"
// @Router /api/v1/country/{code} [put]
// @Security Bearer
func (e TbxCountry) Update(c *gin.Context) {
	s := service.TbxCountry{}
	req := dto.TbxCountryUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除国家编码
// @Summary 删除国家编码
// @Description 删除国家编码
// @Tags 国家编码
// @Param data body dto.TbxCountryDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功", "data": [...]}"
// @Router /api/v1/country [delete]
// @Security Bearer
func (e TbxCountry) Delete(c *gin.Context) {
	s := service.TbxCountry{}
	req := dto.TbxCountryDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
