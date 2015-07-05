package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
	"sort"
	"strings"
	"time"
)

const (
	token          = "qax580"
	qax580_name    = "qax580"
	error_info     = "我们正在努力成为一个有情怀的免费信息网站"
	null_info0     = "未搜索到相关信息"
	null_info1     = "未搜索到有关&的信息"
	subscribe_info = "欢迎关注庆安兄弟微盟，我们正在努力成为一个有情怀的免费信息发布平台，为庆安人服务"
	about_info     = "【客服服务】\n关注我们\n公众号:qax580\n微信:qax580kf\n腾讯微博:庆安兄弟微盟\nQQ : 2063883729\n邮箱：qaxiongdiweimeng@163.com"
	content_url    = "http://www.baoguangguang.cn/content?op=con&id=s%"
)

//接收文本消息
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

//接收音频消息
type VoiceRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MediaId      string
	Format       string
	Recognition  string
	MsgId        int
}

//点击消息
type EventResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Event        string
	EventKey     string
}

//消息类型
type TypeResponseBody struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string
}
type CDATAText struct {
	Text string `xml:",innerxml"`
}

//文本消息
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

func value2CDATA(v string) CDATAText {
	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

//图文消息
type ImageTextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	ArticleCount int
	Articles     []ImageTextResponseItem `xml:"Articles>item"`
}

//图文消息元素
type ImageTextResponseItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATAText
	Description CDATAText
	PicUrl      CDATAText
	Url         CDATAText
}

type WXController struct {
	beego.Controller
}

func (c *WXController) Get() {
	verification(c)
}
func (c *WXController) Post() {
	responseMsg(c)
}

//验证签名
func verification(c *WXController) {
	beego.Debug("valid ")
	echoStr := c.Input().Get("echostr")     //随机字符串
	signature := c.Input().Get("signature") //微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
	timestamp := c.Input().Get("timestamp") //时间戳
	nonce := c.Input().Get("nonce")         //随机数
	token := "qax580"
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := ""
	for i := 0; i < len(tmpArr); i++ {
		tmpStr += tmpArr[i]
	}
	tmpStrSha1 := goSha1(tmpStr)
	respnse := ""
	if strings.EqualFold(tmpStrSha1, signature) {
		respnse = echoStr
	}
	beego.Debug(respnse)
	c.Ctx.WriteString(respnse)
	return
}

//对字符串进行SHA1哈希
func goSha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

//响应消息
func responseMsg(c *WXController) {
	body := c.Ctx.Input.RequestBody
	requestType := &TypeResponseBody{}
	err := xml.Unmarshal(body, requestType)
	response_xml := ""
	if err != nil {
		beego.Debug(err.Error())
	} else {
		beego.Debug(requestType.MsgType)
		response_xml = responseTypeMsg(body, requestType.MsgType)
	}
	beego.Debug(response_xml)
	c.Ctx.WriteString(response_xml)
	return
}

//根据类型解析消息
func responseTypeMsg(body []byte, msgType string) string {
	response_xml := ""
	switch msgType {
	//文字
	case "text":
		requestBody := &TextRequestBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			beego.Debug(err.Error())
			response_xml = responseTextMsg(requestBody.FromUserName, error_info)
		} else {
			requestBody := &TextRequestBody{}
			err := xml.Unmarshal(body, requestBody)
			if err != nil {
				response_xml = responseTextMsg(requestBody.FromUserName, error_info)
			} else {
				beego.Debug(requestBody.Content)
				posts, err := models.QueryFuzzyLimitPost(requestBody.Content, 5)
				if err != nil {
					beego.Error(err)
				}
				beego.Debug(requestBody.FromUserName)
				beego.Debug(requestBody.ToUserName)
				response_xml = responseImageTextXML(requestBody.FromUserName, requestBody.Content, posts)

			}
		}
		//音频
	case "voice":
		requestBody := &VoiceRequestBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			response_xml = responseTextMsg(requestBody.FromUserName, error_info)
		} else {
			beego.Debug(requestBody.Recognition)
			posts, err := models.QueryFuzzyLimitPost(requestBody.Recognition, 5)
			if err != nil {
				beego.Error(err)
			}
			response_xml = responseImageTextXML(requestBody.FromUserName, requestBody.Recognition, posts)

		}
		//点击
	case "event":
		requestBody := &EventResponseBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			beego.Debug(err.Error())
			response_xml = responseTextMsg(requestBody.FromUserName, error_info)
		} else {
			//自定义点击事件
			if requestBody.Event == "CLICK" {
				//推荐
				if requestBody.EventKey == "recommend" {
					posts, err := models.QueryLimitPost(5)
					if err != nil {
						beego.Error(err)
					}
					response_xml = responseImageTextXML(requestBody.FromUserName, "", posts)
					//关于
				} else if requestBody.EventKey == "about" {
					response_xml = responseAbout(requestBody.FromUserName)
				} else {

				}
				//关注
			} else if requestBody.Event == "subscribe" {
				response_xml = responseTextMsg(requestBody.FromUserName, subscribe_info)
			} else {
				//其他类型
				response_xml = responseTextMsg(requestBody.FromUserName, error_info)
			}

		}

	default:
		beego.Debug(msgType)
	}
	beego.Debug(response_xml)
	return response_xml
}

func analysisNull(toUserName string, content string) string {
	null_info := null_info0
	if len(content) > 0 {
		null_info = strings.Replace(null_info1, "&", content, -1)
	}
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(qax580_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(null_info)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	error_info, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(error_info)
}

//返回文本信息
//textMsg 文本信息
func responseTextMsg(toUserName string, textMsg string) string {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(qax580_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(textMsg)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	about_xml, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(about_xml)
}

func responseAbout(toUserName string) string {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(qax580_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(about_info)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	about_xml, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(about_xml)
}

func responseImageTextXML(toUserName string, content string, posts []models.Post) string {
	articles := ""
	if posts != nil && len(posts) > 0 {
		imageTextResponseItems := []ImageTextResponseItem{}
		for i := 0; i < len(posts); i++ {
			imageTextResponseItem := ImageTextResponseItem{}
			imageTextResponseItem.Title = value2CDATA(posts[i].Title)
			imageTextResponseItem.Description = value2CDATA(posts[i].Info)
			imageTextResponseItem.Url = value2CDATA(strings.Replace(content_url, "s%", fmt.Sprintf("%d", posts[i].Id), -1))
			imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)
		}
		textResponseBody := &ImageTextResponseBody{}
		textResponseBody.MsgType = value2CDATA("news")
		textResponseBody.ArticleCount = len(imageTextResponseItems)
		textResponseBody.Articles = imageTextResponseItems
		textResponseBody.FromUserName = value2CDATA(qax580_name)
		textResponseBody.ToUserName = value2CDATA(toUserName)
		res, err := xml.MarshalIndent(textResponseBody, " ", "  ")

		if err != nil {
			beego.Debug(err.Error())
		} else {
			articles = string(res)
		}
	} else {
		articles = analysisNull(toUserName, content)
	}

	return articles
}