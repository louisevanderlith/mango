package routers

import (
	"github.com/louisevanderlith/mango/app/shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
