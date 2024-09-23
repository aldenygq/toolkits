package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_GetRandomString
func Test_GetRandomString(t *testing.T) {
	fmt.Printf("string:%v\n",GetRandomString(32))
}

//go test -v -test.run Test_GetRandomBoth
func Test_GetRandomBoth(t *testing.T) {
	fmt.Printf("string:%v\n",GetRandomBoth(32))
}

//go test -v -test.run Test_GetRandomNum
func Test_GetRandomNum(t *testing.T) {
	fmt.Printf("string:%v\n",GetRandomNum(6))
}

//go test -v -test.run Test_RandInt
func Test_RandInt(t *testing.T) {
	fmt.Printf("string:%v\n",RandInt(456,987))
}
//go test -v -test.run Test_GetRandomBothSpecChar
func Test_GetRandomBothSpecChar(t *testing.T) {
	fmt.Printf("string:%v\n",GetRandomBothSpecChar(32))
}