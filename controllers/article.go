package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Index() {
	this.Data["IsArticle"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/list.html"
}
