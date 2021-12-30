package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common/api"
	svc "github.com/kingwel-xie/k2/common/service"
	"kobh/models"
	"kobh/service"
	"kobh/service/dto"
	"kobh/x/backlog"
)

type TbxDict struct {
	api.Api
}

// GetAll 全部字典数据 业务页面使用
// @Summary 获取全部字典数据
// @Description 获取全部字典数据
// @Tags 字典数据
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/all-dict [get]
// @Security Bearer
func (e TbxDict) GetAll(c *gin.Context) {
	var svc svc.Service
	e.MakeContext(c).MakeService(&svc)

	var count int64
	var data = make(map[string]interface{})

	// SysDict
	var system = service.SysDictData{Service: svc}
	var systemList []models.SysDictData
	var systemReq = dto.SysDictDataGetPageReq{}
	systemReq.PageIndex = -1
	systemReq.PageSize = -1
	err = system.GetPage(&systemReq, &systemList, &count)
	if err != nil {
		e.Error(err)
		return
	}
	data["systemList"] = systemList

	// TbxCountry
	var country = service.TbxCountry{Service: svc}
	var countryList []models.TbxCountry
	var countryReq = dto.TbxCountryGetPageReq{}
	countryReq.PageIndex = -1
	countryReq.PageSize = -1
	err = country.GetPage(&countryReq, &countryList, &count)
	if err != nil {
		e.Error(err)
		return
	}
	data["countryList"] = countryList

	e.OK(data, "成功")
}
