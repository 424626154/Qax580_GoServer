package controllers

/*
后台微信用户列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WxUserListController struct {
	beego.Controller
}

func (c *WxUserListController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplNames = "adminwxuserlist.html"
	admins, err := models.GetAllWxUsers()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(admins)
	c.Data["WxUsers"] = admins
}
