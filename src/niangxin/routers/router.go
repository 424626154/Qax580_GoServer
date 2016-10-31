package routers

import (
	"github.com/astaxie/beego"
	"niangxin/controllers"
)

func init() {
	beego.AutoRouter(&controllers.NiangXinController{}) //web服务器
	beego.AutoRouter(&controllers.NNAdminController{})  //amdin后台
}
