package toolkits
import (
    "fmt"
    "testing"
    //"encoding/json"
    "io/ioutil"
    "io"
)
//go test -v -test.run Test_NewHttpClient
func Test_NewHttpClient(t *testing.T) {
    var (
        timeout int = 60
        keepalive int = 60
    )
    client := NewHttpClient(timeout,keepalive)
    if client == nil {
        fmt.Printf("client is nil\n")
        return
    }
    fmt.Printf("new client success\n")

}
//go test -v -test.run Test_GetUrl
func Test_GetUrl(t *testing.T) {
    var (
        timeout int = 60
        keepalive int = 60
    )
    client := NewHttpClient(timeout,keepalive)
    if client == nil {
        fmt.Printf("client is nil\n")
        return
    }

    var (
		u string =  "https://oapi.dingtalk.com/robot/send"
		paramkey string = "access_token"
		paramvalue string = "ac0066b0a6335d82953934adf6dc32e4d8f539851b7bf788f14e0d744a35e5a1"
		queryparam map[string]string = make(map[string]string,0)
	)
    queryparam[paramkey] = paramvalue
    url,err := client.SetQuery(u,queryparam)
	if err != nil {
		fmt.Printf("url set query param failed:%v\n",err)
		return
	}
    fmt.Printf("url:%v\n",url)
}
//go test -v -test.run Test_CutLastestSlash
func Test_CutLastestSlash(t *testing.T) {
    var (
        timeout int = 60
        keepalive int = 60
    )
    client := NewHttpClient(timeout,keepalive)
    if client == nil {
        fmt.Printf("client is nil\n")
        return
    }

    var (
        u string =  "https://oapi.dingtalk.com/robot/send/"
    )

    fmt.Printf("url:%v\n",CutLastestSlash(u))
}
//go test -v -test.run Test_NewReqByMethod
func Test_NewReqByMethod(t *testing.T) {
    var (
        timeout int = 60
        keepalive int = 60
    )
    client := NewHttpClient(timeout,keepalive)
    if client == nil {
        fmt.Printf("client is nil\n")
        return
    }
    fmt.Printf("client:%v\n",&client)
    req,err := client.NewReqByMethod("GET","https://www.baidu.com",nil,nil)
    if err != nil {
        fmt.Printf("new request failed:%v\n",err)
        return
    }
    if req.Body != nil {
        r, err := io.ReadAll(req.Body)
        if err != nil {
            return
        }
        fmt.Printf("request info:%v\n",string(r))
    }

    fmt.Printf("request info:%v\n",req.Header)
    fmt.Printf("request info:%v\n",req.Method)
    fmt.Printf("request info:%v\n",req.URL)
}

//go test -v -test.run Test_DoReq
func Test_DoReq(t *testing.T) {
    var (
        timeout int = 60
        keepalive int = 60
    )
    client := NewHttpClient(timeout,keepalive)
    if client == nil {
        fmt.Printf("client is nil\n")
        return
    }

    resp,err := client.DoReq("GET","https://www.baidu.com",nil,nil,nil)
    if err != nil {
        fmt.Printf("do response failed:%v\n",err)
        return
    }
   defer resp.Body.Close()
   b, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       return
   }
    fmt.Printf("response info:%v\n",string(b))
}
