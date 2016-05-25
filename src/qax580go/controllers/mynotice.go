package controllers

/*
我的系统消息
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MynoticeController struct {
	beego.Controller
}

func (c *MynoticeController) Get() {
	openid := getMynoticeCookie(c)

	op := c.Input().Get("op")
	beego.Debug("op", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) != 0 {
			err := models.DeleteUserNotice(id)
			if err != nil {
				beego.Error(err)
			}
		}
		beego.Debug("del id", id)
		c.Redirect("/mynotice", 302)
		return
	}

	objs, err := models.GetUeserAllNotice(openid)
	if err != nil {
		beego.Error(err)
	}
	for i := 0; i < len(objs); i++ {
		err := models.UpUeserNoticeRead(objs[i].Id, 1)
		if err != nil {
			beego.Error(err)
		} else {
			objs[i].ToRead = 1
		}
	}
	beego.Debug(objs)
	c.Data["Objs"] = objs
	c.TplName = "mynotice.html"

}

func getMynoticeCookie(c *MynoticeController) string {
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
