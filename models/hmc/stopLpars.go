package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"net/http"
)

func StopLpars(lparName string, lparUri string) (err error) {
	var req *http.Request
	req, err = http.NewRequest("POST", setting.BaseUrl+lparUri+"/operations/stop", nil)
	utils.NetRequest(req)
	return
}
