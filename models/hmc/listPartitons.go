package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func ListPartitions() []interface{} {
	req, err := http.NewRequest("GET", setting.BaseUrl+"/api/cpcs/"+setting.CpcId+"/partitions", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	return resp["partitions"].([]interface{})
}
