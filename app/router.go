package app

import (
	"reflect"
	"strings"
)

//mapping 反射
var mapping = make(map[string]reflect.Type)

func router(pattern string, t reflect.Type) {
	mapping[strings.ToLower(pattern)] = t
}

//Router 指定control注册
func Router(pattern string, app IApp) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	router(pattern, refT)
}

//AutoRouter 自动根据control来注册
func AutoRouter(app IApp) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	refName := strings.TrimSuffix(strings.ToLower(refT.Name()), "controller")
	router(refName, refT)
}
