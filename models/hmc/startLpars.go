package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"net/http"
)

func StartLpars(lparName, lparUri string) (err error) {
	var req *http.Request
	// 已经获取到lunid的才能启动分区
	req, err = http.NewRequest("POST", setting.BaseUrl+lparUri+"/operations/start", nil)
	utils.NetRequest(req)
	return
}
