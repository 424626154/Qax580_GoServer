package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"qax580go/models"
	_ "qax580go/routers"
	"time"
)

func timeFormat(in int64) (out string) {
	minute := 60
	hour := minute * 60
	day := hour * 24
	month := day * 30
	year := month * 12
	now := time.Now().Unix()
	diffValue := now - in
	if diffValue < 0 {
		//若日期不符则弹出窗口告之
	}
	yearC := diffValue / int64(year)
	monthC := diffValue / int64(month)
	weekC := diffValue / int64((7 * day))
	dayC := diffValue / int64(day)
	hourC := diffValue / int64(hour)
	minC := diffValue / int64(minute)
	result := ""

	if yearC >= 1 {
		result = time.Unix(in, 0).Format("2006-01-02 15:04:05")
	} else if monthC >= 1 {
		result = fmt.Sprintf("发表于%d个月前", monthC)
	} else if weekC >= 1 {
		result = fmt.Sprintf("发表于%d周前", weekC)
	} else if dayC >= 1 {
		result = fmt.Sprintf("发表于%d天前", dayC)
	} else if hourC >= 1 {
		result = fmt.Sprintf("发表于%d个小时前", hourC)
	} else if minC >= 1 {
		result = fmt.Sprintf("发表于%d分钟前", minC)
	} else {
		result = "刚刚发表"
	}
	return result
}

func isImgPath(in string) (out string) {
	url := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		if isdebug == "true" {
			url = iniconf.String("qax580::imgurltest")
		} else {
			url = iniconf.String("qax580::imgurl")
		}

	}
	return fmt.Sprintf("%s%s", url, in)
}
func versionInfo() (out string) {
	version := "1.0.0_beta"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		version = iniconf.String("qax580::versioninfo")
	}
	return version
}

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.SetStaticPath("/game", "game")

	beego.AddFuncMap("timeformat", timeFormat)
	beego.AddFuncMap("isImgPath", isImgPath)
	beego.AddFuncMap("versionInfo", versionInfo)
	beego.SetStaticPath("/web", "web")
	beego.Run()
}
