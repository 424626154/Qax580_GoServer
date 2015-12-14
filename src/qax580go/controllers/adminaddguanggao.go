package controllers

/*
添加广告
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

type AdminaAddGuanggaoController struct {
	beego.Controller
}

func (c *AdminaAddGuanggaoController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["Image"] = ""
	c.TplNames = "adminaddguanggao.html"

}
func (c *AdminaAddGuanggaoController) Post() {
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	blink := c.Input().Get("blink")
	link := c.Input().Get("link")
	if len(title) != 0 && len(info) != 0 {
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
		b_link := false
		s_link := ""
		if blink == "true" {
			b_link = true
			s_link = link
		}
		_, err = models.AddGuanggao(title, info, image_name, b_link, s_link)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/guanggaos", 302)
		return
	}
	c.TplNames = "adminaddguanggao.html"

}
