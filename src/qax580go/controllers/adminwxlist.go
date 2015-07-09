package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWcListController struct {
	beego.Controller
}

func (c *AdminWcListController) Get() {
	beego.Debug("AdminWcListController")
	wxnums, err := models.GetAllWxnums()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	c.TplNames = "adminwxlist.html"
	c.Data["Wxnums"] = wxnums
}
