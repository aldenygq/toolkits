package toolkits

import (
    "fmt"
    "time"
    "os"
    "bytes"
    "os/exec"
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
        fmt.Printf("read file content failed:%v\n",err)
        return strExpireTime,unixExpireTime,err
    }
    fmt.Printf("original time:%v",origTime)
    t, err := time.Parse("Jan 02 15:04:05 2006 GMT", strings.TrimSuffix(origTime, "\x0a"))
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return strExpireTime,unixExpireTime,err
    }
    unixExpireTime := t.Unix()
    fmt.Printf("unix timestamp:%v\n",unixExpireTime)
    tm := time.Unix(unixExpireTime,0)
    strExpireTime = tm.Format("2006-01-02 15:04:05")
    fmt.Printf("str time:%v\n",strExpireTime)
    return strExpireTime,unixExpireTime,nil
}

func RunCmd(cmdstring string) (string, error) {
        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd := exec.Command("/bin/sh", "-c", cmdstring)
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        err := cmd.Run()
        if err != nil {
                fmt.Printf("err:%v\n",err)
                return fmt.Sprintf("%s",stderr.String()),err
        }
        return fmt.Sprintf("%v",out.String()),nil
}
