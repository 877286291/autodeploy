package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"net/http"
)

func StartLpars(lparUri string) (err error) {
	var req *http.Request
	req, err = http.NewRequest("POST", setting.BaseUrl+lparUri+"/operations/start", nil)
	utils.NetRequest(req)
	return
}
