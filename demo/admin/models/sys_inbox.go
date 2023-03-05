package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type SysInbox struct {
    Id    int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
    Type    string `json:"type" gorm:"size:63;column:type;comment:消息类型"`
    Sender    string `json:"sender" gorm:"size:63;index;column:sender;comment:发件人"`
    Receiver    string `json:"receiver" gorm:"size:63;index;column:receiver;comment:收件人"`
    OriginId    int `json:"originId" gorm:"index;column:origin_id;comment:起始Id"`
    Title    string `json:"title" gorm:"size:127;column:title;comment:标题"`
    Content    string `json:"content" gorm:"size:511;column:content;comment:内容"`
    Read    bool `json:"read" gorm:"default:false;column:read;comment:已读标志"`
    models.ControlBy
    models.ModelTime
}

func (SysInbox) TableName() string {
    return "sys_inbox"
}
