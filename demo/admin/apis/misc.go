package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common/api"
	svc "github.com/kingwel-xie/k2/common/service"
	"admin/models"
	"admin/service"
	"admin/service/dto"
)

type TbxMisc struct {
	api.Api
}

// GetAll 全部字典数据 业务页面使用
// @Summary 获取全部字典数据
// @Description 获取全部字典数据
// @Tags 其他
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/all-dict [get]
// @Security Bearer
func (e TbxMisc) GetAll(c *gin.Context) {
	var svc svc.Service
	e.MakeContext(c).MakeService(&svc)

	// FIXME
	data := backlog.LoadAllDict()

	// we don't want to leak sysUser
	delete(data, "userList")

	//var count int64
	//var data = make(map[string]interface{})
	//
	//// SysDict
	//var system = service.SysDictData{Service: svc}
	//var systemList []models.SysDictData
	//var systemReq = dto.SysDictDataGetPageReq{}
	//systemReq.PageIndex = -1
	//systemReq.PageSize = -1
	//err = system.GetPage(&systemReq, &systemList, &count)
	//if err != nil {
	//	e.Error(err)
	//	return
	//}
	//data["systemList"] = systemList

	e.OK(data, "成功")
}

// LimitedDownload 受限下载
// @Summary 受限下载
// @Description 受限下载
// @Tags 其他
// @Param id path string true "uuid"
// @Success 200
// @Failure 503
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/limited-download/{uuid} [get]
// @Security Bearer
func (e TbxMisc) LimitedDownload(c *gin.Context) {
	s := service.TbxMisc{}
	req := dto.TbxLimitedDownloadReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		_ = e.Context.AbortWithError(400, err)
		return
	}

	err = s.LimitedDownload(&req)
	if err != nil {
		_ = e.Context.AbortWithError(500, err)
		return
	}
	e.Context.Data(200, req.ContentType, req.Data)
}