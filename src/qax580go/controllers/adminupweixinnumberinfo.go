package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminUpWeixinNumberInfoController struct {
	beego.Controller
}

func (c *AdminUpWeixinNumberInfoController) Get() {
	id := c.Input().Get("id")
	obj, err := models.GetOneWeixinNumber(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["WeixinNumber"] = obj
	c.TplNames = "adminupweixinnumberinfo.html"
}
func (c *AdminUpWeixinNumberInfoController) Post() {
	id := c.Input().Get("id")
	name := c.Input().Get("name")
	info := c.Input().Get("info")
	number := c.Input().Get("number")
	evaluate := c.Input().Get("evaluate")
	if len(id) != 0 && len(name) != 0 && len(info) != 0 && len(number) != 0 {
		err := models.UpdateWeixinNumberInfo(id, name, info, number, evaluate)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/weixinnumberlist", 302)
		return
	}
	c.TplNames = "adminupweixinnumberinfo.html"
}
