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
	imageTextResponseItems := []ImageTextResponseItem{}
	imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	imageTextResponseItems = append(imageTextResponseItems, ImageTextResponseItem{})
	fmt.Println(responseImageText(imageTextResponseItems))
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
