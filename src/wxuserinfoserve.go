package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	jsonStr := ` {"openid":"o3AhEuB_wdTELvlErL4F1Em4Nck4","nickname":"寰","sex":2,"language":"zh_CN","city":"","province":"","country":"涓","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/lQyhZL8HdN24nyDtggslekaRNoUEApk3pNpPUk6Ahw5iadM8CEZU5g7s0wYDY3voMN6jfOGvvtBglicPNYJHFQXqwYGaiaVbN5F\/0","privilege":[]}`
	fmt.Fprintf(w, jsonStr) //这个写入到w的是输出到客户端的
}
func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
