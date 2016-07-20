package routers

import (
	"github.com/astaxie/beego"
	"niangxin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.NiangXinController{}) //冲洗系统
}
