package config

import (
	"github.com/kingwel-xie/k2/core/email"
)

type Email struct {
	Which string `yaml:"which"`
	Smtp  Smtp   `yaml:"smtp"`
}

var EmailConfig = new(Email)

type Smtp struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Identity string `mapstructure:"identity" json:"identity" yaml:"identity"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Sender   string `mapstructure:"sender" json:"sender" yaml:"sender"`
}

func (e Email) Setup() email.Email {
	switch e.Which {
	case "mock":
		return email.NewMock()
	case "smtp":
		return email.NewSmtpEmail(e.Smtp.Address, e.Smtp.Identity, e.Smtp.Username, e.Smtp.Password, e.Smtp.Host, e.Smtp.Sender)
	default:
		return email.NewMock()
	}
}
