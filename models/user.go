package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `orm:"pk"`
	Username string `orm:"size(100);index"`
	Password string
	Email string `orm:"default('');size(100);index"`
	Phone string `orm:"default('');size(11)"`
}
/**
 * 自定义表名
 */
func (this *User) TableName() string {
	return "user"
}

func init()  {
	orm.RegisterModel(new(User))
}