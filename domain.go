package toolkits

import (
    "fmt"
    "net/url"
    "strings"
	"net"
)

// 检查域名格式是否合理
func IsDomainValid(domain string) bool {
    u, err := url.Parse("http://" + domain)
    if err != nil {
        return false
    }
    if u.Host == "" || strings.Contains(u.Host, ":") {
        return false
    }
    return true
}

//检查域名是否存在A记录解析,并输出解析ip列表
func CheckDomainARecord(domain string) ([]string,[]string,error) {
	var (
		ip4Records []string = make([]string,0)
		ip6Records []string = make([]string,0)
	)
	ips, err := net.LookupIP(domain)
    if err != nil {
        fmt.Printf("无法解析域名: %s\n", err)
        return nil,nil,err 
    }
	for _, ip := range ips {
		if ip.To4() != nil {
			ip4Records = append(ip4Records,ip.String())
		} else {
			ip6Records = append(ip6Records,ip.String())
		}
	}

	return ip4Records,ip6Records,nil 
}