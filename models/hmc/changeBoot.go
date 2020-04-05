package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FtpBoot(lparName string, lparObj interface{}) (isSuccess bool, err error) {
	if lparName == lparObj.(map[string]interface{})["name"] {
		lparObjUri := lparObj.(map[string]interface{})["object-uri"].(string)
		form := map[string]string{
			"boot-device":       "ftp",
			"boot-ftp-host":     setting.FtpHost,
			"boot-ftp-username": setting.FtpUserName,
			"boot-ftp-password": setting.FtpPassWord,
			"boot-ftp-insfile":  setting.AutoYastDir + "ins/" + lparName + ".ins",
		}
		var data []byte
		var req *http.Request
		data, err = json.Marshal(form)
		body := strings.NewReader(string(data))
		req, err = http.NewRequest("POST", setting.BaseUrl+lparObjUri, body)
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
		utils.NetRequest(req)
		isSuccess = true
	}
	return
}

func SanBoot(lparName string, lparObj interface{}) (isSuccess bool, err error) {
	if lparName == lparObj.(map[string]interface{})["name"] {
		var sgUri, volumeUri string
		sgUri, err = db.GetSGUri(lparName)
		volumeUri = GetVolumeUri(sgUri)
		lparObjUri := lparObj.(map[string]interface{})["object-uri"].(string)
		form := map[string]string{
			"boot-device":         "storage-volume",
			"boot-storage-volume": volumeUri,
		}
		var data []byte
		var req *http.Request
		data, err = json.Marshal(form)
		body := strings.NewReader(string(data))
		req, err = http.NewRequest("POST", setting.BaseUrl+lparObjUri, body)
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
		utils.NetRequest(req)
		isSuccess = true
	}
	return
}
