package controllers

import (
	"github.com/astaxie/beego"
)

type NiangXinController struct {
	beego.Controller
}

func (c *NiangXinController) App() {
	op := c.Input().Get("op")
	beego.Debug("op ：", op)
	switch op {
	case "register":
		break
	}
	c.Ctx.WriteString("response_json")
}
