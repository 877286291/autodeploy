package autoyast

import (
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/pkg/setting"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	_ = os.MkdirAll(setting.AutoYastDir+"iso/", 0777)
}

func GenerateAutoYast(lparObjList []map[string]string) (err error) {
	for _, lparObj := range lparObjList {
		for lparName, _ := range lparObj {
			// 获取xml模板
			content := GetTemplate()
			lparInfo := db.GetLparInfoByName(lparName)
			// 执行替换
			tmpContent := strings.Replace(content, "example_ip", lparInfo["ip"], -1)
			tmpContent = strings.Replace(tmpContent, "example_netmask", lparInfo["netmask"], -1)
			tmpContent = strings.Replace(tmpContent, "example_prefixlen", lparInfo["prefixlen"], -1)
			tmpContent = strings.Replace(tmpContent, "example_gateway", lparInfo["gateway"], -1)
			tmpContent = strings.Replace(tmpContent, "example_lunid", lparInfo["lunid"], -1)
			err = os.MkdirAll(setting.AutoYastDir+"xml", 0777)
			err = ioutil.WriteFile(setting.AutoYastDir+"xml/"+lparName+".xml", []byte(tmpContent), 0777)
			pmfTemplate := `TERM=dumb manual=0 
ramdisk_size=524288 root=/dev/ram1 ro init=/linuxrc
instnetdev=osa osainterface=qdio layer2=1 portno=0 osahwaddr=
ReadChannel=0.0.3000 WriteChannel=0.0.3001 DataChannel=0.0.3002
hostip=` + lparInfo["ip"] + ` netmask=` + lparInfo["netmask"] + ` gateway=` + lparInfo["gateway"] + `
install=ftp://` + setting.FtpUserName + `:` + setting.FtpPassWord + `@` + setting.FtpHost + `/` + setting.AutoYastDir + `iso
autoyast=ftp://` + setting.FtpUserName + `:` + setting.FtpPassWord + `@` + setting.FtpHost + `/` + setting.AutoYastDir + `xml/` + lparName + `.xml
linuxrclog=/dev/console vnc=1
VNCPassword=` + setting.VNCPassword + "\n"
			insTemplate := `* SUSE Linux for IBM z Systems Installation/Rescue System
../iso/boot/s390x/linux 0x00000000
../iso/boot/s390x/initrd.off 0x0001040c
../iso/boot/s390x/initrd.siz 0x00010414
../iso/boot/s390x/initrd 0x01000000
` + lparName + "_pmf" + " 0x00010480\n"
			err = os.MkdirAll(setting.AutoYastDir+"ins", 0777)
			err = ioutil.WriteFile(setting.AutoYastDir+"ins/"+lparName+".ins", []byte(insTemplate), 0777)
			err = ioutil.WriteFile(setting.AutoYastDir+"ins/"+lparName+"_pmf", []byte(pmfTemplate), 0777)
		}
	}
	return err
}
