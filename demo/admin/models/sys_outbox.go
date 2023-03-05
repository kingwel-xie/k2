package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type SysOutbox struct {
    Id    int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
    Sender    string `json:"sender" gorm:"size:63;index;column:sender;comment:发件人"`
    Receivers    string `json:"receivers" gorm:"column:receiver;comment:收件人"`
    OriginId    int `json:"originId" gorm:"index;column:origin_id;comment:起始Id"`
    Title    string `json:"title" gorm:"size:127;column:title;comment:标题"`
    Content    string `json:"content" gorm:"size:511;column:content;comment:内容"`
    IsDraft    bool `json:"isDraft" gorm:"default:false;column:is_draft;comment:草稿"`
    models.ControlBy
    models.ModelTime
}

func (SysOutbox) TableName() string {
    return "sys_outbox"
}
