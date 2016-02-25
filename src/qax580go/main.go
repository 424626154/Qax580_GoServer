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
func timeFormat1(in int64) (out string) {
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
		result = fmt.Sprintf("%d个月前", monthC)
	} else if weekC >= 1 {
		result = fmt.Sprintf("%d周前", weekC)
	} else if dayC >= 1 {
		result = fmt.Sprintf("%d天前", dayC)
	} else if hourC >= 1 {
		result = fmt.Sprintf("%d小时前", hourC)
	} else if minC >= 1 {
		result = fmt.Sprintf("%d分钟前", minC)
	} else {
		result = "刚刚发表"
	}
	return result
}
func moneyRecord(in int64) (out string) {
	beego.Debug("moneyRecord in", in)
	record := "未知的获得途径"
	if in == 1 {
		record = "关注公众号"
	} else if in == 2 {
		record = "发布信息已审核通过"
	} else if in == 3 {
		record = "发布信息被别人认可"
	} else if in == 4 {
		record = "商城兑换"
	}
	return fmt.Sprintf("%s", record)
}
func moneyRecordInfo(in int64, in1 int64) (out string) {
	beego.Debug("moneyRecord in", in)
	record := "未知"
	if in == 1 {
		record = "获得帮帮币"
	} else if in == 2 {
		record = "获得帮帮币"
	} else if in == 3 {
		record = "获得帮帮币"
	} else if in == 4 {
		record = "消耗帮帮币"
	}
	return fmt.Sprintf("%s%d", record, in1)
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
	beego.AddFuncMap("timeformat1", timeFormat1)
	beego.AddFuncMap("isImgPath", isImgPath)
	beego.AddFuncMap("versionInfo", versionInfo)
	beego.AddFuncMap("moneyrecord", moneyRecord)
	beego.AddFuncMap("moneyrecordinfo", moneyRecordInfo)
	beego.SetStaticPath("/web", "web")
	beego.Run()
}
