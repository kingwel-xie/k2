package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/kingwel-xie/k2/common/api"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)

type SysOutbox struct {
	api.Api
}

// GetPage 获取已发送消息列表
// @Summary 获取已发送消息列表
// @Description 获取已发送消息列表
// @Tags 已发送消息
// @Param sender query string false "发件人"
// @Param receiver query string false "收件人"
// @Param isDraft query bool false "草稿标志"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysOutbox}} "{"code": 200, "data": [...]}"
// @Router /api/v1/message [get]
// @Security Bearer
func (e SysOutbox) GetPage(c *gin.Context) {
    s := service.SysOutbox{}
    req := dto.SysOutboxGetPageReq{}
    err := e.MakeContext(c).
        Bind(&req, binding.Form).
        MakeService(&s.Service).
        Errors
   	if err != nil {
        e.Error(err)
        return
   	}

	list := make([]models.SysOutbox, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
        e.Error(err)
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取已发送消息
// @Summary 获取已发送消息
// @Description 获取已发送消息
// @Tags 已发送消息
// @Param id path string false "ID"
// @Success 200 {object} response.Response{data=models.SysOutbox} "{"code": 200, "data": [...]}"
// @Router /api/v1/message/{id} [get]
// @Security Bearer
func (e SysOutbox) Get(c *gin.Context) {
	s := service.SysOutbox{}
	req := dto.SysOutboxGetReq{}
    err := e.MakeContext(c).
        Bind(&req, nil).
        MakeService(&s.Service).
        Errors
	if err != nil {
        e.Error(err)
        return
	}

	var object models.SysOutbox
	err = s.Get(&req, &object)
	if err != nil {
        e.Error(err)
        return
	}

	e.OK(object, "查询成功")
}

// Delete 删除已发送消息
// @Summary 删除已发送消息
// @Description 删除已发送消息
// @Tags 已发送消息
// @Param data body dto.SysOutboxDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功", "data": [...]}"
// @Router /api/v1/message [delete]
// @Security Bearer
func (e SysOutbox) Delete(c *gin.Context) {
    s := service.SysOutbox{}
    req := dto.SysOutboxDeleteReq{}
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
