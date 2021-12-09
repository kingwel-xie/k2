package service

import (
	cDto "github.com/kingwel-xie/k2/common/dto"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"

	"admin/models"
	"admin/service/dto"
)

type SysDept struct {
	service.Service
}

// GetPage 获取SysDept列表
//func (e *SysDept) GetPage(c *dto.SysDeptGetPageReq, list *[]models.SysDept) error {
//	var err error
//	var data models.SysDept
//
//	err = e.Orm.Model(&data).
//		Scopes(
//			cDto.MakeCondition(c.GetNeedSearch()),
//		).
//		Find(list).Error
//	if err != nil {
//		e.Log.Errorf("db error:%s", err)
//		return err
//	}
//	return nil
//}

// Get 获取SysDept对象
func (e *SysDept) Get(d *dto.SysDeptGetReq, model *models.SysDept) error {
	err := e.Orm.First(model, d.GetId()).Error
	return err
}

// Insert 创建SysDept对象
func (e *SysDept) Insert(c *dto.SysDeptInsertReq) error {
	err := e.Orm.Transaction(func(tx *gorm.DB) error {
		var err error
		var data models.SysDept
		c.Generate(&data)
		data.SetCreateBy(e.Identity.UserId)

		err = tx.Create(&data).Error
		if err != nil {
			return err
		}
		deptPath := utils.IntToString(data.DeptId) + "/"
		if data.ParentId != 0 {
			var deptP models.SysDept
			tx.First(&deptP, data.ParentId)
			deptPath = deptP.DeptPath + deptPath
		} else {
			deptPath = "/0/" + deptPath
		}
		err = tx.Model(&data).Update("dept_path", deptPath).Error
		return err
	})

	return err
}

// Update 修改SysDept对象
func (e *SysDept) Update(c *dto.SysDeptUpdateReq) error {
	var model = models.SysDept{}
	err := e.Orm.First(&model, c.GetId()).Error
	if err != nil {
		return err
	}
	c.Generate(&model)
	model.SetUpdateBy(e.Identity.UserId)

	deptPath := utils.IntToString(model.DeptId) + "/"
	if model.ParentId != 0 {
		var deptP models.SysDept
		err = e.Orm.First(&deptP, model.ParentId).Error
		if err != nil {
			return err
		}
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	model.DeptPath = deptPath
	db := e.Orm.Save(&model)
	if db.Error != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// Remove 删除SysDept
func (e *SysDept) Remove(d *dto.SysDeptDeleteReq) error {
	var data models.SysDept
	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return k2Error.ErrPermissionDenied
	}
	return nil
}

// GetSysDeptList 获取组织数据
func (e *SysDept) getList(c *dto.SysDeptGetPageReq, list *[]models.SysDept) error {
	var data models.SysDept
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error

	return err
}

// SetDeptTree 设置组织数据
func (e *SysDept) SetDeptTree(c *dto.SysDeptGetPageReq) (m []dto.DeptLabel, err error) {
	var list []models.SysDept
	err = e.getList(c, &list)

	m = make([]dto.DeptLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.DeptLabel{}
		e.Id = list[i].DeptId
		e.Label = list[i].DeptName
		deptsInfo := deptTreeCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// Call 递归构造组织数据
func deptTreeCall(deptList *[]models.SysDept, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.DeptLabel{Id: list[j].DeptId, Label: list[j].DeptName, Children: []dto.DeptLabel{}}
		ms := deptTreeCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

// SetDeptPage 设置dept页面数据
func (e *SysDept) SetDeptPage(c *dto.SysDeptGetPageReq) (m []models.SysDept, err error) {
	var list []models.SysDept
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.deptPageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

func (e *SysDept) deptPageCall(deptlist *[]models.SysDept, menu models.SysDept) models.SysDept {
	list := *deptlist
	min := make([]models.SysDept, 0)
	for j := 0; j < len(list); j++ {
		if menu.DeptId != list[j].ParentId {
			continue
		}
		mi := models.SysDept{}
		mi.DeptId = list[j].DeptId
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysDept{}
		ms := e.deptPageCall(deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

// GetRoleDeptId 获取角色的部门ID集合
func (e *SysDept) GetWithRoleId(roleId int) ([]int, error) {
	deptIds := make([]int, 0)
	deptList := make([]dto.DeptIdList, 0)
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, nil
}

func (e *SysDept) SetDeptLabel() (m []dto.DeptLabel, err error) {
	list := make([]models.SysDept, 0)
	err = e.Orm.Find(&list).Error
	if err != nil {
		return
	}
	m = make([]dto.DeptLabel, 0)
	var item dto.DeptLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
		item = dto.DeptLabel{}
		item.Id = list[i].DeptId
		item.Label = list[i].DeptName
		deptInfo := deptLabelCall(&list, item)
		m = append(m, deptInfo)
	}
	return
}

// deptLabelCall
func deptLabelCall(deptList *[]models.SysDept, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	var mi dto.DeptLabel
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi = dto.DeptLabel{Id: list[j].DeptId, Label: list[j].DeptName, Children: []dto.DeptLabel{}}
		ms := deptLabelCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}
