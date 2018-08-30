package controller

import (

	//导入mysql的驱动包
	"bytes"
	"fmt"
	"io"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/model"
)

//SfzController 测试postgresql性能的控制器
type SfzController struct {
	app.App
}

//Data 导身份证数据测试
func (con *SfzController) Data() {
	var buffer bytes.Buffer
	//connection，server=127.0.0.1;user id=sa;password=123456;database=shifenzheng;app name=MyAppName
	engine, err := xorm.NewEngine("mssql", "server=127.0.0.1;user id=sa;password=123456;database=shifenzheng")
	//控制台打印SQL语句
	engine.ShowSQL(true)
	if err != nil {
		fmt.Println(err)
	}
	defer engine.Close()
	buffer.WriteString("hello")
	ids := []model.MsIdcaid{}

	//engine.Where("").Select(" Address ").Find(&ids)
	engine.SQL("SELECT Address from cdsgus where id in (2,3,4,6) ").Find(&ids)
	//counts, _ := engine.Count(&ids)
	buffer.WriteString(fmt.Sprintf("<li>count %v</li>", len(ids)))
	io.WriteString(con.W(), buffer.String())
}
