package controllers

import (
	"beernotes/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/widuu/gojson"
	"gopkg.in/gomail.v2"
	"strings"
)

type AppController struct {
	beego.Controller
}

type AppTokenModels struct {
	Token  string `json:"tonekn,omitempty`
	Ostype string `json:"ostype,omitempty`
}

//注册
func (c *AppController) Register() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP Register Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP Register Post")
		// username := c.Input().Get("username")
		// password := c.Input().Get("password")
		// email := c.Input().Get("email")
		// beego.Debug("username:", c.Input().Get("username"))
		req_json := string(c.Ctx.Input.RequestBody)
		username := gojson.Json(req_json).Get("username").Tostring()
		password := gojson.Json(req_json).Get("password").Tostring()
		email := gojson.Json(req_json).Get("email").Tostring()
		if len(username) != 0 && len(password) != 0 && len(email) != 0 {
			code, user, err := models.RegisterUser(username, password, email)
			if err != nil {
				beego.Error(err)
			}
			if code == 0 {
				res.Errcode = 0
				res.Errmsg = "注册成功"
				data_json, err := json.Marshal(user)
				if err != nil {
					beego.Error(err)
				}
				res.Data = string(data_json)
			} else if code == APP_ERRCODE1004 {
				res.Errcode = APP_ERRCODE1004
				res.Errmsg = APP_ERRMSG1004
			} else if code == APP_ERRCODE1005 {
				res.Errcode = APP_ERRCODE1005
				res.Errmsg = APP_ERRMSG1005
			} else if code == APP_ERRCODE1006 {
				res.Errcode = APP_ERRCODE1006
				res.Errmsg = APP_ERRMSG1006
			}
		} else {
			res.Errcode = APP_ERRCODE1003
			res.Errmsg = APP_ERRMSG1003
		}
	}

	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

//登录
func (c *AppController) Login() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP Register Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP Register Post")
		req_json := string(c.Ctx.Input.RequestBody)
		username := gojson.Json(req_json).Get("username").Tostring()
		password := gojson.Json(req_json).Get("password").Tostring()
		obj, err := models.GetOneUser(username)
		if err != nil {
			beego.Error(err)
		}
		if obj == nil {
			res.Errcode = APP_ERRCODE1007
			res.Errmsg = APP_ERRMSG1007
		} else {
			if obj.Password == password {
				res.Errcode = 0
				res.Errmsg = "登录成功"
				data_json, err := json.Marshal(obj)
				if err != nil {
					beego.Error(err)
				}
				res.Data = string(data_json)
			} else {
				res.Errcode = APP_ERRCODE1008
				res.Errmsg = APP_ERRMSG1008
			}
		}
	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

//修改密码
func (c *AppController) Modify() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP Register Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP Register Post")
		req_json := string(c.Ctx.Input.RequestBody)
		username := gojson.Json(req_json).Get("username").Tostring()
		password := gojson.Json(req_json).Get("password").Tostring()
		newpass := gojson.Json(req_json).Get("newpass").Tostring()
		code, obj, err := models.ModifyPassword(username, password, newpass)
		if err != nil {
			beego.Error(err)
		}
		if code == 0 {
			res.Errcode = 0
			res.Errmsg = "修改密码成功"
			data_json, err := json.Marshal(obj)
			if err != nil {
				beego.Error(err)
			}
			res.Data = string(data_json)
		} else if code == APP_ERRCODE1003 {
			res.Errcode = APP_ERRCODE1003
			res.Errmsg = APP_ERRMSG1003
		} else if code == APP_ERRCODE1007 {
			res.Errcode = APP_ERRCODE1007
			res.Errmsg = APP_ERRMSG1007
		} else if code == APP_ERRCODE1008 {
			res.Errcode = APP_ERRCODE1008
			res.Errmsg = APP_ERRMSG1008
		}

	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

//找回密码
func (c *AppController) Email() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP Register Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP Register Post")
		req_json := string(c.Ctx.Input.RequestBody)
		email := gojson.Json(req_json).Get("email").Tostring()
		if len(email) > 0 {
			code, obj, err := models.EmailUser(email)
			if err != nil {
				beego.Error(err)
			}
			if code == 0 {
				to_email := obj.Email
				to_user := obj.Username
				subject := "精酿啤酒重置密码"
				text := "精酿笔记重置密码成功，新密码为:" + obj.Password
				html := `
				<html>
				<body>
				<h3>
				[TEXT]
				</h3>
				</body>
				</html>
				`
				html = strings.Replace(html, "[TEXT]", text, -1)
				beego.Debug("send email html:", html)
				m := gomail.NewMessage()
				m.SetAddressHeader("From", "13671172337@163.com", "精酿笔记") // 发件人
				m.SetHeader("To",                                         // 收件人
					m.FormatAddress(to_email, to_user),
				)
				m.SetHeader("Subject", subject) // 主题
				m.SetBody("text/html", html)    // 正文

				d := gomail.NewPlainDialer("smtp.163.com", 465, "13671172337@163.com", "s123456") // 发送邮件服务器、端口、发件人账号、发件人密码
				if err := d.DialAndSend(m); err != nil {
					beego.Error(err)
				}
				res.Errcode = 0
				res.Errmsg = "忘记密码成功"
				data_json, err := json.Marshal(obj)
				if err != nil {
					beego.Error(err)
				}
				res.Data = string(data_json)
			} else if code == APP_ERRCODE1003 {
				res.Errcode = APP_ERRCODE1003
				res.Errmsg = APP_ERRMSG1003
			} else if code == APP_ERRCODE1009 {
				res.Errcode = APP_ERRCODE1009
				res.Errmsg = APP_ERRMSG1009
			}

		}
	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

//上传token
func (c *AppController) Uploadtoken() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP UploadToken Get")
	}
	var res = models.ResponseJson{}
	res.Errcode = 0
	res.Errmsg = "上传token成功"
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP UploadToken Post")
		intoken := c.Input().Get("token")
		ostype := c.Input().Get("ostype")
		beego.Debug("token:", intoken, "ostype:", ostype)
		// req_json := string(c.Ctx.Input.RequestBody)
		// var token = AppTokenModels{}
		// err := json.Unmarshal([]byte(req_json), &token)
		// if err != nil {
		// 	beego.Error(err)
		// 	res.Errcode = APP_ERRCODE1001
		// 	res.Errmsg = APP_ERRMSG1001
		// }
		// beego.Debug(token)
		err := models.SaveToken(intoken, ostype)
		if err != nil {
			beego.Error(err)
			res.Errcode = APP_ERRCODE1001
			res.Errmsg = APP_ERRMSG1001
		}
	}

	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

//相关资料
func (c *AppController) Related() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP Related Get")
	}
	var res = models.ResponseJson{}
	res.Errcode = 0
	res.Errmsg = "获取相关资料成功"
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP Related Post")
	}

	objs, err := models.GetAllRelatedState(1)
	if err != nil {
		beego.Error(err)
		res.Errcode = APP_ERRCODE1002
		res.Errmsg = APP_ERRMSG1002
	} else {
		data_json, err := json.Marshal(objs)
		if err != nil {
			beego.Error(err)
		}
		res.Data = string(data_json)
	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}
