package controllers

import (
	"github.com/astaxie/beego"
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
	c.TplNames = "home.html"
	c.Data["Posts"] = posts
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
