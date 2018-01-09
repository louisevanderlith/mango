package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ManufacturerController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ManufacturerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ManufacturerController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ManufacturerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ModelController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ModelController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ModelController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:ModelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:SubCategoryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/things/controllers:SubCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
