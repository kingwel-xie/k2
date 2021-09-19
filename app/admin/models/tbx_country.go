package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type TbxCountry struct {
	Code   string `json:"code" gorm:"size:4;column:code;primaryKey;comment:编码"`
	CName  string `json:"cName" gorm:"size:64;unique;column:c_name;comment:中文名称"`
	EName  string `json:"eName" gorm:"size:64;unique;column:e_name;comment:English Name"`
	EName2 string `json:"eName2" gorm:"size:64;column:e_name2;comment:English Name 2"`
	Code2  string `json:"code2" gorm:"size:4;unique;column:code2;comment:编码2"`
	Code3  string `json:"code3" gorm:"size:4;unique;column:code3;comment:编码3"`
	models.ControlBy
	models.ModelTime
}

func (TbxCountry) TableName() string {
	return "tbx_country"
}
