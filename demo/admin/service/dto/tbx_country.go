package dto

import (
	"admin/models"
	"github.com/kingwel-xie/k2/common/dto"
)

type TbxCountryGetPageReq struct {
	dto.Pagination `search:"-"`
	Code           string `form:"code"  search:"type:exact;column:code;table:tbx_country" comment:"编码"`
	CName          string `form:"cName"  search:"type:contains;column:c_name;table:tbx_country" comment:"中文名称"`
	EName          string `form:"eName"  search:"type:contains;column:e_name;table:tbx_country" comment:"English Name"`
	TbxCountryOrder
}

type TbxCountryOrder struct {
	Code  string `form:"codeOrder"  search:"type:order;column:code;table:tbx_country"`
	CName string `form:"cNameOrder"  search:"type:order;column:c_name;table:tbx_country"`
	EName string `form:"eNameOrder"  search:"type:order;column:e_name;table:tbx_country"`
}

func (m *TbxCountryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbxCountryInsertReq struct {
	Code   string `json:"code" comment:"编码"`
	CName  string `json:"cName" comment:"中文名称" vd:"@:len($)>0; msg:'中文名不能为空'"`
	EName  string `json:"eName" comment:"English Name" vd:"@:len($)>0; msg:'English Name is empty'"`
	EName2 string `json:"eName2" comment:"English Name 2"`
	Code2  string `json:"code2" comment:"编码2"`
	Code3  string `json:"code3" comment:"编码3"`
}

func (s *TbxCountryInsertReq) Generate(model *models.TbxCountry) {
	model.Code = s.Code
	model.CName = s.CName
	model.EName = s.EName
	model.EName2 = s.EName2
	model.Code2 = s.Code2
	model.Code3 = s.Code3
}

func (s *TbxCountryInsertReq) GetId() interface{} {
	return s.Code
}

type TbxCountryUpdateReq struct {
	Code   string `uri:"code" comment:"编码"`
	CName  string `json:"cName" comment:"中文名称" vd:"@:len($)>0; msg:'中文名不能为空'"`
	EName  string `json:"eName" comment:"English Name" vd:"@:len($)>0; msg:'English Name is empty'"`
	EName2 string `json:"eName2" comment:"English Name 2"`
	Code2  string `json:"code2" comment:"编码2"`
	Code3  string `json:"code3" comment:"编码3"`
}

func (s *TbxCountryUpdateReq) Generate(model *models.TbxCountry) {
	model.Code = s.Code
	model.CName = s.CName
	model.EName = s.EName
	model.EName2 = s.EName2
	model.Code2 = s.Code2
	model.Code3 = s.Code3
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
