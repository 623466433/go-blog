package frontend

import "github.com/astaxie/beego"

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index() {
	this.Data["IsAbout"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/about.html"
}
