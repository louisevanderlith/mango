// @APIVersion 1.0.0
// @Title Folio API
// @Description API to control Portfolios
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango/api/folio/controllers"
	"github.com/louisevanderlith/mango/pkg"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/pkg/control"
	"github.com/louisevanderlith/mango/pkg/enums"
)

func Setup(s *util.Service) {
	ctrlmap := EnableFilters(s)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/profile",
			beego.NSInclude(
				controllers.NewProfileCtrl(ctrlmap),
			),
		),
		beego.NSNamespace("/profile/header",
			beego.NSInclude(
				controllers.NewHeaderCtrl(ctrlmap),
			),
		),
		beego.NSNamespace("/profile/portfolio",
			beego.NSInclude(
				controllers.NewPortfolioCtrl(ctrlmap),
			),
		),
		beego.NSNamespace("/profile/social",
			beego.NSInclude(
				controllers.NewSocialCtrl(ctrlmap),
			),
		),
	)

	beego.AddNamespace(ns)
}

func EnableFilters(s *util.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Owner
	emptyMap["PUT"] = enums.Owner

	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("profile/about", emptyMap)
	ctrlmap.Add("profile/header", emptyMap)
	ctrlmap.Add("profile/portfolio", emptyMap)
	ctrlmap.Add("profile/social", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
