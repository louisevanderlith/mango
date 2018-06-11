package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/www/controllers"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/:siteName", &controllers.DefaultController{})
}
