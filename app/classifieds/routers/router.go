package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/classifieds/controllers"
	"github.com/louisevanderlith/mango/util/control"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/create", &controllers.CreateController{})
	beego.Router("/create/:step", &controllers.CreateController{}, "get:GetStep")
}

func setupMapping() {
	uploadMap := make(control.MethodMap)

	control.AddControllerMap("/", uploadMap)
	control.AddControllerMap("/create", uploadMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
