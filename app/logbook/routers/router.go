package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/logbook/controllers"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.DefaultController{})
}

func setupMapping() {
	appName := beego.BConfig.AppName
	ctrlmap := control.CreateControlMap(appName)

	emptyMap := make(control.ActionMap)
	emptyMap["GET"] = enums.User

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
