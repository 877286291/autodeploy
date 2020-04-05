package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CreateLpar(lparInfo map[string]interface{}) (lparObjUri string) {
	data, err := json.Marshal(lparInfo)
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	body := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", setting.BaseUrl+"/api/cpcs/"+setting.CpcId+"/partitions", body)
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	var ok bool
	lparObjUri, ok = resp["object-uri"].(string)
	if ok {
		// 第一次获取，更新数据库中的表
		db.UpdateLparUri(lparObjUri, lparInfo["name"].(string))
	} else {
		// 如果没有获取到lapr对象的uri，说明已经存在该分区，接下来获取这个分区的uri
		lparObjUri = db.GetLparUri(lparInfo["name"].(string))
		if err != nil {
			// 如果lparuri在数据库中找不到，则调用HMC API进行查找并更新到数据库中
			if err.Error() == "sql: Scan error on column index 0, name \"lparuri\": converting NULL to string is unsupported" {
				lparObjUri = findLparByName(lparInfo["name"].(string))
				// 找到lparuri，更新数据库
				db.UpdateLparUri(lparObjUri, lparInfo["name"].(string))
			}
		}
	}
	return lparObjUri
}

func findLparByName(lparName string) (lparObjUri string) {
	allPartitions := ListPartitions()
	if len(allPartitions) > 0 {
	loop:
		for _, lparInfo := range allPartitions {
			if lparInfo.(map[string]interface{})["name"] == lparName {
				lparObjUri = lparInfo.(map[string]interface{})["object-uri"].(string)
				break loop
			}
		}
	} else {
		log.Fatal("没有找到分区...")
	}
	return
}
