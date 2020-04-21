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
func BaseConfig(lparName, ip string) (stdout bytes.Buffer, err error) {
	if stdout, err = utils.LparExecShell(ip, "echo "+lparName+">/etc/hostname"); err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		return
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

func InstallSoftware(ip string) (stdout bytes.Buffer, err error) {
	if err = utils.UploadSoftware(ip); err != nil {
		return
	}
	if stdout, err = utils.LparExecShell(ip, `tar zxvf /root/multipath_s390.tar.gz && cd /root/multipath_s390 && rpm -Uvh kpartx-0.7.3+150+suse.caa50c4-2.11.1.s390x.rpm multipath-tools-0.7.3+150+suse.caa50c4-2.11.1.s390x.rpm`); err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		return
	}
	if stdout, err = utils.LparExecShell(ip, `
cat >> /etc/sysctl.conf <<EOF

kernel.pid_max = 10000
kernel.shmmni = 4096
kernel.shmall = 1073741824
kernel.shmmax = 4398046511104
fs.aio-max-nr = 3145728
kernel.pid_max = 122880
kernel.threads-max = 122880
kernel.sem=250 32000 100 142
EOF`); err != nil {
		return
	}
	return
}
