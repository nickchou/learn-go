package controller

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/model"
)

//PsqlController 测试postgresql性能的控制器
type PsqlController struct {
	app.App
}

//Search 查询性能测试
func (con *PsqlController) Search() {
	times := 10000 //循环查询次数
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//msyql db connection
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	//pgsql db connection
	dbpg, errpg := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test password=123456 sslmode=disable")
	defer dbpg.Close()
	if errpg != nil {
		fmt.Println(errpg)
	}
	//查询mysql
	t1 := time.Now()
	cnt1 := 0
	for i := 0; i < times; i++ {
		db.Model(model.TestArea{}).Where("nodepath  REGEXP ?", "^77.*").Count(&cnt1)
	}
	//db.Where("nodepath  REGEXP ?", "^77.*").Count(&cnt1)
	elapsed1 := time.Since(t1)
	buffer.WriteString(fmt.Sprintf("<li>mysql耗时:%v,数据量：%v</li>", elapsed1, cnt1))

	//查询pgsql
	t2 := time.Now()
	cnt2 := 0
	for i := 0; i < times; i++ {
		dbpg.Model(model.TestArea{}).Where("nodepath<@ ?", "77").Count(&cnt2)
	}
	elapsed2 := time.Since(t2)
	buffer.WriteString(fmt.Sprintf("<li>pgsql耗时:%v,数据量：%v</li>", elapsed2, cnt2))
	io.WriteString(con.W(), buffer.String())
}

//Search2  PG查询性能测试
func (con *PsqlController) Search2() {
	times := 10000 //循环查询次数
	var buffer bytes.Buffer
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//pgsql db connection
	dbpg, errpg := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test password=123456 sslmode=disable")
	defer dbpg.Close()

	dbpg.DB().SetMaxIdleConns(1000)
	dbpg.DB().SetMaxOpenConns(5000)
	if errpg != nil {
		fmt.Println(errpg)
	}
	//设置管道
	// ch := make(chan int, 1) //改动点，管道长度设成了2
	//查询pgsql
	t2 := time.Now()

	for i := 0; i < times; i++ {
		// go func() {
		cnt2 := 0
		dbpg.Model(model.TestArea{}).Where("nodepath<@ ?", "77").Count(&cnt2)
		// 	ch <- 1
		// }()
	}
	// for j := 0; j < times; j++ {
	// 	<-ch
	// }
	elapsed2 := time.Since(t2)
	buffer.WriteString(fmt.Sprintf("<li>pgsql耗时:%v,数据量：%v</li>", elapsed2, "cnt2"))
	io.WriteString(con.W(), buffer.String())
}

//Data 初始化基础数据 ( InitData  MakeData命名为啥不行 ?)
func (con *PsqlController) Data() {
	//msyql db connection
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&model.TestArea{})
	//pgsql db connection
	dbpg, errpg := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test password=123456 sslmode=disable")
	defer dbpg.Close()
	if errpg != nil {
		fmt.Println(errpg)
	}
	dbpg.AutoMigrate(&model.TestArea{})

	//产生随机数据
	rdDatas := []model.TestArea{}
	t1 := time.Now()
	//生成随机数的数据
	for i := 0; i < 100000; i++ {
		//要生成的节点长度
		strNode := getRandomStr()
		entity := model.TestArea{Name: "city", NodePath: strNode}
		rdDatas = append(rdDatas, entity)
		// db.Create(&entity)
		// dbpg.Create(&entity)
	}
	elapsed1 := time.Since(t1)
	fmt.Println("data cost times:", elapsed1)
	//pgsql
	t3 := time.Now()
	for k, data3 := range rdDatas {
		data3.ID = (int64)(300000 + k)
		dbpg.Create(&data3)
	}
	elapsed3 := time.Since(t3)
	fmt.Println("pgsql cost times:", elapsed3)
	//mysql
	t2 := time.Now()
	for _, data2 := range rdDatas {
		db.Create(&data2)
	}
	elapsed2 := time.Since(t2)
	fmt.Println("msyql cost times:", elapsed2)
}

//getRandomStr 生成随机长度的字符串节点信息
func getRandomStr() string {
	var buffer bytes.Buffer
	len := rand.Intn(5) + 1
	for j := 0; j < len; j++ {
		tmpNum := rand.Intn(2000) + 1
		buffer.WriteString(strconv.Itoa(tmpNum) + ".")
	}
	return strings.TrimRight(buffer.String(), ".")
}
