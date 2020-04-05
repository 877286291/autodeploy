package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"net/http"
	"strings"
)

func BindNicToLpar(nicInfo map[string]interface{}) (err error) {
	form := make(map[string]string)
	if strings.HasPrefix(nicInfo["adapterName"].(string), "RoCE") {
		form = map[string]string{
			"name":                     nicInfo["nicName"].(string),
			"network-adapter-port-uri": nicInfo["adapterUri"].(string),
			"device-number":            nicInfo["deviceNumber"].(string),
		}
	} else {
		form = map[string]string{
			"name":               nicInfo["nicName"].(string),
			"virtual-switch-uri": nicInfo["adapterUri"].(string),
			"device-number":      nicInfo["deviceNumber"].(string),
		}
	}
	var data []byte
	var req *http.Request
	data, err = json.Marshal(form)
	body := strings.NewReader(string(data))
	req, err = http.NewRequest("POST", setting.BaseUrl+nicInfo["lparObjUri"].(string)+"/nics", body)
	utils.NetRequest(req)

	return
}
