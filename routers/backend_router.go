package routers

import (
	"blog/controllers/backend"
	"github.com/astaxie/beego"
)

func initRouter() *beego.Namespace {
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &backend.MainController{}, "get:Index"),
		beego.NSRouter("/login", &backend.LoginController{}, "get:Login"),
		beego.NSRouter("/login", &backend.LoginController{}, "post:CheckLogin"),
	)

	return ns
}

func init()  {
	beego.AddNamespace(initRouter())
}