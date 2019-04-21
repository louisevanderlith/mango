package control

import (
	"fmt"
)

type UIController struct {
	APIController
	settings ThemeSetting
}

func (ctrl *UIController) SetTheme(settings ThemeSetting) {
	ctrl.settings = settings
}

func (ctrl *UIController) Prepare() {
	defer ctrl.APIController.Prepare()

	ctrl.Layout = "_shared/master.html"

	output := ctrl.Ctx.Output

	//output.Header("Content-Security-Policy", "script-src 'self' https:")
	output.Header("X-Frame-Options", "SAMEORIGIN")
	output.Header("X-XSS-Protection", "1; mode=block")
	output.Header("X-Content-Type-Options", "nosniff")
}

func (ctrl *UIController) Setup(name, title string, hasScript bool) {
	ctrl.TplName = fmt.Sprintf("%s.html", name)
	ctrl.applySettings(title)

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	ctrl.Data["HasScript"] = hasScript
	ctrl.Data["ScriptName"] = fmt.Sprintf("%s.entry.js", name)
}

func (ctrl *UIController) applySettings(title string) {
	ctrl.Data["Title"] = fmt.Sprintf("%s %s", title, ctrl.settings.Name)
	ctrl.Data["LogoKey"] = ctrl.settings.LogoKey
	ctrl.Data["InstanceID"] = ctrl.settings.InstanceID
	ctrl.Data["Host"] = ctrl.settings.Host
}

//Serve sends the response with 'Error' and 'Data' properties.
func (ctrl *UIController) Serve(data interface{}, err error) {
	if err != nil {
		ctrl.Ctx.Output.SetStatus(500)
	}

	ctrl.Data["Error"] = err
	ctrl.Data["Data"] = data
}

//ServeJSON enables JSON Responses on UI Controllers
func (ctrl *UIController) ServeJSON(data interface{}, err error) {
	ctrl.EnableRender = false

	ctrl.APIController.Serve(data, err)
	ctrl.EnableRender = true
}

func (ctrl *UIController) CreateTopMenu(menu *Menu) {
	ctrl.createMenu("TopMenu", menu)
}

func (ctrl *UIController) CreateSideMenu(menu *Menu) {
	ctrl.createMenu("SideMenu", menu)
}

func (ctrl *UIController) createMenu(name string, menu *Menu) {
	ctrl.Data[name] = menu
}
