package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/auto/controllers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"
)

func Setup(s *util.Service) {
	ctrlmap := EnableFilter(s)

	beego.Router("/", controllers.NewHomeCtrl(ctrlmap))
}

func EnableFilter(s *util.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
