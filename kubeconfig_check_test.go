package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_CheckKubeconfigByFile
func Test_CheckKubeconfigByFile(t *testing.T) {
    filepath := "~/.kube/config"
    strtime,unixtime,err := CheckKubeconfigByFile(filepath)
    if err != nil {
        fmt.Printf("check kubeconfig file failed:%v\n",err)
        return
    }

    fmt.Printf("string time:%v\n",strtime)
    fmt.Printf("unix time:%v\n",unixtime)
}
