// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango/api/comment/controllers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Setup(service *util.Service) {
	EnableFilters(service)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
	)

	beego.AddNamespace(ns)
}

func EnableFilters(service *util.Service) {
	appName := beego.BConfig.AppName
	ctrlmap := control.CreateControlMap(service, appName)

	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.User
	emptyMap["PUT"] = enums.User

	ctrlmap.Add("/message", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterAPI())

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
}
