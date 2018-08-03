package model

//TestArea 测试PG ltree性能
type TestArea struct {
	ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`          //自增ID
	Name     string `gorm:"column:name;size:50;default:'name';not null"`   //节点路径
	NodePath string `gorm:"column:nodepath;size:200;default:'0';not null"` //节点路径
}

//TableName 更改数据库表名
func (TestArea) TableName() string {
	return "test_area"
}
