package controller

import (
	//"html/template"

	"fmt"
	"io"

	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/model"
)

// IndexController 整个项目默认的控制器
type IndexController struct {
	app.App
}

//Index IndexController默认的执行方法名
func (con *IndexController) Index() {
	con.Data["name"] = "张三"
	con.Data["email"] = "zhangsan@163.com"
	//创建slice切片列表
	s1 := model.Student{Name: "张小三", Sex: 1, Age: 20}
	s2 := model.Student{Name: "李小四", Sex: 1, Age: 18}
	s3 := model.Student{Name: "王小五"}
	stus := []model.Student{s1, s2, s3}
	con.Data["stus"] = stus

	//使用template模板
	con.Display("./view/home/index.tpl")
	//con.Display("./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
	//con.DisplayWithFuncs(template.FuncMap{"look": funcs.Lookup}, "./view/info.tpl", "./view/header.tpl", "./view/footer.tpl")
}

//Sec 一个测试方法实际无用
func (con *IndexController) Sec() {
	//设置context-type
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//方式一
	//fmt.Fprintf(i.W(), "index")
	//方式二 ,fmt.Sprintf是格式化字符串
	io.WriteString(con.W(), fmt.Sprintf("sec method 直接返回字符串"))
}
