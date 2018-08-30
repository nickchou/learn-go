package model

//MsIdcaid 测试PG ltree性能
type MsIdcaid struct {
	Id        int    `xorm:"not null pk INT(4) autoincr"`
	Name      string `xorm:"NVARCHAR(4000)"`
	Cardno    string `xorm:"NVARCHAR(4000)"`
	Descriot  string `xorm:"NVARCHAR(4000)"`
	Ctftp     string `xorm:"NVARCHAR(4000)"`
	Ctfid     string `xorm:"NVARCHAR(4000)"`
	Gender    string `xorm:"NVARCHAR(4000)"`
	Birthday  string `xorm:"NVARCHAR(4000)"`
	Address   string `xorm:"NVARCHAR(4000)"`
	Zip       string `xorm:"NVARCHAR(4000)"`
	Dirty     string `xorm:"NVARCHAR(4000)"`
	District1 string `xorm:"NVARCHAR(4000)"`
	District2 string `xorm:"NVARCHAR(4000)"`
	District3 string `xorm:"NVARCHAR(4000)"`
	District4 string `xorm:"NVARCHAR(4000)"`
	District5 string `xorm:"NVARCHAR(4000)"`
	District6 string `xorm:"NVARCHAR(4000)"`
	Firstnm   string `xorm:"NVARCHAR(4000)"`
	Lastnm    string `xorm:"NVARCHAR(4000)"`
	Duty      string `xorm:"NVARCHAR(4000)"`
	Mobile    string `xorm:"NVARCHAR(4000)"`
	Tel       string `xorm:"NVARCHAR(4000)"`
	Fax       string `xorm:"NVARCHAR(4000)"`
	Email     string `xorm:"NVARCHAR(4000)"`
	Nation    string `xorm:"NVARCHAR(4000)"`
	Taste     string `xorm:"NVARCHAR(4000)"`
	Education string `xorm:"NVARCHAR(4000)"`
	Company   string `xorm:"NVARCHAR(4000)"`
	Ctel      string `xorm:"NVARCHAR(4000)"`
	Caddress  string `xorm:"NVARCHAR(4000)"`
	Czip      string `xorm:"NVARCHAR(4000)"`
	Family    string `xorm:"NVARCHAR(4000)"`
	Version   string `xorm:"NVARCHAR(4000)"`
}

//TableName 更改数据库表名
func (MsIdcaid) TableName() string {
	return "cdsgus"
}
