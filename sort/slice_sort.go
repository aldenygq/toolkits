package sort 
import (
  "sort"
)

//针对slice进行正向排序，slice元素中可能存在字母与数字组成的字符串，如：strs := []string{"banana", "apple1","apple3","apple34","apply","attach","cherry"}
func SortSliceByChar(slc []string) []string {
	// 使用sort.Slice进行自定义排序
	sort.Slice(slc, func(i, j int) bool {
		return strs[i] < strs[j] // 按照字典序排序
	})

	return slc
}
