package backend

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type LoginController struct {
	beego.Controller
}
/**
 * 显示登录信息
 */
func (this *LoginController) Login() {
	this.Data["IsLogin"] = true
	this.Layout = "backend/layout.html"
	this.TplName = "backend/login.html"
}
/**
 * 显示登录信息
 */
func (this *LoginController) CheckLogin() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	if username == "haojie" && password == "123456" {
		this.SetSession("name", "haojie");
		this.SetSession("isLogin", true);
		o := orm.NewOrm()
		o.Using("default") // 默认使用 default，你可以指定为其他数据库
		user := new(models.User)
		user.Username = username
		user.Password = "123456"
		i, e := o.Insert(&user)
		if e != nil {
			log.Fatal(e)
		}
		this.Ctx.WriteString(strconv.Itoa(int(i)))
		this.Ctx.WriteString("登录成功")
	} else {
		this.Ctx.WriteString("你没有登录")
	}

	return
}
