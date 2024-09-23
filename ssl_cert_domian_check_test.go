package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_CheckCertValidityByDomain
func Test_CheckCertValidityByDomain(t *testing.T) {
    domain := "www.baidu.com"
    info,err := CheckCertValidityByDomain(domain)
    if err != nil {
        fmt.Printf("check domain ssl cert validity failed:%v\n",err)
        return
    }

    fmt.Printf("info:%v\n",info)
}

//go test -v -test.run Test_CheckCertValidityByPem
func Test_CheckCertValidityByPem(t *testing.T) {
    path := "./*.pem"
    info,err := CheckCertValidityByPem(path)
    if err != nil {
        fmt.Printf("check domain ssl cert validity failed:%v\n",err)
        return
    }

    fmt.Printf("info:%v\n",info)
}
