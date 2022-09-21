package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	cerr "github.com/kingwel-xie/k2/common/error"
)

type PresignTokenRequest struct {
	Directory string `json:"directory"`
	Filename string `json:"filename"`
	Duration int `json:"duration"`
}

type PresignTokenResponse struct {
	Vendor string `json:"vendor"`
	Token interface{} `json:"token"`
}

type PresignToken struct {
	api.Api
}

// PresignToken 预签名令牌
// @Summary 预签名令牌
// @Description 预签名令牌
// @Tags 公共接口
// @Param data body PresignTokenRequest true "data"
// @Failure 500
// @Success 200 {object} PresignTokenResponse "{"code": 200, "data": [...]}"
// @Router /presign-token [post]
// @Security Bearer
func (e PresignToken) PresignToken(c *gin.Context) {
	var req PresignTokenRequest

	oss := common.Runtime.GetOss()
	if oss == nil {
		e.Error(cerr.ErrOssUnavailable)
	}

	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.Error(err)
		return
	}

	//identify := e.GetIdentity()
	token, err := oss.GeneratePresignedToken(req.Directory, req.Filename, int64(req.Duration))
	if err != nil {
		e.Error(err)
		return
	}

	//c.Header("Access-Control-Allow-Methods", "POST")
	//c.Header("Access-Control-Allow-Origin", "*")

	response := PresignTokenResponse {
		Vendor: oss.Name(),
		Token: token,
	}
	e.OK(response, "成功")
}
