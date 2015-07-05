package controllers

import (
	"github.com/astaxie/beego"
	"qax580admin/models"
)

type UplodeController struct {
	beego.Controller
}

func (c *UplodeController) Get() {

	c.TplNames = "uplode.html"

	title := c.Input().Get("title")
	info := c.Input().Get("info")

	if len(title) != 0 && len(info) != 0 {
		err := models.AddPostLabel(title, info, 2, "")
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/", 302)
	}
}
