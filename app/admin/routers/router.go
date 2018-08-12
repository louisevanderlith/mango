package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/admin/controllers"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/comms", &controllers.CommsController{})
	beego.Router("/manufacturer", &controllers.ManufacturerController{})
	beego.Router("/model", &controllers.ModelController{})
	beego.Router("/subcategory", &controllers.SubCategoryController{})
	beego.Router("/site", &controllers.SiteController{})
	beego.Router("/site/:id([0-9]+)", &controllers.SiteController{}, "get:GetEdit")
}

func setupMapping() {
	appName := beego.BConfig.AppName
	control.CreateControllerMap(appName)
	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Admin
	emptyMap["GET"] = enums.Admin

	control.AddControllerMap("/", emptyMap)
	control.AddControllerMap("/category", emptyMap)
	control.AddControllerMap("/comms", emptyMap)
	control.AddControllerMap("/manufacturer", emptyMap)
	control.AddControllerMap("/model", emptyMap)
	control.AddControllerMap("/subcategory", emptyMap)
	control.AddControllerMap("/site", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
