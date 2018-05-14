package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/nickchou/learn-go/app"
)

//ZhuzherController zhuzher控制器
type ZhuzherController struct {
	app.App
}

//Project 获取zhuzher的project信息
func (con *ZhuzherController) Project() {
	var bf bytes.Buffer
	//table相关
	bf.WriteString("<style>table { border-collapse:collapse; } table,th, td  {  border: 1px solid black;</style>")
	bf.WriteString("zhuzher project list<br/>")
	bf.WriteString("<table>")
	bf.WriteString("<tr><td>index</td><td>city</td><td>project code</td><td>project name</td><td>stage</td></tr>")

	//zhuzher所有的城市code 66 items,from=https://flyingdutchman.4009515151.com/api/zhuzher/cities
	var citys = [...]int{210300, 469029, 110000, 220100, 430100, 320400, 510100, 500000, 210200, 469007, 441900, 440600, 210400, 350100, 440100, 520100, 230100, 460100, 330100, 340100, 441300, 330400, 220200, 370100, 140700, 530100, 131000, 469024, 360100, 511300, 320100, 450100, 320600, 330200, 350300, 370200, 441800, 130300, 350500, 350400, 460200, 310000, 330600, 210100, 440300, 320500, 140100, 331000, 130200, 120000, 330300, 420100, 340200, 650100, 320200, 350200, 610100, 320300, 321000, 370600, 210800, 350600, 410100, 321100, 442000, 440400}
	index := 0
	//遍历所有城市
	for _, city := range citys {
		//根据city_code拿到project信息
		resp, err := http.Get(fmt.Sprintf("https://flyingdutchman.4009515151.com/api/zhuzher/projects?city_code=%v", city))
		//判断err
		if err != nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			//json反序列化项目信息
			var pro Project
			if err := json.Unmarshal(body, &pro); err == nil {
				if pro.Code == 0 {
					//遍历json的返回结果
					for _, p := range pro.Results {
						index++
						bf.WriteString(fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>", index, city, p.Code1, p.Name, p.Stage))
					}
				}
			} else {
				fmt.Println(err)
			}
		}
	}
	//table
	bf.WriteString("</table>")
	io.WriteString(con.W(), bf.String())
}

//Project zhuzher的小区实体信息
type Project struct {
	Code    int32    `json:"code"`
	Results []Result `json:"result"`
}

//Result zhuzher接口返回的Result实体定义
type Result struct {
	Code1 string `json:"code"`
	Name  string `json:"name"`
	Stage int32  `json:"stage"`
}
