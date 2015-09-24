package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"qax580go/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {

	posts, err := models.GetAllPosts()
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
