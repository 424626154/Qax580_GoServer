package controllers

/*
后台主页
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminHomeController struct {
	beego.Controller
}

func (c *AdminHomeController) Get() {
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
	num, err := models.GetAllStateNum()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "adminhome.html"
	c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.Data["Num"] = num
	c.Data["All"] = len(posts)
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
		c.Redirect("/admin/home", 302)
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
		//添加审核金钱
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		} else {
			if post.Label == 1 {
				err = models.AddWxUserMoney(post.OpenId, MONEY_SUBSCRIBE_SUM)
				if err != nil {
					beego.Error(err)
				} else {
					_, err = models.AddUserMoneyRecord(post.OpenId, MONEY_EXAMINE_SUM, MONEY_EXAMINE)
					if err != nil {
						beego.Error(err)
					}
				}

			}
		}
		beego.Debug("is admin examine " + id)
		c.Redirect("/admin/home", 302)
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
		c.Redirect("/admin/home", 302)
		return
	case "back":
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/admin", 302)
		return
	}
}
