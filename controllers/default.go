package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsIndex"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/index.html"
}
