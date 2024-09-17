package toolkits
import (
	"gopkg.in/yaml.v2"
	"fmt"
)
func InterfaceToYAML(obj interface{}) (string, error) {
    out, err := yaml.Marshal(obj)
    if err != nil {
		fmt.Printf("parse yaml failed:%v\n",err)
        return "", err
    }
    return string(out), nil
}

func YamlToInterface(data string) (interface{},error) {
	var obj interface{}
	if err := yaml.Unmarshal([]byte(data), &obj); err != nil {
		fmt.Printf("parse yaml failed:%v\n",err)
        return nil,err 
    }
	return obj,nil 
}