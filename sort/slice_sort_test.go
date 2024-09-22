package sort
import (
    "testing"
    "fmt"
)
//go test -v -test.run Test_SortSliceByChar
func Test_SortSliceByChar(t *testing.T) {
	strs := []string{"banana", "apple1","apple3","apple34","apply","attach","cherry"}
	fmt.Printf("slcs:%v\n",SortSliceByChar(strs))
}