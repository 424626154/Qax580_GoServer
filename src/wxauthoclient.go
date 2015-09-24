package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"net/url"
	"strings"
)

func main() {
	GetWeiXinCode()
}

func GetWeiXinCode() {
	wx_url := "[REALM]/?appid=[APPID]&redirect_uri=[REDIRECT_URI]&response_type=code&scope=snsapi_userinfo&state=[STATE]#wechat_redirect"
	// realm_name := "http://localhost:9090"
	realm_name := "https://open.weixin.qq.com/connect/oauth2/authorize"
	appid := "wxf0e81c3bee622d60"
	// redirect_uri := "http://localhost.com/wxuser"
	//授权后重定向的回调链接地址，请使用urlencode对链接进行处理
	redirect_uri := "http://www.baoguanguang.cn/wxuser"
	redirect, _ := url.ParseQuery(redirect_uri)
	redirect_uri = redirect.Encode()
	redirect_uri = Substr(redirect_uri, 0, len(redirect_uri)-len("="))
	state := "qax580" //重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[REDIRECT_URI]", redirect_uri, -1)
	wx_url = strings.Replace(wx_url, "[STATE]", state, -1)
	fmt.Println(wx_url)

	// log.Fatal(wx_url)

	// resp, err := http.Get(wx_url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(string(body))
	// fmt.Println(string(body))
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
