package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_GenerateUniqueID
func Test_GenerateUniqueID(t *testing.T) {
	var digit int = 32
	str,err := GenerateUniqueID(digit)
	if err != nil {
		fmt.Printf("get unique id failed:%v\n",err)
		return 
	}
	fmt.Printf("str:%v\n",str)
}