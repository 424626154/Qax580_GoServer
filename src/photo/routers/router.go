package routers

import (
	"github.com/astaxie/beego"
	"photo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.PhotoController{}) //冲洗系统
}
