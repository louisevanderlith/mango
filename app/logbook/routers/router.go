package routers

import (
	"github.com/louisevanderlith/mango/app/logbook/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
