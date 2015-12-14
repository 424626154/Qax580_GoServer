package controllers

/*
联系我们
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ContactusController struct {
	beego.Controller
}

func (c *ContactusController) Get() {
	getConCookie(c)
	c.TplNames = "contactus.html"
}

func (c *ContactusController) Post() {
	c.TplNames = "contactus.html"
}

func getConCookie(c *ContactusController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
