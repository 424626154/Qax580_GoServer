package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"niangxin/models"
	"strings"
)

type NiangXinController struct {
	beego.Controller
}

func (c *NiangXinController) App() {
	responseJson := ResponseJson{}
	responseJson.ErrCode = ErrCode1001
	responseJson.ErrMsg = ErrCode1001_Msg
	if c.Ctx.Input.IsGet() {
		beego.Debug("App Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("App Post")
		var requestJson RequestJson
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestJson); err == nil {
			op := requestJson.Op
			beego.Debug("requestJson:", requestJson)
			switch op {
			case "register":
				username := requestJson.Username
				password := requestJson.Password
				beego.Debug("username :", username, "password :", password)
				if len(username) != 0 && len(password) != 0 {
					obj, err := models.GetUser(username)
					if err != nil {
						beego.Error(err)
					} else {
						if obj != nil && len(obj.Username) != 0 {
							responseJson.ErrCode = ErrCode1003
							responseJson.ErrMsg = ErrCode1003_Msg
						} else {
							token, err := models.AddUser(username, password)
							if err != nil {
								beego.Error(err)
							} else {
								responseJson.ErrCode = ErrCode0
								responseJson.Data = token
							}
						}
					}
				} else {
					responseJson.ErrCode = ErrCode1002
					responseJson.ErrMsg = ErrCode1002_Msg
				}
				break
			case "login":
				username := requestJson.Username
				password := requestJson.Password
				beego.Debug("username :", username, "password :", password)
				if len(username) != 0 && len(password) != 0 {
					obj, err := models.GetUser(username)
					if err != nil {
						beego.Error(err)
					} else {
						if obj != nil {
							if strings.EqualFold(obj.Password, password) {
								responseJson.ErrCode = ErrCode0
								responseJson.Data = obj.Token
							} else {
								responseJson.ErrCode = ErrCode1005
								responseJson.ErrMsg = ErrCode1005_Msg
							}
						} else {
							responseJson.ErrCode = ErrCode1004
							responseJson.ErrMsg = ErrCode1004_Msg
						}
					}
				} else {
					responseJson.ErrCode = ErrCode1002
					responseJson.ErrMsg = ErrCode1002_Msg
				}
				break
			case "userinfo":
				token := requestJson.Token
				if len(token) != 0 {
					obj, err := models.GetUserInfo(token)
					if err != nil {
						beego.Error(token)
					} else {
						user := UserData{}
						user.Token = obj.Token
						user.UserId = obj.UserId
						user.Username = obj.Username
						body, err := json.Marshal(user)
						if err != nil {
							beego.Error(err)
						} else {
							user_json := string(body)
							responseJson.ErrCode = ErrCode0
							responseJson.Data = user_json
						}
					}
				} else {
					responseJson.ErrCode = ErrCode1006
					responseJson.ErrMsg = ErrCode1006_Msg
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
						responseJson.ErrCode = ErrCode0
						responseJson.Data = string(objs_json)
					}
				}
				break
			}
		} else {
			beego.Error(err)
		}
	}

	body, err := json.Marshal(responseJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	beego.Debug("response_json :", response_json)
	c.Ctx.WriteString(response_json)
}

type ResponseJson struct {
	Data    string `json:"data"`
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type RequestJson struct {
	Op       string `json:"op"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	UserId   string `json:"userid"`
}

const ErrCode0 = 0
const ErrCode1001 = 1001
const ErrCode1001_Msg = "未知错误"
const ErrCode1002 = 1002
const ErrCode1002_Msg = "用户名或密码错误"
const ErrCode1003 = 1003
const ErrCode1003_Msg = "用户名已存在"
const ErrCode1004 = 1004
const ErrCode1004_Msg = "用户名不存在"
const ErrCode1005 = 1005
const ErrCode1005_Msg = "密码错误"
const ErrCode1006 = 1006
const ErrCode1006_Msg = "TOKEN失效"
