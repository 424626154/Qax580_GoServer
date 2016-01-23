package controllers

/*
金钱详情
*/
import (
	"github.com/astaxie/beego"
)

type MoneyHelpController struct {
	beego.Controller
}

func (c *MoneyHelpController) Get() {
	c.TplNames = "moneyhelp.html"
}
func (c *MoneyHelpController) Post() {
	c.TplNames = "moneyhelp.html"
}
