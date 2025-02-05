package toolkits

import (
	"fmt"
	"log"
    "testing"
)

func Test_SendSms(t *testing.T){
    var cnf SmsServiceConfig
    cnf.Cloud = "aliyun"
    cnf.AccessKeyId = "xxxxxx"
    cnf.AccessKeySecret = "xxxxxxxx"
    cnf.Endpoint = "xxxxxxxx"
    cnf.Region = "xxxxxxx"
    cnf.SdkAppId = "xxxxxxx"
    err := InitGlobalSmsClient(cnf)
    if err != nil {
        fmt.Printf("init sms client failed:%v\n",err)
        return
    }

	// 发送短信
	err = SendSms("+8613800138000", "your-sign-name", "your-template-code", "123456")
	if err != nil {
		log.Fatalf("failed to send SMS: %v", err)
	    return
    }
	fmt.Println("SMS sent successfully")
}
