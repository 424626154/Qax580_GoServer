package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWcListController struct {
	beego.Controller
}

func (c *AdminWcListController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	posts, err := models.GetAllPostsAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username

	beego.Debug("AdminWcListController")
	wxnums, err := models.GetAllWxnums()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	c.TplNames = "adminwxlist.html"
	c.Data["Wxnums"] = wxnums
}
