package models

import (
	"github.com/kingwel-xie/k2/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string   `json:"username" gorm:"uniqueIndex;size:64;comment:用户名"`
	Password string   `json:"-" gorm:"size:128;comment:密码"`
	NickName string   `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string   `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int      `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt     string   `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string   `json:"-" gorm:"comment:头像"`
	Sex      string   `json:"sex" gorm:"size:255;comment:性别"`
	Email    string   `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId   int      `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int      `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string   `json:"remark" gorm:"size:255;comment:备注"`
	Status   string   `json:"status" gorm:"size:4;comment:状态"`
	Token    string   `json:"token" gorm:"size:128;comment:Token"`
	DeptIds  []int    `json:"deptIds" gorm:"-"`
	PostIds  []int    `json:"postIds" gorm:"-"`
	RoleIds  []int    `json:"roleIds" gorm:"-"`
	Dept     	*SysDept `json:"dept"`
	Role     *SysRole `json:"role"`
	models.ControlBy
	models.ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

// Encrypt 加密密码
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.DeptIds = []int{e.DeptId}
	e.PostIds = []int{e.PostId}
	e.RoleIds = []int{e.RoleId}
	return nil
}
