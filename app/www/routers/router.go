package routers

import (
	"github.com/louisevanderlith/mango/app/www/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/:siteName", &controllers.DefaultController{})
}
