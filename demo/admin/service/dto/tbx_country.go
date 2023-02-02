package dto

import (
	"github.com/kingwel-xie/k2/common/dto"

	"admin/models"
)

type TbxCountryGetPageReq struct {
    dto.Pagination     `search:"-"`
    Code string `form:"code"  search:"type:exact;column:code;table:tbx_country" comment:"代码"`
    Code2 string `form:"code2"  search:"type:exact;column:code2;table:tbx_country" comment:"三字符代码"`
    DigitCode string `form:"digitCode"  search:"type:exact;column:digit_code;table:tbx_country" comment:"数字代码"`
    TeleCode string `form:"teleCode"  search:"type:exact;column:tele_code;table:tbx_country" comment:"电话代码"`
    Group string `form:"group"  search:"type:exact;column:group;table:tbx_country" comment:"分组"`
    NameCN string `form:"nameCN"  search:"type:contains;column:name_cn;table:tbx_country" comment:"中文简称"`
    NameEN string `form:"nameEN"  search:"type:contains;column:name_en;table:tbx_country" comment:"英文简称"`
    Remark string `form:"remark"  search:"type:contains;column:remark;table:tbx_country" comment:"描述"`
    TbxCountryOrder
}

type TbxCountryOrder struct {
}

func (m *TbxCountryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbxCountryInsertReq struct {
    Code string `json:"code" comment:"代码"`
    Code2 string `json:"code2" comment:"三字符代码"`
    DigitCode string `json:"digitCode" comment:"数字代码"`
    TeleCode string `json:"teleCode" comment:"电话代码"`
    Group string `json:"group" comment:"分组"`
    BelongTo string `json:"belongTo" comment:"从属"`
    NameCN string `json:"nameCN" comment:"中文简称" vd:"@:len($)>0; msg:'中文简称不能为空'"`
    NameEN string `json:"nameEN" comment:"英文简称" vd:"@:len($)>0; msg:'English Name is empty'"`
    DisplaySort int `json:"displaySort" comment:"显示排序"`
    Remark string `json:"remark" comment:"描述"`
}

func (s *TbxCountryInsertReq) Generate(model *models.TbxCountry)  {
    model.Code = s.Code
    model.Code2 = s.Code2
    model.DigitCode = s.DigitCode
    model.TeleCode = s.TeleCode
    model.Group = s.Group
    model.BelongTo = s.BelongTo
    model.NameCN = s.NameCN
    model.NameEN = s.NameEN
    model.DisplaySort = s.DisplaySort
    model.Remark = s.Remark
}

func (s *TbxCountryInsertReq) GetId() interface{} {
	return s.Code
}

type TbxCountryUpdateReq struct {
    Code string `uri:"code" comment:"代码"`
    Code2 string `json:"code2" comment:"三字符代码"`
    DigitCode string `json:"digitCode" comment:"数字代码"`
    TeleCode string `json:"teleCode" comment:"电话代码"`
    Group string `json:"group" comment:"分组"`
    BelongTo string `json:"belongTo" comment:"从属"`
    NameCN string `json:"nameCN" comment:"中文简称" vd:"@:len($)>0; msg:'中文简称不能为空'"`
    NameEN string `json:"nameEN" comment:"英文简称" vd:"@:len($)>0; msg:'English Name is empty'"`
    DisplaySort int `json:"displaySort" comment:"显示排序"`
    Remark string `json:"remark" comment:"描述"`
}

func (s *TbxCountryUpdateReq) Generate(model *models.TbxCountry)  {
    model.Code = s.Code
    model.Code2 = s.Code2
    model.DigitCode = s.DigitCode
    model.TeleCode = s.TeleCode
    model.Group = s.Group
    model.BelongTo = s.BelongTo
    model.NameCN = s.NameCN
    model.NameEN = s.NameEN
    model.DisplaySort = s.DisplaySort
    model.Remark = s.Remark
}

func (s *TbxCountryUpdateReq) GetId() interface{} {
	return s.Code
}

// TbxCountryGetReq 功能获取请求参数
type TbxCountryGetReq struct {
     Code string `uri:"code"`
}

func (s *TbxCountryGetReq) GetId() interface{} {
	return s.Code
}

// TbxCountryDeleteReq 功能删除请求参数
type TbxCountryDeleteReq struct {
	Ids []string `json:"ids"`
}

func (s *TbxCountryDeleteReq) GetId() interface{} {
	return s.Ids
}