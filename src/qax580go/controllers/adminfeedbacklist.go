package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminFeedbackListController struct {
	beego.Controller
}

func (c *AdminFeedbackListController) Get() {
	feedbacks, err := models.GetAllFeedbacks()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "adminfeedbacklist.html"
	c.Data["Feedbacks"] = feedbacks
	beego.Error(feedbacks)
}
