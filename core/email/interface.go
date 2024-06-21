package email

import (
	"github.com/kingwel-xie/k2/core/logger"
	"io"
)

var log = logger.Logger("email")

type Email interface {
	SendSimple(from, to, subject, content string) error
	Send(from string, to []string, subject, templateFilename string, content map[string]any) error
	SendWithAttachment(from string, to []string, subject, templateFilename string, content map[string]any, r io.Reader, filename string, contentType string) error
}
