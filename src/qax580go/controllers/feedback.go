package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type FeedbackController struct {
	beego.Controller
}

func (c *FeedbackController) Get() {

	c.TplNames = "feedback.html"

	info := c.Input().Get("info")
	if len(info) != 0 {
		err := models.AddFeedback(info)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/", 302)
	}
}
