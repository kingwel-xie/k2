package service

import (
	"errors"
	"fmt"
	"github.com/kingwel-xie/k2/common/service"

	"admin/models"
	"admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
)

type SysDictType struct {
	service.Service
}

// GetPage 获取列表
func (e *SysDictType) GetPage(c *dto.SysDictTypeGetPageReq, list *[]models.SysDictType, count *int64) error {
	var data models.SysDictType
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取对象
func (e *SysDictType) Get(d *dto.SysDictTypeGetReq, model *models.SysDictType) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建对象
func (e *SysDictType) Insert(c *dto.SysDictTypeInsertReq) error {
	var data models.SysDictType
	c.Generate(&data)
	data.SetUpdateBy(e.Identity.UserId)

	var count int64
	e.Orm.Model(&data).Where("dict_type = ?", data.DictType).Count(&count)
	if count > 0 {
		return errors.New(fmt.Sprintf("当前字典类型[%s]已经存在！", data.DictType))
	}
	err := e.Orm.Create(&data).Error
	return err
}

// Update 修改对象
func (e *SysDictType) Update(c *dto.SysDictTypeUpdateReq) error {
	var model = models.SysDictType{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	db := e.Orm.Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// Remove 删除
func (e *SysDictType) Remove(d *dto.SysDictTypeDeleteReq) error {
	var data models.SysDictType

	db := e.Orm.Delete(&data, d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// GetAll 获取所有
func (e *SysDictType) GetAll(c *dto.SysDictTypeGetPageReq, list *[]models.SysDictType) error {
	var data models.SysDictType

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error

	return err
}
