package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
	"strings"
)

func Logon(userid string, password string) {
	logonForm := map[string]string{
		"userid":   userid,
		"password": password,
	}
	data, err := json.Marshal(logonForm)
	if err != nil {
		log.Fatal(err)
	}
	body := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", setting.BaseUrl+"/api/sessions", body)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	setting.ApiSession = resp["api-session"].(string)

}
