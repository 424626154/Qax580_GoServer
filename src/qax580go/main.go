package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"qax580go/models"
	_ "qax580go/routers"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
