package controllers

import (
	"github.com/astaxie/beego"
)

type AdminJuheController struct {
	beego.Controller
}

func (c *AdminJuheController) Get() {
	c.TplNames = "adminjuhe.html"
}

func (c *AdminJuheController) Post() {
	c.TplNames = "adminjuhe.html"
}
