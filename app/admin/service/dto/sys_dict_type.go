package dto

import (
	"github.com/kingwel-xie/k2/app/admin/models"

	"github.com/kingwel-xie/k2/common/dto"
)

type SysDictTypeGetPageReq struct {
	dto.Pagination `search:"-"`
	DictId         []int  `form:"dictId" search:"type:in;column:dict_id;table:sys_dict_type"`
	DictName       string `form:"dictName" search:"type:icontains;column:dict_name;table:sys_dict_type"`
	DictType       string `form:"dictType" search:"type:icontains;column:dict_type;table:sys_dict_type"`
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_dict_type"`
}

type SysDictTypeOrder struct {
	DictIdOrder string `search:"type:order;column:dict_id;table:sys_dict_type" form:"dictIdOrder"`
}

func (m *SysDictTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysDictTypeInsertReq struct {
	Id       int    `uri:"id"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
}

func (s *SysDictTypeInsertReq) Generate(model *models.SysDictType) {
	if s.Id != 0 {
		model.ID = s.Id
	}
	model.DictName = s.DictName
	model.DictType = s.DictType
	model.Status = s.Status
	model.Remark = s.Remark

}

func (s *SysDictTypeInsertReq) GetId() interface{} {
	return s.Id
}

type SysDictTypeUpdateReq struct {
	Id       int    `uri:"id"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
}

func (s *SysDictTypeUpdateReq) Generate(model *models.SysDictType) {
	if s.Id != 0 {
		model.ID = s.Id
	}
	model.DictName = s.DictName
	model.DictType = s.DictType
	model.Status = s.Status
	model.Remark = s.Remark

}

func (s *SysDictTypeUpdateReq) GetId() interface{} {
	return s.Id
}

type SysDictTypeGetReq struct {
	Id int `uri:"id"`
}

func (s *SysDictTypeGetReq) GetId() interface{} {
	return s.Id
}

type SysDictTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysDictTypeDeleteReq) GetId() interface{} {
	return s.Ids
}
