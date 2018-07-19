package controller

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/comm"
	//"github.com/nickchou/gocode/dataaccess"
)

//CalendarController 日历控制器
type CalendarController struct {
	app.App
}

//Index 日历控制器默认的执行方法
func (con *CalendarController) Index() {

	//方式一
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2001-02-02",
		comm.ToLunarDate("2001-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2009-02-02",
		comm.ToLunarDate("2009-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2010-02-02",
		comm.ToLunarDate("2010-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2017-02-02",
		comm.ToLunarDate("2017-02-02").Format("2006-01-02")))
	buffer.WriteString(fmt.Sprintf("阳转阴历:%v -> %v <br/>", "2017-09-26",
		comm.ToLunarDate("2017-09-26").Format("2006-01-02")))
	//respnse
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(con.W(), buffer.String())
	//方式二
	//fmt.Fprintf(con.W(), lunarcal.ToLunarDate("2009-02-02").Format("2006-01-02")+"<br/>")
}
