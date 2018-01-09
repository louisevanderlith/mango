package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/folio/controllers:SiteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:site`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
