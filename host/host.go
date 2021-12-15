package host

import (
	"fmt"
	"github.com/cnsilvan/UnblockNeteaseMusic/common"
	"log"
	"net"
)

// Exclude UnblockNetEaseMusic related host
func resolveIps() error {
	for domain, _ := range common.HostDomain {
		rAddr, err := net.ResolveIPAddr("ip", domain)
		if err != nil {
			log.Printf("Fail to resolve %s, %s\n", domain, err)
			return err
		}
		if len(rAddr.IP) == 0 {
			log.Printf("Fail to resolve %s,IP nil\n", domain)
			return fmt.Errorf("Fail to resolve  %s,Ip length==0 \n", domain)
		}
		ip := rAddr.IP.String()
		if ip == "127.0.0.1" {
			panic(fmt.Sprintf("%v ip:%v is error", domain, ip))
		}
		common.HostDomain[domain] = rAddr.IP.String()

	}
	return nil
}
func InitHosts() error {
	log.Println("-------------------Init Host-------------------")
	err := resolveIps()
	if err != nil {
		return err
	}
	log.Println("HostDomain:", common.HostDomain)
	return err
}
