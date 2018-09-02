package controllers

import "github.com/astaxie/beego"

type MessageController struct {
	beego.Controller
}

func (this *MessageController) Index() {
	this.Data["IsMessage"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/message.html"
}
