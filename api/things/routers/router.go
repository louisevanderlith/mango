// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango/api/things/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func init() {
	setupMapping()

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/manufacturer",
			beego.NSInclude(
				&controllers.ManufacturerController{},
			),
		),
		beego.NSNamespace("/model",
			beego.NSInclude(
				&controllers.ModelController{},
			),
		),
		beego.NSNamespace("/subcategory",
			beego.NSInclude(
				&controllers.SubCategoryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

func setupMapping() {
	uploadMap := make(control.MethodMap)
	uploadMap["POST"] = enums.Admin

	control.AddControllerMap("/category", uploadMap)
	control.AddControllerMap("/manufacturer", uploadMap)
	control.AddControllerMap("/model", uploadMap)
	control.AddControllerMap("/subcategory", uploadMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
}
