package controllers

/*
后台修改信息
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminModifyController struct {
	beego.Controller
}

func (c *AdminModifyController) Get() {

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
	case "m":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		post, err := models.GetOnePost(id)
		// err := models.DeletePost(id)
		// if err != nil {
		//  beego.Error(err)
		// }
		if err != nil {
			break
		}
		c.Data["Post"] = post
		beego.Debug("is con " + post.Title)
		c.TplNames = "adminmodify.html"
		return
	default:
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		city := c.Input().Get("city")
		beego.Debug("is con " + title)
		if len(id) != 0 && len(title) != 0 && len(info) != 0 {
			err := models.UpdatePostInfo(id, title, info, city)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin", 302)
		}
		return
	}

}
