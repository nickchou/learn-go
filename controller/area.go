package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"

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
	provs := []model.Area{}
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
		//省code
		provCode, _ := strconv.ParseInt(comm.Substring(url, 0, 2)+"0000", 10, 64)
		url = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2017/" + url
		city, _ := comm.ToGBK(se.Text())
		//插入数据
		area := model.Area{Aid: provCode, Pid: 100000, Level: 1, Name: city, Lng: 2223.322321332, Lat: 32131.12345, URL: url}
		area.SimName = comm.Substring(area.Name, 0, 2)
		if area.SimName == "内蒙" {
			area.SimName = "内蒙古"
		}
		//主键为空返回true
		issucc := db.NewRecord(&area)
		if issucc {
			//插入数据
			db.Create(&area)
			//省份插入数组，方便后面拿市区数据
			provs = append(provs, area)

		}
		//由于插入数据后有主键所以会返回false
		issucc2 := db.NewRecord(&area)
		fmt.Println(fmt.Sprintf("%v,%v,%v", issucc, issucc2, area.ID))
		//造返回字符串
		buffer.WriteString(fmt.Sprintf("%v,%v,%v<br/>", i+1, city, url))
	})
	//根据省份加载城市信息
	City(provs)
	io.WriteString(con.W(), buffer.String())
}

//City  根据省份获取所有城市
func City(provs []model.Area) {
	citys := []model.Area{}
	//数据库连接字符串
	db, _ := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	//遍历省，获取省下面的城市信息
	for i, prov := range provs {
		fmt.Println(i, prov.ID, prov.SimName, prov.URL)
		//抓取市区数据
		doc, _ := goquery.NewDocument(prov.URL)
		doc.Find(".citytr").Each(func(i int, se *goquery.Selection) {
			tagsA := se.Find("a")
			cityCode := tagsA.Eq(0).Text()
			CityName, _ := comm.ToGBK(tagsA.Eq(1).Text())
			cityURL, _ := tagsA.Eq(1).Attr("href")
			cityURL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2017/" + cityURL
			//city编码只取前6位就可以了
			cityID, _ := strconv.ParseInt(comm.Substring(cityCode, 0, 6), 10, 64)
			//省份信息提取前面3位
			provID, _ := strconv.ParseInt(comm.Substring(cityCode, 0, 2)+"0000", 10, 64)
			//实体信息
			city := model.Area{Aid: cityID, Pid: provID, Level: 2, Name: CityName, Lng: 1.111111, Lat: 1.111111, URL: cityURL}
			//入库
			db.Create(&city)
			citys = append(citys, city)
		})
	}
	fmt.Println("load city end")
}
