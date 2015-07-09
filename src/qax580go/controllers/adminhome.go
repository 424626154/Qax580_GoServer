package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminHomeController struct {
	beego.Controller
}

func (c *AdminHomeController) Get() {

	posts, err := models.GetAllPostsAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "adminhome.html"
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
		beego.Debug("is admin del " + id)
		c.Redirect("/admin", 302)
		return
	case "examine":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePost(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin examine " + id)
		c.Redirect("/admin", 302)
		return
	case "examine1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePostExamine(id, 0)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin examine1" + id)
		c.Redirect("/admin", 302)
		return
	}
}
