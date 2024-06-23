package email

import "io"

type Mock struct {
}

func (m Mock) Send(from string, to, cc, bcc []string, subject, content string, isHtml bool) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func (m Mock) SendAttachment(from string, to, cc, bcc []string, subject, content string, r io.Reader, filename string, contentType string) error {
	log.Debugf("MOCK: sending email to %s...", to)
	return nil
}

func NewMock() Email {
	return &Mock{}
}
