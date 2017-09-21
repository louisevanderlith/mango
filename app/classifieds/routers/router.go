package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/classifieds/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
