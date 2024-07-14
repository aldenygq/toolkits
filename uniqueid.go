package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)
func GenerateUniqueID() string {
	// 使用加密安全的随机数生成器生成字节切片
	byteSlice := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, byteSlice); err != nil {
		panic(err)
	}

	// 将字节切片转换为十六进制字符串
	hexStr := hex.EncodeToString(byteSlice)

	// 将十六进制字符串转换为小写
	return hexStr
}