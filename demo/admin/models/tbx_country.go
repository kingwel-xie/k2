package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type TbxCountry struct {
    Code    string `json:"code" gorm:"unique;column:code;primaryKey;comment:编码"`
    NameCN    string `json:"nameCN" gorm:"size:64;unique;column:name_cn;comment:中文名"`
    NameEN    string `json:"nameEN" gorm:"size:64;unique;column:name_en;comment:英文名"`
    TeleCode    string `json:"teleCode" gorm:"column:tele_code;comment:电话代码"`
    Alias    string `json:"alias" gorm:"column:alias;comment:描述"`
    models.ControlBy
    models.ModelTimeHardDelete
}

func (TbxCountry) TableName() string {
    return "tbx_country"
}
