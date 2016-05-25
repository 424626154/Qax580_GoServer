package controllers

/*
商城
*/
import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MallController struct {
	beego.Controller
}

func (c *MallController) Get() {
	beego.Debug("MallController Get")
	getMallInfoCookie(c)
	c.TplName = "mall.html"
}
func (c *MallController) Post() {
	beego.Debug("MallController Post")
	getMallInfoCookie(c)
	commoditys, err := models.GetAllCommoditys()
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("commoditys :", commoditys)
	c.Data["Commoditys"] = commoditys
	op := c.Input().Get("op")
	beego.Debug("op", op)
	if op == "exchange" {
		id := c.Input().Get("id")
		beego.Debug("id", id)
		openid := c.Input().Get("openid")
		beego.Debug("openid", openid)
		if len(id) != 0 && len(openid) != 0 {
			err := models.AddUorder(openid, id)
			if err != nil {
				beego.Error(err)
			}
		}
		if len(openid) != 0 {
			url := fmt.Sprintf("/exchange?openid=%s", openid)
			c.Redirect(url, 302)
			return
		}
	}
	c.TplName = "mall.html"
}
func getMallInfoCookie(c *MallController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
