package controllers

import (
	"github.com/astaxie/beego"
)

type UplodeImgController struct {
	beego.Controller
}

func (c *UplodeImgController) Get() {
	c.TplNames = "uplodeimg.html"
}

func (c *UplodeImgController) Post() {
	c.TplNames = "uplodeimg.html"
}
