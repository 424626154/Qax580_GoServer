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
	beego.Router("/wxautho", &controllers.WxAuthoController{})
	beego.Router("/wxhome", &controllers.WxHomeController{})
	beego.Router("/wxuplode", &controllers.WxUplodeController{})
	beego.Router("/wxfeedback", &controllers.WxFeedbackController{})
	beego.Router("/weather", &controllers.WeatherController{})
	beego.Router("/traintickets", &controllers.TrainTicketsController{})
	beego.Router("/querystation", &controllers.QueryStationController{})
	beego.Router("/querytrain", &controllers.QueryTrainController{})
	beego.Router("/queryrealtime", &controllers.QueryRealTimeController{})
	beego.Router("/queryqutlets", &controllers.QueryQutletsController{})
	beego.Router("/querypeccancy", &controllers.QueryPeccancyController{})
	beego.Router("/history", &controllers.HistoryController{})
	beego.Router("/historycon", &controllers.HistoryConController{})
	beego.Router("/laohuangli", &controllers.LaohuangliController{})
	beego.Router("/zhoubianwifiwx", &controllers.ZhouBianWifiWXController{})
	beego.Router("/kuaidi", &controllers.KuaidiController{})
	beego.Router("/tianqiwx", &controllers.TianqiWXController{})
	beego.Router("/recommend", &controllers.RecommendController{})
	beego.Router("/contactus", &controllers.ContactusController{}) //联系我们
	beego.Router("updatelog", &controllers.UpdateLogController{})  //更新日志
	beego.Router("/guanggaocontent", &controllers.GuanggaoContentController{})
	beego.Router("/waimailist", &controllers.WaimaiListController{}) //外卖订餐
	beego.Router("/caidans", &controllers.CaidansController{})       //菜单
	beego.Router("/weixinnumberlist", &controllers.WeixinNumberListController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/wxgame", &controllers.WeixinGameController{}) //微信游戏
	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})

	beego.Router("/admin/home", &controllers.AdminHomeController{})
	beego.Router("/admin/modify", &controllers.AdminModifyController{})
	beego.Router("/admin/uplode", &controllers.AdminUplodeController{})
	beego.Router("/admin/wxlist", &controllers.AdminWcListController{})
	beego.Router("/admin/feedbacklist", &controllers.AdminFeedbackListController{})
	beego.Router("/admin/feedbackcontent", &controllers.AdminFeedbackContentController{})
	beego.Router("/admin/addwxlist", &controllers.AdminAddPublicNumberController{})
	beego.Router("/admin/upwxnuminfo", &controllers.AdminUpWxnumInfoController{})
	beego.Router("/admin/upwxnumimg", &controllers.AdminUpWxnumImgController{})
	beego.Router("/admin", &controllers.AdminLoginController{})
	beego.Router("/admin/userlist", &controllers.AdminUserListController{})
	beego.Router("/admin/adduser", &controllers.AdminAddUserController{})
	beego.Router("/admin/content", &controllers.AdminContentController{})
	beego.Router("/admin/wxuserlist", &controllers.WxUserListController{})
	beego.Router("/admin/juhe", &controllers.AdminJuheController{})
	beego.Router("/admin/newskey", &controllers.AdminNewsKeyController{}) //新闻关键词
	beego.Router("/admin/addguanggao", &controllers.AdminaAddGuanggaoController{})
	beego.Router("/admin/guanggaos", &controllers.AdminGuanggaosController{})
	beego.Router("/admin/guanggaocontent", &controllers.AdminGuanggaoContentController{})
	beego.Router("/admin/upguanggaoinfo", &controllers.AdminUpGuanggaoInfoController{})
	beego.Router("/admin/upguanggaoimg", &controllers.AdminUpGuanggaoImgController{})
	beego.Router("/admin/waimailist", &controllers.AdminWaimaiListController{}) //外卖列表
	beego.Router("/admin/addwaimai", &controllers.AdminAddWaimaiController{})
	beego.Router("/admin/caidans", &controllers.AdminCaidansController{})
	beego.Router("/admin/addcaidan", &controllers.AdminAddCaidanController{})
	beego.Router("/admin/addweixinnumber", &controllers.AdminAddWeixinNumberController{})
	beego.Router("/admin/weixinnumberlist", &controllers.AdminWeixinNumberListController{})
	beego.Router("/admin/upweixinnumberinfo", &controllers.AdminUpWeixinNumberInfoController{})
	beego.Router("/admin/upweixinnumberimg", &controllers.AdminUpWeixinNumberImgController{})
}
