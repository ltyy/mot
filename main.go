package main

import (
	"fmt"
	"time"

	"github.com/ltyy/mot/ldap"
	"github.com/shirou/gopsutil/host"
)

// 获取启动时间
func getUptime() uint64 {
	up, err := host.BootTime()
	if nil != err {
		return 0
	}
	return uint64(time.Now().Unix()) - up
}

func main() {
	//fmt.Println(getUptime())
	var ad = &ldap.LDAP_CONFIG{
		Addr:       "127.0.0.1:389",
		BaseDn:     "ou=User,dc=wp,dc=cn",
		BindDn:     "cn=admin,dc=wp,dc=cn",
		BindPass:   "king",
		AuthFilter: "(&(uid=%s))",
		Attributes: []string{},
		TLS:        false,
		StartTLS:   false,
	}

	err := ad.Connect()
	defer ad.Close()

	success, err := ad.Auth("hongben", "abc123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)
}
