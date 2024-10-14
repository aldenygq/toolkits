package filemanager

import (
	"os"
    "bufio"
    "io"
    "errors"
    "fmt"
)

//追加写,不会删除已有内容
func AppendWriteFile(filepath,content string) error {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // 确保在函数返回时关闭文件

	// 写入内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

    return nil
}

//清空写,先删除已有内容再写入
func ClearWriteFile(filepath,content string) error {
    file, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close() // 确保在函数返回时关闭文件

    // 写入内容
    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    return nil
}

//缓冲追加写文件，适合高频持续写小文件,不控制写速率
func BufioWriteFile(filepath,content string) error {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
    defer file.Close() // 确保在函数返回时关闭文件

    // 创建一个缓冲写入器，缓冲大小为默认的4096字节
	writer := bufio.NewWriter(file)

	// 写入内容，数据会被写入到缓冲区
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	// 确保缓冲区中的数据被写入到磁盘
	err = writer.Flush()
	if err != nil {
		return err
	}

    return nil
}

//缓冲追加写文件，适合高频持续写,控制速率
func BufioWriteFileByRate(filepath,content string,rate int) error {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
    defer file.Close() // 确保在函数返回时关闭文件

    // 创建一个缓冲写入器，缓冲大小为默认的4096字节
	writer := bufio.NewWriterSize(file, rate)

	// 写入内容，数据会被写入到缓冲区
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	// 确保缓冲区中的数据被写入到磁盘
	err = writer.Flush()
	if err != nil {
		return err
	}

    return nil
}

//拷贝写，复制A文件内容至B文件,追加模式
func CopyWriteFile(srcfile,dstfile string) error {
    src, err := os.Open(srcfile)
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.OpenFile(dstfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()
    _,err = io.Copy(dst, src)
	if err != nil {
		return err
	}
    return nil
}

//删除文件，如文件不存在，则报错
func RemoveFile(filepath string) error {
	// 删除文件
	err := os.Remove(filepath)
	if err != nil {
		// 如果文件不存在，打印错误信息
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("file %v not exist\n", filepath))
		}
		return err
	}

    return nil
}


//删除目录及其子文件
func RemoveDirAndFile(dir string) error {
	// 删除目录及其所有内容
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
    return nil
}
