package dto

import (
	"github.com/kingwel-xie/k2/common/dto"

	"admin/models"
)

type SysInboxGetPageReq struct {
    dto.Pagination     `search:"-"`
    Type string `form:"type"  search:"type:exact;column:type;table:sys_inbox" comment:"消息类型"`
    Sender string `form:"sender"  search:"type:exact;column:sender;table:sys_inbox" comment:"发件人"`
    Receiver string `form:"receiver"  search:"type:exact;column:receiver;table:sys_inbox" comment:"收件人"`
    Unread bool `form:"unread"  search:"type:not-exact;column:read;table:sys_inbox" comment:"未读标志"`
    BeginTime      string `form:"beginTime" search:"type:gte;column:createdAt;table:sys_inbox" comment:"起始时间"`
    EndTime        string `form:"endTime" search:"type:lte;column:createdAt;table:sys_inbox" comment:"截止时间"`
    SysInboxOrder
}

type SysInboxOrder struct {
    IdOrder         string `search:"type:order;column:id;table:sys_inbox" form:"idOrder"`
}

func (m *SysInboxGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysInboxInsertReq struct {
    Id int `json:"-" comment:"ID"`
    Receiver []string `json:"receiver" comment:"收件人"`
    OriginId int `json:"originId" comment:"起始Id"`
    Title string `json:"title" comment:"标题"`
    Content string `json:"content" comment:"内容"`
}

type SysInboxUpdateReq struct {
    Id int `uri:"id" comment:"ID"`
    Sender string `json:"sender" comment:"发件人"`
    Receiver string `json:"receiver" comment:"收件人"`
    Title string `json:"title" comment:"标题"`
    Content string `json:"content" comment:"内容"`
    Read bool `json:"read" comment:"已读标志"`
}

func (s *SysInboxUpdateReq) Generate(model *models.SysInbox)  {
    model.Id = s.Id
    model.Sender = s.Sender
    model.Receiver = s.Receiver
    model.Title = s.Title
    model.Content = s.Content
    model.Read = s.Read
}

func (s *SysInboxUpdateReq) GetId() interface{} {
	return s.Id
}

// SysInboxGetReq 功能获取请求参数
type SysInboxGetReq struct {
     Id int `uri:"id"`
}

func (s *SysInboxGetReq) GetId() interface{} {
	return s.Id
}

// SysInboxDeleteReq 功能删除请求参数
type SysInboxDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysInboxDeleteReq) GetId() interface{} {
	return s.Ids
}

type SysUnread struct {
    NumNotices int                  `json:"numNotices"`
    NoticeList *[]models.SysInbox `json:"noticeList"`
    NumMessages int                  `json:"numMessages"`
    MessageList *[]models.SysInbox `json:"messageList"`
}

// SysInboxMarkReadReq 标记已读请求参数
type SysInboxMarkReadReq struct {
    Ids []int `json:"ids"`
    Read bool `json:"read"`
}

func (s *SysInboxMarkReadReq) GetId() interface{} {
    return s.Ids
}

// SysInboxSendMessageReq 发送站内信请求参数
type SysInboxSendMessageReq struct {
    TargetType string `json:"targetType" comment:"接收人类别"`
    Targets string `json:"targets" comment:"接收人"`
    Title string `json:"title" comment:"标题"`
    Content string `json:"content" comment:"内容"`
}

