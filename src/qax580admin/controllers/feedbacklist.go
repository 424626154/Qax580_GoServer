package controllers

import (
	"github.com/astaxie/beego"
	"qax580admin/models"
)

type FeedbackListController struct {
	beego.Controller
}

func (c *FeedbackListController) Get() {
	feedbacks, err := models.GetAllFeedbacks()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "feedbacklist.html"
	c.Data["Feedbacks"] = feedbacks
	beego.Error(feedbacks)
}
