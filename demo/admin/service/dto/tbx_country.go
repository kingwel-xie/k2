package dto

import (
	"github.com/kingwel-xie/k2/common/dto"

	"admin/models"
)

type TbxCountryGetPageReq struct {
	dto.Pagination     `search:"-"`
    Code string `form:"code"  search:"type:exact;column:code;table:tbx_country" comment:"编码"`
    NameCN string `form:"nameCN"  search:"type:contains;column:name_cn;table:tbx_country" comment:"中文名称"`
    NameEN string `form:"nameEN"  search:"type:contains;column:name_en;table:tbx_country" comment:"English Name"`
    Alias string `form:"alias"  search:"type:contains;column:alias;table:tbx_country" comment:"描述"`
    TbxCountryOrder
}

type TbxCountryOrder struct {
}

func (m *TbxCountryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbxCountryInsertReq struct {
    Code string `json:"code" comment:"编码"`
    NameCN string `json:"nameCN" comment:"中文名称" vd:"@:len($)>0; msg:'中文名不能为空'"`
    NameEN string `json:"nameEN" comment:"English Name" vd:"@:len($)>0; msg:'English Name is empty'"`
    Alias string `json:"alias" comment:"描述"`
}

func (s *TbxCountryInsertReq) Generate(model *models.TbxCountry)  {
    model.Code = s.Code
    model.NameCN = s.NameCN
    model.NameEN = s.NameEN
    model.Alias = s.Alias
}

func (s *TbxCountryInsertReq) GetId() interface{} {
	return s.Code
}

type TbxCountryUpdateReq struct {
    Code string `uri:"code" comment:"编码"`
    NameCN string `json:"nameCN" comment:"中文名称" vd:"@:len($)>0; msg:'中文名不能为空'"`
    NameEN string `json:"nameEN" comment:"English Name" vd:"@:len($)>0; msg:'English Name is empty'"`
    Alias string `json:"alias" comment:"描述"`
}

func (s *TbxCountryUpdateReq) Generate(model *models.TbxCountry)  {
    model.Code = s.Code
    model.NameCN = s.NameCN
    model.NameEN = s.NameEN
    model.Alias = s.Alias
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