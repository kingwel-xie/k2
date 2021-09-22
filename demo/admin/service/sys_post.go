package service

import (
	"github.com/kingwel-xie/k2/common/service"

	"admin/models"
	"admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
)

type SysPost struct {
	service.Service
}

// GetPage 获取SysPost列表
func (e *SysPost) GetPage(c *dto.SysPostPageReq, list *[]models.SysPost, count *int64) error {
	var data models.SysPost

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysPost对象
func (e *SysPost) Get(d *dto.SysPostGetReq, model *models.SysPost) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建SysPost对象
func (e *SysPost) Insert(c *dto.SysPostInsertReq) error {
	var data models.SysPost
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)

	err := e.Orm.Create(&data).Error
	return err
}

// Update 修改SysPost对象
func (e *SysPost) Update(c *dto.SysPostUpdateReq) error {
	var model = models.SysPost{}
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
		return service.ErrPermissionDenied
	}
	return nil
}

// Remove 删除SysPost
func (e *SysPost) Remove(d *dto.SysPostDeleteReq) error {
	var data models.SysPost

	db := e.Orm.Delete(&data, d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}
