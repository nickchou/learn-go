package controller

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nickchou/learn-go/app"
)

//PgsqlController 测试postgresql性能的控制器
type PsqlController struct {
	app.App
}

//InitData 初始化基础数据 ( InitData  MakeData命名为啥不行 ?)
func (con *PsqlController) Data() {
	//msyql db
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//pgsql db
	dbpg, errpg := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test password=123456 sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	defer dbpg.Close()
	if errpg != nil {
		fmt.Println(errpg)
	}
}
