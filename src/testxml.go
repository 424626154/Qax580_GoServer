// <xml>
// <ToUserName><![CDATA[toUser]]></ToUserName>
// <FromUserName><![CDATA[fromUser]]></FromUserName>
// <CreateTime>12345678</CreateTime>
// <MsgType><![CDATA[news]]></MsgType>
// <ArticleCount>2</ArticleCount>
// <Articles>
// <item>
// <Title><![CDATA[title1]]></Title>
// <Description><![CDATA[description1]]></Description>
// <PicUrl><![CDATA[picurl]]></PicUrl>
// <Url><![CDATA[url]]></Url>
// </item>
// <item>
// <Title><![CDATA[title]]></Title>
// <Description><![CDATA[description]]></Description>
// <PicUrl><![CDATA[picurl]]></PicUrl>
// <Url><![CDATA[url]]></Url>
// </item>
// </Articles>
// </xml>
package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type ImageTextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	ArticleCount int
	Articles     CDATAText
}
type ImageTextResponseItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATAText
	Description CDATAText
	PicUrl      CDATAText
	Url         CDATAText
}

func main() {
	// text := ""
	// imageTextResponseItem := &ImageTextResponseItem{}
	// articles := ""
	// res, err := xml.MarshalIndent(imageTextResponseItem, " ", "  ")

	// if err != nil {

	// } else {
	// 	articles += "\n" + string(res) + "\n"
	// }
	// res, err = xml.MarshalIndent(imageTextResponseItem, " ", "  ")

	// if err != nil {

	// } else {
	// 	articles += "\n" + string(res) + "\n"
	// }

	// textResponseBody := &ImageTextResponseBody{}
	// // items := []ImageTextResponseItem{}
	// // items = append(items, ImageTextResponseItem{})
	// // items = append(items, ImageTextResponseItem{})
	// // textResponseBody.Articles = items
	// textResponseBody.Articles.Text = "\n" + articles + "\n"
	// textResponseBody.ToUserName = CDATAText{"abc"}
	// res, err = xml.MarshalIndent(textResponseBody, " ", "  ")

	// if err != nil {

	// } else {
	// 	text = string(res)
	// }

	// fmt.Println(text)
	// textResponseItem := &TextResponseItem{}
	// res, err = xml.MarshalIndent(textResponseItem, " ", "  ")
	// if err != nil {

	// } else {
	// 	text = string(res)
	// }

	// fmt.Println(text)
	// textResponseItems := &TextResponseItems{}
	// res, err = xml.MarshalIndent(textResponseItems, " ", "  ")
	// if err != nil {

	// } else {
	// 	text = string(res)
	// }

	// fmt.Println(text)
	// imageTextResponseItems := []ImageTextResponseItem{}
	// imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	// imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	// imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	// fmt.Println(responseImageText(imageTextResponseItems))
	// fmt.Println(len("https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx570bbcc8cf9fdd80redirect_uri=http%3a%2f%2fwww.baoguangguang.cn%2fmymessageresponse_type=codescope=snsapi_userinfostate=STATE#wechat_redirect"))
	// fmt.Println(len("https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx570bbcc8cf9fdd80redirect_uri=http%3a%2f%2fwww.baoguangguang.cn%2fwxuploderesponse_type=codescope=snsapi_userinfostate=STATE#wechat_redirect"))
	// fmt.Println(len("https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx570bbcc8cf9fdd80redirect_uri=http%3a%2f%2fwww.baoguangguang.cn%2fwxhomeresponse_type=codescope=snsapi_userinfostate=STATE#wechat_redirect"))
	fmt.Println(len("1940年11月27日，他出生在美国三蕃市，英文名叫布鲁斯?李。因为父亲是演员，他从小就有了跑龙套的机会，于是产生想当一名演员的梦想。可由于身体虚弱，父亲便让他拜师习武来强身。1961年，他考入华盛顿州立大学主修哲学，后来，他像所有正常人一样结婚生子。但在他内心深处，一刻也不曾放弃当一名演员的梦想。　　一天，他与一位朋友谈到梦想时，随手在一张便笺上写下了自己的人生目标：“我，布鲁斯?李，将会成为全美国最高薪酬的超级巨星。作为回报，我将奉献出最激动人心、最具震撼力的演出。从1970年开始，我将会赢得世界性声"))
}

func responseImageText(imageTextResponseItems []ImageTextResponseItem) string {

	articles := "\n"
	if len(imageTextResponseItems) > 0 {
		for i := 0; i < len(imageTextResponseItems); i++ {
			res, err := xml.MarshalIndent(imageTextResponseItems[i], " ", "  ")

			if err != nil {

			} else {
				articles += string(res) + "\n"
			}
		}
	}
	// articles += "\n"
	textResponseBody := &ImageTextResponseBody{}
	textResponseBody.ArticleCount = len(imageTextResponseItems)
	textResponseBody.Articles.Text = articles
	textResponseBody.ToUserName = CDATAText{"abc"}
	res, err := xml.MarshalIndent(textResponseBody, " ", "  ")

	if err != nil {

	} else {
		articles = string(res)
	}
	return articles

	// text := ""
	// imageTextResponseItem := &ImageTextResponseItem{}
	// articles := ""
	// res, err := xml.MarshalIndent(imageTextResponseItem, " ", "  ")

	// if err != nil {

	// } else {
	// 	articles += "\n" + string(res) + "\n"
	// }
	// res, err = xml.MarshalIndent(imageTextResponseItem, " ", "  ")

	// if err != nil {

	// } else {
	// 	articles += "\n" + string(res) + "\n"
	// }

	// textResponseBody := &ImageTextResponseBody{}
	// // items := []ImageTextResponseItem{}
	// // items = append(items, ImageTextResponseItem{})
	// // items = append(items, ImageTextResponseItem{})
	// // textResponseBody.Articles = items
	// textResponseBody.Articles.Text = "\n" + articles + "\n"
	// textResponseBody.ToUserName = CDATAText{"abc"}
	// res, err = xml.MarshalIndent(textResponseBody, " ", "  ")

	// if err != nil {

	// } else {
	// 	text = string(res)
	// }
}
