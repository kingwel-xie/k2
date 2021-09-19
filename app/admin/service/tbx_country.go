package service

import (
	"github.com/kingwel-xie/k2/app/admin/models"
	"github.com/kingwel-xie/k2/app/admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
	"github.com/kingwel-xie/k2/common/service"
)

type TbxCountry struct {
	service.Service
}

// GetPage 获取TbxCountry列表
func (e *TbxCountry) GetPage(c *dto.TbxCountryGetPageReq, list *[]models.TbxCountry, count *int64) error {
	var data models.TbxCountry
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取TbxCountry对象
func (e *TbxCountry) Get(d *dto.TbxCountryGetReq, model *models.TbxCountry) error {
	var data models.TbxCountry
	err := e.Orm.Model(&data).First(model, "code = ?", d.GetId()).Error
	return err
}

// Insert 创建TbxCountry对象
func (e *TbxCountry) Insert(c *dto.TbxCountryInsertReq) error {
	var data models.TbxCountry
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)

	err := e.Orm.Create(&data).Error
	if err == nil {
		c.Code = data.Code
	}
	return err
}

// Update 修改TbxCountry对象
func (e *TbxCountry) Update(c *dto.TbxCountryUpdateReq) error {
	var data = models.TbxCountry{}
	err := e.Orm.First(&data, "code = ?", c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&data)
	data.SetUpdateBy(e.Identity.UserId)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// Remove 删除TbxCountry
func (e *TbxCountry) Remove(d *dto.TbxCountryDeleteReq) error {
	var data models.TbxCountry
	db := e.Orm.Model(&data).Delete(&data, "code in ?", d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}
