package service

import (
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"gorm.io/gorm"

	"admin/models"
	"admin/service/dto"
)

type SysOutbox struct {
	service.Service
}

// Permission check permission of message 'Sender' field
func senderPermission(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sys_outbox.sender = ? ", username, username)
	}
}

// GetPage 获取SysOutbox列表
func (e *SysOutbox) GetPage(c *dto.SysOutboxGetPageReq, list *[]models.SysOutbox, count *int64) error {
	var data models.SysOutbox
	err := e.Orm.Model(&data).
		Scopes(
			senderPermission(e.Identity.Username),
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.OrderDest("id", true),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysOutbox对象
func (e *SysOutbox) Get(d *dto.SysOutboxGetReq, model *models.SysOutbox) error {
	err := e.Orm.
		Scopes(
			senderPermission(e.Identity.Username),
		).
        First(model, "id = ?", d.GetId()).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}
	return nil
}

// Remove 删除SysOutbox
func (e *SysOutbox) Remove(d *dto.SysOutboxDeleteReq) error {
	var data models.SysOutbox
	db := e.Orm.
		Scopes(
			senderPermission(e.Identity.Username),
		).
	    Delete(&data, "id in ?", d.GetId())
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}
