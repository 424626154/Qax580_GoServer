package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"niangxin/models"
	"strings"
)

type NNAdminController struct {
	beego.Controller
}

func (c *NNAdminController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = true
		c.Data["User"] = username
	} else {
		c.Redirect("/nnadmin/login", 302)
		return
	}
	c.TplName = "adminhome.html"
}

func (c *NNAdminController) Login() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Login Get")
		bool, _ := chackAccount(c.Ctx)
		if bool {
			c.Redirect("/nnadmin/home", 302)
			return
		} else {

		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Login Post")
		username := c.Input().Get("user")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetAdminUser(username)
			if err != nil {
				c.Redirect("/nnadmin/login", 302)
				return
			}
			if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
				maxAge := 0
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie("username", username, maxAge, "/")
				c.Ctx.SetCookie("password", password, maxAge, "/")
				beego.Debug("login ok")
				c.Redirect("/nnadmin/home", 302)
				return
			} else {
				c.Redirect("/nnadmin/login", 302)
				return
			}
		} else {
			c.Redirect("/nnadmin/login", 302)
			return
		}
	}
	c.TplName = "adminlogin.html"
}

func (c *NNAdminController) BeerVideo() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = true
		c.Data["User"] = username
	} else {
		c.Redirect("/nnadmin/login", 302)
		return
	}
	objs, err := models.GetAllBeerVideo()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	c.TplName = "beervideo.html"
}

func (c *NNAdminController) Post() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Post Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Post Post")
	}
	op := c.Input().Get("op")
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	beego.Debug("op:", op)
	switch op {
	case "addvideo":
		title := c.Input().Get("title")
		url := c.Input().Get("url")
		if len(title) != 0 && len(url) != 0 {
			err := models.AddBeerVideo(title, url)
			if err != nil {
				beego.Debug(err)
			} else {
				request_json = `{"errcode":0,"errmsg":""}`
			}
		}
		break
	case "beervideolist":
		objs, err := models.GetAllBeerVideo()
		if err != nil {
			beego.Error(err)
		} else {
			objs_json, err := json.Marshal(objs)
			if err != nil {
				beego.Error(err)
			} else {
				request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, objs_json)
			}
		}
		break
	case "del":
		id := c.Input().Get("id")
		if len(id) != 0 {
			err := models.DelBeerVideo(id)
			if err != nil {
				beego.Debug(err)
			} else {
				request_json = `{"errcode":0,"errmsg":""}`
			}
		}
		break
	}
	beego.Debug("request_json:", request_json)
	c.Ctx.WriteString(request_json)
}

func chackAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetAdminUser(username)
	if err != nil {
		return false, ""
	}
	if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
		beego.Debug(" cookie username ", username)
		return true, username
	} else {
		return false, username
	}

}
