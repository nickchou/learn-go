package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/nickchou/learn-go/app"
)

//CommController 公共控制器，执行通用方法
type CommController struct {
	app.App
}

//Topic 测试方法，goquery获取文章列表
func (con *CommController) Topic() {
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//goquery http
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		buffer.WriteString(fmt.Sprintf("第%v个帖子标题：%v<br/>", i+1, title))
	})
	io.WriteString(con.W(), buffer.String())
}

//Country 获取维基百科里的国家信息
func (con *CommController) Country() {
	//wiki下载国家图片地址参考https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//要抓的网页
	doc, err := goquery.NewDocument("http://en.wikipedia.org/wiki/Continent")
	if err != nil {
		log.Fatal(err)
	}
	buffer.WriteString("<style>table { border-collapse:collapse; } table,th, td  {  border: 1px solid black;</style>")
	buffer.WriteString("七大洲信息<br/>")
	doc.Find(".wikitable").Each(func(i int, sel *goquery.Selection) {
		//table包含样式才继续
		if sel.HasClass("sortable") && i == 1 {
			//找到表格body里的tr
			trs := sel.Find("tbody").Find("tr")
			//调试tbody
			//sshtml, _ := sel.Find("tbody").Html()
			//fmt.Println(sshtml)
			//遍历trs
			buffer.WriteString("<table>")
			buffer.WriteString("<tr><td>洲名称</td><td>国家信息明细</td><td>陆地面积km²</td><td>面积占比</td><td>人口数量</td><td>人口占比</td><td>人口最多城市</td></tr>")
			for j := range trs.Nodes {
				//j==0是过滤第一个tr表头，Find("tbody tr")达不到过滤效果
				if j == 0 {
					continue
				}
				tr := trs.Eq(j) //找到第j个tr，这里要注意和trs.Get(j)的区别
				buffer.WriteString("<tr>")
				//找到洲的名称
				ctrlA := tr.Find("th a")
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", ctrlA.Text()))
				ctrlAHref, _ := ctrlA.Attr("href") //拿洲的url连接
				buffer.WriteString(fmt.Sprintf("<td>https://en.wikipedia.org%v</td>", ctrlAHref))
				//找到洲的其他属性列
				tds := tr.Find("td")
				//陆地面积
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", tds.Eq(1).Text()))
				//陆地占比
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", tds.Eq(3).Text()))
				//人口
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", tds.Eq(4).Text()))
				//人口占比
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", tds.Eq(5).Text()))
				//人口最多城市
				buffer.WriteString(fmt.Sprintf("<td>%v</td>", tds.Eq(6).Text()))
				buffer.WriteString("</tr>")

			}
			buffer.WriteString("</table>")
		}
	})
	io.WriteString(con.W(), buffer.String())
}
