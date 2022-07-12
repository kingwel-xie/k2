package models

import (
	"github.com/kingwel-xie/k2/common/models"
)

type TbxCountry struct {
    Code    string `json:"code" gorm:"size:63;column:code;primaryKey;comment:代码"`
    Code2    string `json:"code2" gorm:"size:63;unique_index;column:code2;comment:三字符代码"`
    DigitCode    string `json:"digitCode" gorm:"size:63;column:digit_code;comment:数字代码"`
    TeleCode    string `json:"teleCode" gorm:"size:63;column:tele_code;comment:电话代码"`
    Group    string `json:"group" gorm:"size:63;column:group;comment:分组"`
    BelongTo    string `json:"belongTo" gorm:"size:63;column:belong_to;comment:从属"`
    NameCN    string `json:"nameCN" gorm:"size:63;unique_index;column:name_cn;comment:中文简称"`
    NameEN    string `json:"nameEN" gorm:"size:63;unique_index;column:name_en;comment:英文简称"`
    DisplaySort    int `json:"displaySort" gorm:"size:16;column:display_sort;comment:显示排序"`
    Remark    string `json:"remark" gorm:"column:remark;comment:描述"`
    Children   []TbxCountry `json:"children,omitempty" gorm:"-"`
    models.ControlBy
    models.ModelTimeHardDelete
}

func (TbxCountry) TableName() string {
    return "tbx_country"
}
