package sysconfig

import (
	"bytes"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
)

var IpList []map[string]string

func init() {
	IpList = db.GetIp()
}

// 更改hostname
func BaseConfig() (stdout bytes.Buffer, err error) {
	//var stdout bytes.Buffer
	for _, obj := range IpList {
		for lparName, ip := range obj {
			if stdout, err = utils.LparExecShell(ip, "echo "+lparName+">/etc/hostname"); err != nil {
				_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
				return
			}
		}
	}
	return
}
func CheckConn(ip string) (stdout bytes.Buffer, err error) {
	if stdout, err = utils.LparExecShell(ip, "ping -c 1  -i  2 -s 2048 -W 1 "+ip); err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		return
	}
	return
}
