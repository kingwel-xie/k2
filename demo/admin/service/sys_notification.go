package service

import (
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"gorm.io/gorm"

	"admin/models"
	"admin/service/dto"
)

type SysNotification struct {
	service.Service
}

// GetPage 获取SysNotification列表
func (e *SysNotification) GetPage(c *dto.SysNotificationGetPageReq, list *[]models.SysNotification, count *int64) error {
	var data models.SysNotification
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.OrderDest("id", true),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysNotification对象
func (e *SysNotification) Get(d *dto.SysNotificationGetReq, model *models.SysNotification) error {
	err := e.Orm.
        First(model, "id = ?", d.GetId()).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}
	return nil
}

// Insert 创建SysNotification对象
func (e *SysNotification) Insert(c *dto.SysNotificationInsertReq) error {
	var data models.SysNotification
	c.Generate(&data)
	data.SetCreateBy(e.Identity.Username)

	return e.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&data).Error
		if err != nil {
			return k2Error.ErrDatabase.Wrap(err)
		}
		c.Id = data.Id

		var sendMessage = dto.SysInboxSendMessageReq{
			TargetType: data.TargetType,
			Targets: data.Targets,
			Title: data.Title,
			Content: data.Content,
		}
		return SendMessage(tx, &sendMessage, e.Identity.Username)
	})
}

// Update 修改SysNotification对象
func (e *SysNotification) Update(c *dto.SysNotificationUpdateReq) error {
    var data = models.SysNotification{}
    err := e.Orm.
        First(&data, "id = ?", c.GetId()).Error
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

// Remove 删除SysNotification
func (e *SysNotification) Remove(d *dto.SysNotificationDeleteReq) error {
	var data models.SysNotification
	db := e.Orm.
	    Delete(&data, "id in ?", d.GetId())
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}
