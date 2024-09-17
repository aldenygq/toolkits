package toolkits
import (
    "fmt"
    "testing"
	"strings"
)
//go test -v -test.run Test_GenerateUuid
func Test_GenerateUuid(t *testing.T) {
	fmt.Printf("str:%v\n",strings.ToUpper(GenerateUuid()))
}