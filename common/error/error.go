package error

import (
	"fmt"
	"github.com/kingwel-xie/k2/core/utils"
)

var (
	BadRequestError = New(400, "错误的请求参数", "bad request")
	InternalServerError = New(500, "系统内部错误", "internal server error")
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
			"zh-CN": cnMsg,
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
	msg, ok := e.messages[lang];
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
