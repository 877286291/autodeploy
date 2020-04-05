package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func GetCPCId() {
	req, err := http.NewRequest("GET", setting.BaseUrl+"/api/cpcs", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	for _, cpc := range resp["cpcs"].([]interface{}) {
		if cpc.(map[string]interface{})["name"] == setting.SN {
			// 获取配置文件中的SN相应的cpcId
			setting.CpcId = cpc.(map[string]interface{})["object-uri"].(string)[10:]
		}
	}
}
