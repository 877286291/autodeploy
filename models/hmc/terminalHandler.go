package hmc

import (
	"fmt"
	"github.com/aurora/autodeploy/db"
	"github.com/aurora/autodeploy/models/autoyast"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/aurora/autodeploy/utils"
	"github.com/rivo/tview"
	"io/ioutil"
	"strings"
	"time"
)

var (
	lparInfoList   []map[string]interface{}
	nicInfoList    []map[string]interface{}
	lparObjUriList []map[string]string
	sgObjUriList   []map[string]string
	App            *tview.Application
)
var logText = tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWordWrap(true).
	SetChangedFunc(func() {
		App.Draw()
	})

func init() {
	// 从数据库读取lparObjUriList和sgObjUriList
	lparObjUriList, sgObjUriList = db.GetLparAndSGInfoFromDb()
	// 读取lpar配置信息
	lparInfoList = utils.HandleLparInfo("lpar_sg_info.csv")
	// 读取nic卡配置信息
	nicInfoList = utils.HandleNicInfo("lpar_sg_info.csv")
}
func MenuList(height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(p, height, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 1, true).
		AddItem(logText, 0, 1, false)
}

func Partition() (title string, content tview.Primitive) {
	listView := tview.NewList().
		AddItem("创建分区", "", '1', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行创建分区\n", time.Now().Format("2006-01-02 15:04:05"))
				// 创建分区
				for _, lparInfo := range lparInfoList {
					CreateLpar(lparInfo)
					_, _ = fmt.Fprintf(logText, "[#008000]分区[%s]创建完成\n", lparInfo["name"].(string))
				}
			}()
		}).
		AddItem("删除分区", "", '2', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行删除分区\n", time.Now().Format("2006-01-02 15:04:05"))
				for _, lparUri := range db.GetLparUriInDb() {
					if DeleteLpar(lparUri) {
						_, _ = fmt.Fprintf(logText, "[#008000][%s]删除成功\n", lparUri)
					} else {
						_, _ = fmt.Fprintf(logText, "[#ff0000][%s]删除失败\n", lparUri)
					}
				}
			}()
		}).
		AddItem("启动分区", "", '3', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行启动分区\n", time.Now().Format("2006-01-02 15:04:05"))
				lparObjUriList, sgObjUriList = db.GetLparAndSGInfoFromDb()
				for _, lparObj := range lparObjUriList {
					for lparName, lparUri := range lparObj {
						if db.GetLunid(lparName) != "" {
							err := StartLpars(lparUri)
							if err != nil {
								_, _ = fmt.Fprintf(logText, "[#ff0000]分区[%s]发送启动分区指令失败\n", lparName)
							} else {
								_, _ = fmt.Fprintf(logText, "[#008000]分区[%s]发送启动分区指令成功\n", lparName)
							}
						} else {
							_, _ = fmt.Fprintf(logText, "[#ff0000]分区[%s]请先获取该分区的lunid再进行启动！\n", lparName)
						}
					}
				}
			}()
		}).
		AddItem("停止分区", "", '4', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行停止分区\n", time.Now().Format("2006-01-02 15:04:05"))
				lparObjUriList, sgObjUriList = db.GetLparAndSGInfoFromDb()
				for _, lparObj := range lparObjUriList {
					for lparName, lparUri := range lparObj {
						err := StopLpars(lparUri)
						if err != nil {
							_, _ = fmt.Fprintf(logText, "[#ff0000]分区[%s]发送停止分区指令失败\n", lparName)
						} else {
							_, _ = fmt.Fprintf(logText, "[#008000]分区[%s]发送停止分区指令成功\n", lparName)
						}
					}
				}
			}()
		}).
		AddItem("绑定网卡", "", '5', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行绑定网卡到分区\n", time.Now().Format("2006-01-02 15:04:05"))
				// 获取所有Lpar
				allPartitions := ListPartitions()
				// 获取所有Adapters
				allAdapters := ListAdapters()
				// 获取所有virtual-switches
				virtualSwitchesMap := ListVirtualSwitches()
				// 绑定NIC卡到lpar
				var nicList []map[string]interface{}
				for _, nicInfo := range nicInfoList {
					nicName := nicInfo["nicName"]
					deviceNumber := nicInfo["deviceNumber"]
					adapterPort := nicInfo["adapterPort"].(string)
					lparName := nicInfo["lparName"]
					adapterName := nicInfo["adapterName"].(string)
					var lparObjUri string
				Loop:
					for _, lparInfo := range allPartitions {
						if lparName == lparInfo.(map[string]interface{})["name"] {
							lparObjUri = lparInfo.(map[string]interface{})["object-uri"].(string)
							break Loop
						}
					}
					nicData := make(map[string]interface{})
					if strings.HasPrefix(adapterName, "OS") {
						adapterId := strings.Split(adapterName, " ")[1]
						adapterPort = "P" + adapterPort
						virtualSwitchUri := virtualSwitchesMap[adapterId+"."+adapterPort]
						nicData = map[string]interface{}{
							"nicName":      nicName,
							"deviceNumber": deviceNumber,
							"adapterUri":   virtualSwitchUri,
							"lparObjUri":   lparObjUri,
							"adapterName":  adapterName,
							"adapterPort":  adapterPort,
							"lparName":     lparName,
						}
					} else {
						var adapterUri string
						for _, adapterInfo := range allAdapters {
							if adapterInfo.(map[string]interface{})["name"].(string) == adapterName {
								adapterUri = adapterInfo.(map[string]interface{})["object-uri"].(string) + "/network-ports/" + adapterPort
							}
						}
						nicData = map[string]interface{}{
							"nicName":      nicName,
							"deviceNumber": deviceNumber,
							"adapterUri":   adapterUri,
							"lparObjUri":   lparObjUri,
							"adapterName":  adapterName,
							"adapterPort":  adapterPort,
							"lparName":     lparName,
						}
					}
					nicList = append(nicList, nicData)
				}
				for _, nicInfo := range nicList {
					err := BindNicToLpar(nicInfo)
					if err != nil {
						_, _ = fmt.Fprintf(logText, "[#ff0000]网卡[%s]绑定到分区[%s]失败\n", nicInfo["nicName"].(string), nicInfo["lparName"].(string))
					} else {
						_, _ = fmt.Fprintf(logText, "[#008000]网卡[%s]绑定到分区[%s]成功\n", nicInfo["nicName"].(string), nicInfo["lparName"].(string))
					}
				}
			}()
		}).
		AddItem("更改启动项为FTP", "", '6', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行更改分区从FTP启动\n", time.Now().Format("2006-01-02 15:04:05"))
				allPartitions := ListPartitions()
				for _, lparInfo := range lparInfoList {
					lparName := lparInfo["name"]
					for _, lparObj := range allPartitions {
						isSuccess, err := FtpBoot(lparName.(string), lparObj)
						if err == nil && isSuccess {
							_, _ = fmt.Fprintf(logText, "[#008000]更改分区[%s]从FTP启动成功\n", lparName.(string))
						} else if err != nil {
							_, _ = fmt.Fprintf(logText, "[#ff0000]更改分区[%s]从FTP启动失败\n", lparName.(string))
						}
					}
				}
			}()
		}).
		AddItem("更改启动项为SAN", "", '7', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行更改分区从SAN启动\n", time.Now().Format("2006-01-02 15:04:05"))
				allPartitions := ListPartitions()
				for _, lparInfo := range lparInfoList {
					lparName := lparInfo["name"]
					for _, lparObj := range allPartitions {
						isSuccess, err := SanBoot(lparName.(string), lparObj)
						if err == nil && isSuccess {
							_, _ = fmt.Fprintf(logText, "[#008000]更改分区[%s]从SAN启动成功\n", lparName.(string))
						} else if err != nil {
							_, _ = fmt.Fprintf(logText, "[#ff0000]更改分区[%s]从SAN启动失败\n", lparName.(string))
						}
					}
				}
			}()
		})
	return "分区管理", MenuList(20, listView)
}

func SG() (title string, content tview.Primitive) {
	listView := tview.NewList().
		AddItem("创建SG", "", '1', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行创建SG\n", time.Now().Format("2006-01-02 15:04:05"))
				// 读取lpar配置信息
				lparInfoList = utils.HandleLparInfo("lpar_sg_info.csv")
				for _, lparInfo := range lparInfoList {
					CreateSG(lparInfo)
					_, _ = fmt.Fprintf(logText, "[#008000]SG[%s]创建完成\n", lparInfo["name"].(string))
				}
			}()
		}).
		AddItem("绑定SG到分区", "", '2', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行绑定SG到分区\n", time.Now().Format("2006-01-02 15:04:05"))
				lparObjUriList, sgObjUriList = db.GetLparAndSGInfoFromDb()
				for _, lparObjUriMap := range lparObjUriList {
					for lparName, lparObjUri := range lparObjUriMap {
						for _, sgObjUriMap := range sgObjUriList {
							for sgName, sgObjUri := range sgObjUriMap {
								if BindSGToLpar(lparName, lparObjUri, sgName, sgObjUri) {
									_, _ = fmt.Fprintf(logText, "[#008000]分区[%s]绑定SG[%s]完成\n", lparName, sgName)
								}
							}
						}
					}
				}
			}()
		}).
		AddItem("导出WWPN", "", '3', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行导出WWPN\n", time.Now().Format("2006-01-02 15:04:05"))
				ExportWWPN(sgObjUriList)
				_, _ = fmt.Fprintln(logText, "[#008000]WWPN导出完成！文件名为wwpn.csv")
			}()
		}).
		AddItem("HMC扫盘", "", '4', func() {
			go func() {
				for _, SGObjUri := range sgObjUriList {
					for sgName, sgUri := range SGObjUri {
						err := ConnDiscovery(sgUri)
						if err != nil {
							_, _ = fmt.Fprintf(logText, "[#ff0000]SG[%s]扫盘指令发送失败！\n", sgName)
						} else {
							_, _ = fmt.Fprintf(logText, "[#008000]SG[%s]扫盘指令发送成功，请等待扫盘结束\n", sgName)
						}
					}
				}
			}()
		}).
		AddItem("获取Lunid", "", '5', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行获取lunid\n", time.Now().Format("2006-01-02 15:04:05"))
				_, sgObjUriList = db.GetLparAndSGInfoFromDb()
			Loop:
				for _, SGObjUri := range sgObjUriList {
					for sgName, sgUri := range SGObjUri {
						if sgUri == "" {
							_, _ = fmt.Fprintf(logText, "[#ff0000]数据库中[%s]sguri为空，请手动执行[创建SG]指令\n", sgName)
							continue Loop
						}
						lunId := GetLunId(sgUri)
						if lunId != "" {
							_, _ = fmt.Fprintf(logText, "[#008000]SG[%s]lunid为%s\n", sgName, lunId)
						} else {
							_, _ = fmt.Fprintf(logText, "[#ff0000]SG[%s]没有获取到lunid\n", sgName)
						}
					}
				}
			}()
		})
	return "存储管理", MenuList(12, listView)
}

// 安装操作系统
func OS() (title string, content tview.Primitive) {
	listView := tview.NewList().
		AddItem("生成配置文件(ins/pmf/xml)", "", '1', func() {
			go func() {
				// 更新lparObjUriList
				lparObjUriList, sgObjUriList = db.GetLparAndSGInfoFromDb()
				err := autoyast.GenerateAutoYast(lparObjUriList)
				if err != nil {
					_, _ = fmt.Fprintln(logText, "[#ff0000]生成配置文件发生错误！")
				} else {
					_, _ = fmt.Fprintln(logText, "[#008000]生成配置文件成功")
				}
			}()
		}).
		AddItem("上传配置文件到FTP", "", '2', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行上传配置文件到FTP\n", time.Now().Format("2006-01-02 15:04:05"))
				err := utils.Upload(setting.AutoYastDir, "/home/"+setting.FtpUserName+"/"+setting.AutoYastDir)
				if err != nil {
					if err.Error() == "CreateFile "+setting.AutoYastDir+": The system cannot find the file specified." {
						_, _ = fmt.Fprintf(logText, "[#ff0000]%s文件路径不存在\n", setting.AutoYastDir)
					} else {
						_, _ = fmt.Fprintf(logText, "[#ff0000]%s\n", setting.AutoYastDir)
					}
				} else {
					_, _ = fmt.Fprintln(logText, "[#008000]上传配置文件到FTP成功！")
				}
			}()
		}).
		AddItem("上传iso镜像到FTP", "", '3', func() {
			go func() {
				_, _ = fmt.Fprintf(logText, "%s 执行上传ISO镜像到FTP,请稍等\n", time.Now().Format("2006-01-02 15:04:05"))
				err := utils.Upload("SLE12_SP4.iso", "/home/"+setting.FtpUserName+"/"+setting.AutoYastDir)
				if err != nil {
					_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
					_, _ = fmt.Fprintln(logText, "[#ff0000]ISO镜像上传失败")
				} else {
					_, _ = fmt.Fprintln(logText, "[#008000]ISO镜像上传成功！")
					// todo 挂载镜像到autoyast/iso文件夹
					utils.ExecShell("mount " + "/home/" + setting.FtpUserName + "/" + setting.AutoYastDir + "SLE12_SP4.iso " + "/home/" + setting.FtpUserName + "/" + setting.AutoYastDir + "iso")
					_, _ = fmt.Fprintln(logText, "[#008000]挂载ISO成功！")
				}
			}()
		})
	return "安装系统", MenuList(10, listView)
}

// 系统标准化
func OsStandardization() (title string, content tview.Primitive) {
	listView := tview.NewList().
		AddItem("host节点基础配置", "", '1', func() {
			go func() {
				// todo 系统标准化

			}()
		}).AddItem("软件安装", "", '2', func() {
		go func() {
			// todo 系统标准化

		}()
	}).AddItem("网络连通性测试", "", '3', func() {
		go func() {

		}()
	}).AddItem("roce_bond配置文件创建", "", 'a', func() {
		go func() {

		}()
	}).AddItem("osa_bond配置文件创建", "", 'b', func() {
		go func() {

		}()
	}).AddItem("roce_vlan配置文件创建", "", 'c', func() {
		go func() {

		}()
	}).AddItem("osa_vlan配置文件创建", "", 'd', func() {
		go func() {

		}()
	}).AddItem("网卡配置文件发送/启用", "", 'e', func() {
		go func() {

		}()
	}).AddItem("删除网卡", "", 'f', func() {
		go func() {

		}()
	}).AddItem("内核参数修改", "", 'g', func() {
		go func() {

		}()
	}).AddItem("host节点网络重启", "", 'h', func() {
		go func() {

		}()
	})
	return "系统标准化", MenuList(22, listView)
}

//func Database() (title string, content tview.Primitive) {
//	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
//	})
//	return "查看数据库", MenuList(10, textView)
//}
//func AutoDeploy() (title string, content tview.Primitive) {
//	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
//	}).SetBorder(true)
//	return "自动部署", MenuList(10, textView)
//}
func Help() (title string, content tview.Primitive) {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			App.Draw()
		})
	go func() {
		_, _ = fmt.Fprintln(textView, "Ctrl+N 切换下一页面\n\n"+
			"Ctrl+P 切换上一页面\n\n"+
			"方向键或数字键进行选择\n\n"+
			"AutoYast模板系统SLE12_SP4密码为P1cc@xfzy\n\n其余配置请对conf/app.ini配置文件进行修改\n\n"+
			"要上传的ISO镜像位置为当前目录,文件名必须为SLE12_SP4.iso\n\n"+
			"[#ff0000]必须修改ssh配置文件PasswordAuthentication的值为yes")
	}()
	return "帮助？", MenuList(20, textView)
}
