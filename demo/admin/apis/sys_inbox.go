package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/kingwel-xie/k2/common/api"

	"admin/models"
	"admin/service"
	"admin/service/dto"
)

type SysInbox struct {
	api.Api
}

// GetPage 获取消息列表
// @Summary 获取消息列表
// @Description 获取消息列表
// @Tags 消息
// @Param type query string false "消息类型"
// @Param sender query string false "发件人"
// @Param receiver query string false "收件人"
// @Param unread query bool false "未读标志"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysInbox}} "{"code": 200, "data": [...]}"
// @Router /api/v1/message [get]
// @Security Bearer
func (e SysInbox) GetPage(c *gin.Context) {
    s := service.SysInbox{}
    req := dto.SysInboxGetPageReq{}
    err := e.MakeContext(c).
        Bind(&req, binding.Form).
        MakeService(&s.Service).
        Errors
   	if err != nil {
        e.Error(err)
        return
   	}

	list := make([]models.SysInbox, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
        e.Error(err)
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取消息
// @Summary 获取消息
// @Description 获取消息
// @Tags 消息
// @Param id path string false "ID"
// @Success 200 {object} response.Response{data=models.SysInbox} "{"code": 200, "data": [...]}"
// @Router /api/v1/message/{id} [get]
// @Security Bearer
func (e SysInbox) Get(c *gin.Context) {
	s := service.SysInbox{}
	req := dto.SysInboxGetReq{}
    err := e.MakeContext(c).
        Bind(&req, nil).
        MakeService(&s.Service).
        Errors
	if err != nil {
        e.Error(err)
        return
	}

	var object models.SysInbox
	err = s.Get(&req, &object)
	if err != nil {
        e.Error(err)
        return
	}

	e.OK(object, "查询成功")
}

// Insert 创建消息
// @Summary 创建消息
// @Description 创建消息
// @Tags 消息
// @Accept application/json
// @Product application/json
// @Param data body dto.SysInboxInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "创建成功", "data": id}"
// @Router /api/v1/message [post]
// @Security Bearer
func (e SysInbox) Insert(c *gin.Context) {
    s := service.SysInbox{}
    req := dto.SysInboxInsertReq{}
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
	e.OK(nil, "创建成功")
}

// Update 修改消息
// @Summary 修改消息
// @Description 修改消息
// @Tags 消息
// @Accept application/json
// @Product application/json
// @Param id path string true "ID"
// @Param data body dto.SysInboxUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功", "data": id}"
// @Router /api/v1/message/{id} [put]
// @Security Bearer
func (e SysInbox) Update(c *gin.Context) {
    s := service.SysInbox{}
    req := dto.SysInboxUpdateReq{}
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

// Delete 删除消息
// @Summary 删除消息
// @Description 删除消息
// @Tags 消息
// @Param data body dto.SysInboxDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功", "data": [...]}"
// @Router /api/v1/message [delete]
// @Security Bearer
func (e SysInbox) Delete(c *gin.Context) {
    s := service.SysInbox{}
    req := dto.SysInboxDeleteReq{}
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

// GetUnread 获取未读消息
// @Summary 获取未读消息
// @Description 获取未读消息
// @Tags Inbox
// @Success 200 {object} response.Response{data=dto.SysUnread} "{"code": 200, "data": [...]}"
// @Router /api/v1/inbox/unread [get]
// @Security Bearer
func (e SysInbox) GetUnread(c *gin.Context) {
	s := service.SysInbox{}
	err := e.MakeContext(c).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	var object dto.SysUnread
	err = s.GetUnread(&object)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(object, "查询成功")
}

// MarkRead 标记消息已读
// @Summary 标记消息已读
// @Description 标记消息已读
// @Tags Inbox
// @Param data body dto.SysInboxMarkReadReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "标记成功", "data": [...]}"
// @Router /api/v1/inbox/read [post]
// @Security Bearer
func (e SysInbox) MarkRead(c *gin.Context) {
	s := service.SysInbox{}
	req := dto.SysInboxMarkReadReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	err = s.MarkRead(&req)
	if err != nil {
		e.Error(err)
		return
	}
	e.OK(req.GetId(), "标记成功")
}
