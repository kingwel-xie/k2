package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunSMS struct{
	*dysmsapi20170525.Client
}

func NewAliyunSms(accessKeyId, accessKeySecret string) Sms {
	client, err := createClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		panic(err)
	}
	return &AliyunSMS{
		client,
	}
}


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func _main (args []*string) (_err error) {
	client, _err := createClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err
	}
	return _err
}


func (ali *AliyunSMS) SendSms(phoneNumber, signName, templateCode, templateParam string) error {
	// 1.发送短信
	sendReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers: &phoneNumber,
		SignName: &signName,
		TemplateCode: &templateCode,
		TemplateParam: &templateParam,
	}
	_, err := ali.Client.SendSms(sendReq)
	return err
}