package controllers

import (
	"github.com/astaxie/beego"
)

type WcListController struct {
	beego.Controller
}

func (this *WcListController) Get() {

	this.TplNames = "wxlist.html"
}
