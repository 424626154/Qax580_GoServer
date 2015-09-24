package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminFeedbackContentController struct {
	beego.Controller
}

func (c *AdminFeedbackContentController) Get() {
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
