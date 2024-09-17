package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_RunCmd
func Test_RunCmd(t *testing.T) {
	script := fmt.Sprintf("ls -al")
	output,err := RunCmd(script)
	if err != nil {
		fmt.Printf("run cmd failed:%v\n",err)
		return 
	}
	fmt.Printf("run cmd success,output:%v\n",output)
}