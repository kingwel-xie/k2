package email

import (
	"io"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SmtpEmail struct {
	Address string
	Auth    smtp.Auth
}

func NewSmtpEmail(address string, identity, username, password, host string) Email {
	return SmtpEmail{
		Address: address,
		Auth:    smtp.PlainAuth(identity, username, password, host),
	}
}

func (s SmtpEmail) SendText(from string, to []string, subject, content string) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Subject = subject
	e.Text = []byte(content)

	// just send
	return e.Send(s.Address, s.Auth)
}

func (s SmtpEmail) SendHtml(from string, to []string, subject, content string) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Subject = subject
	e.HTML = []byte(content)

	// just send
	return e.Send(s.Address, s.Auth)
}

func (s SmtpEmail) SendAttachment(from string, to []string, subject, content string, r io.Reader, filename string, contentType string) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Subject = subject
	e.HTML = []byte(content)

	// attachment
	_, err := e.Attach(r, filename, contentType)
	if err != nil {
		return err
	}
	return e.Send(s.Address, s.Auth)
}
