package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/www/controllers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/control"
)

func Setup(s *util.Service) {
	ctrlmap := control.CreateControlMap(s)
	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
	beego.Router("/:siteName", controllers.NewDefaultCtrl(ctrlmap))
}
