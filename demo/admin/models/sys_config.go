package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type SysConfig struct {
	models.Model
	ConfigName  string `json:"configName" gorm:"comment:ConfigName"`   //
	ConfigKey   string `json:"configKey" gorm:"comment:ConfigKey"`     //
	ConfigValue string `json:"configValue" gorm:"comment:ConfigValue"` //
	ConfigType  string `json:"configType" gorm:"size:64;comment:ConfigType"`
	IsFrontend  int    `json:"isFrontend" gorm:"size:64;comment:是否前台"` //
	Remark      string `json:"remark" gorm:"comment:Remark"`  //
	models.ControlBy
	models.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}
