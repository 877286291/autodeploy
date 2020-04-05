package hmc

import (
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"log"
	"net/http"
)

func GetVolumeUri(sgUri string) (volumeUri string) {
	req, err := http.NewRequest("GET", setting.BaseUrl+sgUri, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp := utils.NetRequest(req)
	volumeUriList, ok := resp["storage-volume-uris"]
	if ok {
		volumeUri = volumeUriList.([]interface{})[0].(string)
	}
	return
}
