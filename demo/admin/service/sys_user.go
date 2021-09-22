package service

import (
	"errors"
	"fmt"
	"admin/models"
	"admin/service/dto"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/utils"

	cDto "github.com/kingwel-xie/k2/common/dto"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	err := e.Orm.Preload("Dept").
		Scopes(
			service.Permission(models.SysUser{}.TableName(), e.Identity),
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, model *models.SysUser) error {
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(model, d.GetId()).Error
	return err
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		return err
	}
	if i > 0 {
		return errors.New("用户名已存在！")
	}
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)

	err = e.Orm.Create(&data).Error
	return err
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq) error {
	var model models.SysUser
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(&model, c.GetId()).Error
	if err != nil {
		return err
	}

	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	update := e.Orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if update.Error != nil {
		return update.Error
	}
	if update.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var model models.SysUser
	db := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(&model, c.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	err := e.Orm.Save(&model).Error
	return err
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var model models.SysUser
	db := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(&model, c.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}

	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	err := e.Orm.Save(&model).Error
	return err
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var model models.SysUser
	db := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(&model, c.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}

	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)
	err := e.Orm.Save(&model).Error

	return err
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *dto.SysUserById) error {
	var data models.SysUser

	db := e.Orm.Model(&data).Scopes(
		service.Permission(data.TableName(), e.Identity),
	).Delete(&data, c.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) error {
	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err := e.Orm.Model(c).Scopes(
		service.Permission(c.TableName(), e.Identity),
	).Select("UserId", "Password", "Salt").First(c, id).Error
	if err != nil {
		return err
	}
	var ok bool
	ok, err = utils.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		return fmt.Errorf("CompareHashAndPassword error, %s", err.Error())
	}
	if !ok {
		return fmt.Errorf("user[%d] incorrect password", id)
	}
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).Select("Password", "Salt").Updates(c)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := e.Orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(posts, user.PostIds).Error
	if err != nil {
		return err
	}

	return nil
}
