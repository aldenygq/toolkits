package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_StrToInt
func Test_StrToInt(t *testing.T) {
    fmt.Printf("num:%v\n",StrToInt("4"))
}

//go test -v -test.run Test_StrToInt64
func Test_StrToInt64(t *testing.T) {
    num,err := StrToInt64("400")
    if err != nil {
        fmt.Printf("string to int64 failed:%v\n",err)
        return
    }
    fmt.Printf("num:%v\n",num)
}
