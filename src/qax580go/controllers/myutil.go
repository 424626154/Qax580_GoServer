package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func getImageUrl() string {
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
	beego.Debug("111111", url)
	return url
}
