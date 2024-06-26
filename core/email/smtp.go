package email

import (
	"crypto/tls"
	"io"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SmtpEmail struct {
	Address string
	Host    string
	Auth    smtp.Auth
}

func NewSmtpEmail(address string, identity, username, password, host string) Email {
	return SmtpEmail{
		Address: address,
		Host:    host,
		Auth:    smtp.PlainAuth(identity, username, password, host),
	}
}

func (s SmtpEmail) Send(from string, to []string, cc []string, bcc []string, subject, content string, isHtml bool) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	if isHtml {
		e.HTML = []byte(content)
	} else {
		e.Text = []byte(content)
	}
	// just send
	return e.SendWithTLS(s.Address, s.Auth, &tls.Config{
		ServerName: s.Host,
	})
}

func (s SmtpEmail) SendAttachment(from string, to []string, cc []string, bcc []string, subject, content string, r io.Reader, filename string, contentType string) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(content)

	// attachment
	_, err := e.Attach(r, filename, contentType)
	if err != nil {
		return err
	}
	return e.SendWithTLS(s.Address, s.Auth, &tls.Config{
		ServerName: s.Host,
	})
}
