package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type SysConfig struct {
	models.Model
	ConfigName  string `json:"configName" gorm:"size:128;comment:ConfigName"`   //
	ConfigKey   string `json:"configKey" gorm:"size:128;comment:ConfigKey"`     //
	ConfigValue string `json:"configValue" gorm:"size:255;comment:ConfigValue"` //
	ConfigType  string `json:"configType" gorm:"size:64;comment:ConfigType"`
	IsFrontend  int    `json:"isFrontend" gorm:"size:64;comment:是否前台"` //
	Remark      string `json:"remark" gorm:"size:128;comment:Remark"`  //
	models.ControlBy
	models.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}
