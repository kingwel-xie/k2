package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/core/captcha"
)

// GenerateCaptchaHandler 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登陆
// @Success 200 {object} response.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/captcha [get]
func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "验证码获取失败",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 200,
			"data": b64s,
			"id":   id,
			"msg":  "success",
		})
	}
}
