package main

import (
	// "bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "time"
	"strings"
)

var client = &http.Client{}

func Get() {
	//向服务端发送get请求
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8080/wx?timestamp=adas&nonce=ccc&echostr=123456&signature=4890bc5f00fe125eb71acde77f2282f80a568ec5", nil)
	// request, _ := http.NewRequest("GET", "http://127.0.0.1:8080/wx", nil
	response, _ := client.Do(request)
	fmt.Println(response)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		fmt.Println(bodystr)
	}
}

func Post() {
	//post请求
	postValues := url.Values{}
	// postValues.Add("MsgType", "text")
	// postValues.Add("CreateTime", time.Now().String())
	// postValues.Add("body", "<xml><URL><![CDATA[http://www.baoguangguang.cn/qax580/wx_sample.php]]></URL><ToUserName><![CDATA[sbb001]]></ToUserName><FromUserName><![CDATA[abc]]></FromUserName><CreateTime>123456</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[abc]]></Content><MsgId>123</MsgId></xml>")
	body_xml := "<xml><URL><![CDATA[http://www.baoguangguang.cn/qax580/wx_sample.php]]></URL><ToUserName><![CDATA[sbb001]]></ToUserName><FromUserName><![CDATA[abc]]></FromUserName><CreateTime>123456</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[1]]></Content><MsgId>123</MsgId></xml>"
	// body_xml := "<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[voice]]></MsgType><MediaId><![CDATA[media_id]]></MediaId><Format><![CDATA[Format]]></Format><Recognition><![CDATA[腾讯微信团队]]></Recognition><MsgId>1234567890123456</MsgId></xml>"
	postValues.Add("body", body_xml)
	// body_xml = "<xml><URL><![CDATA[http://www.baoguangguang.cn/qax580/wx_sample.php]]></URL><ToUserName><![CDATA[sbb]]></ToUserName><FromUserName><![CDATA[12]]></FromUserName><CreateTime>123</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[CLICK]]></Event><EventKey><![CDATA[recommend]]></EventKey><Latitude></Latitude><Longitude></Longitude><Precision></Precision><MsgId>1</MsgId></xml>"
	// body_xml = "<xml><URL><![CDATA[http://www.baoguangguang.cn/qax580/wx_sample.php]]></URL><ToUserName><![CDATA[sbb]]></ToUserName><FromUserName><![CDATA[12]]></FromUserName><CreateTime>123</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[about]]></Event><Latitude></Latitude><Longitude></Longitude><Precision></Precision><MsgId>1</MsgId></xml>"
	// c := []byte{60, 120, 109, 108, 62, 60, 84, 111, 85, 115, 101, 114, 78, 97, 109, 101, 62, 60, 33, 91, 67, 68, 65, 84, 65, 91, 103, 104, 95, 56, 98, 97, 56, 99, 48, 51, 48, 57, 97, 52, 100, 93, 93, 62, 60, 47, 84, 111, 85, 115, 101, 114, 78, 97, 109, 101, 62, 10, 60, 70, 114, 111, 109, 85, 115, 101, 114, 78, 97, 109, 101, 62, 60, 33, 91, 67, 68, 65, 84, 65, 91, 111, 51, 65, 104, 69, 117, 66, 95, 119, 100, 84, 69, 76, 118, 108, 69, 114, 76, 52, 70, 49, 69, 109, 52, 78, 99, 107, 52, 93, 93, 62, 60, 47, 70, 114, 111, 109, 85, 115, 101, 114, 78, 97, 109, 101, 62, 10, 60, 67, 114, 101, 97, 116, 101, 84, 105, 109, 101, 62, 49, 52, 51, 52, 48, 51, 57, 57, 55, 48, 60, 47, 67, 114, 101, 97, 116, 101, 84, 105, 109, 101, 62, 10, 60, 77, 115, 103, 84, 121, 112, 101, 62, 60, 33, 91, 67, 68, 65, 84, 65, 91, 101, 118, 101, 110, 116, 93, 93, 62, 60, 47, 77, 115, 103, 84, 121, 112, 101, 62, 10, 60, 69, 118, 101, 110, 116, 62, 60, 33, 91, 67, 68, 65, 84, 65, 91, 67, 76, 73, 67, 75, 93, 93, 62, 60, 47, 69, 118, 101, 110, 116, 62, 10, 60, 69, 118, 101, 110, 116, 75, 101, 121, 62, 60, 33, 91, 67, 68, 65, 84, 65, 91, 97, 98, 111, 117, 116, 93, 93, 62, 60, 47, 69, 118, 101, 110, 116, 75, 101, 121, 62, 10, 60, 47, 120, 109, 108, 62}
	// body_xml = string(c)
	fmt.Println("_____________")
	fmt.Println(body_xml)
	fmt.Println("_____________")
	// resp, err := client.PostForm("http://127.0.0.1:8080/wx", postValues)
	resp, err := client.Post("http://127.0.0.1:8080/wx", "application/x-www-form-urlencoded", strings.NewReader(body_xml))
	// fmt.Println(resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
func test() {
	v := url.Values{}
	v.Set("huifu", "hello world")
	v.Set("huifu1", "hello world")
	// body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/wx", strings.NewReader("abc"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构
	fmt.Printf("%+v\n", req.Body)
	// fmt.Printf("%+v\n", req.URL)
	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}

func main() {
	// Get()
	Post()
	// test()

}
