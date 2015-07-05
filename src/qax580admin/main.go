package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"qax580admin/models"
	_ "qax580admin/routers"
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
