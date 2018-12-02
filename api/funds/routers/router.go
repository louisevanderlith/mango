// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango/api/funds/controllers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/util/control"
)

func Setup(s *util.Service) {
	ctrlmap := EnableFilter(s)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/credit",
			beego.NSInclude(
				controllers.NewCreditCtrl(ctrlmap),
			),
		),
		beego.NSNamespace("/requisition",
			beego.NSInclude(
				controllers.NewRequisitionCtrl(ctrlmap),
			),
		),
	)
	beego.AddNamespace(ns)
}

func EnableFilter(s *util.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)
	emptyMap["GET"] = enums.User
	emptyMap["POST"] = enums.User
	emptyMap["PUT"] = enums.User

	ctrlmap.Add("/credit", emptyMap)
	ctrlmap.Add("/register", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
