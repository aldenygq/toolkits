package sort 
import (
  "sort"
)
func SortSliceByChar(slc []string) []string {
	// 使用sort.Slice进行自定义排序
	sort.Slice(slc, func(i, j int) bool {
		return strs[i] < strs[j] // 按照字典序排序
	})

	return slc
}
