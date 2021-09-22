package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"admin/models"
	"admin/service/dto"
	"github.com/kingwel-xie/k2/common/casbin"
	cDto "github.com/kingwel-xie/k2/common/dto"
	"github.com/kingwel-xie/k2/common/service"
)

type SysRole struct {
	service.Service
}

// GetPage 获取SysRole列表
func (e *SysRole) GetPage(c *dto.SysRoleGetPageReq, list *[]models.SysRole, count *int64) error {
	var data models.SysRole

	err := e.Orm.Model(&data).Preload("SysMenu").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	return err
}

// Get 获取SysRole对象
func (e *SysRole) Get(d *dto.SysRoleGetReq, model *models.SysRole) error {
	err := e.Orm.First(model, d.GetId()).Error
	if err != nil {
		return err
	}
	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		return err
	}
	return nil
}

// Insert 创建SysRole对象
func (e *SysRole) Insert(c *dto.SysRoleInsertReq) error {
	var data models.SysRole
	var dataMenu []models.SysMenu
	err := e.Orm.Preload("SysApi").Where("menu_id in ?", c.MenuIds).Find(&dataMenu).Error
	if err != nil {
		return err
	}
	c.SysMenu = dataMenu
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)

	tx := e.Orm.Begin()

	// setup casbin with the TX
	cb := casbin.Setup(tx)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(&data).Error
	if err != nil {
		return err
	}

	for _, menu := range dataMenu {
		for _, api := range menu.SysApi {
			_, err = cb.AddNamedPolicy("p", data.RoleKey, api.Path, api.Action)
		}
	}
	_ = cb.SavePolicy()
	//if len(c.MenuIds) > 0 {
	//	s := SysRoleMenu{}
	//	s.Orm = e.Orm
	//	s.Log = e.Log
	//	err = s.ReloadRule(tx, c.RoleId, c.MenuIds)
	//	if err != nil {
	//		e.Log.Errorf("reload casbin rule error, %", err.Error())
	//		return err
	//	}
	//}
	return nil
}

// Update 修改SysRole对象, Menu & API
func (e *SysRole) Update(c *dto.SysRoleUpdateReq) error {
	var err error
	tx := e.Orm.Begin()
	// setup casbin with the TX
	cb := casbin.Setup(tx)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	err = tx.Preload("SysMenu").First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	err = tx.Model(&model).Association("SysMenu").Delete(model.SysMenu)
	if err != nil {
		return err
	}

	var mlist = make([]models.SysMenu, 0)
	err = tx.Preload("SysApi").Where("menu_id in ?", c.MenuIds).Find(&mlist).Error
	if err != nil {
		return err
	}
	model.SysMenu = &mlist

	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}

	_, err = cb.RemoveFilteredPolicy(0, model.RoleKey)
	if err != nil {
		return err
	}

	for _, menu := range mlist {
		for _, api := range menu.SysApi {
			_, err = cb.AddNamedPolicy("p", model.RoleKey, api.Path, api.Action)
		}
	}
	_ = cb.SavePolicy()
	return nil
}

// Remove 删除SysRole
func (e *SysRole) Remove(c *dto.SysRoleDeleteReq) error {
	var err error
	tx := e.Orm.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.Preload("SysMenu").Preload("SysDept").First(&model, c.GetId())
	db := tx.Select(clause.Associations).Delete(&model)
	err = db.Error

	if err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// GetRoleMenuId 获取角色对应的菜单ids
func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil
}

// UpdateDataScope 修改数据权限, datascope => 自定义部门权限
func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) error {
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var model = models.SysRole{}
	err = tx.Preload("SysDept").First(&model, c.RoleId).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	err = tx.Model(&model).Association("SysDept").Delete(model.SysDept)
	if err != nil {
		return err
	}

	var dlist = make([]models.SysDept, 0)
	err = tx.Where("dept_id in ?", c.DeptIds).Find(&dlist).Error
	if err != nil {
		return err
	}

	model.SysDept = dlist
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// UpdateStatus 修改SysRole对象status
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) error {
	var model = models.SysRole{}
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

// GetWithName 获取SysRole对象
func (e *SysRole) GetWithName(d *dto.SysRoleByName, model *models.SysRole) error {
	var err error
	db := e.Orm.Where("role_name = ?", d.RoleName).First(model)
	err = db.Error
	if err != nil {
		return err
	}

	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	return err
}

// GetById 获取SysRole对象
func (e *SysRole) GetById(roleId int) ([]string, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, nil
}
