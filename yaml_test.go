package toolkits
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_InterfaceToYAML
func Test_InterfaceToYAML(t *testing.T) {
	var a map[string]interface{} = make(map[string]interface{},0)
	a["name"] = "alden"
	a["age"] = 29
	a["friend"] = []string{"summer","ycg"}
	strYaml,err := InterfaceToYAML(a)
	if err != nil {
		fmt.Printf("interface to ymal failed:%v\n",err)
		return 
	}
	fmt.Printf("yaml info:\n%v\n",strYaml)
}