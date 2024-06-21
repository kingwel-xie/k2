package email

import (
	"bytes"
	"github.com/jordan-wright/email"
	"html/template"
	"io"
	"net/smtp"
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

func (s SmtpEmail) SendSimple(from, to, subject, content string) error {
	e := email.NewEmail()

	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(content)

	// just send
	return e.Send(s.Address, s.Auth)
}

func (s SmtpEmail) SendWithAttachment(from string, to []string, subject, templateFilename string, content map[string]any, r io.Reader, filename string, contentType string) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Subject = subject

	// parse html template
	t, err := template.ParseFiles(templateFilename)
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	err = t.Execute(body, content)
	if err != nil {
		return err
	}
	// html形式的消息
	e.HTML = body.Bytes()

	// attachment
	_, err = e.Attach(r, filename, contentType)
	if err != nil {
		return err
	}
	return e.Send(s.Address, s.Auth)
}

func (s SmtpEmail) Send(from string, to []string, subject, templateFilename string, content map[string]any) error {
	e := email.NewEmail()

	e.From = from
	e.To = to
	e.Subject = subject

	// parse html template
	t, err := template.ParseFiles(templateFilename)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = t.Execute(body, content)
	if err != nil {
		return err
	}
	// html body
	e.HTML = body.Bytes()

	return e.Send(s.Address, s.Auth)
}

//func _main(args []*string) (_err error) {
//	client, _err := NewSmtpEmail()
//	if _err != nil {
//		return _err
//	}
//
//	_, _err = client.SendEmail()
//	if _err != nil {
//		return _err
//	}
//	return _err
//}
