package routers

import (
	"github.com/astaxie/beego"
	"os"
	"qax580go/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})                             //主页
	beego.Router("/uplode", &controllers.UplodeController{})                     //发布消息
	beego.Router("/content", &controllers.ContentController{})                   //消息详情
	beego.Router("/login", &controllers.LoginController{})                       //登录
	beego.Router("/register", &controllers.RegisterController{})                 //注册
	beego.Router("/wxlist", &controllers.WcListController{})                     //微推荐公众号列表
	beego.Router("/wx", &controllers.WXController{})                             //微信公众号服务器
	beego.Router("/feedback", &controllers.FeedbackController{})                 //意见反馈
	beego.Router("/wxautho", &controllers.WxAuthoController{})                   //微信用户信息
	beego.Router("/wxhome", &controllers.WxHomeController{})                     //微信回调主页
	beego.Router("/wxuplode", &controllers.WxUplodeController{})                 //微信回调上传
	beego.Router("/wxfeedback", &controllers.WxFeedbackController{})             //微信回调意见反馈
	beego.Router("/weather", &controllers.WeatherController{})                   //天气预报
	beego.Router("/traintickets", &controllers.TrainTicketsController{})         //火车票
	beego.Router("/querystation", &controllers.QueryStationController{})         //火车票起点终点查询
	beego.Router("/querytrain", &controllers.QueryTrainController{})             //火车票车次查询
	beego.Router("/queryrealtime", &controllers.QueryRealTimeController{})       //火车票实时查询
	beego.Router("/queryqutlets", &controllers.QueryQutletsController{})         //火车票代售点查询
	beego.Router("/querypeccancy", &controllers.QueryPeccancyController{})       //违章查询
	beego.Router("/history", &controllers.HistoryController{})                   //历史今天
	beego.Router("/historycon", &controllers.HistoryConController{})             //历史今天
	beego.Router("/laohuangli", &controllers.LaohuangliController{})             //老黄历
	beego.Router("/zhoubianwifiwx", &controllers.ZhouBianWifiWXController{})     //周边Wi-Fi
	beego.Router("/kuaidi", &controllers.KuaidiController{})                     //快递查询
	beego.Router("/tianqiwx", &controllers.TianqiWXController{})                 //天气查询
	beego.Router("/recommend", &controllers.RecommendController{})               //推荐
	beego.Router("/contactus", &controllers.ContactusController{})               //联系我们
	beego.Router("updatelog", &controllers.UpdateLogController{})                //更新日志
	beego.Router("/guanggaocontent", &controllers.GuanggaoContentController{})   //广告详情
	beego.Router("/waimailist", &controllers.WaimaiListController{})             //外卖订餐
	beego.Router("/caidans", &controllers.CaidansController{})                   //菜单
	beego.Router("/weixinnumberlist", &controllers.WeixinNumberListController{}) //推荐微信号
	beego.Router("/about", &controllers.AboutController{})                       //关于
	beego.Router("/wxgame", &controllers.WeixinGameController{})                 //微信游戏
	beego.Router("/mymessage", &controllers.MyMessageController{})               //我的发布
	beego.Router("/wxmymessage", &controllers.WxMyMessageController{})           //微信回调我的发布
	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})

	beego.Router("/admin/home", &controllers.AdminHomeController{})                             //后台主页
	beego.Router("/admin/modify", &controllers.AdminModifyController{})                         //修改信息
	beego.Router("/admin/uplode", &controllers.AdminUplodeController{})                         //后台上传
	beego.Router("/admin/wxlist", &controllers.AdminWcListController{})                         //公众号列表
	beego.Router("/admin/feedbacklist", &controllers.AdminFeedbackListController{})             //意见反馈列表
	beego.Router("/admin/feedbackcontent", &controllers.AdminFeedbackContentController{})       //意见反馈内容
	beego.Router("/admin/addwxlist", &controllers.AdminAddPublicNumberController{})             //添加微信公众号
	beego.Router("/admin/upwxnuminfo", &controllers.AdminUpWxnumInfoController{})               //修改微信公众号内容
	beego.Router("/admin/upwxnumimg", &controllers.AdminUpWxnumImgController{})                 //修改微信公众号图片
	beego.Router("/admin", &controllers.AdminLoginController{})                                 //后台登陆
	beego.Router("/admin/userlist", &controllers.AdminUserListController{})                     //后台可登录用户列表
	beego.Router("/admin/adduser", &controllers.AdminAddUserController{})                       //添加后台用户
	beego.Router("/admin/content", &controllers.AdminContentController{})                       //后台消息内容
	beego.Router("/admin/wxuserlist", &controllers.WxUserListController{})                      //后台统计公众号登录用户列表
	beego.Router("/admin/juhe", &controllers.AdminJuheController{})                             //聚合数据主页
	beego.Router("/admin/newskey", &controllers.AdminNewsKeyController{})                       //新闻关键词
	beego.Router("/admin/addguanggao", &controllers.AdminaAddGuanggaoController{})              //后台添加广告
	beego.Router("/admin/guanggaos", &controllers.AdminGuanggaosController{})                   //后台广告列表
	beego.Router("/admin/guanggaocontent", &controllers.AdminGuanggaoContentController{})       //后台广告内容
	beego.Router("/admin/upguanggaoinfo", &controllers.AdminUpGuanggaoInfoController{})         //后台修改广告内容
	beego.Router("/admin/upguanggaoimg", &controllers.AdminUpGuanggaoImgController{})           //后台修改广告图片
	beego.Router("/admin/waimailist", &controllers.AdminWaimaiListController{})                 //外卖列表
	beego.Router("/admin/addwaimai", &controllers.AdminAddWaimaiController{})                   //后台添加外卖
	beego.Router("/admin/caidans", &controllers.AdminCaidansController{})                       //后台菜单列表
	beego.Router("/admin/addcaidan", &controllers.AdminAddCaidanController{})                   //后台添加菜单
	beego.Router("/admin/addweixinnumber", &controllers.AdminAddWeixinNumberController{})       //后台添加推荐微信号
	beego.Router("/admin/weixinnumberlist", &controllers.AdminWeixinNumberListController{})     //后台推荐微信号列表
	beego.Router("/admin/upweixinnumberinfo", &controllers.AdminUpWeixinNumberInfoController{}) //后台修改微信号内容
	beego.Router("/admin/upweixinnumberimg", &controllers.AdminUpWeixinNumberImgController{})   //后台修改微信号图片
}
