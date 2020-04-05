package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func ListVirtualSwitches() map[string]string {
	req, err := http.NewRequest("GET", setting.BaseUrl+"/api/cpcs/"+setting.CpcId+"/virtual-switches", nil)
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	virtualSwitchesMap := make(map[string]string)
	for _, virtualSwitch := range resp["virtual-switches"].([]interface{}) {
		virtualSwitchesMap["0"+virtualSwitch.(map[string]interface{})["name"].(string)[:6]] = virtualSwitch.(map[string]interface{})["object-uri"].(string)
	}
	return virtualSwitchesMap
}
