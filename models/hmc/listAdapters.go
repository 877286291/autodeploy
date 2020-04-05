package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func ListAdapters() []interface{} {
	req, err := http.NewRequest("GET", setting.BaseUrl+"/api/cpcs/"+setting.CpcId+"/adapters", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	return resp["adapters"].([]interface{})
}
