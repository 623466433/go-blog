package frontend

import "github.com/astaxie/beego"

type TimelineController struct {
	beego.Controller
}

func (this *TimelineController) Index() {
	this.Data["IsTimeline"] = true
	this.Layout = "frontend/frontend-layout.html"
	this.TplName = "frontend/timeline.html"
}
