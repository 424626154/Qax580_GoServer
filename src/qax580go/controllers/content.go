package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ContentController struct {
	beego.Controller
}

func (c *ContentController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Post"] = post
		beego.Debug("is con " + post.Title)
		c.TplNames = "content.html"
		return
	}
	c.TplNames = "content.html"

}
