package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_CheckMobile
func Test_CheckMobile(t *testing.T) {
    phone := "01023456"
    fmt.Printf("result:%v\n",CheckMobile(phone))
}
//go test -v -test.run Test_CheckEmail
func Test_CheckEmail(t *testing.T) {
    email := "123@qq.com"
    fmt.Printf("result:%v\n",CheckEmail(email))
}
//go test -v -test.run Test_ValidatePassword
func Test_ValidatePassword(t *testing.T) {
    pwd := "ydgw574**"
    err := ValidatePassword(pwd,8)
    if err != nil {
        fmt.Printf("pwd does not meet the requirements:%v\n",err)
        return
    }
    fmt.Printf("pwd avlid")
}
