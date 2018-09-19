// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/api/artifact/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func Setup(instanceKey husk.Key) {
	EnableFilters(instanceKey)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			),
		),
	)

	beego.AddNamespace(ns)
}

func EnableFilters(instanceKey husk.Key) {
	appName := beego.BConfig.AppName

	ctrlmap := control.CreateControlMap(appName)

	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Owner

	ctrlmap.Add("/upload", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
}
