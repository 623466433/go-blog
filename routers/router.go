package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{}, "get:Index")
	beego.Router("/books", &controllers.BooksController{}, "get:Index")
	beego.Router("/message", &controllers.MessageController{}, "get:Index")
	beego.Router("/timeline", &controllers.TimelineController{}, "get:Index")
	beego.Router("/article", &controllers.ArticleController{}, "get:Index")
}
