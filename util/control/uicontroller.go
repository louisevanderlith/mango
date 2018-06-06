package control

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/util"
)

type UIController struct {
	APIController
	HasScript  bool
	ScriptName string
}

func (ctrl *UIController) Prepare() {
	ctrl.Layout = "master.html"
}

func (ctrl *UIController) Setup(name string) {
	ctrl.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	ctrl.Data["HasScript"] = true
	ctrl.Data["ScriptName"] = name + ".entry.js"
	ctrl.Data["InstanceKey"] = util.GetInstanceKey()
	ctrl.Data["RunModeDEV"] = beego.BConfig.RunMode == "dev"
}

func (ctrl *UIController) Serve(err error, data interface{}) {
	if err != nil {
		ctrl.Ctx.Output.SetStatus(500)
		ctrl.Data["error"] = err
	} else {
		ctrl.Data["data"] = data
	}
}
