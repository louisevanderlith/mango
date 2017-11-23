package util

type UIController struct {
	SecureController
	HasScript  bool
	ScriptName string
}

func (ctrl *UIController) Prepare() {
	ctrl.SecureController.Prepare()
	ctrl.Layout = "master.html"
}

func (ctrl *UIController) Setup(name string) {
	ctrl.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	ctrl.Data["HasScript"] = true
	ctrl.Data["ScriptName"] = name + ".entry.js"
}