package model

//Area 国家行政区划相关
type Area struct {
	ID           int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"`             //自增ID
	Aid          int64   `gorm:"column:aid;not null"`                              //国家地区ID
	Pid          int64   `gorm:"column:pid;not null"`                              //父PID,确定归属
	Level        int     `gorm:"column:level;not null"`                            //层级 0:国家 1：省  2：市 3：区/县  4:镇  5：村委会
	Name         string  `gorm:"type:varchar(50);column:name;not null"`            //名字
	SimName      string  `gorm:"type:varchar(50);column:sim_name;not null"`        //名字简写
	NamesPath    string  `gorm:"type:varchar(200);column:names_path;not null"`     //层级全名
	SimNamesPath string  `gorm:"type:varchar(200);column:sim_names_path;not null"` //缩写层级全名
	TelCode      string  `gorm:"type:varchar(50);column:tel_code;not null"`        //电话区号
	ZipCode      string  `gorm:"type:varchar(50);column:zip_code;not null"`        //邮编
	Pinyin       string  `gorm:"type:varchar(100);column:pinyin;not null"`         //拼音
	SimPinyin    string  `gorm:"type:varchar(100);column:sim_pinyin;not null"`     //拼音缩写
	FirstPinyin  string  `gorm:"type:varchar(50);column:first_pinyin;not null"`    //拼音首字母
	Lng          float32 `gorm:"type:decimal(18,7);column:lng;not null"`           //经度
	Lat          float32 `gorm:"type:decimal(18,7);column:lat;not null"`           //维度
	LoadChild    int     `gorm:"column:load_child;not null"`                       //是否加载完子级,默认false
	URL          string  `gorm:"type:varchar(100);column:url;not null"`            //官方url地址
}

//TableName 更改数据库表名
func (Area) TableName() string {
	return "area_asia"
}
