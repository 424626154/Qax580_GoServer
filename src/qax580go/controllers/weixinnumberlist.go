package controllers

/*
推荐微信公众号
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WeixinNumberListController struct {
	beego.Controller
}

func (c *WeixinNumberListController) Get() {
	getWeixinNumberListUser(c)
	objs, err := models.GetAllWeixinNumbersState1()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(objs)
	c.TplName = "weixinnumberlist.html"
	c.Data["WeixinNumbers"] = objs
}
func getWeixinNumberListUser(c *WeixinNumberListController) {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
}
