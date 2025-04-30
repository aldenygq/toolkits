package slice 

import (
  
)

//字符串切片，根据字符串值查找字符串在slice的位置
func IndexOfSliace(arr []string, target string) int {
	for index, value := range arr {
		if value == target {
			return index
		}
	}
	return -1 // 未找到
}
