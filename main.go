package main

import (
	"fmt"
	"time"

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
	fmt.Println(getUptime())
}
