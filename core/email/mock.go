package email

import "io"

type Mock struct {
}

func (m Mock) SendText(from string, to []string, subject, content string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func (m Mock) SendHtml(from string, to []string, subject, content string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func (m Mock) SendAttachment(from string, to []string, subject, content string, r io.Reader, filename string, contentType string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func NewMock() Email {
	return &Mock{}
}
