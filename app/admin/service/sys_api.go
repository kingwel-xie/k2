package service

import (
	"github.com/kingwel-xie/k2/app/admin/models"
	"github.com/kingwel-xie/k2/app/admin/service/dto"

	"github.com/kingwel-xie/k2/common"
	cDto "github.com/kingwel-xie/k2/common/dto"
	"github.com/kingwel-xie/k2/common/service"
)

type SysApi struct {
	service.Service
}

// GetPage 获取SysApi列表
func (e *SysApi) GetPage(c *dto.SysApiGetPageReq, list *[]models.SysApi, count *int64) error {
	var err error
	var data models.SysApi

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysApi对象with id
func (e *SysApi) Get(d *dto.SysApiGetReq, model *models.SysApi) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Update 修改SysApi对象
func (e *SysApi) Update(c *dto.SysApiUpdateReq) error {
	var model = models.SysApi{}
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

// Remove 删除SysApi
func (e *SysApi) Remove(d *dto.SysApiDeleteReq) error {
	var data models.SysApi

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// CheckStorageSysApi 创建SysApi对象
func (e *SysApi) CheckStorageSysApi(c *[]common.Router) error {
	for _, v := range *c {
		err := e.Orm.Where(models.SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
			Attrs(models.SysApi{Handle: v.Handler}).
			FirstOrCreate(&models.SysApi{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
