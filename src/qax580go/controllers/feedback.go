package controllers

/*
意见反馈
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type FeedbackController struct {
	beego.Controller
}

func (c *FeedbackController) Get() {
	openid := getFeedbackCookie(c)
	c.TplNames = "feedback.html"

	info := c.Input().Get("info")
	if len(info) != 0 {
		beego.Debug("------------AddFeedback--------")
		beego.Debug(openid)
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		}
		err = models.AddFeedback(info, openid, wxuser.NickeName, wxuser.Sex, wxuser.HeadImgurl)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/", 302)
	}
}

func getFeedbackCookie(c *FeedbackController) string {
	isUser := false
	openid := c.Ctx.GetCookie("wx_openid")
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
