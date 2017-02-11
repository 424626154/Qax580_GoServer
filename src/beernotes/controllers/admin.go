package controllers

import (
	"beernotes/models"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Home() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
		op := c.Input().Get("op")
		switch op {
		case "back":
			c.Ctx.SetCookie(BN_USERNAME, "", -1, "/")
			c.Ctx.SetCookie(BN_PASSWORD, "", -1, "/")
			c.Redirect("/admin/login", 302)
			return
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	c.TplName = "ahome.html"
}

func (c *AdminController) Login() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Login Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Login Post")
		username := c.Input().Get("user")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		beego.Debug("username:", username)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneAdmin(username)
			beego.Debug("admin:", admin)
			if err != nil {
				c.Redirect("/admin/login", 302)
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
				c.Redirect("/admin/home", 302)
				return
			} else {
				c.Redirect("/admin/home", 302)
				return
			}
		} else {
			c.Redirect("/admin/home", 302)
			return
		}
	}
	c.TplName = "alogin.html"
}

//相关资料
func (c *AdminController) Related() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		op := c.Input().Get("op")
		id := c.Input().Get("id")
		beego.Debug("Related id:", id, "op:", op)
		switch op {
		case "state0":
			err := models.UpdateRelatedState(id, 1)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/related", 302)
				return
			}
			return
		case "state1":
			err := models.UpdateRelatedState(id, 0)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/related", 302)
				return
			}
			return
		case "del":
			err := models.DeleteRelated(id)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/related", 302)
				return
			}
			return
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Related Post")
	}
	objs, err := models.GetAllRelated()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(objs)
	c.Data["Objes"] = objs
	c.TplName = "arelated.html"
}

func (c *AdminController) AddRelated() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("AddRelated Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AddRelated Post")
		name := c.Input().Get("name")
		brief := c.Input().Get("brief")
		link := c.Input().Get("link")
		if len(name) != 0 && len(brief) != 0 && len(link) != 0 {
			err := models.AddRelated(name, brief, link)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/related", 302)
				return
			}
		}
	}

	c.TplName = "aaddrelated.html"
}

func (c *AdminController) UpRelated() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("AddRelated Get")
		id := c.Input().Get("pid")
		obj, err := models.GetOneRelated(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Obj"] = obj
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AddRelated Post")
		id := c.Input().Get("id")
		name := c.Input().Get("name")
		brief := c.Input().Get("brief")
		link := c.Input().Get("link")
		if len(name) != 0 && len(brief) != 0 && len(link) != 0 {
			err := models.UpdateRelated(id, name, brief, link)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/related", 302)
				return
			}
		}
	}

	c.TplName = "auprelated.html"
}

func (c *AdminController) AppPush() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("AppPush Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AppPush Post")
	}
	objs, err := models.GetAllPush()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("objs:", objs)
	c.Data["Objs"] = objs
	c.TplName = "aapppush.html"
}

func (c *AdminController) SendPush() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("SendPush Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("SendPush Post")
		title := c.Input().Get("title")
		content := c.Input().Get("content")
		if len(title) != 0 && len(content) != 0 {
			android_ostype := "android"
			android_tokens := getTokensStr(android_ostype)
			beego.Debug("android_tokens:", android_tokens)
			ios_ostype := "ios"
			ios_tokens := getTokensStr(ios_ostype)
			beego.Debug("ios_tokens:", ios_tokens)
			if len(android_tokens) > 0 {
				sendPush(title, content, android_tokens, android_ostype)
			}
			if len(ios_tokens) > 0 {
				sendPush(title, content, ios_tokens, ios_ostype)
			}
			users, err := models.GetAllUser()
			if err != nil {
				beego.Error(err)
			}
			if len(users) > 0 {
				for i := 0; i < len(users); i++ {
					_, err = models.AddMessage(users[i].Uid, title, content)
					if err != nil {
						beego.Error(err)
					}
				}
			}
			c.Redirect("/admin/apppush", 302)
			return
		}
	}
	c.TplName = "asendpush.html"
}

func (c *AdminController) Admin() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin Post")
	}
	c.TplName = "aadmin.html"
}

func chackAccount(ctx *context.Context) (bool, string) {
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

	admin, err := models.GetOneAdmin(username)
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

type UMPushModels struct {
	Appkey         string          `json:"appkey"`
	Timestamp      int64           `json:"timestamp"`
	Pushtype       string          `json:"type"`
	DeviceTokens   string          `json:"device_tokens,omitempty"`
	ProductionMode string          `json:"production_mode"`
	Description    string          `json:"description"`
	Payload        UMPayloadModels `json:"payload,omitempty"`
	Body           UMBodyModels    `json:"body,omitempty"`
}

type UMPayloadModels struct {
	DisplayType string      `json:"display_type,omitempty"`
	Aps         UMApsModels `json:"aps,omitempty"`
}

type UMBodyModels struct {
	Ticker    string `json:"ticker,omitempty"`
	Title     string `json:"title,omitempty"`
	Text      string `json:"text,omitempty"`
	AfterOpen string `json:"after_open,omitempty"`
}

type UMApsModels struct {
	Alert string `json:"alert,omitempty"`
}

type ResPushDataModels struct {
	TackId    string `json:"task_id,omitempty"`
	MsgId     string `json:"msg_id,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}

type ResPushModels struct {
	Ret  string            `json:"ret,omitempty"`
	Data ResPushDataModels `json:"data,omitempty"`
}

func sendPush(title string, content string, tokens string, ostype string) {
	appkey := ""
	app_master_secret := ""
	timestamp := time.Now().Unix()
	pushtype := "listcast"
	production_mode := "false"
	description := ""
	push_str := ""
	method := "POST"
	// device_token := ""
	url := "http://msg.umeng.com/api/send"
	if ostype == "ios" {
		appkey = "5846ef9fa325116db20017b3"
		app_master_secret = "5yzwppsoszclwakverdoiihpc8dhbrme"
		description = "测试消息-IOS"
		var aps = UMApsModels{
			Alert: content,
		}
		var payload = UMPayloadModels{
			Aps: aps,
		}
		push := UMPushModels{
			Appkey:         appkey,
			Timestamp:      timestamp,
			Pushtype:       pushtype,
			DeviceTokens:   tokens,
			ProductionMode: production_mode,
			Description:    description,
			Payload:        payload,
		}
		push_json, err := json.Marshal(push)
		if err != nil {
			beego.Error(err)
		}
		push_str = string(push_json)
	} else if ostype == "android" {
		appkey = "5846ec8eae1bf8125b000d7d"
		app_master_secret = "vxutm65kn7st3kdnlimltlkyvv1xssgw"
		description = "测试消息-Android"
		var payload = UMPayloadModels{
			DisplayType: "notification",
		}
		var body = UMBodyModels{
			Ticker:    title,
			Title:     title,
			Text:      content,
			AfterOpen: "go_app",
		}
		push := UMPushModels{
			Appkey:         appkey,
			Timestamp:      timestamp,
			Pushtype:       pushtype,
			DeviceTokens:   tokens,
			ProductionMode: production_mode,
			Description:    description,
			Payload:        payload,
			Body:           body,
		}
		push_json, err := json.Marshal(push)
		if err != nil {
			beego.Error(err)
		}
		push_str = string(push_json)
	}
	// beego.Debug("push_str:", push_str)
	// beego.Debug("app_master_secret", app_master_secret)
	params_str := push_str
	var base_str = method + url + params_str + app_master_secret
	// beego.Debug("base_str", base_str)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(base_str))
	cipherStr := md5Ctx.Sum(nil)
	md5_str := hex.EncodeToString(cipherStr)
	sign := md5_str
	url = url + "?sign=" + sign
	beego.Debug("url:", url)
	post_body := bytes.NewBuffer([]byte(params_str))
	resp, err := http.Post(url, "application/json;charset=utf-8", post_body)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var respush = ResPushModels{}
	err = json.Unmarshal([]byte(string(body)), &respush)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("respush:", respush)
	if respush.Ret == "SUCCESS" {

	}
	err = models.AddPush(title, content, tokens, ostype)
	if err != nil {
		beego.Error(err)
	}
}

func getTokensStr(ostype string) string {
	// ios_ostype := "ios"
	objs, err := models.GetOstypeAllTokesn(ostype)
	if err != nil {
		beego.Error(err)
	}
	tokens := ""
	if len(objs) > 0 {
		for i := 0; i < len(objs); i++ {
			if i == len(objs)-1 {
				tokens += objs[i].Token
			} else {
				tokens += objs[i].Token + ","
			}
		}
	}
	// beego.Debug("tokens:", tokens)
	return tokens
}

func (c *AdminController) WNumber() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("WechatNumber Get")
		op := c.Input().Get("op")
		id := c.Input().Get("id")
		switch op {
		case "state0":
			err := models.UpdateWNumberState(id, 1)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/wnumber", 302)
				return
			}
			return
		case "state1":
			err := models.UpdateWNumberState(id, 0)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/wnumber", 302)
				return
			}
			return
		case "del":
			err := models.DeleteWNumber(id)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/wnumber", 302)
				return
			}
			return
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("WechatNumber Post")
	}
	objs, err := models.GetAllWNumber()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(objs)
	c.Data["Objes"] = objs
	c.TplName = "awnumber.html"
}

func (c *AdminController) AddWNumber() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin AddWNumber Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin AddWNumber Post")
		image_name := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		number := c.Input().Get("number")
		if len(title) != 0 && len(info) != 0 && len(number) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("title:", title, "info:", info, "number:", number, "image_name:", image_name)
			_, err = models.AddWNumber(title, info, number, image_name)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wnumber", 302)
		} else {
			c.Redirect("/admin/addwnumber", 302)
		}
	}
	c.TplName = "aaddwnumber.html"
}

func (c *AdminController) UpWnumberInfo() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin UpWnumberInfo Get")
		pid := c.Input().Get("pid")
		if len(pid) != 0 {
			obj, err := models.GetWNumber(pid)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Wxnum"] = obj
		} else {
			c.Redirect("/admin/wnumber", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin UpWnumberInfo Post")
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		number := c.Input().Get("number")
		beego.Debug(len(id) != 0, len(title) != 0, len(info) != 0, len(number) != 0)
		if len(id) != 0 && len(title) != 0 && len(info) != 0 && len(number) != 0 {
			beego.Debug(id)
			_, err := models.UpdateWNumberInfo(id, title, info, number)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wnumber", 302)
			return
		}
	}
	c.TplName = "aupwnumberinfo.html"
}
func (c *AdminController) UpWnumberImg() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin UpWnumberImg Get")
		pid := c.Input().Get("pid")
		if len(pid) != 0 {
			obj, err := models.GetWNumber(pid)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Wxnum"] = obj
		} else {
			c.Redirect("/admin/wnumber", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin UpWnumberImg Post")
		id := c.Input().Get("id")
		image_name := ""
		if len(id) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			if image_name != "" {
				_, err := models.UpdateWNumberImg(id, image_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/admin/wnumber", 302)
					return
				}
			}
		}
	}
	c.TplName = "aupwnumberimg.html"
}

func (c *AdminController) UpdateLog() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin/login", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin UpdateLog Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin UpdateLog Post")
	}
	c.TplName = "aupdatelog.html"
}
