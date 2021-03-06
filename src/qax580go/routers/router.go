package routers

import (
	"github.com/astaxie/beego"
	"os"
	"qax580go/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})                                    //主页
	beego.Router("/uplode", &controllers.UplodeController{})                            //发布消息
	beego.Router("/content", &controllers.ContentController{})                          //消息详情
	beego.Router("/login", &controllers.LoginController{})                              //登录
	beego.Router("/register", &controllers.RegisterController{})                        //注册
	beego.Router("/wxlist", &controllers.WcListController{})                            //微推荐公众号列表
	beego.Router("/wx", &controllers.WXController{})                                    //微信公众号服务器
	beego.Router("/wxtl", &controllers.WXTlController{})                                //微信公众号服务器铁力
	beego.Router("/feedback", &controllers.FeedbackController{})                        //意见反馈
	beego.Router("/wxautho", &controllers.WxAuthoController{})                          //微信用户信息
	beego.Router("/wxhome", &controllers.WxHomeController{})                            //微信回调主页
	beego.Router("/wxuplode", &controllers.WxUplodeController{})                        //微信回调上传
	beego.Router("/wxfeedback", &controllers.WxFeedbackController{})                    //微信回调意见反馈
	beego.Router("/weather", &controllers.WeatherController{})                          //天气预报
	beego.Router("/traintickets", &controllers.TrainTicketsController{})                //火车票
	beego.Router("/querystation", &controllers.QueryStationController{})                //火车票起点终点查询
	beego.Router("/querytrain", &controllers.QueryTrainController{})                    //火车票车次查询
	beego.Router("/queryrealtime", &controllers.QueryRealTimeController{})              //火车票实时查询
	beego.Router("/queryqutlets", &controllers.QueryQutletsController{})                //火车票代售点查询
	beego.Router("/querypeccancy", &controllers.QueryPeccancyController{})              //违章查询
	beego.Router("/history", &controllers.HistoryController{})                          //历史今天
	beego.Router("/historycon", &controllers.HistoryConController{})                    //历史今天
	beego.Router("/laohuangli", &controllers.LaohuangliController{})                    //老黄历
	beego.Router("/zhoubianwifiwx", &controllers.ZhouBianWifiWXController{})            //周边Wi-Fi
	beego.Router("/kuaidi", &controllers.KuaidiController{})                            //快递查询
	beego.Router("/tianqiwx", &controllers.TianqiWXController{})                        //天气查询
	beego.Router("/recommend", &controllers.RecommendController{})                      //推荐
	beego.Router("/contactus", &controllers.ContactusController{})                      //联系我们
	beego.Router("updatelog", &controllers.UpdateLogController{})                       //更新日志
	beego.Router("/guanggaocontent", &controllers.GuanggaoContentController{})          //广告详情
	beego.Router("/waimailist", &controllers.WaimaiListController{})                    //外卖订餐
	beego.Router("/caidans", &controllers.CaidansController{})                          //菜单
	beego.Router("/weixinnumberlist", &controllers.WeixinNumberListController{})        //推荐微信号
	beego.Router("/about", &controllers.AboutController{})                              //关于
	beego.Router("/wxgame", &controllers.WeixinGameController{})                        //微信游戏
	beego.Router("/mymessage", &controllers.MyMessageController{})                      //我的发布
	beego.Router("/wxmymessage", &controllers.WxMyMessageController{})                  //微信回调我的发布
	beego.Router("/myextension", &controllers.MyExtensionController{})                  //我的推广
	beego.Router("/mymoney", &controllers.MymoneyController{})                          //我的帮帮币
	beego.Router("/mmyextensionresponse", &controllers.MyExtensionResponseController{}) //我的推广返回测试
	beego.Router("/moneyinfo", &controllers.MoneyInfoController{})                      //我的金钱详情
	beego.Router("/moneyhelp", &controllers.MoneyHelpController{})                      //我的金钱帮助
	beego.Router("/mall", &controllers.MallController{})                                //商城
	beego.Router("/exchange", &controllers.ExchangeController{})                        //兑换
	beego.Router("/subsribe", &controllers.SubsribeController{})                        //关注与取消关注
	beego.Router("/shanghus", &controllers.ShangHusController{})                        //商户列表
	beego.Router("/shanghulist", &controllers.ShangHuListController{})                  //商户子列表
	beego.Router("/mynotice", &controllers.MynoticeController{})                        //我的消息

	beego.AutoRouter(&controllers.WxqaxController{}) //微信http自动匹配

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})

	os.Mkdir("imageserver", os.ModePerm)
	beego.Router("/imageserver/:all", &controllers.ImageHostingController{})

	os.Mkdir("filehosting", os.ModePerm)
	beego.Router("/filehosting/:all", &controllers.FileHostingController{})

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
	beego.Router("/admin/upusermoney", &controllers.AdminUpUserMoneyController{})               //后台用户金钱
	beego.Router("/admin/moneyinfo", &controllers.AdminMoneyInfoController{})                   //后台用户金钱记录
	beego.Router("admin/importuser", &controllers.AdminImportUserController{})                  //后台导入微信用户
	beego.Router("admin/upwxuserinfo", &controllers.AdminUpWxuserInfoController{})              //后台导入微信用户
	beego.Router("admin/mall", &controllers.AdminMallController{})                              //后台商城
	beego.Router("admin/addcommodity", &controllers.AdminaAddCommodityController{})             //添加商品
	beego.Router("admin/upcommodityinfo", &controllers.AdminUpCommodityInfoController{})        //修改商品信息
	beego.Router("admin/upcommodityimg", &controllers.AdminUpCommodityImgController{})          //修改商品图片
	beego.Router("admin/exchange", &controllers.AdminExchangeController{})                      //用户兑换
	beego.Router("admin/shanghus", &controllers.AdminShanghusController{})                      //后台商户
	beego.Router("admin/addshanghu", &controllers.AdminAddShanghuController{})                  //添加商户
	beego.Router("admin/upshanghuinfo", &controllers.AdminUpShangHuInfoController{})            //修改商户信息
	beego.Router("admin/upshanghuimg", &controllers.AdminUpShangHuImgController{})              //修改商户图片
	beego.Router("admin/keywords", &controllers.AdminKeywordsController{})                      //关键字列表
	beego.Router("admin/addkeywords", &controllers.AdminaAddKeywordsController{})               //添加关键字
	beego.Router("admin/keyobj", &controllers.AdminKeyobjController{})                          //关键字内容
	beego.Router("admin/addkeyobj", &controllers.AdminaAddKeyobjController{})                   //添加关键字内容
	beego.Router("admin/wxtest", &controllers.AdminWxTestController{})                          //添加关键字内容
	beego.Router("admin/updatelog", &controllers.AdminUpdateLogController{})                    //后台更新日志
	beego.Router("admin/notice", &controllers.AdminNoticeController{})                          //后台通知管理

	beego.AutoRouter(&controllers.PollController{})      //投票系统
	beego.AutoRouter(&controllers.RinseController{})     //冲洗系统
	beego.AutoRouter(&controllers.WptController{})       //微信平台
	beego.AutoRouter(&controllers.AdminPostController{}) //后台提交post
	beego.AutoRouter(&controllers.PhotoController{})     //洗相系统

	beego.AutoRouter(&controllers.ImageController{}) //图床

	beego.AutoRouter(&controllers.WeiZhanController{}) //微站
	beego.AutoRouter(&controllers.DqsjController{})    //大签世界

	beego.Router("/admin/dqsj", &controllers.AdminDqsjUserListController{})
	beego.Router("/admin/adddqsjuser", &controllers.AdminAddDqsjUserController{})

	beego.AutoRouter(&controllers.WxAppController{}) //微信小程序

	beego.AutoRouter(&controllers.BeerMapController{}) //大签世界

}
