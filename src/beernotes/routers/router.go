package routers

import (
	"beernotes/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.WebController{})   //web服务器
	beego.AutoRouter(&controllers.AdminController{}) //amdin后台
	beego.AutoRouter(&controllers.AppController{})   //amdin后台

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})
}
