package model

//Area 国家行政区划相关
type Area struct {
	ID          int     //国家地区ID
	Pid         int     //父PID,确定归属
	level       int     //层级 0:国家 1：省  2：市 3：区/县  4:镇  5：村委会
	Name        string  //名字
	SimName     string  //名字简写
	NamePath    string  //层级全名
	SimNamePath string  //缩写层级全名
	TelCode     string  //电话区号
	ZipCode     string  //邮编
	Pinyin      string  //拼音
	SimPinyin   string  //拼音缩写
	FirstPinyin string  //拼音首字母
	Lng         float32 //经度
	Lat         float32 //维度
}
