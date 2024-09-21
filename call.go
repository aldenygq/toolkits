package toolkits

import (
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  dyvmsapi20170525  "github.com/alibabacloud-go/dyvmsapi-20170525/v3/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  "github.com/alibabacloud-go/tea/tea"
)
type DyvmsConfig struct {
    Ak string
    Sk string
    DyvmsHost string
}
type DyvmsTtl struct {
    CalledShowNumber string
    CalledNumber string
    TtsCode string
    Message string
    PlayTimes int
}

func NewDyvmsClient(dyvms DyvmsConfig) (*dyvmsapi20170525.Client, error) {
    var (
        err error
        client *dyvmsapi20170525.Client = &dyvmsapi20170525.Client{}
    )
    config := &openapi.Config{
        // 必填，您的 AccessKey ID
        AccessKeyId: &dyvms.Ak,
        // 必填，您的 AccessKey Secret
        AccessKeySecret: &dyvms.Sk,
    }
     // Endpoint 请参考 https://api.aliyun.com/product/Dyvmsapi
     config.Endpoint = tea.String(dyvms.DyvmsHost)
     client, err = dyvmsapi20170525.NewClient(config)
     if err != nil {
        return nil,err
     }
     return client, nil
}
func SingleCallByTts(client *dyvmsapi20170525.Client,param DyvmsTtl) error {
    var times int32 = int32(param.PlayTimes)
    req := &dyvmsapi20170525.SingleCallByTtsRequest{
        CalledShowNumber: &param.CalledShowNumber,
        CalledNumber: &param.CalledNumber,
        TtsCode: &param.TtsCode,
        TtsParam: &param.Message,
        PlayTimes: &times,
    }
    resp, err := client.SingleCallByTts(req)
    if err != nil {
        return err
    }
    if !tea.BoolValue(util.EqualString(resp.Body.Code, tea.String("OK"))) {
        err = tea.NewSDKError(map[string]interface{}{
            "code": tea.StringValue(resp.Body.Code),
            "message": tea.StringValue(resp.Body.Message),
        })
        return err
    }
    return  nil
}
