package filemanager

import (
	"os"
    "bufio"
    "fmt"
    "io/ioutil"
)

//逐行读
func ScannerReadFile(filepath string) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

//一次性读取,适用于小文件
func FullReadFile(filepath string) (string,error) {
    content, err := ioutil.ReadFile(filepath)
    if err != nil {
        return "",err
    }

    return string(content),nil
}

//缓冲持续读文件，适合大文件持续读,控制速率
func ReadFileByRate(filepath string,rate int) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := make([]byte, rate)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            return err
        }

        fmt.Print(string(buf[:n]))
    }
    return nil
}
