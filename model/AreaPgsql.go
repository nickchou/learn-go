package model

//AreaPgsql 测试PG ltree性能
type AreaPgsql struct {
	ID       int64  `gorm:"column:id;type:serial;primary_key"`               //自增ID
	Name     string `gorm:"column:name;size:50;default:'name';not null"`     //节点路径
	NodePath string `gorm:"column:nodepath;type:ltree;default:'0';not null"` //节点路径
}

//TableName 更改数据库表名
func (AreaPgsql) TableName() string {
	return "test_area"
}
