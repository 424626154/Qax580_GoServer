package controllers

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

type AdminUplodeController struct {
	beego.Controller
}

func (c *AdminUplodeController) Get() {

	c.TplNames = "adminuplode.html"
}

func (c *AdminUplodeController) Post() {
	image_name := ""
	beego.Debug("AdminUplodeController-------------------1")
	title := c.Input().Get("title")
	info := c.Input().Get("info")
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
		c.Redirect("/admin/uplode", 302)
	}

	if len(title) != 0 && len(info) != 0 {
		beego.Debug("AdminUplodeController-------------------2")
		beego.Debug(image_name)
		err := models.AddPostLabel(title, info, 2, image_name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin", 302)
	}
}
