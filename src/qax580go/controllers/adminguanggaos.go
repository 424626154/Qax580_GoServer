package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminGuanggaosController struct {
	beego.Controller
}

func (c *AdminGuanggaosController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateGuanggao(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateGuanggao(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	case "up":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		url := fmt.Sprintf("/admin/upguanggao?id=%s", id)
		beego.Debug("up_rul", url)
		c.Redirect(url, 302)
		return
	}

	guanggaos, err := models.GetAllGuanggaos()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "adminguanggaos.html"
	c.Data["Guanggaos"] = guanggaos
	// beego.Debug(guanggaos)
}
