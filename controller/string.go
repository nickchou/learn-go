package controller

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/nickchou/learn-go/app"
	//"github.com/nickchou/gocode/dataaccess"
)

//StringController 字符串测试控制器
type StringController struct {
	app.App
}

//Index 日历控制器默认的执行方法
func (con *StringController) Index() {
	con.W().Header().Set("content-type", "text/html; charset=utf-8")
	//方式一：bytes.Buffer
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("当前日期:%v<br/>", time.Now().Format("2006-01-02")))
	io.WriteString(con.W(), buffer.String())
	//方式二：strings.Builder 这个是go1.10的版本
	//参考文章 https://studygolang.com/articles/12796?fr=sidebar
	var sb strings.Builder
	sb.WriteString("hahaha<br/>")
	sb.WriteString(fmt.Sprintf("当前时间:%v<br/>", time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")))
	io.WriteString(con.W(), sb.String())
}
