package routers

import (
	"github.com/astaxie/beego"
	"os"
	"qax580go/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/uplode", &controllers.UplodeController{})
	beego.Router("/content", &controllers.ContentController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/wxlist", &controllers.WcListController{})
	beego.Router("/wx", &controllers.WXController{})
	beego.Router("/feedback", &controllers.FeedbackController{})

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})

	beego.Router("/admin", &controllers.AdminHomeController{})
	beego.Router("/admin/modify", &controllers.AdminModifyController{})
	beego.Router("/admin/uplode", &controllers.AdminUplodeController{})
	beego.Router("/admin/wxlist", &controllers.AdminWcListController{})
	beego.Router("/admin/feedbacklist", &controllers.AdminFeedbackListController{})
	beego.Router("/admin/feedbackcontent", &controllers.AdminFeedbackContentController{})
	beego.Router("/admin/addwxlist", &controllers.AdminAddPublicNumberController{})
}
