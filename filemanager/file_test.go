package filemanager
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_ToString
func Test_ToString(t *testing.T) {
	file := "ip_test.go"
    info,err := ReadFileInfo(file)
	if err != nil {
		fmt.Printf("Read file failed:%v\n",err)
		return
	}
	fmt.Printf("file info:%v\n",info)
}
//go test -v -test.run Test_CreateFile
func Test_CreateFile(t *testing.T) {
    filepath := "./yaml_test.go"
    err := CreateFile(filepath)
    if err != nil {
        fmt.Printf("err:%v\n",err)
        return
    }

    fmt.Printf("create file success")
}

//go test -v -test.run Test_AppendWriteFile
func Test_AppendWriteFile(t *testing.T) {
    filepath := "/Users/mingyu/code/src/test.txt"
    content := fmt.Sprintf("hello world")
    err := AppendWriteFile(filepath,content)
    if err != nil {
        fmt.Printf("write file failed:%v\n",err)
        return
    }

    fmt.Printf("write file success")
}

//go test -v -test.run Test_ClearWriteFile
func Test_ClearWriteFile(t *testing.T) {
     filepath := "/Users/mingyu/code/src/test.txt"
     content := fmt.Sprintf("hello world")
     err := ClearWriteFile(filepath,content)
     if err != nil {
         fmt.Printf("write file failed:%v\n",err)
         return
     }

     fmt.Printf("write file success")
}

//go test -v -test.run Test_CopyWriteFile
func Test_CopyWriteFile(t *testing.T) {
    src := "/Users/mingyu/code/src/test.txt"
    dst := "/Users/mingyu/code/src/test1.txt"
    err := CopyWriteFile(src,dst)
    if err != nil {
        fmt.Printf("copy file %v to file %v failed:%v\n",src,dst,err)
        return
    }

    fmt.Printf("copy file content success")

}

//go test -v -test.run Test_BufioWriteFileByRate
func Test_BufioWriteFileByRate(t *testing.T) {
    file := "/Users/mingyu/code/src/test1.txt"
    err := BufioWriteFileByRate(file,"\nni hao",1024)
    if err != nil {
        fmt.Printf("write file failed:%v\n",err)
        return
    }

    fmt.Printf("write file success")
}
//go test -v -test.run Test_RemoveFile
func Test_RemoveFile(t *testing.T) {
     file := "/Users/mingyu/code/src/test1.txt"
     err := RemoveFile(file)
     if err != nil {
        fmt.Printf("delete file failed:%v\n",err)
        return
     }

     fmt.Printf("delete file success")
}
//go test -v -test.run Test_RemoveDirAndFile
func Test_RemoveDirAndFile(t *testing.T) {
    file := "/Users/mingyu/code/src/text"
    err := RemoveDirAndFile(file)
    if err != nil {
        fmt.Printf("delete dir failed:%v\n",err)
        return
    }

    fmt.Printf("delete dir success")
}

//go test -v -test.run Test_ScannerReadFile
func Test_ScannerReadFile(t *testing.T) {
    file := "write_file.go"
    err := ScannerReadFile(file)
    if err != nil {
        fmt.Printf("read file failed:%v\n",err)
        return
    }
}

// go test -v -test.run Test_FullReadFile
func Test_FullReadFile(t *testing.T) {
     file := "write_file.go"
      content,err := FullReadFile(file)
      if err != nil {
        fmt.Printf("read file failed:%v\n",err)
        return
      }

      fmt.Printf("content:%v\n",content)
}
// go test -v -test.run Test_ReadFileByRate
func Test_ReadFileByRate(t *testing.T) {
     file := "write_file.go"
      err := ReadFileByRate(file,1024)
      if err != nil {
        fmt.Printf("read file failed:%v\n",err)
        return
      }

}
