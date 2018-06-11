package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/shop/controllers"
	"github.com/louisevanderlith/mango/util/control"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.DefaultController{})
}

func setupMapping() {
	uploadMap := make(control.MethodMap)

	control.AddControllerMap("/", uploadMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
