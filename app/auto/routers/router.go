package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/auto/controllers"
	"github.com/louisevanderlith/mango/util/control"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.HomeController{})
}

func setupMapping() {
	appName := beego.BConfig.AppName
	control.CreateControllerMap(appName)
	emptyMap := make(control.ActionMap)

	control.AddControllerMap("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
