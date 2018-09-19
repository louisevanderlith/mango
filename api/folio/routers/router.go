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

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/util/control"
	"github.com/louisevanderlith/mango/util/enums"
)

func init() {
	setupMapping()

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/profile",
			beego.NSInclude(
				&controllers.ProfileController{},
			),
		),
		beego.NSNamespace("/profile/about",
			beego.NSInclude(
				&controllers.AboutController{},
			),
		),
		beego.NSNamespace("/profile/header",
			beego.NSInclude(
				&controllers.HeaderController{},
			),
		),
		beego.NSNamespace("/profile/portfolio",
			beego.NSInclude(
				&controllers.PortfolioController{},
			),
		),
		beego.NSNamespace("/profile/social",
			beego.NSInclude(
				&controllers.SocialController{},
			),
		),
	)

	beego.AddNamespace(ns)
}

func setupMapping() {
	appName := beego.BConfig.AppName
	ctrlmap := control.CreateControlMap(appName)

	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Owner
	emptyMap["PUT"] = enums.Owner

	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("profile/about", emptyMap)
	ctrlmap.Add("profile/header", emptyMap)
	ctrlmap.Add("profile/portfolio", emptyMap)
	ctrlmap.Add("profile/social", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
}
