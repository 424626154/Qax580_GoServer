package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminFeedbackContentController struct {
	beego.Controller
}

func (c *AdminFeedbackContentController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		feedback, err := models.GetOneFeedback(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Feedback"] = feedback
		beego.Debug("is adminfeedbackcontent " + feedback.Info)
		c.TplNames = "adminfeedbackcontent.html"
		return
	}
	c.TplNames = "adminfeedbackcontent.html"

}
