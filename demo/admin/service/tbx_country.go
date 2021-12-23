package service

import (
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"

	"admin/models"
	"admin/service/dto"
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
	err := e.Orm.
        First(model, "code = ?", d.GetId()).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}
	return nil
}

// Insert 创建TbxCountry对象
func (e *TbxCountry) Insert(c *dto.TbxCountryInsertReq) error {
    var err error
	var list []models.TbxCountry
	err = e.Orm.
		Find(&list, "code = ?", c.GetId()).Error
	if err != nil {
		return k2Error.ErrDatabase.Wrap(err)
	}
	if len(list) > 0 {
		return k2Error.ErrCodeExisted
	}

    var data models.TbxCountry
    c.Generate(&data)
    data.SetCreateBy(e.Identity.Username)

	err = e.Orm.Create(&data).Error
	if err != nil {
	    return k2Error.ErrDatabase.Wrap(err)
	}
	c.Code = data.Code
	return nil
}

// Update 修改TbxCountry对象
func (e *TbxCountry) Update(c *dto.TbxCountryUpdateReq) error {
    var data = models.TbxCountry{}
    err := e.Orm.
        First(&data, "code = ?", c.GetId()).Error
    if err != nil {
    	return k2Error.ErrCodeNotFound.Wrap(err)
	}
    c.Generate(&data)
    data.SetUpdateBy(e.Identity.Username)

    db := e.Orm.Save(&data)
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
    return nil
}

// Remove 删除TbxCountry
func (e *TbxCountry) Remove(d *dto.TbxCountryDeleteReq) error {
	var data models.TbxCountry
	db := e.Orm.
	    Delete(&data, "code in ?", d.GetId())
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}