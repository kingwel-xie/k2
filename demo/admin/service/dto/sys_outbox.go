package dto

import (
    "github.com/kingwel-xie/k2/common/dto"
)

type SysOutboxGetPageReq struct {
    dto.Pagination     `search:"-"`
    Sender string `form:"sender"  search:"type:exact;column:sender;table:sys_outbox" comment:"发件人"`
    Receiver string `form:"receiver"  search:"type:exact;column:receiver;table:sys_outbox" comment:"收件人"`
    IsDraft bool `form:"isDraft"  search:"type:exact;column:is_draft;table:sys_outbox" comment:"草稿"`
    BeginTime      string `form:"beginTime" search:"type:gte;column:createdAt;table:sys_outbox" comment:"起始时间"`
    EndTime        string `form:"endTime" search:"type:lte;column:createdAt;table:sys_outbox" comment:"截止时间"`
    SysOutboxOrder
}

type SysOutboxOrder struct {
    IdOrder         string `search:"type:order;column:id;table:sys_outbox" form:"idOrder"`
}

func (m *SysOutboxGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysOutboxGetReq 功能获取请求参数
type SysOutboxGetReq struct {
     Id int `uri:"id"`
}

func (s *SysOutboxGetReq) GetId() interface{} {
	return s.Id
}

// SysOutboxDeleteReq 功能删除请求参数
type SysOutboxDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysOutboxDeleteReq) GetId() interface{} {
	return s.Ids
}
