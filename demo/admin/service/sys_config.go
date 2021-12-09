package service

import (
	"github.com/kingwel-xie/k2/common/service"

	"admin/models"
	"admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
)

type SysConfig struct {
	service.Service
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(c *dto.SysConfigGetPageReq, list *[]models.SysConfig, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysConfig对象
func (e *SysConfig) Get(d *dto.SysConfigGetReq, model *models.SysConfig) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *dto.SysConfigControl) error {
	var data models.SysConfig
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)
	err := e.Orm.Create(&data).Error
	return err
}

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *dto.SysConfigControl) error {
	var model = models.SysConfig{}
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

// SetSysConfig 修改SysConfig对象
func (e *SysConfig) SetSysConfig(c *[]dto.GetSetSysConfigReq) error {
	var err error
	for _, req := range *c {
		var model = models.SysConfig{}
		e.Orm.Where("config_key = ?", req.ConfigKey).First(&model)
		if model.Id != 0 {
			req.Generate(&model)
			db := e.Orm.Save(&model)
			err = db.Error
			if err != nil {
				return err
			}
			if db.RowsAffected == 0 {
				return k2Error.ErrPermissionDenied
			}
		}
	}
	return nil
}

func (e *SysConfig) GetForSet(c *[]dto.GetSetSysConfigReq) error {
	var data models.SysConfig
	err := e.Orm.Model(&data).
		Find(c).Error

	return err
}

func (e *SysConfig) UpdateForSet(c *[]dto.GetSetSysConfigReq) error {
	m := *c
	for _, req := range m {
		var data models.SysConfig
		if err := e.Orm.Where("config_key = ?", req.ConfigKey).
			First(&data).Error; err != nil {
			return err
		}
		if data.ConfigValue != req.ConfigValue {
			data.ConfigValue = req.ConfigValue
			if err := e.Orm.Save(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Remove 删除SysConfig
func (e *SysConfig) Remove(d *dto.SysConfigDeleteReq) error {
	var data models.SysConfig

	db := e.Orm.Delete(&data, d.Ids)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// GetWithKey 根据Key获取SysConfig
func (e *SysConfig) GetWithKey(c *dto.SysConfigByKeyReq, resp *dto.GetSysConfigByKEYForServiceResp) error {
	var data models.SysConfig
	err := e.Orm.Table(data.TableName()).Where("config_key = ?", c.ConfigKey).First(resp).Error
	return err
}

func (e *SysConfig) GetWithKeyList(c *dto.SysConfigGetToSysAppReq, list *[]models.SysConfig) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	return err
}
