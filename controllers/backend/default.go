package backend

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Index()  {
	this.Layout = "backend/layout.html"
	this.TplName = "backend/index.html"
}
