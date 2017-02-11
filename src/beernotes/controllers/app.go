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

//上传配方
func (c *AppController) AddFormula() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP AddFormula Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP AddFormula Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		fid := gojson.Json(req_json).Get("fid").Tostring()
		fname := gojson.Json(req_json).Get("fname").Tostring()
		malts := gojson.Json(req_json).Get("malts").Tostring()
		hopss := gojson.Json(req_json).Get("hopss").Tostring()
		yeasts := gojson.Json(req_json).Get("yeasts").Tostring()
		water := gojson.Json(req_json).Get("water").Tostring()
		accessoriess := gojson.Json(req_json).Get("accessoriess").Tostring()
		if len(token) != 0 && len(fid) != 0 && len(fname) != 0 && len(malts) != 0 && len(hopss) != 0 && len(yeasts) != 0 && len(water) != 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				code, obj, err := models.AddFormula(user.Uid, user.Username, fid, fname, malts, hopss, yeasts, water, accessoriess)
				if err != nil {
					beego.Error(err)
				}
				if code == 0 {
					res.Errcode = 0
					res.Errmsg = "上传配方成功"
					data_json, err := json.Marshal(obj)
					if err != nil {
						beego.Error(err)
					}
					res.Data = string(data_json)
				} else if code == APP_ERRCODE1003 {
					res.Errcode = APP_ERRCODE1003
					res.Errmsg = APP_ERRMSG1003
				} else if code == APP_ERRCODE1010 {
					res.Errcode = APP_ERRCODE1010
					res.Errmsg = APP_ERRMSG1010
				}
			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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

//删除配方
func (c *AppController) DelFormula() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP AddFormula Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP AddFormula Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		fid := gojson.Json(req_json).Get("fid").Tostring()
		if len(token) != 0 && len(fid) != 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				err := models.DeleteFormula(fid, user.Uid)
				if err != nil {
					beego.Error(err)
				}
				res.Errcode = 0
				res.Errmsg = "删除配方成功"

			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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

func (c *AppController) GetFormula() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP GetFormula Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		beego.Debug("APP GetFormula Post")
		beego.Debug("token:", token)
		objs, err := models.GetFormula()
		if len(objs) > 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			for i := 0; i < len(objs); i++ {
				num, err := models.GetFormulaLoveNum(objs[i].Fid, 0)
				if err != nil {
					beego.Error(err)
				}
				objs[i].Lovenum = int64(num)
				// beego.Debug("lovenum:", objs[i].Lovenum)
				if user != nil {
					islove, err := models.GetFormulaLove(user.Uid, objs[i].Fid, 0)
					if err != nil {
						beego.Error(err)
					}
					objs[i].Islove = islove
					// beego.Debug("islove:", objs[i].Islove)
					if user.Uid == objs[i].Uid {
						objs[i].Isdele = true
					}
					beego.Debug("Isdele:", objs[i].Isdele, user.Uid, objs[i].Uid)
				}
			}
		}
		if err != nil {
			beego.Error(err)
			res.Errcode = APP_ERRCODE1012
			res.Errmsg = APP_ERRMSG1012
		} else {
			res.Errcode = 0
			res.Errmsg = "获取配方成功"
			data_json, err := json.Marshal(objs)
			if err != nil {
				beego.Error(err)
			}
			res.Data = string(data_json)
		}

	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString(string(res_json))
}

func (c *AppController) CommentFormula() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP CommentFormula Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP CommentFormula Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		fid := gojson.Json(req_json).Get("fid").Tostring()
		fcid := gojson.Json(req_json).Get("fcid").Tostring()
		comment := gojson.Json(req_json).Get("comment").Tostring()
		beego.Debug("req_json:", req_json)
		if len(token) != 0 && len(fid) != 0 && len(fcid) != 0 && len(comment) != 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				obj, err := models.AddFormulaComment(fid, user.Uid, user.Username, comment, fcid)
				if err != nil {
					beego.Error(err)
					res.Errcode = APP_ERRCODE1003
					res.Errmsg = APP_ERRMSG1003
				} else {
					res.Errcode = 0
					res.Errmsg = "评论成功"
					data_json, err := json.Marshal(obj)
					if err != nil {
						beego.Error(err)
					}
					res.Data = string(data_json)
				}
			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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

func (c *AppController) DelFComment() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP GetFormulaComment Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP GetFormulaComment Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		id := gojson.Json(req_json).Get("id").Tostring()
		if len(token) != 0 && len(id) != 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				err := models.DeleteFormulaComment(id)
				if err != nil {
					beego.Error(err)
					res.Errcode = APP_ERRCODE1003
					res.Errmsg = APP_ERRMSG1003
				} else {
					res.Errcode = 0
					res.Errmsg = "删除评论成功"
				}
			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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

func (c *AppController) GetFormulaComment() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP GetFormulaComment Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP GetFormulaComment Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		objs, err := models.GetFormulaComment()
		if err != nil {
			beego.Error(err)
			res.Errcode = APP_ERRCODE1013
			res.Errmsg = APP_ERRMSG1013
		} else {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if len(objs) > 0 {
				for i := 0; i < len(objs); i++ {
					num, err := models.GetFormulaLoveNum(objs[i].Fid, objs[i].Id)
					if err != nil {
						beego.Error(err)
					}
					objs[i].Lovenum = int64(num)
					// beego.Debug("lovenum:", objs[i].Lovenum)
					if user != nil {
						islove, err := models.GetFormulaLove(user.Uid, objs[i].Fid, objs[i].Id)
						if err != nil {
							beego.Error(err)
						}
						objs[i].Islove = islove
						// beego.Debug("islove:", objs[i].Islove)
						if user.Uid == objs[i].Uid {
							objs[i].Isdele = true
						}
						beego.Debug("Isdele:", objs[i].Isdele)
					}
				}
			}

			res.Errcode = 0
			res.Errmsg = "获取配方评论成功"
			data_json, err := json.Marshal(objs)
			if err != nil {
				beego.Error(err)
			}
			res.Data = string(data_json)
		}

	}
	res_json, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("LoveFormula res_json:", string(res_json))
	c.Ctx.WriteString(string(res_json))
}

func (c *AppController) LoveFormula() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP LoveFormula Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP LoveFormula Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		fid := gojson.Json(req_json).Get("fid").Tostring()
		fcid := gojson.Json(req_json).Get("fcid").Tostring()
		islove := gojson.Json(req_json).Get("islove").Tostring()
		if len(token) != 0 && len(fid) != 0 && len(fcid) != 0 {
			beego.Debug("islove:", islove)
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				love := false
				if islove == "true" {
					love = true
				}
				obj, err := models.AddFormulaLove(user.Uid, fid, fcid, love)
				if err != nil {
					beego.Error(err)
					res.Errcode = APP_ERRCODE1003
					res.Errmsg = APP_ERRMSG1003
				} else {
					res.Errcode = 0
					res.Errmsg = "点赞成功"
					data_json, err := json.Marshal(obj)
					if err != nil {
						beego.Error(err)
					}
					res.Data = string(data_json)
				}
			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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
	// beego.Debug("LoveFormula res_json:", string(res_json))
	c.Ctx.WriteString(string(res_json))
}

func (c *AppController) GetMessage() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP GetMessage Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP GetMessage Post")
		req_json := string(c.Ctx.Input.RequestBody)
		token := gojson.Json(req_json).Get("token").Tostring()
		if len(token) != 0 {
			user, err := models.GetOneUserToken(token)
			if err != nil {
				beego.Error(err)
			}
			if user != nil {
				objs, err := models.GetReadMessage(user.Uid)
				if err != nil {
					beego.Error(err)
				}
				if len(objs) > 0 {
					for i := 0; i < len(objs); i++ {
						_, err := models.UpMessage(objs[i].Id, true)
						if err != nil {
							beego.Error(err)
						}
					}
				}
				res.Errcode = 0
				res.Errmsg = "获取消息成功"
				data_json, err := json.Marshal(objs)
				if err != nil {
					beego.Error(err)
				}
				res.Data = string(data_json)
			} else {
				res.Errcode = APP_ERRCODE1011
				res.Errmsg = APP_ERRMSG1011
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
	beego.Debug("GetMessage res_json:", string(res_json))
	c.Ctx.WriteString(string(res_json))
}

func (c *AppController) WNumber() {
	var res = models.ResponseJson{}
	if c.Ctx.Input.IsGet() {
		beego.Debug("APP WNumber Get")
		res.Errcode = APP_ERRCODE1003
		res.Errmsg = APP_ERRMSG1003
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("APP WNumber Post")
		// req_json := string(c.Ctx.Input.RequestBody)
		// token := gojson.Json(req_json).Get("token").Tostring()
		objs, err := models.GetWNumberState(1)
		if err != nil {
			beego.Error(err)
		}
		res.Errcode = 0
		res.Errmsg = "获取公众号列表成功"
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
	beego.Debug("GetMessage res_json:", string(res_json))
	c.Ctx.WriteString(string(res_json))
}
