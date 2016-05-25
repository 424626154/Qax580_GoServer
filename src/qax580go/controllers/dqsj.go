package controllers

/**
大签世界
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"
	"time"
)

type DqsjController struct {
	beego.Controller
}

func (c *DqsjController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}

	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, "http://www.baoguangguang.cn/dqsj/home")
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = "大签世界火盆烤肉欢迎您的到来！"
	wxShareCon.Link = "http://www.baoguangguang.cn/dqsj/home"
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)
	c.TplName = "dqsjhome.html"
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
			if current_time-wxAttribute.AccessTokenTime < 7000 {
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
			if current_time-wxAttribute.TicketTime < 7000 {
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
