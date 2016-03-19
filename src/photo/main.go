package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "photo/routers"
	"qax580go/models"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
