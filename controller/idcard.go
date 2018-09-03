package controller

import (

	//导入mysql的驱动包
	"bytes"
	"fmt"
	"io"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
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
	//MSSQL
	engine, err := xorm.NewEngine("mssql", "server=127.0.0.1;user id=sa;password=123456;database=shifenzheng")
	if err != nil {
		fmt.Println(err)
	}
	//PGSQL
	enginePG, err := xorm.NewEngine("postgres", "user=postgres password=123456 dbname=test host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	//控制台打印SQL语句
	// engine.ShowSQL(true)

	//定期自动释放
	defer engine.Close()
	defer enginePG.Close()
	//
	ids := []model.MsIdcaid{}
	//同步postgresql的表结构
	//fmt.Println("start sync table")
	err = enginePG.Sync(new(model.PgIdcard))
	if err != nil {
		fmt.Println(err)
	}

	engine.SQL("SELECT Id,Name from cdsgus where id in (2,3,4,5,6)and id>? order by id asc OFFSET ? ROW FETCH NEXT ? ROW ONLY;", 0, 2, 10).Find(&ids)
	// if len(ids) > 0 {
	// 	pgs := []model.PgIdcard{}
	// 	//转换为Postgresql 的model
	// 	for _, c := range ids {
	// 		idcard := model.PgIdcard{ID: c.Id, Name: c.Name, CardType: c.Ctftp, CardNum: c.Ctfid, Gender: c.Gender, Birthday: c.Birthday, Mobile: c.Mobile, EMail: c.Email, Address: c.Address}
	// 		pgs = append(pgs, idcard)
	// 	}
	// 	//postgre 批量入库
	// 	enginePG.Insert(pgs)
	// }

	//engine.Where("").Select(" Address ").Find(&ids)
	//engine.Desc("id").Cols("Id", "Address").Where(" id in(2,3,4,5,6)").Find(&ids)

	//engine.Cols("Id", "Address").Where("id in(2,3,4,5,6)").OrderBy("id desc,address asc").Find(&ids)
	//engine.Cols("Id", "Name").Where("id in(2,3,4,5,6)").OrderBy("id desc,address asc").Limit(10, 2).Find(&ids)
	//engine.SQL("SELECT Address from cdsgus where id in (2,3,4,6) order by id desc ").Find(&ids)
	//engine.SQL("SELECT Id,Name from cdsgus where id in (2,3,4,5,6) order by id desc").Limit(10, 2).Find(&ids)
	//data, _ := engine.Sql("SELECT Id,Name from cdsgus where id in (2,3,4,5,6) ").OrderBy("id").Limit(10, 2).Query()
	//buffer.WriteString(fmt.Sprintf("<li>count %v</li>", len(data)))
	buffer.WriteString(fmt.Sprintf("<li>count %v</li>", len(ids)))
	for i, card := range ids {
		buffer.WriteString(fmt.Sprintf("<li>i= %v,id= %v,name=%v</li>", i, card.Id, card.Name))
	}
	//response
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(con.W(), buffer.String())
}
