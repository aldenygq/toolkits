package toolkits

import (
	"fmt"
	"net"
)
func LocalIP() (net.IP, error) {
	tables, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		addrs, err := t.Addrs()
		if err != nil {
			return nil, err
		}
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}
			if v4 := ipnet.IP.To4(); v4 != nil {
				return v4, nil
			}
		}
	}
	return nil, fmt.Errorf("cannot find local IP address")
}

//计算网段中可用IP的数量
func CountUsableIPs(cidr string) (int,error) {
     _, ipnet, err := net.ParseCIDR(cidr)
     if err != nil {
         return 0, err
     }

     // 获取IP地址的位数
     ip := ipnet.IP.To4()
     if ip == nil {
         ip = ipnet.IP.To16()
     }
     mask := ipnet.Mask

     // 计算掩码中1的数量
     ones, _ := mask.Size()

     // 计算总IP数量
     totalIPs := 1 << (len(ip) * 8 - ones)

     // 减去网络地址和广播地址
     usableIPs := totalIPs - 2

     return usableIPs, nil
}

//计算网段起始ip与最后一个ip
func CidrFirstAndLastIp(cidr string) (string,string,error) {
  var err error
  // 解析CIDR
  _, ipnet, err := net.ParseCIDR(cidr)
  if err != nil {
      //fmt.Printf("Error parsing CIDR: %v\n", err)
      return "","",err
  }

  // 获取网络的起始IP（第一个IP）
  firstip := ipnet.IP.String()

  // 计算网络的结束IP（最后一个IP）
  // 通过将网络掩码的位反转并应用到网络的广播地址上来实现
  lastips := make(net.IP, len(ipnet.IP))
  for i := range lastips {
      lastips[i] = ipnet.IP[i] | ^ipnet.Mask[i]
  }
  return firstip,lastips.String(),nil
}

//根据基础网段/子网掩码/子网个数输出子网网段
func CalculateSubnets(baseNetwork string, subnetMaskPrefixLen int, numSubnets int) ([]net.IPNet, error) {
    // 解析基础网段
    _, baseIPNet, err := net.ParseCIDR(baseNetwork)
    if err!= nil {
        return nil, err
    }
    baseIP := baseIPNet.IP.To4()
    if baseIP == nil {
        return nil, fmt.Errorf("invalid IPv4 address in base network")
    }
    newPrefixLen := subnetMaskPrefixLen
    blockSize := 1 << (32 - newPrefixLen)
    subnets := make([]net.IPNet, numSubnets)
    for i := 0; i < numSubnets; i++ {
        // 计算每个子网的网络地址
        subnetIP := make(net.IP, 4)
        copy(subnetIP, baseIP)
        for j := 0; j < 4; j++ {
            subnetIP[j] += byte((i * blockSize) >> (8 * (3 - j)))
        }

        subnets[i].IP = subnetIP
        subnets[i].Mask = net.CIDRMask(newPrefixLen, 32)
    }

    return subnets, nil
}
