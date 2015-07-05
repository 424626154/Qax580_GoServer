package routers

import (
	"github.com/astaxie/beego"
	"qax580admin/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/uplode", &controllers.UplodeController{})
	beego.Router("/content", &controllers.ContentController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/wxlist", &controllers.WcListController{})
	beego.Router("/feedbacklist", &controllers.FeedbackListController{})
	beego.Router("/modify", &controllers.ModifyController{})
}
