package sms

import (
	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("sms")


type Sms interface {
	SendSms(phoneNumber, signName, templateCode, templateParam string) error
}
