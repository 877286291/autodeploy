package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg         *ini.File
	ApiSession  string
	BaseUrl     string
	HmcUserName string
	HmcPassWord string
	CpcId       string
	SN          string
	FtpHost     string
	FtpUserName string
	FtpPassWord string
	AutoYastDir string
	OSUser      string
	OSPassword  string
	Truncate    bool
	VNCPassword string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadHMC()
	LoadFtp()
	LoadOS()
	LoadDB()
	LoadVNC()
}

func LoadHMC() {
	sec, err := Cfg.GetSection("hmc")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	BaseUrl = sec.Key("BASE_URL").MustString("")
	HmcUserName = sec.Key("HMC_USERNAME").MustString("")
	HmcPassWord = sec.Key("HMC_PASSWORD").MustString("")
	SN = sec.Key("SN").MustString("")
}
func LoadFtp() {
	sec, err := Cfg.GetSection("ftp")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	FtpHost = sec.Key("FTP_HOST").MustString("")
	FtpUserName = sec.Key("FTP_USER").MustString("")
	FtpPassWord = sec.Key("FTP_PASSWORD").MustString("")
	AutoYastDir = sec.Key("AUTOYAST_DIR").MustString("")
	//InsDir = sec.Key("INS_DIR").MustString("")
	//XmlDir = sec.Key("XML_DIR").MustString("")

}

func LoadOS() {
	sec, err := Cfg.GetSection("os")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	OSUser = sec.Key("OS_USERNAME").MustString("")
	OSPassword = sec.Key("OS_PASSWORD").MustString("")
}
func LoadDB() {
	sec, err := Cfg.GetSection("db")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Truncate = sec.Key("TRUNCATE").MustBool(false)
}
func LoadVNC() {
	sec, err := Cfg.GetSection("vnc")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	VNCPassword = sec.Key("VNC_PASSWORD").MustString("")
}
