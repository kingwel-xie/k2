package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy string `json:"createBy" gorm:"size:63;index;comment:创建者"`
	UpdateBy string `json:"updateBy" gorm:"size:63;index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy string) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy string) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"index;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index;comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

// ModelTimeHardDelete We'd like to hard delete...
type ModelTimeHardDelete struct {
	CreatedAt time.Time `json:"createdAt" gorm:"index;comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index;comment:最后更新时间"`
}
