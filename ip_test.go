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

//go test -v -test.run Test_CountUsableIPs
func Test_CountUsableIPs(t *testing.T) {
   cidr := "172.16.0.0/16"
   count,err := CountUsableIPs(cidr)
   if err != nil {
        fmt.Printf("get ip count failed:%v\n",err)
        return
   }

   fmt.Printf("count;%v\n",count)
}
//go test -v -test.run Test_CalculateSubnets
func Test_CalculateSubnets(t *testing.T) {
   cidr := "172.16.0.0/16"
   subnets,err := CalculateSubnets(cidr,20,5)
   if err != nil {
        fmt.Printf("get ip count failed:%v\n",err)
        return
   }
   for _, subnet := range subnets {
       fmt.Println(subnet.String())
   }
}

//go test -v -test.run Test_CidrFirstAndLastIp
func Test_CidrFirstAndLastIp(t *testing.T) {
    cidr := "172.16.0.0/16"
    firstip,lastip,err := CidrFirstAndLastIp(cidr)
    if err != nil {
        fmt.Printf("get first ip and last ip failed:%v\n",err)
        return
    }
    fmt.Printf("first ip:%v,last ip:%v\n",firstip,lastip)
}
