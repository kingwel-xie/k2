package email

import (
	"github.com/kingwel-xie/k2/core/logger"
	"io"
)

var log = logger.Logger("email")

type Email interface {
	Send(from string, to []string, cc []string, bcc []string, subject, content string, isHtml bool) error
	SendAttachment(from string, to []string, cc []string, bcc []string, subject, content string, r io.Reader, filename string, contentType string) error
}
