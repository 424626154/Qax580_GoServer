package controllers

/*
主页
*/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"qax580go/models"
	"strconv"
	"strings"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	getCookie(c)
	setUrl(c)
	city := getSelectCity(c)
	CurrentPage := int32(1)
	count, err := models.GetPostCount()
	if strings.Contains(city, CITY_ALL) {
	} else {
		count, err = models.GetCityPostCount(city)
	}
	NumberofPages := int32(10)
	temp := count / NumberofPages
	if (count % NumberofPages) != 0 {
		temp = temp + 1
	}
	CotalPages := temp
	pagetype := c.Input().Get("type")
	page := c.Input().Get("page")
	// beego.Debug("pagetype:", pagetype)

	guanggaos, err := models.GetAllGuanggaosState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["City"] = city
	c.Data["Guanggaos"] = guanggaos

	if len(pagetype) != 0 && len(page) != 0 {
		switch pagetype {
		case "first": //首页
			CurrentPage = 1
		case "prev": //上一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "next": //下一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "last": //尾页
			CurrentPage = CotalPages
		case "page": //页码
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		}
	}
	c.Data["CurrentPage"] = CurrentPage
	c.Data["CotalPages"] = CotalPages
	c.Data["NumberofPages"] = NumberofPages
	if strings.Contains(city, CITY_ALL) {
		posts, err := models.QueryPagePost(CurrentPage-1, NumberofPages)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Posts"] = posts
	} else {
		posts, err := models.QueryCityPagePost(CurrentPage-1, NumberofPages, city)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Posts"] = posts
	}

	// beego.Debug(posts)
	c.TplNames = "home.html"

	isdebug := "true"
	iscanting := "false"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		iscanting = iniconf.String("qax580::iscanting")
	}
	// beego.Debug(isdebug)
	beego.Debug("IsCanting", iscanting)
	c.Data["IsDebug"] = isdebug
	c.Data["IsCanting"] = iscanting
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeletePost(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is del " + id)
		c.Redirect("/", 302)
		return
	}
}

func (c *HomeController) Post() {
	op := c.Input().Get("op")
	switch op {
	case "city":
		city := c.Input().Get("city")
		maxAge := 1<<31 - 1
		// // beego.Debug(COOKIE_CITY, city)
		// citys := [12]City{City{"庆安县", "qinganxian"}, City{"绥化市", "suihuashi"}, City{"北林市", "beilinshi"}, City{"安达市", "andashi"},
		// 	City{"肇东市", "zhaodongshi"}, City{"海伦市", "hailunshu"}, City{"望奎市", "wangkuishi"}, City{"兰西县", "lanxixian"},
		// 	City{"青冈县", "qinggangxian"}, City{"明水县", "mingshuixian"}, City{"绥棱县", "suilingxian"}, City{CITY_ALL, PY_ALL}}
		citys := [12]City{City{CITY_QA, PY_QA}, City{CITY_SH, PY_SH}, City{CITY_BL, PY_BL}, City{CITY_AD, PY_AD},
			City{CITY_SD, PY_SD}, City{CITY_HL, PY_HL}, City{CITY_WK, PY_WK}, City{CITY_LX, PY_LX},
			City{CITY_QG, PY_QG}, City{CITY_MS, PY_MS}, City{CITY_SL, PY_SL}, City{CITY_ALL, PY_ALL}}
		city_name := ""
		for i := 0; i < len(citys); i++ {
			if citys[i].City == city {
				city_name = citys[i].Name
			}
		}
		if len(city_name) != 0 {
			c.Ctx.SetCookie(COOKIE_CITY, city_name, maxAge, "/")
		}
		c.Redirect("/", 302)
		return
	default:
		c.Redirect("/", 302)
		return
	}
}

func getCookie(c *HomeController) {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	// beego.Debug("------------openid--------")
	// beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			// beego.Debug("--------------wxuser----------")
			// beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
}

func setUrl(c *HomeController) {
	c.Data["ImgUrlPath"] = getImageUrl()
}

type City struct {
	City string
	Name string
}

func getSelectCity(c *HomeController) string {
	// citys := [12]City{City{"庆安县", "qinganxian"}, City{"绥化市", "suihuashi"}, City{"北林市", "beilinshi"}, City{"安达市", "andashi"},
	// 	City{"肇东市", "zhaodongshi"}, City{"海伦市", "hailunshu"}, City{"望奎市", "wangkuishi"}, City{"兰西县", "lanxixian"},
	// 	City{"青冈县", "qinggangxian"}, City{"明水县", "mingshuixian"}, City{"绥棱县", "suilingxian"}, City{CITY_ALL, PY_ALL}}
	citys := [12]City{City{CITY_QA, PY_QA}, City{CITY_SH, PY_SH}, City{CITY_BL, PY_BL}, City{CITY_AD, PY_AD},
		City{CITY_SD, PY_SD}, City{CITY_HL, PY_HL}, City{CITY_WK, PY_WK}, City{CITY_LX, PY_LX},
		City{CITY_QG, PY_QG}, City{CITY_MS, PY_MS}, City{CITY_SL, PY_SL}, City{CITY_ALL, PY_ALL}}
	city_default := CITY_ALL
	city := ""
	city_name := c.Ctx.GetCookie(COOKIE_CITY)
	for i := 0; i < len(citys); i++ {
		if citys[i].Name == city_name {
			city = citys[i].City
		}
	}
	if len(city) == 0 {
		city = city_default
	}
	// beego.Debug("getSelectCity", city)
	return city
}
