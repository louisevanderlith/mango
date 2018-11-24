package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:CreditController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:CreditController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/funds/controllers:RequisitionController"],
		beego.ControllerComments{
			Method: "GetByID",
			Router: `/:requisitionID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
