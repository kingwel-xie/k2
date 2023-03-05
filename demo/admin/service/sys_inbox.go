package service

import (
	"encoding/json"
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"gorm.io/gorm"
	"admin/models"
	"admin/service/dto"
)

type SysInbox struct {
	service.Service
}

// Permission check permission of message 'Receiver' field
func messagePermission(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sys_inbox.receiver = ? ", username, username)
	}
}

// GetPage 获取SysInbox列表
func (e *SysInbox) GetPage(c *dto.SysInboxGetPageReq, list *[]models.SysInbox, count *int64) error {
	var data models.SysInbox
	err := e.Orm.Model(&data).
		Scopes(
			messagePermission(e.Identity.Username),
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.OrderDest("id", true),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysInbox对象
func (e *SysInbox) Get(d *dto.SysInboxGetReq, model *models.SysInbox) error {
	err := e.Orm.
		Scopes(
			messagePermission(e.Identity.Username),
		).
        First(model, "id = ?", d.GetId()).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}
	return nil
}

func inUserArray(u string, users ...models.SysUser) bool {
	for _, s := range users {
		if u == s.Username {
			return true
		}
	}
	return false
}


// Insert 创建SysInbox对象
func (e *SysInbox) Insert(c *dto.SysInboxInsertReq) error {
    var err error

	if len(c.Receiver) == 0 {
		return k2Error.ErrCodeNotFound.Wrapf("no receivers")
	}
	// check if Receiver exists
	var users []models.SysUser
	err = e.Orm.
		Find(&users, "username in ?", c.Receiver).Error
	if err != nil {
		return k2Error.ErrCodeNotFound.Wrap(err)
	}
	if len(c.Receiver) != len(users) {
		var diff []string
		for _, v := range c.Receiver {
			if !inUserArray(v, users...) {
				diff = append(diff, v)
			}
		}
		return k2Error.ErrCodeNotFound.Wrapf("mismatch receivers, %v", diff)
	}

	// generate the messages
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		var list []models.SysInbox
		for _, receiver := range c.Receiver {
			var data models.SysInbox
			data.Receiver = receiver
			data.OriginId = c.OriginId
			data.Title = c.Title
			data.Content = c.Content
			data.SetCreateBy(e.Identity.Username)
			data.Sender = e.Identity.Username
			data.Type = "message"
			data.Read = false
			list = append(list, data)
		}
		err = tx.Create(&list).Error
		if err != nil {
			return k2Error.ErrDatabase.Wrap(err)
		}

		// then put into Outbox
		receiversInJsonString, _ := json.Marshal(c.Receiver)
		var outbox = models.SysOutbox{
			Sender: e.Identity.Username,
			Receivers: string(receiversInJsonString),
			OriginId: c.OriginId,
			Title: c.Title,
			Content: c.Content,
			IsDraft: false,
		}
		err = tx.Create(&outbox).Error
		if err != nil {
			return k2Error.ErrDatabase.Wrap(err)
		}

		return nil
	})
}

// Update 修改SysInbox对象
func (e *SysInbox) Update(c *dto.SysInboxUpdateReq) error {
    var data = models.SysInbox{}
    err := e.Orm.
		Scopes(
			messagePermission(e.Identity.Username),
		).
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

// Remove 删除SysInbox
func (e *SysInbox) Remove(d *dto.SysInboxDeleteReq) error {
	var data models.SysInbox
	db := e.Orm.
		Scopes(
			messagePermission(e.Identity.Username),
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

// GetUnread 返回未读消息数量及列表
func (e *SysInbox) GetUnread(object *dto.SysUnread) error {
	var err error
	var req dto.SysInboxGetPageReq
	req.Type = "notification"
	req.Unread = true
	req.IdOrder = "desc"
	req.Pagination.PageSize = 10
	list := make([]models.SysInbox, 0)
	var count int64
	err = e.GetPage(&req, &list, &count)
	if err != nil {
		return err
	}
	object.NumNotices = int(count)
	object.NoticeList = &list

	req.Type = "message"
	req.Unread = true
	req.IdOrder = "desc"
	req.Pagination.PageSize = 10
	list2 := make([]models.SysInbox, 0)
	var count2 int64
	err = e.GetPage(&req, &list2, &count2)
	if err != nil {
		return err
	}
	object.NumMessages = int(count2)
	object.MessageList = &list2

	return nil
}

// MarkRead 标记已读SysInbox
func (e *SysInbox) MarkRead(d *dto.SysInboxMarkReadReq) error {
	var data models.SysInbox
	db := e.Orm.Model(&data).
		Scopes(
			messagePermission(e.Identity.Username),
		).
		Where("id in ?", d.GetId()).
		Update("read", d.Read)
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

func SendMessage(db *gorm.DB, req *dto.SysInboxSendMessageReq, from string) error {
	var list []models.SysUser
	switch req.TargetType {
	case models.TargetTypeAll:
		db.Find(&list, "role_id <> 10")
	case models.TargetTypeRole:
		var roleKeys []string
		_ = json.Unmarshal([]byte(req.Targets), &roleKeys)
		roleIds := db.Model(&models.SysRole{}).Distinct("role_id").Where("role_key in ?", roleKeys)
		db.Find(&list, "role_id in (?)", roleIds)
	case models.TargetTypeDept:
		var deptIdList []int
		_ = json.Unmarshal([]byte(req.Targets), &deptIdList)
		db.Find(&list, "dept_id in ?", deptIdList)
	case models.TargetTypeUser:
		var usernames []string
		_ = json.Unmarshal([]byte(req.Targets), &usernames)
		db.Find(&list, "username in ?", usernames)
	}
	var messages []models.SysInbox
	for _, user := range list {
		m := models.SysInbox{
			Type: "notification",
			Sender: from,
			Receiver: user.Username,
			Title: req.Title,
			Content: req.Content,
			Read: false,
		}
		m.SetCreateBy(from)
		messages = append(messages, m)
	}
	err := db.Create(messages).Error
	if err != nil {
		return k2Error.ErrDatabase.Wrap(err)
	}
	return nil
}

