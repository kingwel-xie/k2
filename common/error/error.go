package error

import (
	"fmt"
	"github.com/kingwel-xie/k2/core/utils"
	"strings"
)

var (
	ErrBadRequest = New(400, "错误的请求参数", "bad request")
	ErrInternal = New(500, "系统内部错误", "internal server error")
	ErrUnimplemented = New(501, "功能尚未实现", "not implemented yet")
	ErrOssUnavailable = New(505, "OSS不可用", "OSS not available")
	ErrCodeExisted = New(550, "已存在相同编码", "code existed")
	ErrCodeNotFound = New(551, "未找到该编码", "code not found")
	ErrWrongPassword = New(560, "原密码错误", "wrong password")
	ErrMismatchPassword = New(561, "两次输入的密码不匹配", "passwords mismatch")
	ErrDatabase = New(562, "数据库内部错误", "db error")
	ErrNoSuchObject = New(563, "对象不存在", "no such object")
	ErrPermissionDenied = New(564, "对象不存在或无权查看", "no such object or no permission")
	ErrNoPermission = New(403, "对不起，您没有该接口访问权限，请联系管理员", "You don't have the permission to access the interface. Please contact the administrator.")
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
	if mode != utils.ModeProd.String() {
		msg = fmt.Sprintf("%s : %d", msg, e.code)
		if e.err != nil {
			msg += ": " + e.err.Error()
		}
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

func (e bizError) Wrapf(format string, a ...interface{}) *bizError {
	e.err = fmt.Errorf(format, a...)
	return &e
}

func (e bizError) Is(err error) bool {
	if x, ok := err.(*bizError); ok && x.code == e.code {
		return true
	}
	return false
}