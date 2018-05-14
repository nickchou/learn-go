package app

import (
	"net/http"
)

//IContext 定义一个接口
type IContext interface {
	Config(w http.ResponseWriter, r *http.Request)
}

//Context 结构体
type Context struct {
	w http.ResponseWriter
	r *http.Request
}

//Config 配置ResponseWriter和Request
func (c *Context) Config(w http.ResponseWriter, r *http.Request) {
	c.w = w
	c.r = r
}
