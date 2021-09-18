package service

import (
	"github.com/kingwel-xie/k2/common/service"

	"github.com/kingwel-xie/k2/app/admin/models"
	"github.com/kingwel-xie/k2/app/admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
)

type SysOperaLog struct {
	service.Service
}

// GetPage 获取SysOperaLog列表
func (e *SysOperaLog) GetPage(c *dto.SysOperaLogGetPageReq, list *[]models.SysOperaLog, count *int64) error {
	var data models.SysOperaLog

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysOperaLog对象
func (e *SysOperaLog) Get(d *dto.SysOperaLogGetReq, model *models.SysOperaLog) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建SysOperaLog对象
func (e *SysOperaLog) Insert(model *models.SysOperaLog) error {
	var data models.SysOperaLog

	err := e.Orm.Model(&data).
		Create(model).Error

	return err
}

// Remove 删除SysOperaLog
func (e *SysOperaLog) Remove(d *dto.SysOperaLogDeleteReq) error {
	var data models.SysOperaLog

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}
