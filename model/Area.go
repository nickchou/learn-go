package model

//Area 国家行政区划相关
type Area struct {
	ID           int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"` //自增ID
	Aid          int64   `gorm:"column:aid;not null"`                  //国家地区ID
	Pid          int64   `gorm:"column:pid;not null"`                  //父PID,确定归属
	Level        int     `gorm:"column:level;not null"`                //层级 0:国家 1：省  2：市 3：区/县  4:镇  5：村委会
	Name         string  `gorm:"column:name;not null"`                 //名字
	SimName      string  `gorm:"column:sim_name;not null"`             //名字简写
	NamesPath    string  `gorm:"column:names_path;not null"`           //层级全名
	SimNamesPath string  `gorm:"column:sim_names_path;not null"`       //缩写层级全名
	TelCode      string  `gorm:"column:tel_code;not null"`             //电话区号
	ZipCode      string  `gorm:"column:zip_code;not null"`             //邮编
	Pinyin       string  `gorm:"column:pinyin;not null"`               //拼音
	SimPinyin    string  `gorm:"column:sim_pinyin;not null"`           //拼音缩写
	FirstPinyin  string  `gorm:"column:first_pinyin;not null"`         //拼音首字母
	Lng          float32 `gorm:"column:lng;not null"`                  //经度
	Lat          float32 `gorm:"column:lat;not null"`                  //维度
	URL          string  `gorm:"column:url;not null"`                  //官方url地址
}

//TableName 更改数据库表名
func (Area) TableName() string {
	return "area_asia"
}
