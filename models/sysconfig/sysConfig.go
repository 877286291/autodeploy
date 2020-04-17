package sysconfig

import (
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
	"log"
	"os/exec"
)

var ipList []map[string]string

func init() {
	ipList = db.GetIp()
}

// 更改hostname
func BaseConfig() (err error) {
	for _, obj := range ipList {
		for lparName, ip := range obj {
			if err = utils.LparExecShell(ip, "echo "+lparName+">/etc/hostname"); err != nil {
				_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
				log.Fatal(err)
			}
		}
	}
	return
}
func CheckConn() {
	for _, obj := range ipList {
		for _, ip := range obj {
			command := exec.Command("cmd", "ping", ip)
			command.Run()
		}
	}
}
func ConnHost() {
	for _, obj := range ipList {
		for _, ip := range obj {
			if err := utils.LparExecShell(ip, "pwd"); err != nil {
				_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
				log.Fatal(err)
			}
		}
	}
}
