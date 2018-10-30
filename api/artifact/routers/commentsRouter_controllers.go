package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"],
		beego.ControllerComments{
			Method: "GetByID",
			Router: `/:uploadKey`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/all/:pagesize`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/louisevanderlith/mango/api/artifact/controllers:UploadController"],
		beego.ControllerComments{
			Method: "GetFileBytes",
			Router: `/file/:uploadKey`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
