package logic

import (
	"github.com/louisevanderlith/mango/util/control"
)

type MenuController struct {
	control.UIController
}

func (ctrl *MenuController) Prepare() {
	ctrl.Data["menu"] = getMenu("/")
	ctrl.UIController.Prepare()
}

func (ctrl *MenuController) Setup(name string) {
	ctrl.UIController.Setup(name)
}

type Menu map[int64]*menuItem

type menuItem struct {
	State     bool
	Text      string
	Path      string
	IconClass string
	Children  Menu
}

var _path string

func getMenu(path string) Menu {
	_path = path

	return getItems()
}

func getItems() Menu {
	result := make(Menu)

	result[0] = newItem("Home", "/", "fa-home")
	result[0].setActive()

	result[1] = artifactMenu()
	result[2] = commsMenu()
	result[3] = folioMenu()
	result[4] = userMenu()

	return result
}

func artifactMenu() *menuItem {

	result := newItem("Artifact API", "#", "fa-ban")

	result.Children[0] = newItem("Uploads", "/uploads", "fa-ban")

	result.setActive()
	return result
}

func commsMenu() *menuItem {
	result := newItem("Comms API", "#", "fa-mail")

	result.Children[0] = newItem("Messages", "/comms", "fa-mail")

	result.setActive()
	return result
}

func folioMenu() *menuItem {
	result := newItem("Folio API", "/site", "fa-web")

	result.setActive()
	return result
}

func userMenu() *menuItem {
	result := newItem("Secure API", "#", "fa-lock")

	result.Children[0] = newItem("Users", "/user", "fa-user")

	result.setActive()
	return result
}

func newItem(text, path, iconClass string) *menuItem {
	return &menuItem{
		Text:      text,
		Path:      path,
		IconClass: iconClass,
		State:     false,
		Children:  make(Menu),
	}
}

func (item *menuItem) setActive() {
	if len(item.Children) > 0 {
		for _, v := range item.Children {
			if v.Path == _path {
				v.State = true
				item.State = true
			}
		}
	} else if item.Path == _path {
		item.State = true
	}
}
