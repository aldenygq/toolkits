package filemanager

import (
	"io/ioutil"
	"os"
	"encoding/json"
    "errors"
    "fmt"
)

//目录是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/**
创建文件夹
*/
func CreateDir(dirName string) error {
    if IsDir(dirName) {
        return errors.New(fmt.Sprintf("dir %v is exist",dirName))
    }
	err := os.Mkdir(dirName, 755)
	if err != nil {
		return err
	}
	return nil
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func ReadFileInfo(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(fd), nil
}

func ReadFileToStruct(file string, data interface{}) error {
	filecontent, err := ReadFileInfo(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(filecontent), &data)
	if err != nil {
		return err
	}

	return nil
}

//创建空文件
func CreateFile(filepath string) error {
    if IsExist(filepath) {
        return errors.New(fmt.Sprintf("file %v is exist",filepath))
    }
    file, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer file.Close() // 确保文件在函数结束时关闭
    return nil
}
