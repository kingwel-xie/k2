package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type TbxCountry struct {
    Code    string `json:"code" gorm:"unique;column:code;primaryKey;comment:编码"`
    Code2    string `json:"code2" gorm:"column:code2;comment:编码2"`
    NameCN    string `json:"nameCN" gorm:"size:64;unique;column:name_cn;comment:中文名称"`
    NameEN    string `json:"nameEN" gorm:"size:64;unique;column:name_en;comment:English Name"`
    Alias    string `json:"alias" gorm:"column:alias;comment:描述"`
    models.ControlBy
    models.ModelTime
}

func (TbxCountry) TableName() string {
    return "tbx_country"
}
