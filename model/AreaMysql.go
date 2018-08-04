package model

//AreaMysql 测试PG ltree性能，mysql字符串类型长度用size设置好像有问题
type AreaMysql struct {
	ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`                   //自增ID
	Name     string `gorm:"column:name;type:varchar(50);default:'name';not null"`   //节点路径
	NodePath string `gorm:"column:nodepath;type:varchar(200);default:'0';not null"` //节点路径
}

//TableName 更改数据库表名
func (AreaMysql) TableName() string {
	return "test_area"
}
