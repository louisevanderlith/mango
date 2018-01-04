package routers

import (
	"github.com/louisevanderlith/mango/app/logbook/controllers"
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.DefaultController{})
}

func setupMapping() {
	uploadMap := make(control.MethodMap)
	uploadMap["GET"] = enums.User

	control.AddControllerMap("/", uploadMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
