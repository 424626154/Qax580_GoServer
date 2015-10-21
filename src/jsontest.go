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

type WeatherJson struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     Result `json:"result"`
	ErrorCode  int64  `json:"error_code"`
}

type Result struct {
	Today Today `json:"today"`
}
type Today struct {
	City           string `json:"city"`
	DateY          string `json:"date_y"`
	Week           string `json:"week"`
	Temperature    string `json:"temperature"`
	Weather        string `json:"weather"`
	Wind           string `json:"wind"`
	Dressingindex  string `json:"dressing_index"`
	DressingAdvice string `json:"dressing_advice"`
	UvIndex        string `json:"uv_index"`
	WashIndex      string `json:"wash_index"`
	WravelIndex    string `json:"wravel_index"`
	WxerciseIndex  string `json:"wxercise_index"`
}

func main() {
	// jsonStr := `{"city":"CITY","country":"COUNTRY","headimgurl":"http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46","nickname":"NICKNAME","openid":"OPENID","privilege":["PRIVILEGE1","PRIVILEGE2"],"province":"PROVINCE","sex":"1","unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL"}`
	// jsonStr := `{"openid":"o3AhEuB_wdTELvlErL4F1Em4Nck4","nickname":"寰","sex":2,"language":"zh_CN","city":"","province":"","country":"涓","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/lQyhZL8HdN24nyDtggslekaRNoUEApk3pNpPUk6Ahw5iadM8CEZU5g7s0wYDY3voMN6jfOGvvtBglicPNYJHFQXqwYGaiaVbN5F\/0","privilege":[]}`
	// var uij UserInfoJson
	// if err := json.Unmarshal([]byte(jsonStr), &uij); err == nil {
	// 	fmt.Println("================json str 转struct==")
	// 	fmt.Println(uij)
	// } else {
	// 	fmt.Println(err)
	// }

	jsonStr := `{"resultcode":"200","reason":"successed!","result":{"sk":{"temp":"15","wind_direction":"东北风","wind_strength":"1级","humidity":"76%","time":"23:17"},"today":{"temperature":"10℃~22℃","weather":"多云","weather_id":{"fa":"01","fb":"01"},"wind":"微风","week":"星期三","city":"北京","date_y":"2015年10月14日","dressing_index":"较舒适","dressing_advice":"建议着薄外套或牛仔衫裤等服装。年老体弱者宜着夹克衫、薄毛衣等。昼夜温差较大，注意适当增减衣服。","uv_index":"强","comfort_index":"","wash_index":"较适宜","travel_index":"较适宜","exercise_index":"较适宜","drying_index":""},"future":{"day_20151014":{"temperature":"10℃~22℃","weather":"多云","weather_id":{"fa":"01","fb":"01"},"wind":"微风","week":"星期三","date":"20151014"},"day_20151015":{"temperature":"10℃~25℃","weather":"晴","weather_id":{"fa":"00","fb":"00"},"wind":"微风","week":"星期四","date":"20151015"},"day_20151016":{"temperature":"12℃~25℃","weather":"霾转阴","weather_id":{"fa":"53","fb":"02"},"wind":"微风","week":"星期五","date":"20151016"},"day_20151017":{"temperature":"13℃~23℃","weather":"小雨转多云","weather_id":{"fa":"07","fb":"01"},"wind":"北风3-4 级","week":"星期六","date":"20151017"},"day_20151018":{"temperature":"11℃~18℃","weather":"阴转小雨","weather_id":{"fa":"02","fb":"07"},"wind":"微风","week":"星期日","date":"20151018"},"day_20151019":{"temperature":"11℃~21℃","weather":"多云转阴","weather_id":{"fa":"01","fb":"02"},"wind":"微风","week":"星期一","date":"20151019"},"day_20151020":{"temperature":"12℃~17℃","weather":"小雨","weather_id":{"fa":"07","fb":"07"},"wind":"微风","week":"星期二","date":"20151020"}}},"error_code":0}`
	var uij WeatherJson
	if err := json.Unmarshal([]byte(jsonStr), &uij); err == nil {
		fmt.Println("================json str 转struct==")
		today := uij.Result.Today
		fmt.Println(today)
	} else {
		fmt.Println(err)
	}
}
