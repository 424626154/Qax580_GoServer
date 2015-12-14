package controllers

/*
我的消息
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MyMessageController struct {
	beego.Controller
}

func (c *MyMessageController) Get() {
	openid := getMyMessageCookie(c)

	op := c.Input().Get("op")
	switch op {
	case "del":
		if len(openid) != 0 {
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.DeletePostOpenid(id, openid)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("MyMessageController delete id = %s openid = %s ", id, openid)
			c.Redirect("/mymessage", 302)
		}
		return
	}

	if len(openid) != 0 {
		posts, err := models.GetAllPostsOpenid(openid)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Posts"] = posts
	}
	c.TplNames = "mymessage.html"
}

func (c *MyMessageController) Post() {
	c.TplNames = "mymessage.html"
}
func getMyMessageCookie(c *MyMessageController) string {
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
