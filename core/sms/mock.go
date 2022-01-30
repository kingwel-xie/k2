package sms

type Mock struct{
}

func (m Mock) SendSms(phoneNumber, signName, templateCode, templateParam string) error {
	log.Debugf("MOCK: sending sms to %s...", phoneNumber)
	return nil
}

func NewMock() Sms {
	return &Mock{  }
}
