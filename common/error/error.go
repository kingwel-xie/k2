package error

type bizError struct {
	messages map[string]string
	code int
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

func (e *bizError) Message(lang string) string {
	if msg, ok := e.messages[lang]; ok {
		return msg
	}
	return e.messages["__default__"]
}

func (e *bizError) Error() string {
	return e.messages["__default__"]
}
