package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminFeedbackListController struct {
	beego.Controller
}

func (c *AdminFeedbackListController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	posts, err := models.GetAllPostsAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username

	feedbacks, err := models.GetAllFeedbacks()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "adminfeedbacklist.html"
	c.Data["Feedbacks"] = feedbacks
	beego.Error(feedbacks)
}
