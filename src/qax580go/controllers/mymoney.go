package controllers

/*
我的帮帮币
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MymoneyController struct {
	beego.Controller
}

func (c *MymoneyController) Get() {
	getMymoneyCookie(c)
	c.TplName = "mymoney.html"

}

func getMymoneyCookie(c *MymoneyController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
