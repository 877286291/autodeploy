package db

import (
	"database/sql"
	"github.com/aurora/autodeploy/pkg/setting"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var database *sql.DB

func init() {
	if setting.Truncate {
		err := os.Remove("./autodeploy.db")
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
	}
	database, _ = sql.Open("sqlite3", "./autodeploy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS lpar_info (id INTEGER PRIMARY KEY, lparname TEXT unique , ip TEXT,netmask TEXT,prefixlen TEXT,gateway TEXT,hostname TEXT,lparuri TEXT,sguri,lunid TEXT,user TEXT,password TEXT)")
	_, err := statement.Exec()
	if err != nil {
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
}

func InsertLparInfo(lparname, ip, netmask, perfixlen, gateway, hostname, user, password string) {
	statement, _ := database.Prepare("INSERT INTO lpar_info(lparname,ip,netmask,prefixlen,gateway,hostname,user,password) VALUES (?,?,?,?,?,?,?,?)")
	_, err := statement.Exec(lparname, ip, netmask, perfixlen, gateway, hostname, user, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: lpar_info.lparname" {
			//log.Printf("分区[%s]在数据库中已存在", lparname)
		} else {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
	}

}

func UpdateLparUri(lparUri, lparName string) {
	statement, _ := database.Prepare("UPDATE lpar_info SET lparuri=? WHERE lparname=?")
	_, err := statement.Exec(lparUri, lparName)
	if err != nil {
		//log.Println(err)
	}
}

func GetLparUri(lparName string) (lparUri string, err error) {
	rows, _ := database.Query("SELECT lparuri from lpar_info WHERE lparname='" + lparName + "'")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&lparUri)
	}
	return
}
func GetLparUriInDb() (lparUriList []string) {
	rows, _ := database.Query("SELECT lparuri from lpar_info")
	var lparuri string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&lparuri)
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Fatal(err)
		}
		lparUriList = append(lparUriList, lparuri)
	}
	return
}

func UpdateSGUri(SGUri, lparName string) {
	statement, _ := database.Prepare("UPDATE lpar_info SET sguri=? WHERE lparname=?")
	_, err := statement.Exec(SGUri, lparName)
	if err != nil {
		//log.Println(err)
	}
}
func GetSGUri(lparName string) (SGUri string, err error) {
	rows, _ := database.Query("SELECT sguri from lpar_info WHERE lparname='" + lparName + "'")
	// 提前注册关闭数据库，默认情况rows.Next在结束给时候会调用close，但是在发生错误时直接return不会close会发生数据库死锁，所以在此提前注册
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&SGUri)
		if err != nil {
			return
		}
	}
	defer rows.Close()
	return
}

func GetLunid(lparName string) (lunId string) {
	rows, _ := database.Query("SELECT lunid from lpar_info WHERE lparname='" + lparName + "'")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&lunId)
		if err != nil {
			if err.Error() == "sql: Scan error on column index 0, name \"lunid\": converting NULL to string is unsupported" {
				//log.Printf("[%s]分区未获取到lunid，请检查！", lparName)
			} else {
				_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
				log.Fatal(err)
			}
		}
	}
	return
}
func UpdateLunid(lunId, sgUri string) {
	statement, _ := database.Prepare("UPDATE lpar_info SET lunid=? WHERE sguri=?")
	_, err := statement.Exec(lunId, sgUri)
	if err != nil {
		//log.Println(err)
	}
}

func GetLparInfoByName(lparName string) map[string]string {
	rows, _ := database.Query("SELECT ip,netmask,prefixlen,gateway,lunid from lpar_info WHERE lparname='" + lparName + "'")
	result := make(map[string]string)
	var ip, netmask, prefixlen, gateway, lunid string
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&ip, &netmask, &prefixlen, &gateway, &lunid)
		result["ip"] = ip
		result["netmask"] = netmask
		result["prefixlen"] = prefixlen
		result["gateway"] = gateway
		result["lunid"] = strings.ToLower(lunid)
	}
	return result
}
func GetLparAndSGInfoFromDb() (lparObjUriList, sgObjUriList []map[string]string) {
	rows, _ := database.Query("SELECT lparname,lparuri,sguri from lpar_info")
	var lparname, lparuri, sguri string
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&lparname, &lparuri, &sguri)
		lparResult := map[string]string{lparname: lparuri}
		sgResult := map[string]string{lparname: sguri}
		lparObjUriList = append(lparObjUriList, lparResult)
		sgObjUriList = append(sgObjUriList, sgResult)
	}
	return
}
