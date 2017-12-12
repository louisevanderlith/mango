package routers

import (
	"github.com/louisevanderlith/mango/app/admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/comms", &controllers.CommsController{})
	beego.Router("/manufacturer", &controllers.ManufacturerController{})
	beego.Router("/model", &controllers.ModelController{})
	beego.Router("/subcategory", &controllers.SubCategoryController{})
}
