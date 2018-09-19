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
	ctrlmap := control.CreateControlMap(appName)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
