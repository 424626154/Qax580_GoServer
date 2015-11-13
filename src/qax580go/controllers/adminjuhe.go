package controllers

import (
	"github.com/astaxie/beego"
)

type AdminJuheController struct {
	beego.Controller
}

func (c *AdminJuheController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.TplNames = "adminjuhe.html"
}

func (c *AdminJuheController) Post() {
	c.TplNames = "adminjuhe.html"
}
