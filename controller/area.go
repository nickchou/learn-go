package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/comm"
	"github.com/nickchou/learn-go/model"
)

//AreaController 国家统计局数据获取控制器
type AreaController struct {
	app.App
}

//Goo 数据库一些初始化操作
func (con *AreaController) Goo() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
		//检查表
		if db.HasTable(&model.Area{}) {
			fmt.Println("存在表Area")
		} else {
			fmt.Println("不存在表Area")
			//AutoMigrate只会创建表、缺失的列和索引。不会改变现有列的数据类型或者删除数据
			//http://doc.gorm.io/database.html#migration
			db.AutoMigrate(&model.Area{})
			fmt.Println("创建表Area")
		}
	}

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
	//创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	//AutoMigrate只会创建表、缺失的列和索引。不会改变现有列的数据类型或者删除数据
	//http://doc.gorm.io/database.html#migration
	db.AutoMigrate(&model.Area{})
	//找文档
	doc.Find(".provincetr td a").Each(func(i int, se *goquery.Selection) {
		url, _ := se.Attr("href")
		url = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2017/" + url
		city, _ := comm.ToGBK(se.Text())
		//插入数据
		area := model.Area{Aid: 0, Pid: 100000, Level: 1, Name: city, Lng: 2223.322321332, Lat: 32131.12345, URL: url}
		//主键为空返回true
		issucc := db.NewRecord(&area)
		if issucc {
			//插入数据
			db.Create(&area)
		}
		//由于插入数据后有主键所以会返回false
		issucc2 := db.NewRecord(&area)
		fmt.Println(fmt.Sprintf("%v,%v,%v", issucc, issucc2, area.ID))
		//造返回字符串
		buffer.WriteString(fmt.Sprintf("%v,%v,%v<br/>", i+1, city, url))
	})
	io.WriteString(con.W(), buffer.String())
}
