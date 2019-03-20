package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func LoadConfig(filename string) (ldap.LDAP_CONFIG, error) {
	var config ldap.LDAP_CONFIG
	//configFile, err := ioutil.ReadFile(filename) //slice
	configFile, err := os.Open(filename) // * File
	defer configFile.Close()

	if err != nil {
		return config, err
	}
	//json.Marshal --- slice
	//json.Unmarshal  --- slice
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config) //store it in struct config
	return config, err
}

func main() {
	//fmt.Println(getUptime())

	ad, err := LoadConfig("config.json")
	if err != nil {
		os.Exit(1)
	}

	e := ad.Connect()
	if e != nil {
		os.Exit(1)
	}
	defer ad.Close()

	success, err := ad.Auth("zhangsan", "123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)
}
