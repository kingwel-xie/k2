package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/kingwel-xie/k2/common/api"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)

type SysNotification struct {
	api.Api
}

// GetPage 获取通知列表
// @Summary 获取通知列表
// @Description 获取通知列表
// @Tags 通知
// @Param targetType query string false "接收人类别"
// @Param targets query string false "接收人"
// @Param title query string false "标题"
// @Param content query string false "内容"
// @Param importance query string false "重要"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysNotification}} "{"code": 200, "data": [...]}"
// @Router /api/v1/notification [get]
// @Security Bearer
func (e SysNotification) GetPage(c *gin.Context) {
    s := service.SysNotification{}
    req := dto.SysNotificationGetPageReq{}
    err := e.MakeContext(c).
        Bind(&req, binding.Form).
        MakeService(&s.Service).
        Errors
   	if err != nil {
        e.Error(err)
        return
   	}

	list := make([]models.SysNotification, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
        e.Error(err)
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取通知
// @Summary 获取通知
// @Description 获取通知
// @Tags 通知
// @Param id path string false "ID"
// @Success 200 {object} response.Response{data=models.SysNotification} "{"code": 200, "data": [...]}"
// @Router /api/v1/notification/{id} [get]
// @Security Bearer
func (e SysNotification) Get(c *gin.Context) {
	s := service.SysNotification{}
	req := dto.SysNotificationGetReq{}
    err := e.MakeContext(c).
        Bind(&req, nil).
        MakeService(&s.Service).
        Errors
	if err != nil {
        e.Error(err)
        return
	}

	var object models.SysNotification
	err = s.Get(&req, &object)
	if err != nil {
        e.Error(err)
        return
	}

	e.OK(object, "查询成功")
}

// Insert 创建通知
// @Summary 创建通知
// @Description 创建通知
// @Tags 通知
// @Accept application/json
// @Product application/json
// @Param data body dto.SysNotificationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "创建成功", "data": id}"
// @Router /api/v1/notification [post]
// @Security Bearer
func (e SysNotification) Insert(c *gin.Context) {
    s := service.SysNotification{}
    req := dto.SysNotificationInsertReq{}
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

// Update 修改通知
// @Summary 修改通知
// @Description 修改通知
// @Tags 通知
// @Accept application/json
// @Product application/json
// @Param id path string true "ID"
// @Param data body dto.SysNotificationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功", "data": id}"
// @Router /api/v1/notification/{id} [put]
// @Security Bearer
func (e SysNotification) Update(c *gin.Context) {
    s := service.SysNotification{}
    req := dto.SysNotificationUpdateReq{}
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

// Delete 删除通知
// @Summary 删除通知
// @Description 删除通知
// @Tags 通知
// @Param data body dto.SysNotificationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功", "data": [...]}"
// @Router /api/v1/notification [delete]
// @Security Bearer
func (e SysNotification) Delete(c *gin.Context) {
    s := service.SysNotification{}
    req := dto.SysNotificationDeleteReq{}
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
