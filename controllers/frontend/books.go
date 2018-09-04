package frontend

import (
	"github.com/astaxie/beego"
)

type BooksController struct {
	beego.Controller
}

func (this *BooksController) Index() {
	this.Data["IsBooks"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/photo.html"
}
