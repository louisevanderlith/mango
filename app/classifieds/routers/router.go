package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/classifieds/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
