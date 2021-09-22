package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kingwel-xie/k2/core/storage"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/models"
)

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:标题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Type   string `json:"type" gorm:"size:16;comment:接口类型"`
	Action string `json:"action" gorm:"size:16;comment:请求类型"`
	models.ModelTime
	models.ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}

func SaveSysApi(message storage.Messager) (err error) {
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		fmt.Errorf("json Marshal error, %s", err.Error())
		return err
	}

	var l common.Routers
	err = json.Unmarshal(rb, &l)
	if err != nil {
		fmt.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	db := common.Runtime.GetDb()
	for _, v := range l.List {
		if v.HttpMethod != "HEAD" ||
			strings.Contains(v.RelativePath, "/swagger/") ||
			strings.Contains(v.RelativePath, "/static/") ||
			strings.Contains(v.RelativePath, "/form-generator/") {
			err := db.Where(SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
				Attrs(SysApi{Handle: v.Handler, Type: "CHECK"}).
				FirstOrCreate(&SysApi{}).
				//Update("handle", v.Handler).
				Error
			if err != nil {
				err := fmt.Errorf("Models SaveSysApi error: %s \r\n ", err.Error())
				return err
			}
		}
	}

	return nil
}
