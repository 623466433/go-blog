package routers

import (
	"blog/controllers/frontend"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由
	beego.Router("/", &frontend.MainController{})
	beego.Router("/about", &frontend.AboutController{}, "get:Index")
	beego.Router("/books", &frontend.BooksController{}, "get:Index")
	beego.Router("/message", &frontend.MessageController{}, "get:Index")
	beego.Router("/timeline", &frontend.TimelineController{}, "get:Index")
	beego.Router("/article", &frontend.ArticleController{}, "get:Index")
}
