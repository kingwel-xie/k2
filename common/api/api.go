package api

import (
	"errors"
	"fmt"
	"io"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"github.com/kingwel-xie/k2/common/response"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/logger"
	"github.com/kingwel-xie/k2/core/tools/language"
	"github.com/kingwel-xie/k2/core/utils"
)

var DefaultLanguage = "zh-CN"

func init() {
	vd.SetErrorFactory(func(failPath, msg string) error {
		return fmt.Errorf(`validation failed: %s %s`, failPath, msg)
	})
}

type Api struct {
	Context *gin.Context
	Logger  *logger.KLogger
	Orm     *gorm.DB
	Errors  error
}

func (e *Api) SetError(err error) {
	e.Errors = err
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = MustGetLogger(c)
	e.Orm = MustGetOrm(c)
	return e
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && errors.Is(err, io.EOF) {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			e.SetError(err)
			return e
		}
	}
	if err1 := vd.Validate(d); err1 != nil {
		e.SetError(err1)
	}
	return e
}

func (e *Api) MakeService(c *service.Service) *Api {
	c.Orm = e.Orm
	c.Identity = service.GetIdentity(e.Context)
	return e
}

func (e *Api) GetIdentity() *service.AuthIdentity {
	return service.GetIdentity(e.Context)
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	// 500 means something wrong inside, otherwise, it is a binging error
	if code == 400 {
		e.Logger.Warnw(msg, "code", code, "error", err)
	} else {
		e.Logger.Errorw(msg, "code", code, "error", err)
	}
	response.Error(e.Context, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custom(e.Context, data)
}

func (e Api) Translate(form, to interface{}) {
	utils.Translate(form, to)
}

// getAcceptLanguage 获取当前语言
func (e *Api) getAcceptLanguage() string {
	languages := language.ParseAcceptLanguage(e.Context.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}

// MustGetLogger 获取上下文提供的日志
func MustGetLogger(c *gin.Context) *logger.KLogger {
	var log *logger.KLogger
	l, ok := c.Get(utils.LoggerKey)
	if !ok {
		panic("no logger found in gin.context, shouldn't happen!!!")
	}

	log, ok = l.(*logger.KLogger)
	if !ok {
		panic("no logger found in gin.context, shouldn't happen!!!")
	}
	return log
}

// GetOrm 获取orm连接
func MustGetOrm(c *gin.Context) *gorm.DB {
	idb := c.MustGet("db")
	switch idb.(type) {
	case *gorm.DB:
		return idb.(*gorm.DB)
	default:
		panic("WTF, bad db??")
	}
}
