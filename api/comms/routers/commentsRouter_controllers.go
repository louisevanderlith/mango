package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comms/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comms/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comms/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comms/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:pageData[A-Z](?:_?[0-9]+)*`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
