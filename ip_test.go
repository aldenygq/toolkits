package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_LocalIP
func Test_LocalIP(t *testing.T) {
	ipinfo,err := LocalIP()
	if err != nil {
		fmt.Printf("get local ip failed:%v\n",err)
		return 
	}
	fmt.Printf("local ip:%v\n",ipinfo.String())
}