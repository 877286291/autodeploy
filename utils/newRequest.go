package utils

import (
	"crypto/tls"
	"encoding/json"
	"github.com/aurora/autodeploy/pkg/setting"
	"io/ioutil"
	"log"
	"net/http"
)

func NetRequest(req *http.Request) (responseBody map[string]interface{}) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("X-API-Session", setting.ApiSession)
	resp, err := client.Do(req)
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(content, &responseBody)
	//ParseResponse(responseBody, resp.StatusCode)
	//fmt.Println(string(content),resp.StatusCode)
	return
}
