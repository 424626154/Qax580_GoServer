package controllers

/**
大签世界
*/
import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"net/http"
	"path"
	"qax580go/models"
	"strings"
	"time"
)

type DqsjController struct {
	beego.Controller
}

//后台
func (c *DqsjController) Admin() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin Post")
	}
	bool, username := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	c.Data["User"] = username
	c.Data["isUser"] = bool
	beego.Debug("username:", username)
	op := c.Input().Get("op")
	switch op {
	case "back":
		c.Ctx.SetCookie(DQSJ_USERNAME, "", -1, "/")
		c.Ctx.SetCookie(DQSJ_PASSWORD, "", -1, "/")
		c.Redirect("/dqsj/admin", 302)
		return
	}

	c.TplName = "dqsjadmin.html"
}

//后台登录
func (c *DqsjController) AdminLogin() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminLogin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminLogin Post")
		username := c.Input().Get("user")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		beego.Debug("AdminLogin Post user:", username, "password:", password)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneDqsjAdmin(username)
			if err != nil {
				c.Redirect("/dqsj/adminlogin", 302)
				return
			}
			if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
				maxAge := 0
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie(DQSJ_USERNAME, username, maxAge, "/")
				c.Ctx.SetCookie(DQSJ_PASSWORD, password, maxAge, "/")
				beego.Debug("login ok------")
				c.Redirect("/dqsj/admin", 302)
				return
			} else {
				c.Redirect("/dqsj/adminlogin", 302)
				return
			}
		} else {
			c.Redirect("/dqsj/adminlogin", 302)
			return
		}
	}
	c.TplName = "dqsjadminlogin.html"
}
func (c *DqsjController) AdminCai() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminCai Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminCai Post")
	}
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}

	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeleteAllCaiItem(id)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteCaiGroup(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "state0":
		id := c.Input().Get("id")
		err := models.UpdateCaiGroup(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		err := models.UpdateCaiGroup(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemdel":
		id := c.Input().Get("id")
		err := models.DeleteCaiGrItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemstate0":
		id := c.Input().Get("id")
		err := models.UpdateCaiItem(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemstate1":
		id := c.Input().Get("id")
		err := models.UpdateCaiItem(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsdel":
		id := c.Input().Get("id")
		err := models.DeleteCaiTips(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsstate0":
		id := c.Input().Get("id")
		err := models.UpdateCaiTips(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsstate1":
		id := c.Input().Get("id")
		err := models.UpdateCaiTips(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	}

	objs, err := models.GetAllCaiGroup()
	if err != nil {
		beego.Error(err)
	}
	var showCaiGroup []models.DqsjShowCaiGroup
	for i := 0; i < len(objs); i++ {
		objitem, err := models.GetAllCaiItem(objs[i].Id)
		if err != nil {
			beego.Error(err)
		} else {
			obgshow := models.DqsjShowCaiGroup{Id: objs[i].Id, Name: objs[i].Name,
				OrderId: objs[i].OrderId, State: objs[i].State, Time: objs[i].Time, CaiItems: objitem}
			showCaiGroup = append(showCaiGroup, obgshow)
		}

	}

	c.Data["ShowCaiGroup"] = showCaiGroup

	tips, err := models.GetAllCaiTips()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CaiTips"] = tips
	c.TplName = "dqsjadmincai.html"
}
func (c *DqsjController) AdminAddCaiGroup() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiGroup Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiGroup Post")
		name := c.Input().Get("name")
		orderid := c.Input().Get("orderid")
		if len(name) != 0 && len(orderid) != 0 {
			err := models.AddCaiGroup(name, orderid)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
		} else {
			c.Redirect("/dqsj/adminaddcaigroup", 302)
		}

	}
	c.TplName = "dqsjadminaddcaigroup.html"
}
func (c *DqsjController) AdminAddCaiItem() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiItem Get")
		groupid := c.Input().Get("groupid")
		c.Data["GroupId"] = groupid
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiItem Post")
		image_name := ""
		name := c.Input().Get("name")
		price := c.Input().Get("price")
		pricedesc := c.Input().Get("pricedesc")
		groupid := c.Input().Get("groupid")
		if len(name) != 0 && len(price) != 0 {
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
				// beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			err = models.AddCaiItem(name, image_name, groupid, price, pricedesc)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
			return
		}
	}
	c.TplName = "dqsjadminaddcaiitem.html"
}

func (c *DqsjController) AdminPan() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminPan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminPan Post")
	}
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}

	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeletePanItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	case "state0":
		id := c.Input().Get("id")
		err := models.UpdatePanItem(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		err := models.UpdatePanItem(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	}
	panitem, err := models.GetAllPanItem()
	if err != nil {
		beego.Error(err)
	}
	//计算总概率
	allProbability := int64(0)
	for i := 0; i < len(panitem); i++ {
		if panitem[i].State == 1 {
			allProbability += panitem[i].Probability
		}
	}
	for i := 0; i < len(panitem); i++ {
		panitem[i].AllProbability = allProbability
	}
	c.Data["PanItem"] = panitem
	c.TplName = "dqsjadminpan.html"
}
func (c *DqsjController) AdminAddPan() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddPanItem Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddPanItem Post")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		probability := c.Input().Get("probability")
		if len(name) != 0 && len(info) != 0 {
			err := models.AddPanItem(name, info, probability)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/adminpan", 302)
		} else {
			c.Redirect("/dqsj/adminaddpan", 302)
		}

	}
	c.TplName = "dqsjadminaddpan.html"
}
func (c *DqsjController) AdminAddCaiTips() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiTips Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiTips Post")
		info := c.Input().Get("info")
		if len(info) != 0 {
			err := models.AddCaiTips(info)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
		} else {
			c.Redirect("/dqsj/adminaddcaitips", 302)
		}

	}
	c.TplName = "dqsjadminaddcaitips.html"
}

func (c *DqsjController) AdminGg() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminGuangGao Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminGuangGao Post")
	}
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteDqsjGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjGuanggao(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjGuanggao(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	}
	guanggaos, err := models.GetAllDqsjGuanggaos()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggaos"] = guanggaos
	c.TplName = "dqsjadminguanggao.html"
}

func (c *DqsjController) AdminAddGg() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddGg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddGg Post")
		image_name := ""
		imageitem0 := ""
		imageitem1 := ""
		imageitem2 := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		blink := c.Input().Get("blink")
		link := c.Input().Get("link")
		bimg := c.Input().Get("bimg")
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
			//上传imageitem0
			_, fh, err = c.GetFile("imageitem0")
			beego.Debug("上传imageitem0:", fh)
			if err != nil {
				beego.Error(err)
			}
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d%s", t, "imageitem0")
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				imageitem0 = hex.EncodeToString(h.Sum(nil))
				beego.Info(imageitem0) // 输出加密结果
				err = c.SaveToFile("imageitem0", path.Join("imagehosting", imageitem0))
				if err != nil {
					beego.Error(err)
					imageitem0 = ""
				}
			}
			//上传imageitem1
			_, fh, err = c.GetFile("imageitem1")
			beego.Debug("上传imageitem1:", fh)
			if err != nil {
				beego.Error(err)
			}
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d%s", t, imageitem1)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				imageitem1 = hex.EncodeToString(h.Sum(nil))
				beego.Info(imageitem1) // 输出加密结果
				err = c.SaveToFile("imageitem1", path.Join("imagehosting", imageitem1))
				if err != nil {
					beego.Error(err)
					imageitem1 = ""
				}
			}
			//上传imageitem2
			_, fh, err = c.GetFile("imageitem2")
			beego.Debug("上传imageitem2:", fh)
			if err != nil {
				beego.Error(err)
			}
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d%s", t, "imageitem2")
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				imageitem2 = hex.EncodeToString(h.Sum(nil))
				beego.Info(imageitem2) // 输出加密结果
				err = c.SaveToFile("imageitem2", path.Join("imagehosting", imageitem2))
				if err != nil {
					beego.Error(err)
					imageitem2 = ""
				}
			}

			b_link := false
			s_link := ""
			if blink == "true" {
				b_link = true
				s_link = link
			}
			b_img := false
			if bimg == "true" {
				b_img = true
			}
			beego.Debug("info", info)
			_, err = models.AddDqsjGuanggao(title, info, image_name, b_link, s_link, b_img, imageitem0, imageitem1, imageitem2)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admingg", 302)
			return
		}
	}
	c.TplName = "dqsjadminaddgg.html"
}

func (c *DqsjController) AdminHuoDong() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminHuoDong Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminHuoDong Post")
	}
	op := c.Input().Get("op")
	switch op {
	case "uphuodong":
		huodong := c.Input().Get("huodong")
		if len(huodong) == 0 {
			break
		}
		err := models.ModifyDqsjHomeHD(huodong)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteDqsjHD(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjHD(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjHD(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "additem":
		content := c.Input().Get("content")
		if len(content) == 0 {
			break
		}
		err := models.AddDqsjHD(content)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	}

	obj, err := models.GetOneDqsjHome()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["DqsjHome"] = obj
	obj1, err := models.GetAllDqsjHD()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["DqsjHuoDong"] = obj1
	c.TplName = "dqsjadminhuodong.html"
}

//主页
func (c *DqsjController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/home"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/home"
	}

	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = "大签世界火盆烤肉欢迎您的到来！"
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	//广告栏
	c.Data["ImgUrlPath"] = getImageUrl()
	guanggaos, err := models.GetAllDqsjGuanggaosState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggaos"] = guanggaos

	obj, err := models.GetOneDqsjHome()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["DqsjHome"] = obj
	obj1, err := models.GetAllDqsjHDState1()
	if err != nil {
		beego.Debug(err)
	}
	if obj1 != nil {
		for i := 0; i < len(obj1); i++ {
			obj1[i].ShowId = int64(i + 1)
		}
	}
	c.Data["DqsjHuoDong"] = obj1

	c.TplName = "dqsjhome.html"
}

//菜单
func (c *DqsjController) Cai() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Cai Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Cai Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/cai"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/cai"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = "大签世界火盆烤肉欢迎您的到来！"
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	objs, err := models.GetAllCaiGroupState1()
	if err != nil {
		beego.Error(err)
	}
	var showCaiGroup []models.DqsjShowCaiGroup
	for i := 0; i < len(objs); i++ {
		objitem, err := models.GetAllCaiItemState1(objs[i].Id)
		if err != nil {
			beego.Error(err)
		} else {
			obgshow := models.DqsjShowCaiGroup{Id: objs[i].Id, Name: objs[i].Name,
				OrderId: objs[i].OrderId, State: objs[i].State, Time: objs[i].Time, CaiItems: objitem}
			showCaiGroup = append(showCaiGroup, obgshow)
		}

	}

	c.Data["ShowCaiGroup"] = showCaiGroup
	tips, err := models.GetAllCaiTipsState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CaiTips"] = tips
	c.TplName = "dqsjcai.html"
}

//幸运盘
func (c *DqsjController) Pan() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Pan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Pan Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/pan"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/pan"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = "大签世界火盆烤肉欢迎您的到来！"
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)
	panitem, err := models.GetAllPanItemState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["PanItem"] = panitem
	beego.Debug("panitem :", panitem)
	c.TplName = "dqsjpan.html"
}

func (c *DqsjController) GuangGao() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("GuangGao Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("GuangGao Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/guanggao"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/guanggao"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = "大签世界火盆烤肉欢迎您的到来！"
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		guangao, err := models.GetOneDqsjGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Guanggao"] = guangao
		beego.Debug("guangao :", guangao)
		c.TplName = "dqsjguanggao.html"
		return
	}

	c.TplName = "dqsjguanggao.html"
}

func getDqsjToken() string {
	//https://api.weixin.qq.com/cgi-bin/token?&appid=APPID&secret=APPSECRET
	wxAttribute, err := models.GetWxAttribute()
	if err != nil {
		beego.Debug(err)
	}
	if wxAttribute != nil {
		if len(wxAttribute.AccessToken) != 0 {
			current_time := time.Now().Unix()
			beego.Debug("current_time:", current_time, "wxAttribute.AccessTokenTime :", wxAttribute.AccessTokenTime, "current_time-wxAttribute.AccessTokenTime:", current_time-wxAttribute.AccessTokenTime)
			if current_time-wxAttribute.AccessTokenTime < 6000 {
				return wxAttribute.AccessToken
			}
		}
	}
	wx_url := "[REALM]?grant_type=client_credential&appid=[APPID]&secret=[SECRET]"
	realm_name := "https://api.weixin.qq.com/cgi-bin/token"
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("http_wx_token_url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug("http_wx_token_err :", err)
	} else {
		beego.Debug("http_wx_token_body :", string(body))
	}

	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("http_wx_token_json :", atj)
		if atj.ErrCode == 0 {
			_, err = models.AddWxAttributeToken(atj.AccessToken)
			if err != nil {
				beego.Debug(err)
			}
			return atj.AccessToken
		} else {
			return ""
		}
	} else {
		beego.Debug("http_wx_token_err :", err)
		return ""
	}
}

func getDqsjTicket(access_toke string) string {
	wxAttribute, err := models.GetWxAttribute()
	if err != nil {
		beego.Debug(err)
	}
	if wxAttribute != nil {
		if len(wxAttribute.Ticket) != 0 {
			current_time := time.Now().Unix()
			beego.Debug("current_time:", current_time, "wxAttribute.TicketTime :", wxAttribute.TicketTime, "current_time-wxAttribute.TicketTime:", current_time-wxAttribute.TicketTime)
			if current_time-wxAttribute.TicketTime < 6000 {
				return wxAttribute.Ticket
			}
		}
	}

	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	realm_name := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	beego.Debug("http_wx_ticket_url :", wx_url)
	resp, err := http.Get(wx_url)
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
	var ticket models.JsApiTicketJson
	if err := json.Unmarshal(body, &ticket); err == nil {
		beego.Debug("http_wx_ticket_ticketobj :", ticket)
		if ticket.ErrCode == 0 {
			_, err = models.AddWxAttributeTicket(ticket.Ticket)
			if err != nil {
				beego.Debug(err)
			}
			return ticket.Ticket
		}

		return ""
	} else {
		beego.Debug("http_wx_ticket_ticke :", err)
		return ""
	}
}

func chackDqsjAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie(DQSJ_USERNAME)
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie(DQSJ_PASSWORD)
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneDqsjAdmin(username)
	beego.Debug("GetOneDqsjAdmin admin:", admin)
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
