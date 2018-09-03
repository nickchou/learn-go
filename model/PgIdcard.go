package model

//PgIdcard 测试PG ltree性能
type PgIdcard struct {
	Nid      int    `xorm:"'nid' not null BIGINT pk autoincr"`
	OID      int    `xorm:"'oid' BIGINT not null"`
	Name     string `xorm:"'name' NVARCHAR(20)"`
	CardType string `xorm:"'cardtype' NVARCHAR(10)"`
	CardNum  string `xorm:"'cardnum' NVARCHAR(10)"`
	Gender   string `xorm:"'gender' NVARCHAR(10)"`
	Birthday string `xorm:"'birthday' NVARCHAR(20)"`
	Mobile   string `xorm:"'mobile' NVARCHAR(20)"`
	EMail    string `xorm:"'email' NVARCHAR(50)"`
	Address  string `xorm:"'address' NVARCHAR(200)"`
}

//TableName 更改数据库表名
func (PgIdcard) TableName() string {
	return "idcard"
}
