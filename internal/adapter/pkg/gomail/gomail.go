package gomail

import (
	"bytes"
	"html/template"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

type Gomail struct {
	message  *gomail.Message
	dialer   *gomail.Dialer
	htmlPath string
}

func NewGomail(conf env.Email) *Gomail {
	return &Gomail{
		message:  gomail.NewMessage(),
		dialer:   gomail.NewDialer(conf.Host, conf.Port, conf.Email, conf.Password),
		htmlPath: conf.HtmlPath,
	}
}

func (g *Gomail) SetSender(sender string) {
	g.message.SetHeader("From", sender)
}

func (g *Gomail) SetReciever(to ...string) {
	g.message.SetHeader("To", to...)
}

func (g *Gomail) SetSubject(subject string) {
	g.message.SetHeader("Subject", subject)
}

func (g *Gomail) SetBodyHTML(path string, data interface{}) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(g.htmlPath + path)
	if err != nil {
		return errx.New(fiber.StatusInternalServerError, "Failed to parse template", err)
	}

	err = t.Execute(&body, data)
	if err != nil {
		return errx.New(fiber.StatusInternalServerError, "Failed to execute template", err)
	}

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) Send() error {
	if err := g.dialer.DialAndSend(g.message); err != nil {
		return errx.New(fiber.StatusInternalServerError, "Failed to send email", err)
	}
	return nil
}
