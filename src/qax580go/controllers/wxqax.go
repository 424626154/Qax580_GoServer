package controllers

/*
微信http服务器
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"
)

type WxqaxController struct {
	beego.Controller
}

/**************请求接口*******************/
/*
关注与取消关注请求
／wxqax/sunscribe
//http请求参数 subscribe_type from_openid to_user create_time
*/
func (c *WxqaxController) Sunscribe() {
	response_json := `{"errcode":1,"errmsg":"Sunscribe error"}`
	subscribe_type := c.Input().Get("subscribe_type")
	from_openid := c.Input().Get("from_openid")
	appid := c.Input().Get("appid")
	secret := c.Input().Get("secret")
	if len(subscribe_type) != 0 && len(from_openid) != 0 {
		response_json = getWxAccessToken(c, from_openid, appid, secret)

		var user models.Wxuserinfo
		if err := json.Unmarshal([]byte(response_json), &user); err == nil {
			beego.Debug("----------------get Wxuserinfo json--------------------")
			beego.Debug(user)
			if user.ErrCode == 0 {
				state, err := models.SunscribeWxUserInfo(user)
				if err != nil {
					beego.Error(err)
					response_json = `{"errcode":1,"errmsg":"AddWxUserInfo error"}`
				} else {
					beego.Debug("SunscribeWxUserInfo state", state)
					if subscribe_type == "subscribe" && state == 0 {
						err = models.AddWxUserMoney(user.OpenId, 4)
						if err != nil {
							beego.Error(err)
							response_json = `{"errcode":1,"errmsg":"AddWxUserMoney error"}`
						} else {
							_, err = models.AddUserMoneyRecord(user.OpenId, MONEY_SUBSCRIBE_SUM, MONEY_SUBSCRIBE)
						}
					}
				}
			}
		} else {
			beego.Debug("----------------get Token json error--------------------")
			beego.Debug(err)
		}

	}
	c.Ctx.WriteString(response_json)

}

/**************请求接口*******************/
/*
 openid=
{
    "subscribe": 1,
    "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
    "nickname": "Band",
    "sex": 1,
    "language": "zh_CN",
    "city": "广州",
    "province": "广东",
    "country": "中国",
    "headimgurl":    "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
   "subscribe_time": 1382694957,
   "unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
   "remark": "",
   "groupid": 0
}

{"errcode":40013,"errmsg":"invalid appid"}
*/
func (c *WxqaxController) Getuserinfo() {
	response_json := `{"errcode":1,"errmsg":"Getuserinfo error"}`
	openid := c.Input().Get("openid")
	if len(openid) != 0 {
		response_json = getWxAccessToken(c, openid, qa_appid, qa_secret)
	} else {
		response_json = `{"errcode":1,"errmsg":"parameter error"}`
	}
	c.Ctx.WriteString(response_json)
}

func getWxAccessToken(c *WxqaxController, openid string, appid string, secret string) string {
	// https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
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
		realm_name = "http://localhost:9090"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	}
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
			response_json = getWxUserInfo(atj.AccessToken, openid, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
	return response_json
}

func getWxUserInfo(access_toke, openid string, c *WxqaxController) string {
	response_json := `{"errcode":1,"errmsg":"getWxUserInfo error"}`
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&openid=[OPENID]&lang=zh_CN"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9091"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/user/info"
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	wx_url = strings.Replace(wx_url, "[OPENID]", openid, -1)
	beego.Debug("----------------get UserInfo --------------------")
	beego.Debug(wx_url)

	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------get UserInfo body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var uij models.Wxuserinfo
	if err := json.Unmarshal(body, &uij); err == nil {
		beego.Debug("----------------get UserInfo json--------------------")
		beego.Debug(uij)
		if uij.ErrCode == 0 {
			response_json = string(body)
		}
	} else {
		beego.Debug("----------------get UserInfo json error--------------------")
		beego.Debug(err)
	}
	return response_json
}

func getToken() (errcode int64, token string) {
	r_errcode := int64(0)
	r_token := ""
	// https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&grant_type=client_credential"
	realm_name := "https://api.weixin.qq.com/cgi-bin/token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("token url:", wx_url)
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
	// body := []byte(`{"access_token":"ACCESS_TOKEN","expires_in":7200}`)
	var atj AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug(atj)
		if atj.ErrCode == 0 {
			r_token = atj.AccessToken
		} else {
			r_errcode = atj.ErrCode
		}
	} else {
		beego.Debug(err)
	}
	return r_errcode, r_token
}

func getWxUser(openid string, access_token string) (models.Wxuserinfo, error) {
	user := models.Wxuserinfo{}
	// ?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	query_url := "[REALM]?access_token=[ACCESS_TOKEN]&openid=[OPENID]&lang=zh_CN"
	query_url = strings.Replace(query_url, "[REALM]", "https://api.weixin.qq.com/cgi-bin/user/info", -1)
	query_url = strings.Replace(query_url, "[ACCESS_TOKEN]", access_token, -1)
	query_url = strings.Replace(query_url, "[OPENID]", openid, -1)
	beego.Debug("importUser url:", query_url)

	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------get UserInfo body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var uij models.Wxuserinfo
	if err := json.Unmarshal(body, &uij); err == nil {
		beego.Debug("----------------get UserInfo json--------------------")
		beego.Debug(uij)
		if uij.ErrCode == 0 {
			user = uij
		}

	} else {
		beego.Debug("----------------get UserInfo json error--------------------")
		beego.Debug(err)
	}
	return user, err
}
