package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:AboutController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:AboutController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:HeaderController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:HeaderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:PortfolioController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:PortfolioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:pageData[A-Z](?:_?[0-9]+)*`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:ProfileController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:site`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SocialController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SocialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
