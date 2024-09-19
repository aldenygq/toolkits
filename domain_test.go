package toolkits
import (
    "testing"
    "fmt"
)
//go test -v -test.run Test_CheckDomainValid
func Test_CheckDomainValid(t *testing.T) {
    domain := "a.b.c"
    fmt.Printf("result:%v\n",CheckDomainValid(domain))
}
