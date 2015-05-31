package routers

import (
	"github.com/astaxie/beego"
	"qax580go/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/uplode", &controllers.UplodeController{})
	beego.Router("/content", &controllers.ContentController{})
}
