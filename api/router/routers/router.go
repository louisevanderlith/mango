// @APIVersion 1.0.0
// @Title Router API
// @Description API for the Router
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango/api/router/controllers"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/control"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Setup(s *util.Service) {
	ctrlmap := EnableFilter(s)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/discovery",
			beego.NSInclude(
				controllers.NewDiscoveryCtrl(ctrlmap),
			),
		),
		beego.NSNamespace("/memory",
			beego.NSInclude(
				controllers.NewMemoryCtrl(ctrlmap),
			),
		),
	) 

	beego.AddNamespace(ns)
}

func EnableFilter(s *util.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/discovery", emptyMap)
	ctrlmap.Add("/memory", emptyMap)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
