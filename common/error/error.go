package error

import (
	"fmt"
	"github.com/kingwel-xie/k2/core/utils"
	"strings"
)

var (
	ErrBadRequest = New(400, "错误的请求参数", "bad request")
	ErrInternal = New(500, "系统内部错误", "internal server error")
	ErrCodeExisted = New(550, "已存在相同编码", "code existed")
	ErrCodeNotFound = New(551, "未找到该编码", "code not found")
	ErrWrongPassword = New(560, "原密码错误", "wrong password")
	ErrMismatchPassword = New(561, "两次输入的密码不匹配", "passwords mismatch")
	ErrDatabase = New(562, "数据库内部错误", "db error")
	ErrNoSuchObject = New(563, "对象不存在", "no such object")
	ErrPermissionDenied = New(564, "对象不存在或无权查看", "no such object or no permission")
)

type bizError struct {
	messages map[string]string
	code int
	err error
}

func New(code int, cnMsg string, enMsg string) *bizError {
	return &bizError{
		messages: map[string]string{
			"__default__": enMsg,
			"zh-cn": cnMsg,
			"en": enMsg,
		},
		code:     code,
	}
}

func (e *bizError) AddMessage(lang, msg string)  {
	e.messages[lang] = msg
}

func (e *bizError) Code() int {
	return e.code
}

func (e *bizError) Message(mode string, lang string) string {
	msg, ok := e.messages[strings.ToLower(lang)]
	if !ok {
		msg = e.messages["__default__"]
	}
	if mode != utils.ModeProd.String() && e.err != nil {
		msg = fmt.Sprintf("%s: %s", msg, e.err)
	}
	return msg
}

func (e *bizError) Error() string {
	return e.Message(utils.ModeDev.String(), "__default__")
}

func (e bizError) Wrap(err error) *bizError {
	e.err = err
	return &e
}
