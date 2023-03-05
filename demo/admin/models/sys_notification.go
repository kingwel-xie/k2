package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

const (
    TargetTypeAll = "all"
    TargetTypeRole = "role"
    TargetTypeUser = "user"
    TargetTypeDept = "dept"
)

type SysNotification struct {
    Id    int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
    TargetType    string `json:"targetType" gorm:"size:63;column:target_type;comment:接收人类别"`
    Targets    string `json:"targets" gorm:"size:63;column:targets;comment:接收人"`
    Title    string `json:"title" gorm:"size:127;column:title;comment:标题"`
    Content    string `json:"content" gorm:"size:511;column:content;comment:内容"`
    Remark    string `json:"remark" gorm:"column:remark;comment:备注"`
    models.ControlBy
    models.ModelTimeHardDelete
}

func (SysNotification) TableName() string {
    return "sys_notification"
}
