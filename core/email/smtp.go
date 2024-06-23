package email

import (
	"io"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SmtpEmail struct {
	Address string
	Sender  string
	Auth    smtp.Auth
}

func NewSmtpEmail(address string, identity, username, password, host, sender string) Email {
	return SmtpEmail{
		Address: address,
		Sender:  sender,
		Auth:    smtp.PlainAuth(identity, username, password, host),
	}
}

func (s SmtpEmail) Send(from string, to []string, cc []string, bcc []string, subject, content string, isHtml bool) error {
	e := email.NewEmail()

	if from != "" {
		e.From = from
	} else {
		e.From = s.Sender
	}
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
	return e.Send(s.Address, s.Auth)
}

func (s SmtpEmail) SendAttachment(from string, to []string, cc []string, bcc []string, subject, content string, r io.Reader, filename string, contentType string) error {
	e := email.NewEmail()

	if from != "" {
		e.From = from
	} else {
		e.From = s.Sender
	}
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
	return e.Send(s.Address, s.Auth)
}
