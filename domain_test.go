package toolkits
import (
    "testing"
    "fmt"
)
//go test -v -test.run Test_IsDomainValid
func Test_IsDomainValid(t *testing.T) {
    domain := "a.b.c"
    fmt.Printf("result:%v\n",IsDomainValid(domain))
}
