package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
	"time"
)

/**
冲洗模块
*/

type PhotoController struct {
	beego.Controller
}

/******前台******/
/**
主页
*/
func (c *PhotoController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
		account := c.Ctx.GetCookie("cookei_photo_account")
		if len(account) != 0 {
			obj, err := models.GetRUserAccount(account)
			if err != nil {
				c.Redirect("/rense/login", 302)
				return
			} else {
				beego.Debug(obj)
			}
		} else {
			c.Redirect("/photo/login", 302)
			return
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
		phone := c.Input().Get("phone")
		pwd := c.Input().Get("pwd")
		// beego.Debug("rhome c.Input():", c.Input())
		// beego.Debug("rhome phone:", phone)
		// beego.Debug("rhome pwd:", pwd)
		if len(phone) != 0 && len(pwd) != 0 {
			obj, err := models.GetRUserPP(phone)
			if err != nil {
				c.Redirect("/rense/login", 302)
				return
			} else {
				beego.Debug(obj)
				if obj.Pwd == pwd {

				} else {
					c.Redirect("/photo/login", 302)
					return
				}
			}
		} else {
			c.Redirect("/photo/login", 302)
			return
		}
	}

	c.TplNames = "rhome.html"
}

/**
注册
*/
func (c *PhotoController) Register() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Register Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Register Post")
	}
	c.TplNames = "rregister.html"
}

/**
登录
*/
func (c *PhotoController) Login() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Login Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Login Post")
	}
	c.TplNames = "rlogin.html"
}

/**
上传
*/
func (c *PhotoController) Upload() {

	c.TplNames = "rupload.html"
}

/******后台******/

/******AJAX请求******/

func (c *PhotoController) RequestAjax() {
	response_obj := models.RinseJson{}
	response_obj.ErrCode = 1
	response_obj.ErrMsg = "未知错误"
	response_json := "response_json"
	if c.Ctx.Input.IsGet() {
		beego.Debug("RequestAjax Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("RequestAjax Post")
	}
	rtype := c.Input().Get("rtype")
	// beego.Debug("RequestAjax rtype :", rtype)
	response_obj.Rtype = rtype
	switch rtype {
	case "register":
		phone := c.Input().Get("phone")
		pwd := c.Input().Get("pwd")
		// verify := c.Input().Get("verify")
		// beego.Debug(c.Input())
		// beego.Debug("RequestAjax phone", phone)
		// beego.Debug("RequestAjax pwd", pwd)
		if len(phone) != 0 && len(pwd) != 0 {
			obj, err := models.GetRUser(phone)
			if err != nil {
				beego.Error(err)
			} else {
				if obj == nil {
					// beego.Debug("RequestAjax AddRUserPhone", phone)
					account := createAccount(phone, pwd)
					rid, err := createId()
					if err != nil {
						beego.Error(err)
					} else {
						err = models.AddRUser(phone, pwd, account, rid)
						if err != nil {
							beego.Error(err)
							response_obj.ErrCode = 1
							response_obj.ErrMsg = "注册：注册失败"
						} else {
							response_obj.ErrCode = 0
							data := fmt.Sprintf(`{"phone":%s,"pwd":%s}`, phone, pwd)
							response_obj.Data = data
						}
					}
				} else {
					response_obj.ErrCode = 1
					response_obj.ErrMsg = "注册：用户已注册"
				}
			}
		} else {
			response_obj.ErrCode = 1
			response_obj.ErrMsg = "注册：用户名或者密码错误"
		}

		break
	case "login":
		phone := c.Input().Get("phone")
		pwd := c.Input().Get("pwd")
		if len(phone) != 0 && len(pwd) != 0 {
			obj, err := models.GetRUserPP(phone)
			if err != nil {
				beego.Error(err)
				response_obj.ErrCode = 1
				response_obj.ErrMsg = "登录：查询用户错误"
			} else {
				if pwd == obj.Pwd {
					response_obj.ErrCode = 0
					data := fmt.Sprintf(`{"phone":%s,"pwd":%s}`, phone, pwd)
					response_obj.Data = data
					response_obj.Phone = phone
					response_obj.Pwd = pwd
				} else {
					response_obj.ErrCode = 1
					response_obj.ErrMsg = "登录：密码错误"
				}
			}
		} else {
			response_obj.ErrCode = 1
			response_obj.ErrMsg = "登录：用户名或者密码错误"
		}
		break
	}
	//struct 到json str
	body, err := json.Marshal(response_obj)
	if err != nil {
		beego.Error(err)
	}
	response_json = string(body)
	beego.Debug("RequestAjax response_json:", response_json)
	c.Ctx.WriteString(response_json)
}

func createAccount(phone string, pwd string) string {
	create_time := time.Now().Unix()
	base_account := fmt.Sprintf("%s_%s_%d", phone, pwd, create_time)
	return base_account
}
func createId() (int64, error) {
	count, err := models.GetUserCount()
	if err != nil {
		beego.Error(err)
		return 0, err
	} else {
		return int64(10000) + count, nil
	}
}
