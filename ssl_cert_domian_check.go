package toolkits

import (
    "crypto/tls"
    "crypto/x509"
	"encoding/pem"
    "fmt"
    "net/http"
    "time"
    //"encoding/json"
    "strings"
    "errors"
    //"os"
    "io/ioutil"
)

// 检查证书有效期的函数
func CheckCertValidityByDomain(domain string) (map[string]interface{},error) {
    var (
        err error
        domainmap map[string]interface{} = make(map[string]interface{},0)
    )
    // 创建一个新的HTTPS客户端
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true, // 为了简单起见，这里忽略证书验证
        },
    }
    client := &http.Client{Transport: tr}

    // 尝试获取证书信息
    resp, err := client.Get("https://" + domain)
    if err != nil {
        fmt.Printf("Error checking certificate for %s: %v\n", domain, err)
        return nil,err
    }
    defer resp.Body.Close()

    // 获取证书链
    certs := resp.TLS.PeerCertificates
    if len(certs) == 0 {
        fmt.Printf("no certificates found for %s\n", domain)
        return nil,errors.New(fmt.Sprintf("no certificates found for %s\n", domain))
    }
    // 打印证书有效期
    for _, cert := range certs {
        if !strings.Contains(cert.Subject.CommonName,strings.Split(domain,".")[len(certs)-2]) {
            continue
        }
        strtime,_ := time.Parse("2006-01-02T15:04:05Z",cert.NotAfter.Format(time.RFC3339))
        domainmap["domain"] = domain
        domainmap["unixtime"] = strtime.Unix()
        domainmap["strtime"] = strtime.Format("2006-01-02 15:04:05")
    }
    return domainmap,nil
}

func CheckCertValidityByPem(filepath string) (map[string]interface{},error) {
     var (
         err error
         domainmap map[string]interface{} = make(map[string]interface{},0)
     )
    // 读取PEM格式的证书文件
	certPEM, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading certificate file:", err)
		return nil ,err
	}

	// 解码PEM数据
	block, _ := pem.Decode(certPEM)
	if block == nil {
		fmt.Println("Failed to parse certificate PEM")
		return nil,err
	}

	// 解析证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing certificate:", err)
		return nil,err
	}
    strtime,_ := time.Parse("2006-01-02T15:04:05Z",cert.NotAfter.Format(time.RFC3339))
    domainmap["pem_file"] = filepath
    domainmap["unixtime"] = strtime.Unix()
    domainmap["strtime"] = strtime.Format("2006-01-02 15:04:05")
    return domainmap,nil
}
