package controllers

import (
	"beernotes/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Web Home Post")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Web Home Post")
	}
	op := c.Input().Get("op")
	if op == "logout" {
		c.Ctx.SetCookie(BN_USERNAME, "", -1, "/")
		c.Ctx.SetCookie(BN_PASSWORD, "", -1, "/")
		c.Redirect("/web/home", 302)
		return
	}
	mytype := c.Input().Get("type")
	beego.Debug("mytype:", mytype)
	c.Data["Type"] = mytype
	if mytype == "wnumber" {
		objs, err := models.GetAllWNumber()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Objs"] = objs
	} else if mytype == "knowbrew" {
		objs, err := models.GetKnowBrewSAE(1, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Objs"] = objs
	} else if mytype == "website" {
		objs, err := models.GetAllRelatedState(1)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Objs"] = objs
	} else {
		objs, err := models.GetKnowBrewSAE(1, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Objs"] = objs
	}
	bool, username := chackUserAccount(c.Ctx)
	beego.Debug("bool:", bool)
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "home.html"
}

func (c *WebController) UpdateLog() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Web UpdateLog Post")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Web UpdateLog Post")
	}
	c.TplName = "updatelog.html"
}

func (c *WebController) App() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Web App Post")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Web App Post")
	}
	c.Data["Android"] = "android_download_512"
	c.Data["Ios"] = "ios_download_512"
	c.TplName = "app.html"
}
func (c *WebController) Login() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Web Login Post")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Web Login Post")
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		beego.Debug("username:", username)
		beego.Debug("password:", password)
		beego.Debug("autologin:", autologin)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneUser(username)
			beego.Debug("user:", admin)
			if err != nil {
				c.Redirect("/web/login", 302)
				return
			}
			if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
				maxAge := 0
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie(BN_USERNAME, username, maxAge, "/")
				c.Ctx.SetCookie(BN_PASSWORD, password, maxAge, "/")
				beego.Debug("login ok")
				c.Redirect("/web/home", 302)
				return
			} else {
				c.Redirect("/web/login", 302)
				return
			}
		} else {
			c.Redirect("/web/login", 302)
			return
		}
	}
	c.TplName = "login.html"
}
func (c *WebController) Register() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Web Register Post")
		err := c.Input().Get("err")
		if len(err) > 0 {
			if err == "1004" {
				c.Data["Error"] = "用户名已存在"
			} else if err == "1005" {
				c.Data["Error"] = "邮箱已注册"
			} else if err == "1006" {
				c.Data["Error"] = "数据库操作失败"
			} else {
				c.Data["Error"] = ""
			}
		} else {
			c.Data["Error"] = ""
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Web Register Post")
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		email := c.Input().Get("email")
		beego.Debug("username:", username)
		beego.Debug("password:", password)
		beego.Debug("email:", email)
		if len(username) != 0 && len(password) != 0 && len(email) != 0 {
			code, _, err := models.RegisterUser(username, password, email)
			if err != nil {
				beego.Error(err)
			}
			if code == 0 {
				maxAge := 0
				autologin := true
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie(BN_USERNAME, username, maxAge, "/")
				c.Ctx.SetCookie(BN_PASSWORD, password, maxAge, "/")
				c.Redirect("/web/home", 302)
				return
			} else if code == APP_ERRCODE1004 {
				c.Redirect("/web/register?err=1004", 302)
				return
			} else if code == APP_ERRCODE1005 {
				c.Redirect("/web/register?err=1005", 302)
				return
			} else if code == APP_ERRCODE1006 {
				c.Redirect("/web/register?err=1006", 302)
				return
			}
		} else {
			c.Data["Error"] = "参数错误"
			c.Redirect("/web/register", 302)
			return
		}
	}
	c.TplName = "register.html"
}

func chackUserAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie(BN_USERNAME)
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie(BN_PASSWORD)
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneUser(username)
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
