package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func DeleteLpar(lparUri string) (isDelete bool) {
	req, err := http.NewRequest("DELETE", setting.BaseUrl+lparUri, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	if len(resp) == 0 {
		isDelete = true
	} else {
		isDelete = false
	}
	return isDelete
}
