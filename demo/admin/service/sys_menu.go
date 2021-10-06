package service

import (
	"errors"
	"fmt"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"

	"admin/models"
	"admin/service/dto"
	cDto "github.com/kingwel-xie/k2/common/dto"
	cModels "github.com/kingwel-xie/k2/common/models"
)

type SysMenu struct {
	service.Service
}

// GetPage 获取SysMenu列表
func (e *SysMenu) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) error {
	var menu = make([]models.SysMenu, 0)
	err := e.getPage(c, &menu)
	if err != nil {
		return err
	}
	for i := 0; i < len(menu); i++ {
		if menu[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menu, menu[i])
		*menus = append(*menus, menusInfo)
	}
	return nil
}

// getPage 菜单分页列表
func (e *SysMenu) getPage(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) error {
	var data models.SysMenu

	err := e.Orm.Model(&data).Preload("SysApi").
		Scopes(
			cDto.OrderDest("sort", false),
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error

	return err
}

// Get 获取SysMenu对象
func (e *SysMenu) Get(d *dto.SysMenuGetReq, model *models.SysMenu) error {
	var data models.SysMenu

	err := e.Orm.Model(&data).Preload("SysApi").
		First(model, d.GetId()).Error
	if err != nil {
		return err
	}
	apis := make([]int, 0)
	for _, v := range model.SysApi {
		apis = append(apis, v.Id)
	}
	model.Apis = apis
	return nil
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(c *dto.SysMenuInsertReq) error {
	var data models.SysMenu

	// load SysApis specified by c.Apis
	var sysApis []models.SysApi
	err := e.Orm.Where("id in ?", c.Apis).Find(&sysApis).Error
	if err != nil {
		return err
	}

	// prepare the Menu data, with the SysApis
	c.Generate(&data)
	data.SetCreateBy(e.Identity.UserId)
	data.SysApi = sysApis

	err = e.Orm.Transaction(func(tx *gorm.DB) (err error) {
		// this will create the menu and the related APIs
		if err = tx.Create(&data).Error; err != nil {
			return
		}
		var paths string
		if data.ParentId == 0 {
			paths = fmt.Sprintf("/0/%d", data.MenuId)
		} else {
			var parent models.SysMenu
			if err = tx.Where("menu_id", data.ParentId).First(&parent).Error; err != nil {
				return
			}
			paths = fmt.Sprintf("%s/%d", parent.Paths, data.MenuId)
		}
		return tx.Model(&data).Update("paths", paths).Error
	})
	if err != nil {
		return err
	}
	c.MenuId = data.MenuId
	return nil
}

func (e *SysMenu) initPaths(menu *models.SysMenu) error {
	var data models.SysMenu
	parentMenu := new(models.SysMenu)
	if menu.ParentId != 0 {
		e.Orm.Model(&data).First(parentMenu, menu.ParentId)
		if parentMenu.Paths == "" {
			return errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
		}
		menu.Paths = parentMenu.Paths + "/" + utils.IntToString(menu.MenuId)
	} else {
		menu.Paths = "/0/" + utils.IntToString(menu.MenuId)
	}
	e.Orm.Model(&data).Where("menu_id = ?", menu.MenuId).Update("paths", menu.Paths)
	return nil
}

// Update 修改SysMenu对象
func (e *SysMenu) Update(c *dto.SysMenuUpdateReq) error {
	return e.Orm.Transaction(func(tx *gorm.DB) error {

		// first, load menu and []SysApi with the given menu_id
		var model = models.SysMenu{}
		err := e.Orm.Preload("SysApi").
			First(&model, "menu_id = ?", c.GetId()).Error
		if err != nil {
			return err
		}

		// delete its associated []SysApi
		err = tx.Model(&model).Association("SysApi").Delete(model.SysApi)
		if err != nil {
			return err
		}

		// try to construct the new []SysApi from req.ids
		var apiList = make([]models.SysApi, 0)
		err = tx.Where("id in ?", c.Apis).Find(&apiList).Error
		if err != nil {
			return err
		}

		// save the new CBs
		c.Generate(&model)
		model.SetUpdateBy(e.Identity.UserId)
		model.SysApi = apiList
		db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model)
		if db.Error != nil {
			return db.Error
		}
		if db.RowsAffected == 0 {
			return service.ErrPermissionDenied
		}

		return RebuildMenuPathsIfNeeded(tx, model.MenuId)
	})
}

func RebuildMenuPathsIfNeeded(tx *gorm.DB, menuIds ...int) (err error) {
	var menus []*models.SysMenu
	if err = tx.Find(&menus).Error; err != nil {
		return
	}

	menuTree := map[int][]*models.SysMenu{}
	menux := map[int]*models.SysMenu{}

	for _, menu := range menus {
		children, ok := menuTree[menu.ParentId]
		if !ok {
			children = []*models.SysMenu{}
		}
		children = append(children, menu)
		menuTree[menu.ParentId] = children
		menux[menu.MenuId] = menu
	}

	var buildPaths func(id int) string
	buildPaths = func(id int) string {
		if id == 0 {
			return "/0"
		}
		m := menux[id]
		p := menux[m.ParentId]

		if p != nil && len(p.Paths) > 0 {
			// untrustworthy, may be bad
			// return fmt.Sprintf("%s/%d", p.Paths, id)
		}
		return fmt.Sprintf("%s/%d", buildPaths(m.ParentId), id)
	}

	var batch []*models.SysMenu

	var rebuild func(newPaths string, id int)

	rebuild = func(newPaths string, id int) {

		if id != 0 {
			batch = append(batch, &models.SysMenu{
				MenuId: id,
				Paths:  newPaths,
			})
		}

		if children, ok := menuTree[id]; ok {
			for _, child := range children {
				childNewPaths := fmt.Sprintf("%s/%d", newPaths, child.MenuId)
				rebuild(childNewPaths, child.MenuId)
			}
		}
	}

	containsRoot := false

	for _, id := range menuIds {
		if id == 0 {
			containsRoot = true
			break
		}
	}

	if containsRoot || len(menuIds) == 0 {
		// rebuild all
		rebuild("/0", 0)
	} else {
		for _, id := range menuIds {

			m := menux[id]
			// /0/2/4/230
			paths := strings.Split(m.Paths, "/")

			var newPaths string

			if len(paths) < 3 {
				// invalid paths
				newPaths = buildPaths(id)
			} else {

				parentId, err := strconv.Atoi(paths[len(paths)-2])
				if err != nil {
					// ignore, shouldn't happen
					continue
				}

				if parentId == m.ParentId {
					// unchanged, skip
					continue
				}

				parent := menux[m.ParentId]

				newPaths = fmt.Sprintf("%s/%d", parent.Paths, id)
			}

			rebuild(newPaths, id)
		}
	}

	if len(batch) > 0 {
		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "menu_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"paths"}),
		}).Create(batch).Error
	}
	return
	/*
			Pure SQL implementation
		WITH RECURSIVE menu_tree (id, parent_id, pathx) AS (
		    SELECT
		           0 AS id,
		           NULL AS parent_id,
		           '/0' AS pathx
		    UNION ALL
		    SELECT
		           m.menu_id,
		           m.parent_id,
		           mt.pathx || '/' || m.menu_id as pathx
		    FROM
		         sys_menu m, menu_tree mt
		    WHERE m.parent_id = mt.id
		)
		SELECT id, pathx FROM menu_tree
		-- WHERE parent_id IS NOT NULL
		-- WHERE id = 52 OR pathx LIKE '%/52/%'
	*/
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(d *dto.SysMenuDeleteReq) error {
	var data models.SysMenu

	db := e.Orm.Model(&data).Delete(&data, d.Ids)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return service.ErrPermissionDenied
	}
	return nil
}

// GetList 获取菜单数据
func (e *SysMenu) GetList(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) error {
	var data models.SysMenu

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		return err
	}
	return nil
}

// SetLabel 修改角色中 设置菜单基础数据
func (e *SysMenu) SetLabel() (m []dto.MenuLabel, err error) {
	var list []models.SysMenu
	err = e.GetList(&dto.SysMenuGetPageReq{}, &list)

	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].MenuId
		e.Label = list[i].Title
		deptsInfo := menuLabelCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// GetSysMenuByRoleName 左侧菜单
func (e *SysMenu) GetSysMenuByRoleName(roleName ...string) ([]models.SysMenu, error) {
	var MenuList []models.SysMenu
	var role models.SysRole
	var err error
	admin := false
	for _, s := range roleName {
		if s == "admin" {
			admin = true
		}
	}

	if len(roleName) > 0 && admin {
		var data []models.SysMenu
		err = e.Orm.Where(" menu_type in ('M','C')").
			Order("sort").
			Find(&data).
			Error
		MenuList = data
	} else {
		err = e.Orm.Model(&role).Preload("SysMenu", func(db *gorm.DB) *gorm.DB {
			return db.Where(" menu_type in ('M','C')").Order("sort")
		}).Where("role_name in ?", roleName).Find(&role).
			Error
		MenuList = *role.SysMenu
	}

	return MenuList, err
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != cModels.Button {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.SysApi = list[j].SysApi
		mi.Children = []models.SysMenu{}

		if mi.MenuType != cModels.Button {
			ms := menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

// GetMenuRole 获取左侧菜单树使用
func (e *SysMenu) GetMenuRole(roleKey string) (m []models.SysMenu, err error) {
	menus, err := e.getByRoleKey(roleKey)
	m = make([]models.SysMenu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}

func (e *SysMenu) getByRoleKey(roleKey string) ([]models.SysMenu, error) {
	var MenuList []models.SysMenu
	var role models.SysRole
	var err error

	if roleKey == "admin" {
		var data []models.SysMenu
		err = e.Orm.Where(" menu_type in ('M','C')").Order("sort").Find(&data).Error
		MenuList = data
	} else {
		role.RoleKey = roleKey
		err = e.Orm.Model(&role).Where("role_key = ? ", roleKey).Preload("SysMenu", func(db *gorm.DB) *gorm.DB {
			return db.Where(" menu_type in ('C')").Order("sort")
		}).Find(&role).Error
		if role.SysMenu != nil {
			MenuList = *role.SysMenu
		}
		mIds := make([]int, 0)
		for _, menu := range MenuList {
			if menu.ParentId != 0 {
				mIds = append(mIds, menu.ParentId)
			}
		}
		var data []models.SysMenu
		err = e.Orm.Where(" menu_type in ('M') and menu_id in ?", mIds).Order("sort").Find(&data).Error
		if err != nil {
			return nil, err
		}
		for _, datum := range data {
			MenuList = append(MenuList, datum)
		}
	}

	return MenuList, err
}
