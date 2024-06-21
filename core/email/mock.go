package email

import "io"

type Mock struct {
}

func (m Mock) SendSimple(from, to, subject, content string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func (m Mock) Send(from string, to []string, subject, templateFilename string, content map[string]any) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func (m Mock) SendWithAttachment(from string, to []string, subject, templateFilename string, content map[string]any, r io.Reader, filename string, contentType string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func NewMock() Email {
	return &Mock{}
}
