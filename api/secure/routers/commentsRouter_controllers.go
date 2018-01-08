package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"],
		beego.ControllerComments{
			Method: "GetCookie",
			Router: `/avo/:sessionID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/secure/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
