package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func BindSGToLpar(lparName, lparObjUri, sgName, sgObjUri string) (isSuccess bool) {
	if lparName == sgName {
		data, _ := json.Marshal(map[string]string{
			"storage-group-uri": sgObjUri,
		})
		body := strings.NewReader(string(data))
		req, err := http.NewRequest("POST", setting.BaseUrl+lparObjUri+"/operations/attach-storage-group", body)
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
		utils.NetRequest(req)
		isSuccess = true
	}
	return
}
