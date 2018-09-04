package main

import (
	_ "blog/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db_name := beego.AppConfig.String("db_name")
	db_password := beego.AppConfig.String("db_password")
	db_user := beego.AppConfig.String("db_user")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSourceName := db_user + ":" + db_password + "@/"+db_name+"?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dataSourceName)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()
}

