package utils

import (
	"encoding/csv"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/pkg/setting"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func HandleLparInfo(fileName string) []map[string]interface{} {
	var lparInfoList []map[string]interface{}
	// 获取数据，按照文件
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	allLparInfo, _ := r2.ReadAll()
	// 循环取数据
	for i := 0; i < len(allLparInfo); i++ {
		if allLparInfo[i][0] == "crt_lpar" {
			initMemory, _ := strconv.Atoi(allLparInfo[i][6])
			maxMemory, _ := strconv.Atoi(allLparInfo[i][7])
			ifl, _ := strconv.Atoi(allLparInfo[i][5])
			lparInfoList = append(lparInfoList, map[string]interface{}{
				"name":           allLparInfo[i][2],
				"short-name":     allLparInfo[i][3],
				"description":    allLparInfo[i][4],
				"processor-mode": "shared",
				"ifl-processors": ifl,
				"initial-memory": initMemory * 1024,
				"maximum-memory": maxMemory * 1024,
			})
			db.InsertLparInfo(allLparInfo[i][2], allLparInfo[i][13], allLparInfo[i][15], "24", allLparInfo[i][14], allLparInfo[i][12], setting.OSUser, setting.OSPassword)
		}
	}
	return lparInfoList
}

func HandleNicInfo(fileName string) []map[string]interface{} {
	var nicInfoList []map[string]interface{}
	// 获取数据，按照文件
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	allNicInfo, _ := r2.ReadAll()
	var lparName string
	for i := 1; i < len(allNicInfo); i++ {
		if allNicInfo[i][0] == "crt_lpar" {
			lparName = allNicInfo[i][2]
		}
		info := map[string]interface{}{
			"lparName":     lparName,
			"nicName":      allNicInfo[i][8],
			"deviceNumber": allNicInfo[i][9],
			"adapterName":  allNicInfo[i][10],
			"adapterPort":  allNicInfo[i][11],
		}
		nicInfoList = append(nicInfoList, info)
	}
	return nicInfoList
}
func WriteCSV(input map[string][]string) {
	f, err := os.OpenFile("wwpn.csv", os.O_CREATE|os.O_TRUNC, 0777) //创建文件
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	defer f.Close()
	_, _ = f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)                //创建一个新的写入文件流
	for key, value := range input {
		data := append([]string{key}, value...)
		_ = w.Write(data)
		w.Flush()
	}
}
