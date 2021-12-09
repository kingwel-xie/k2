package service

import (
	"admin/models"
	"admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
	"github.com/kingwel-xie/k2/common/service"
)

type SysDictData struct {
	service.Service
}

// GetPage 获取列表
func (e *SysDictData) GetPage(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData, count *int64) error {
	var data models.SysDictData
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
func (e *SysDictData) Get(d *dto.SysDictDataGetReq, model *models.SysDictData) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建对象
func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) error {
	var data = new(models.SysDictData)
	c.Generate(data)
	data.SetCreateBy(e.Identity.UserId)
	err := e.Orm.Create(data).Error

	return err
}

// Update 修改对象
func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq) error {
	var model = models.SysDictData{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)
	db := e.Orm.Save(model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// Remove 删除
func (e *SysDictData) Remove(c *dto.SysDictDataDeleteReq) error {
	var data models.SysDictData
	db := e.Orm.Delete(&data, c.GetId())

	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// GetAll 获取所有
func (e *SysDictData) GetAll(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData) error {
	var data models.SysDictData
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error

	return err
}
