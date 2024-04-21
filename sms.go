package toolkits

import (
	"errors"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tentcent "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)
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

type TencentSmsClient struct {
	TClient *tentcent.Client
}
func NewTencentSmsClient(ak,sk,endpoint,region string) (*TencentSmsClient,error) {
	credential := common.NewCredential(ak, sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, err := tencent.NewClient(credential, region, cpf)
	if err != nil {
		retturn nil,err
	}
	return &TencentSmsClient{
		TClient: client,
	},err
}
func (t *TencentSmsClient) TencentSendSmsCode(appid,signName,templateId,phoneNumbers code string) error {
	request := t.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(appid)
	request.SignName = common.StringPtr(signName)
	request.TemplateId = common.StringPtr(templateId)
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumbers})
	request.TemplateParamSet = common.StringPtrs([]string{code})
	
	response, err := t.SendSms(request)
	if err != nil {
		return err
	}
	if *response.Response.SendStatusSet[0].Code != "Ok" {
		return errors.New(*response.Response.SendStatusSet[0].Message)
	}
	
	return nil
}
