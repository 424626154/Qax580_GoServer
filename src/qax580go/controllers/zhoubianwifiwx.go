package controllers

/*
周边wifi
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

type Wifi struct {
	Reason    string  `json:"reason"`
	ErrorCode int64   `json:"error_code"`
	WResult   WResult `json:"result"`
}
type WResult struct {
	WFData []WFData `json:"data"`
}
type WFData struct {
	Name      string `json:"name"`
	Intro     string `json:"intro"`
	Address   string `json:"address"`
	GoogleLat string `json:"google_lat"`
	GoogleLon string `json:"google_lon"`
	BaiduLat  string `json:"baidu_lat"`
	BaiduLon  string `json:"baidu_lon"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Distance  int32  `json:"distance"`
}

type ZhouBianWifiWXController struct {
	beego.Controller
}

func (c *ZhouBianWifiWXController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "location":
		latitude := c.Input().Get("latitude")
		longitude := c.Input().Get("longitude")
		c.Data["latitude"] = latitude
		c.Data["longitude"] = longitude
		beego.Debug("latitude:", latitude)
		beego.Debug("longitude:", longitude)
		getWifi(longitude, latitude, c)
		c.TplNames = "zhoubianwifi.html"
		return
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
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	getWifiWxJsToken(noncestr, timestamp, c)
	c.TplNames = "zhoubianwifiwx.html"
	// getWifi("116.366324", "39.905859", c)
	// c.TplNames = "zhoubianwifi.html"
}

func (c *ZhouBianWifiWXController) Post() {
	c.TplNames = "zhoubianwifiwx.html"
}

func getWifiWxJsToken(noncestr string, timestamp int64, c *ZhouBianWifiWXController) {
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&grant_type=client_credential"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9093"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	}
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("----------------get Token --------------------")
	beego.Debug(wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------get Token body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("----------------get Token json--------------------")
		beego.Debug(atj)
		if atj.ErrCode == 0 {
			getWifiJsapiTicket(atj.AccessToken, noncestr, timestamp, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
}

func getWifiJsapiTicket(access_toke string, noncestr string, timestamp int64, c *ZhouBianWifiWXController) {
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9092"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	beego.Debug("getWifiJsapiTicket", wx_url)

	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------getJsapiTicket body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var ticket models.JsApiTicketJson
	if err := json.Unmarshal(body, &ticket); err == nil {
		beego.Debug("----------------getJsapiTicket json--------------------")
		beego.Debug(ticket)
		if ticket.ErrCode == 0 {
			c.Data["Ticket"] = signatureWxJs(ticket.Ticket, noncestr, timestamp, "http://www.baoguangguang.cn/zhoubianwifiwx")
		}

		return
	} else {
		beego.Debug("----------------getJsapiTicket error--------------------")
		beego.Debug(err)
	}
}

func getWifi(lon string, lat string, c *ZhouBianWifiWXController) {
	// http://apis.juhe.cn/wifi/local?key=appkey&lon=116.366324&lat=39.905859&r=3000&type=1
	url := "[REALM]?key=[KEY]&lon=[LON]&lat=[LAT]&r=3000&type=1"
	url = strings.Replace(url, "[REALM]", "http://apis.juhe.cn/wifi/local", -1)
	url = strings.Replace(url, "[KEY]", "e70a87d1214dbfa21215137920739239", -1)
	url = strings.Replace(url, "[LON]", lon, -1)
	url = strings.Replace(url, "[LAT]", lat, -1)
	beego.Debug("url:", url)

	resp, err := http.Get(url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------getWeather body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var obj Wifi
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getWiFiError(obj.ErrorCode)
		} else {
			c.Data["WResults"] = obj.WResult.WFData
		}
	} else {
		beego.Debug(err)
	}
}

func getWiFiError(errorcode int64) string {
	error_info := ""
	if errorcode == 201801 {
		error_info = "错误的经纬度"
	}
	if errorcode == 201802 {
		error_info = "城市区号不能为空"
	}
	if errorcode == 201803 {
		error_info = "查询无结果"
	}
	if errorcode == 201805 {
		error_info = "周边查询无结果"
	}
	return error_info
}
