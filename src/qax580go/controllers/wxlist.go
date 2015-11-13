package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WcListController struct {
	beego.Controller
}

func (c *WcListController) Get() {
	getWXListUser(c)
	wxnums, err := models.GetAllWxnumsState1()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	c.TplNames = "wxlist.html"
	c.Data["Wxnums"] = wxnums
}
func getWXListUser(c *WcListController) {
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
