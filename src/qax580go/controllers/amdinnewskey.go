package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminNewsKeyController struct {
	beego.Controller
}

func (c *AdminNewsKeyController) Get() {

	op := c.Input().Get("op")
	switch op {
	case "add":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateNewsKey(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/newskey", 302)
		return
	case "remove":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateNewsKey(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/newskey", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteNewsKey(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/newskey", 302)
		return
	}
	var keys []string
	opnewskey, err := models.GetOpNewsKey(1)
	if len(opnewskey) != 0 {
		for i := 0; i < len(opnewskey); i++ {
			keys = append(keys, opnewskey[i].Info)
		}
		beego.Debug(keys, "-----------")
	}
	c.Data["NewsKey"] = keys

	newskey, err := models.GetAllNewsKey()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("newskey", newskey)
	c.Data["NewsKeys"] = newskey
	c.TplNames = "adminnewskey.html"
}

func (c *AdminNewsKeyController) Post() {
	key := c.Input().Get("key")
	if len(key) != 0 {
		beego.Debug("news key :", key)
		err := models.AddNewsKey(key)
		if err != nil {
			beego.Debug("addNewsKey error :", err)
		}
	}
	var keys []string
	opnewskey, err := models.GetOpNewsKey(1)
	if len(opnewskey) != 0 {
		for i := 0; i < len(opnewskey); i++ {
			keys = append(keys, opnewskey[i].Info)
		}
		beego.Debug(keys, "-----------")
	}
	c.Data["NewsKey"] = keys

	newskey, err := models.GetAllNewsKey()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("newskey", newskey)
	c.Data["NewsKeys"] = newskey
	c.TplNames = "adminnewskey.html"
}
