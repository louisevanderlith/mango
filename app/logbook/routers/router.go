package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/logbook/controllers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func Setup(s *util.Service) {
	ctrlmap := EnableFilter(s)

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
}

func EnableFilter(s *util.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)
	emptyMap["GET"] = enums.User

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
