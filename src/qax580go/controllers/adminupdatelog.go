package controllers

import (
	"github.com/astaxie/beego"
)

type AdminUpdateLogController struct {
	beego.Controller
}

func (c *AdminUpdateLogController) Get() {
	c.TplNames = "adminupdatelog.html"
}

func (c *AdminUpdateLogController) Post() {
	c.TplNames = "adminupdatelog.html"
}
