package email

import (
	"github.com/kingwel-xie/k2/core/logger"
	"io"
)

var log = logger.Logger("email")

type Email interface {
	SendText(from string, to []string, subject, content string) error
	SendHtml(from string, to []string, subject, content string) error
	SendAttachment(from string, to []string, subject, content string, r io.Reader, filename string, contentType string) error
}
