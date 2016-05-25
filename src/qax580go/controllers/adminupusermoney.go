package controllers

/*
发布消息
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
	"strconv"
)

type AdminUpUserMoneyController struct {
	beego.Controller
}

func (c *AdminUpUserMoneyController) Get() {
	id := c.Input().Get("id")
	c.Data["IsId"] = false
	if len(id) != 0 {
		user, err := models.GetOneWxUserInfoId(id)
		if err != nil {

		} else {
			c.Data["User"] = user
			c.Data["IsId"] = true
		}
	}
	c.TplName = "adminupusermoney.html"
}

func (c *AdminUpUserMoneyController) Post() {
	id := c.Input().Get("id")
	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "up":
		beego.Debug("op  ok")
		if len(id) != 0 {
			user, err := models.GetOneWxUserInfoId(id)
			if err != nil {
				beego.Error(err)
			} else {
				moneytype := c.Input().Get("moneytype")
				if len(moneytype) != 0 {
					addmoney := int64(0)
					if moneytype == "1" {
						addmoney = MONEY_SUBSCRIBE_SUM
					} else if moneytype == "2" {
						addmoney = MONEY_EXAMINE_SUM
					} else if moneytype == "3" {
						addmoney = MONEY_BELIKE_SUM
					}
					err = models.AddWxUserMoney(user.OpenId, addmoney)
					moneytype_i, err := strconv.ParseInt(moneytype, 10, 64)
					if err != nil {
						beego.Debug("err :", err)
					}
					_, err = models.AddUserMoneyRecord(user.OpenId, addmoney, moneytype_i)
					if err != nil {
						beego.Error(err)
					}
				}
			}
		}
		c.Redirect("/admin/wxuserlist", 302)
		return
	}
	c.TplName = "adminupusermoney.html"
}
