package main

import (
	"encoding/json"
	"fmt"
)

type UserInfoJson struct {
	OpenId     string   `json:"openid"`
	NickeName  string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func main() {
	// jsonStr := `{"city":"CITY","country":"COUNTRY","headimgurl":"http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46","nickname":"NICKNAME","openid":"OPENID","privilege":["PRIVILEGE1","PRIVILEGE2"],"province":"PROVINCE","sex":"1","unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL"}`
	jsonStr := `{"openid":"o3AhEuB_wdTELvlErL4F1Em4Nck4","nickname":"寰","sex":2,"language":"zh_CN","city":"","province":"","country":"涓","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/lQyhZL8HdN24nyDtggslekaRNoUEApk3pNpPUk6Ahw5iadM8CEZU5g7s0wYDY3voMN6jfOGvvtBglicPNYJHFQXqwYGaiaVbN5F\/0","privilege":[]}`
	var uij UserInfoJson
	if err := json.Unmarshal([]byte(jsonStr), &uij); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(uij)
	} else {
		fmt.Println(err)
	}
}
