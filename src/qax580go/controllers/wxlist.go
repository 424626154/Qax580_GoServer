package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WcListController struct {
	beego.Controller
}

func (this *WcListController) Get() {
	wxnums, err := models.GetAllWxnums()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	this.TplNames = "wxlist.html"
	this.Data["Wxnums"] = wxnums
}
