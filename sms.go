package toolkits

import (
	"errors"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)
type AliyunSmsClient struct {
	AClient *dysmsapi20170525.Client
}
func NewAliyunSmsClient(ak,sk,endpoint string) (*AliyunSmsClient,error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(ak),
		AccessKeySecret: tea.String(sk),
		Endpoint:        tea.String(endpoint),
	}
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return nil,err
	}
	return &AliyunSmsClient{
		AClient:client,
	},nil
}
func (a *AliyunSmsClient) SendSms(phoneNumbers,signName,TemplateCode,conntent string) error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumbers),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(TemplateCode),
		TemplateParam: tea.String(conntent),
	}
	
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _err := a.AClient.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		if *result.Body.Code != "OK" {
			return errors.New(*result.Body.Message)
		}
		return nil
	}()
	
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		result, _err := util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
		return errors.New(*result)
	}
	
	return nil
}
