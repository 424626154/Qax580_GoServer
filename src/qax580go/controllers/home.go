package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"qax580go/models"
	"strconv"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	getCookie(c)
	CurrentPage := int32(1)
	count, err := models.GetPostCount()
	NumberofPages := int32(10)
	temp := count / NumberofPages
	if (count % NumberofPages) != 0 {
		temp = temp + 1
	}
	CotalPages := temp
	pagetype := c.Input().Get("type")
	page := c.Input().Get("page")
	beego.Debug("pagetype:", pagetype)

	guanggaos, err := models.GetAllGuanggaosState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggaos"] = guanggaos

	if len(pagetype) != 0 && len(page) != 0 {
		switch pagetype {
		case "first": //首页
			CurrentPage = 1
		case "prev": //上一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "next": //下一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "last": //尾页
			CurrentPage = CotalPages
		case "page": //页码
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		}
	}
	c.Data["CurrentPage"] = CurrentPage
	c.Data["CotalPages"] = CotalPages
	c.Data["NumberofPages"] = NumberofPages
	posts, err := models.QueryPagePost(CurrentPage-1, NumberofPages)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(posts)
	c.TplNames = "home.html"
	c.Data["Posts"] = posts
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	beego.Debug(isdebug)
	c.Data["IsDebug"] = isdebug
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeletePost(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is del " + id)
		c.Redirect("/", 302)
		return
	}
}

func getCookie(c *HomeController) {
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
}
