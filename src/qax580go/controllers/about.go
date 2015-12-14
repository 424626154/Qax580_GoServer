package controllers

/*
关于
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	getAboutCookie(c)
	c.TplNames = "about.html"
}

func (c *AboutController) Post() {
	c.TplNames = "about.html"
}

func getAboutCookie(c *AboutController) string {
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
