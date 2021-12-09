package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kingwel-xie/k2/common/config"
	cerr "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/core/tools/language"
	"github.com/kingwel-xie/k2/core/utils"

)

// var DefaultLanguage = "zh-CN"
var DefaultLanguage = "en"

var Default = &response{}

type bizError interface {
	Code() int
	Message(mode string, lang string) string
}

// Error 失败数据处理
func Error(c *gin.Context, err error) {
	res := Default.Clone()
	res.SetTraceID(utils.GenerateMsgIDFromContext(c))

	e, ok := err.(bizError)
	if !ok {
		e = cerr.InternalServerError.Wrap(err)
	}

	res.SetMsg(e.Message(config.ApplicationConfig.Mode, getAcceptLanguage(c)))
	res.SetCode(int32(e.Code()))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", e.Code())
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(utils.GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custom 兼容函数
func Custom(c *gin.Context, data gin.H) {
	data["requestId"] = utils.GenerateMsgIDFromContext(c)
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}
// getAcceptLanguage 获取当前语言
func getAcceptLanguage(c *gin.Context) string {
	languages := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}

