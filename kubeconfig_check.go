package toolkits

import (
    "fmt"
    "time"
    //"os"
    //"bytes"
    //"os/exec"
    "strings"
)

func CheckKubeconfigByFile(filepath string) (string,int64,error) {
    var (
      err error
      strExpireTime string
      unixExpireTime int64
    )
    // 加载Kubeconfig文件
    fileContentCmd := fmt.Sprintf("cat %v | grep client-certificate-data | awk -F ' ' '{print $2}' |base64 -d| openssl x509 -text -noout -dates | grep After |awk -F '=' '{print $2}' | grep -v '^$'",filepath)
    origTime,err := RunCmd(fileContentCmd)
    if err != nil {
        return strExpireTime,unixExpireTime,err
    }
    t, err := time.Parse("Jan 2 15:04:05 2006 GMT", strings.TrimSuffix(origTime, "\x0a"))
    if err != nil {
        return strExpireTime,unixExpireTime,err
    }
    unixExpireTime = t.Unix()
    tm := time.Unix(unixExpireTime,0)
    strExpireTime = tm.Format("2006-01-02 15:04:05")
    return strExpireTime,unixExpireTime,nil
}
