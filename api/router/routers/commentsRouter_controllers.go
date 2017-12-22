package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/router/controllers:DiscoveryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/router/controllers:DiscoveryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/router/controllers:DiscoveryController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/router/controllers:DiscoveryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:appID/:serviceName/:clean`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
