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
)

func init() {
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
