package controllers

/*
发布消息
*/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"qax580go/models"
	"strings"
	"time"
)

type UplodeController struct {
	beego.Controller
}

func (c *UplodeController) Get() {
	getUplodeCookie(c)
	c.Data["FromType"] = getUplodeFromType(c)
	c.TplName = "uplode.html"
}

func (c *UplodeController) Post() {
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	city := c.Input().Get("city")
	beego.Debug("city", city)
	if len(title) != 0 && len(info) != 0 {
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
	} else {
		c.Redirect("/uplode", 302)
	}
	if len(title) != 0 && len(info) != 0 {
		openid := getUplodeCookie(c)
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("----------AddPostLabelWx--------")
		beego.Debug(openid)
		beego.Debug(wxuser)
		err = models.AddPostLabelWx(title, info, 1, image_name, openid, wxuser.NickeName, wxuser.Sex, wxuser.HeadImgurl, city)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("inset ok")
		c.Redirect("/mymessage", 302)
	}
}

func getUplodeCookie(c *UplodeController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}

/**
*来源类型
 */
func getUplodeFromType(c *UplodeController) string {
	from_type := c.Ctx.GetCookie(COOKIE_FROM_TYPE)
	if len(from_type) == 0 {
		from_type = COOKIE_FROM_ALL
	}
	return from_type
}
