package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common"
	"net/http"
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

// PresignToken 预签名令牌
// @Summary 预签名令牌
// @Description 预签名令牌
// @Tags 公共接口
// @Param data body PresignTokenRequest true "data"
// @Failure 500
// @Success 200 {object} PresignTokenResponse "{"code": 200, "data": [...]}"
// @Router /presign-token [post]
// @Security Bearer
func PresignToken(c *gin.Context) {
	var req PresignTokenRequest

	// return 400 if binding fails
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	oss := common.Runtime.GetOss()
	if oss == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	token, err := oss.GeneratePresignedToken(req.Directory, req.Filename, int64(req.Duration))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Origin", "*")

	response := PresignTokenResponse {
		Vendor: oss.Name(),
		Token: token,
	}

	c.AbortWithStatusJSON(http.StatusOK, &response)
}
