package execl
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_CreateSheet
func Test_CreateSheet(t *testing.T) {
    file := NewExeclFile()
    err := file.CreateSheet("测试")
    if err != nil {
        fmt.Printf("create sheet fialed:%v\n",err)
        return
    }

    err = file.SaveFile("test.xlsx")
    if err != nil {
        fmt.Printf("save execl file fialed:%v\n",err)
        return
    }

    fmt.Printf("create sheet success")
}

//go test -v -test.run Test_SetSheetHeader
func Test_SetSheetHeader(t *testing.T) {
    f := NewExeclFile()
    err := f.CreateSheet("测试")
    if err != nil {
        fmt.Printf("create sheet fialed:%v\n",err)
        return
    }
    headers := []string{"云厂商","环境","状态"}
    err = f.SetSheetHeader("测试",headers)
    if err != nil {
        fmt.Printf("set sheet header failed:%v\n",err)
        return
    }
    err = f.SaveFile("test.xlsx")
    if err != nil {
        fmt.Printf("save execl file fialed:%v\n",err)
        return
    }

    fmt.Printf("create sheet success")
}
