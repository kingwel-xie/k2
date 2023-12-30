package dto

import (
	"github.com/kingwel-xie/k2/common/dto"
	"admin/models"
)

type SysNotificationGetPageReq struct {
    dto.Pagination     `search:"-"`
    TargetType string `form:"targetType"  search:"type:exact;column:target_type;table:sys_notification" comment:"接收人类别"`
    Targets string `form:"targets"  search:"type:exact;column:targets;table:sys_notification" comment:"接收人"`
    Title string `form:"title"  search:"type:contains;column:title;table:sys_notification" comment:"标题"`
    Content string `form:"content"  search:"type:contains;column:content;table:sys_notification" comment:"内容"`
    Importance string `form:"importance"  search:"type:exact;column:importance;table:sys_notification" comment:"重要"`
    SysNotificationOrder
}

type SysNotificationOrder struct {
	IdOrder         string `search:"type:order;column:id;table:sys_notification" form:"idOrder"`
}

func (m *SysNotificationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysNotificationInsertReq struct {
    Id int `json:"-" comment:"ID"`
    TargetType string `json:"targetType" comment:"接收人类别"`
    Targets string `json:"targets" comment:"接收人"`
    Title string `json:"title" comment:"标题"`
    Content string `json:"content" comment:"内容"`
    Importance string `json:"importance" comment:"重要"`
    Remark string `json:"remark" comment:"备注"`
}

func (s *SysNotificationInsertReq) Generate(model *models.SysNotification)  {
    model.TargetType = s.TargetType
    model.Targets = s.Targets
    model.Title = s.Title
    model.Content = s.Content
    model.Importance = s.Importance
    model.Remark = s.Remark
}

func (s *SysNotificationInsertReq) GetId() interface{} {
	return s.Id
}

type SysNotificationUpdateReq struct {
    Id int `uri:"id" comment:"ID"`
    Targets string `json:"targets" comment:"接收人"`
    Title string `json:"title" comment:"标题"`
    Content string `json:"content" comment:"内容"`
    Importance string `json:"importance" comment:"重要"`
    Remark string `json:"remark" comment:"备注"`
}

func (s *SysNotificationUpdateReq) Generate(model *models.SysNotification)  {
    model.Id = s.Id
    model.Targets = s.Targets
    model.Title = s.Title
    model.Content = s.Content
    model.Importance = s.Importance
    model.Remark = s.Remark
}

func (s *SysNotificationUpdateReq) GetId() interface{} {
	return s.Id
}

// SysNotificationGetReq 功能获取请求参数
type SysNotificationGetReq struct {
     Id int `uri:"id"`
}

func (s *SysNotificationGetReq) GetId() interface{} {
	return s.Id
}

// SysNotificationDeleteReq 功能删除请求参数
type SysNotificationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysNotificationDeleteReq) GetId() interface{} {
	return s.Ids
}
