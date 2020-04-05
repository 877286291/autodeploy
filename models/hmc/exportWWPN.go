package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func ExportWWPN(sgObjUriList []map[string]string) {
	wwpnMap := make(map[string][]string)
	for _, sgObjUriMap := range sgObjUriList {
		for sgName, sgObjUri := range sgObjUriMap {
			req, err := http.NewRequest("GET", setting.BaseUrl+sgObjUri, nil)
			if err != nil {
				log.Fatal(err)
			}
			resp := utils.NetRequest(req)
			virtualStorgeUriList, ok := resp["virtual-storage-resource-uris"]
			if ok {
				for _, virtualStorgeUri := range virtualStorgeUriList.([]interface{}) {
					req, err := http.NewRequest("GET", setting.BaseUrl+virtualStorgeUri.(string), nil)
					if err != nil {
						log.Fatal(err)
					}
					resp := utils.NetRequest(req)
					wwpnMap[sgName] = append(wwpnMap[sgName], resp["world-wide-port-name"].(string))
				}
			}
		}
	}
	utils.WriteCSV(wwpnMap)
}
