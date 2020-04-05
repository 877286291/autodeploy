package hmc

import (
	"encoding/json"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
	"strings"
)

func CreateSG(lparInfo map[string]interface{}) (sgObjUri string) {
	sgName := lparInfo["name"].(string)
	volume := map[string]interface{}{
		"operation":   "create",
		"size":        200,
		"usage":       "boot",
		"description": "san_root",
	}
	var volumes []map[string]interface{}
	volumes = append(volumes, volume)
	form := map[string]interface{}{
		"cpc-uri":         "/api/cpcs/" + setting.CpcId,
		"type":            "fcp",
		"shared":          false,
		"max-partitions":  1,
		"connectivity":    4,
		"storage-volumes": volumes,
		"name":            sgName,
		"description":     "sg_" + sgName,
	}
	data, err := json.Marshal(form)
	if err != nil {
		log.Fatal(err)
	}
	body := strings.NewReader(string(data))
	req, err2 := http.NewRequest("POST", setting.BaseUrl+"/api/storage-groups", body)
	if err2 != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	var ok bool
	sgObjUri, ok = resp["object-uri"].(string)
	if ok {
		db.UpdateSGUri(sgObjUri, lparInfo["name"].(string))
	} else {
		sgObjUri, err = db.GetSGUri(lparInfo["name"].(string))
		if err != nil {
			// 如果sguri在数据库中找不到，则调用HMC API进行查找并更新到数据库中
			if err.Error() == "sql: Scan error on column index 0, name \"sguri\": converting NULL to string is unsupported" {
				sgObjUri = findSGByName(sgName)
				// 找到sgUri，更新数据库
				db.UpdateSGUri(sgObjUri, lparInfo["name"].(string))
			}
		}
	}
	return
}

func findSGByName(SGName string) (SGObjUri string) {
	req, err := http.NewRequest("GET", setting.BaseUrl+"/api/storage-groups", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
loop:
	for _, sgInfo := range resp["storage-groups"].([]interface{}) {
		if sgInfo.(map[string]interface{})["name"] == SGName {
			SGObjUri = sgInfo.(map[string]interface{})["object-uri"].(string)
			break loop
		}
	}
	return
}

// HMC扫盘
func ConnDiscovery(sgUri string) (err error) {
	//  /api/storage-groups/{storage-group-id}/operations/start-fcp-storage-discovery
	requireForm := map[string]bool{
		"force-restart": false,
	}
	var data []byte
	var req *http.Request
	data, err = json.Marshal(requireForm)
	body := strings.NewReader(string(data))
	req, err = http.NewRequest("POST", setting.BaseUrl+sgUri+"/operations/start-fcp-storage-discovery", body)
	utils.NetRequest(req)
	return
}

// 获取lunid
func GetLunId(sgUri string) (lunId string) {
	req, err := http.NewRequest("GET", setting.BaseUrl+sgUri+"/storage-volumes", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	storageVolumes := resp["storage-volumes"].([]interface{})
	for _, volumeUri := range storageVolumes {
		req, err := http.NewRequest("GET", setting.BaseUrl+volumeUri.(map[string]interface{})["element-uri"].(string), nil)
		if err != nil {
			log.Fatal(err)
		}
		resp := utils.NetRequest(req)
		uuid, ok := resp["uuid"]
		if ok && uuid != nil {
			lunId = "3" + strings.ToLower(uuid.(string))
			// 更新到数据库
			db.UpdateLunid(lunId, sgUri)
		}
	}
	return
}
