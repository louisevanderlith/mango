package control

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango"
)

type UIController struct {
	APIController
	settings mango.ThemeSetting
}

func (ctrl *UIController) SetTheme(settings mango.ThemeSetting) {
	ctrl.settings = settings
}

func (ctrl *UIController) Prepare() {
	defer ctrl.APIController.Prepare()

	ctrl.Layout = "_shared/master.html"

	output := ctrl.Ctx.Output

	//output.Header("Content-Security-Policy", "script-src 'self' https:")
	output.Header("X-Frame-Options", "SAMEORIGIN")
	output.Header("X-XSS-Protection", "1; mode=block")
}

func (ctrl *UIController) Setup(name, title string, hasScript bool) {
	ctrl.TplName = fmt.Sprintf("%s.html", name)
	ctrl.applySettings(title)

	ctrl.Data["HasScript"] = hasScript
	ctrl.Data["ScriptName"] = fmt.Sprintf("%s.entry.dart.js", name)
	ctrl.Data["ShowSave"] = false
}

func (ctrl *UIController) EnableSave() {
	ctrl.Data["ShowSave"] = true
}

func (ctrl *UIController) applySettings(title string) {
	ctrl.Data["Title"] = fmt.Sprintf("%s %s", title, ctrl.settings.Name)
	ctrl.Data["LogoKey"] = ctrl.settings.LogoKey
	ctrl.Data["InstanceID"] = ctrl.settings.InstanceID
	ctrl.Data["Host"] = ctrl.settings.Host
	ctrl.Data["Crumbs"] = decipherURL(ctrl.Ctx.Request.URL.RequestURI())
	ctrl.Data["GTag"] = ctrl.settings.GTag

	//User Details
	avoc, err := GetAvoCookie(ctrl.GetMyToken(), ctrl.ctrlMap.GetPublicKeyPath())
	loggedIn := err == nil
	ctrl.Data["LoggedIn"] = loggedIn

	if loggedIn {
		ctrl.Data["Username"] = avoc.Username
	}
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
func (ctrl *UIController) ServeJSON(statuscode int, err error, data interface{}) {
	ctrl.EnableRender = false

	ctrl.APIController.Serve(statuscode, err, data)
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

func (ctrl *UIController) GetMyToken() string {
	return ctrl.Ctx.GetCookie("avosession")
}

func decipherURL(url string) []string {
	var result []string
	qryIndx := strings.Index(url, "?")

	if qryIndx != -1 {
		url = url[:qryIndx]
	}

	parts := strings.Split(url, "/")

	for _, v := range parts {
		if len(v) > 0 {
			result = append(result, v)
		}
	}

	return result
}
