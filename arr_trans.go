package toolkits
import (
  "strings"
  ""
)
//检测数组中是否存在某元素，如存在，返回true,，不存在返回false
func CheckSlcHasStr(a string, b []string) bool {
	for _, v := range b {
		if strings.Contains(v, a) {
			return true
		}
	}

	return false
}
//去除数组中的空元素
func RemovEmpty(str []string) []string {
	var result []string = make([]string, 0)
	for i, _ := range str {
		if str[i] != "" {
			result = append(result, str[i])
		} else {
			continue
		}
	}

	return result
}
//校验二维slce存在重复元素
func CheckDuplicates(slice [][]string) bool {
	allElements := make([]string, 0)
	for _, innerSlice := range slice {
		allElements = append(allElements, innerSlice...)
	}
	sort.Strings(allElements)

	for i := 1; i < len(allElements); i++ {
		if allElements[i-1] == allElements[i] {
			return true
		}
	}
	return false
}
