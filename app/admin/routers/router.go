package routers

import (
	"github.com/louisevanderlith/mango/app/admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/comms", &controllers.CommsController{})
}
