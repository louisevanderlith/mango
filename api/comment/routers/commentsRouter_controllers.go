package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/comment/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:type/:nodeID[get]`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
