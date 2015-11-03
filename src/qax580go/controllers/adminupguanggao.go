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

type AdminUpGuanggaoController struct {
	beego.Controller
}

func (c *AdminUpGuanggaoController) Get() {
	id := c.Input().Get("id")
	// beego.Debug(id)
	guangao, err := models.GetOneGuanggao(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggao"] = guangao
	c.TplNames = "adminupguanggao.html"
}
func (c *AdminUpGuanggaoController) Post() {
	image_name := ""
	c.Data["Image"] = image_name
	op := c.Input().Get("op")
	id := c.Input().Get("id")
	switch op {
	case "upimg": //上传图片

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
			err := models.UpdateGuanggaoImg(id, image_name)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/guanggaos", 302)
				return
			}
		}
		guangao, err := models.GetOneGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Guanggao"] = guangao
	case "add":
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		blink := c.Input().Get("blink")
		link := c.Input().Get("link")
		if len(title) != 0 && len(info) != 0 {
			b_link := false
			s_link := ""
			if blink == "true" {
				b_link = true
				s_link = link
			}

			err := models.UpdateGuanggaoInfo(id, title, info, b_link, s_link)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/guanggaos", 302)
			return
		}
	}
	c.TplNames = "adminupguanggao.html"
}
