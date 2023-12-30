package service

import (
	"errors"
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/utils"

	"admin/models"
	"admin/service/dto"
)

const RoleIdAdmin = 1

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	err := e.Orm.Preload("Dept").Preload("Role").
		Scopes(
			service.Permission(models.SysUser{}.TableName(), e.Identity),
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// ListNoCheck 获取SysUser列表, NoCheck
func (e *SysUser) ListNoCheck(c *dto.SysUserGetPageReq, list *[]models.SysUser) error {
	err := e.Orm.
		Scopes(
			service.Permission(models.SysUser{}.TableName(), e.Identity),
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Select("UserId", "Username", "NickName").
		Find(list).Error

	return err
}

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, model *models.SysUser) error {
	err := e.Orm.Preload("Dept").Preload("Role").
		Scopes(
		service.Permission(model.TableName(), e.Identity),
	).First(model, d.GetId()).Error
	return err
}

func (e *SysUser) GetCurrentUser() (model models.SysUser, err error) {
	err = e.Orm.First(&model, e.Identity.UserId).Error
	return
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	if c.RoleId == RoleIdAdmin {
		return errors.New("'admin' can not be specified")
	}

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
	err = data.Encrypt()
	if err != nil {
		return k2Error.ErrInternal
	}
	data.SetCreateBy(e.Identity.Username)

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

	if model.RoleId == RoleIdAdmin {
		// 'admin' can not be downgraded
		c.RoleId = RoleIdAdmin
	} else {
		// upgrade to 'admin' is not allowed
		if c.RoleId == RoleIdAdmin {
			return errors.New("'admin' can not be specified")
		}
	}

	c.Generate(&model)
	model.SetUpdateBy(e.Identity.Username)

	db := e.Orm.Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var model models.SysUser
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).Select("UserId", "Avatar").First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.Username)

	db := e.Orm.Select("Avatar").Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var model models.SysUser
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).Select("UserId", "Status").First(&model, c.GetId()).Error
	if err != nil {
		return err
	}

	c.Generate(&model)
	model.SetUpdateBy(e.Identity.Username)

	db := e.Orm.Select("Status").Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var model models.SysUser
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).Select("UserId", "Password", "Salt").First(&model, c.GetId()).Error
	if err != nil {
		return err
	}

	c.Generate(&model)
	err = model.Encrypt()
	if err != nil {
		return k2Error.ErrInternal
	}
	model.SetUpdateBy(e.Identity.Username)
	db := e.Orm.Select("Password", "Salt").Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// ResetToken 重置用户Token
func (e *SysUser) ResetToken(c *dto.ResetSysUserTokenReq) error {
	var model models.SysUser
	err := e.Orm.Scopes(
		service.Permission(model.TableName(), e.Identity),
	).Select("UserId", "Token").First(&model, c.GetId()).Error
	if err != nil {
		return err
	}

	c.Generate(&model)
	if err != nil {
		return k2Error.ErrInternal
	}
	model.SetUpdateBy(e.Identity.Username)
	db := e.Orm.Select("Token", ).Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}

	// refresh cache, as data changed
	backlog.ReloadCacheAsync()
	return nil
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
		return k2Error.ErrPermissionDenied
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
		return k2Error.ErrDatabase.Wrap(err)
	}
	var ok bool
	ok, err = utils.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		return k2Error.ErrWrongPassword
	}
	if !ok {
		return k2Error.ErrWrongPassword
	}

	c.Password = newPassword
	err = c.Encrypt()
	if err != nil {
		return k2Error.ErrInternal
	}
	db := e.Orm.Select("Password", "Salt").Save(c)
	if db.Error != nil {
		return k2Error.ErrDatabase.Wrap(db.Error)
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}


// UpdateProfile 修改Profile
func (e *SysUser) UpdateProfile(id int, req *dto.SysUserUpdateProfileReq) error {
	c := &models.SysUser{}

	err := e.Orm.Model(c).Scopes(
		service.Permission(c.TableName(), e.Identity),
	).First(c, id).Error
	if err != nil {
		return err
	}

	c.NickName = req.NickName
	c.Email = req.Email
	c.Phone = req.Phone
	c.Sex = req.Sex
	c.Remark = req.Remark

	// update the profile, only the specified fields
	db := e.Orm.Select("NickName", "Email", "Phone", "Sex", "UpdateBy").Save(c)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
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

func (e *SysUser) CheckExistence(s *dto.SysUserCheckExistenceReq) error {
	var list []models.SysUser
	err := e.Orm.
		Find(&list, "username = ?", s.GetId()).Limit(1).Error
	if err == nil {
		if len(list) > 0 {
			return nil
		} else {
			err = k2Error.ErrCodeNotFound
		}
	}
	return err
}
