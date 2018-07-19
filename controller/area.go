package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/comm"
)

//AreaController 国家统计局数据获取控制器
type AreaController struct {
	app.App
}

//Prov 获取国家统计局的行政规划数据
func (con *AreaController) Prov() {
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//goquery http
	doc, err := goquery.NewDocument("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2017/index.html")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".provincetr td a").Each(func(i int, se *goquery.Selection) {
		url, _ := se.Attr("href")
		url = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2017/" + url
		city, _ := comm.ToGBK(se.Text())
		//造返回字符串
		buffer.WriteString(fmt.Sprintf("%v,%v,%v<br/>", i+1, city, url))
	})
	io.WriteString(con.W(), buffer.String())
}
