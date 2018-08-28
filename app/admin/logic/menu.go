package logic

import (
	"github.com/louisevanderlith/mango/util/control"
)

func GetMenu(path string) *control.Menu {
	return getItems(path)
}

func getItems(path string) *control.Menu {
	result := control.NewMenu(path)

	result.AddItem("#", "Artifact API", "fa-cloud", artifactChlidren(path))
	result.AddItem("#", "Comms API", "fa-fax", commsChildren(path))
	result.AddItem("#", "Folio API", "fa-users", folioChildren(path))
	result.AddItem("#", "Secure API", "fa-user-secret", secureChildren(path))

	return result
}

func artifactChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/uploads", "Uploads", "fa-newspaper-o", nil)

	return children
}

func commsChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/comms", "Messages", "fa-newspaper-o", nil)

	return children
}

func folioChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/sites", "Sites", "fa-globe", nil)

	return children
}

func secureChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/user", "Users", "fa-user", nil)

	return children
}
