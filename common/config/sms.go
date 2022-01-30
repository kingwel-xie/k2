package config

import (
	"github.com/kingwel-xie/k2/core/sms"
)

type Sms struct {
	Which string 			`yaml:"which"`
	Mock   Mock      		`yaml:"mock"`
	Aliyun AliyunSMS 		`yaml:"aliyun"`
}

var SmsConfig = new(Sms)


type Mock struct {
	SSS string  				`yaml:"sss"`
}

type AliyunSMS struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
}

func (e Sms) Setup() sms.Sms {
	switch e.Which {
	case "mock":
		return sms.NewMock()
	case "aliyun":
		return sms.NewAliyunSms(e.Aliyun.AccessKeyId, e.Aliyun.AccessKeySecret)
	default:
		return sms.NewMock()
	}
}