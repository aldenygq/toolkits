package toolkits

import (
	"errors"
	"sync"
	"fmt"
    openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencent "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// SmsClient 定义通用短信服务接口
type SmsClient interface {
	SendSms(phoneNumbers, signName, templateCode, templateParam string) error
}


// AliyunSmsClient 阿里云短信服务客户端
type AliyunSmsClient struct {
	 AClient *dysmsapi20170525.Client
}

// NewAliyunSmsClient 创建阿里云短信服务客户端
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
// Send 发送短信
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
	TClient *tencent.Client
    SdkAppId string
}
func NewTencentSmsClient(ak,sk,endpoint,region,sdkAppId string) (*TencentSmsClient,error) {
	credential := common.NewCredential(ak, sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, err := tencent.NewClient(credential, region, cpf)
	if err != nil {
		return nil,err
	}
	return &TencentSmsClient{
		TClient: client,
        SdkAppId: sdkAppId,
	},err
}
func (t *TencentSmsClient) SendSms(signName,templateId,phoneNumbers,content string) error {
	request := tencent.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(t.SdkAppId)
	request.SignName = common.StringPtr(signName)
	request.TemplateId = common.StringPtr(templateId)
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumbers})
	request.TemplateParamSet = common.StringPtrs([]string{content})

	response, err := t.TClient.SendSms(request)
	if err != nil {
		return err
	}
	if *response.Response.SendStatusSet[0].Code != "Ok" {
		return errors.New(*response.Response.SendStatusSet[0].Message)
	}

	return nil
}

// SmsServiceConfig 短信服务配置
type SmsServiceConfig struct {
	//Type SmsServiceType
	Type string
	// AK/SK配置
	AccessKeyId     string
	AccessKeySecret string
	// 其他配置
	Region    string
    Endpoint string
    SdkAppId string
}

// GlobalSmsClient 全局短信服务客户端
var (
	GlobalSmsClient SmsClient
	GlobalMutex     sync.Mutex
)

// InitGlobalSmsClient 初始化全局短信服务客户端
func InitGlobalSmsClient(config SmsServiceConfig) error {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	var err error
	switch config.Type {
	case "aliyun":
		GlobalSmsClient, err = NewAliyunSmsClient(config.AccessKeyId, config.AccessKeySecret,config.Endpoint)
	case "tencent":
		GlobalSmsClient, err = NewTencentSmsClient(config.AccessKeyId, config.AccessKeySecret,config.Endpoint, config.Region,config.SdkAppId)
	default:
		return fmt.Errorf("unsupported SMS service type: %s", config.Type)
	}
	if err != nil {
		return fmt.Errorf("failed to initialize global SMS client: %w", err)
	}
	return nil
}

// SendSms 使用全局短信服务客户端发送短信
func SendSms(phoneNumbers, signName, templateCode, templateParam string) error {
	if GlobalSmsClient == nil {
		return errors.New("global SMS client is not initialized")
	}
    if phoneNumbers == "" {
        return errors.New("phone numbers cannot be empty")
    }
    if signName == "" {
        return errors.New("sign name cannot be empty")
    }
    if templateCode == "" {
        return errors.New("template code cannot be empty")
    }
    if templateParam == "" {
        return errors.New("template param cannot be empty")
    }
	return GlobalSmsClient.SendSms(phoneNumbers, signName, templateCode, templateParam)
}
