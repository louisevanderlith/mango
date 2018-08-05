package control

import (
	"github.com/louisevanderlith/mango/util/enums"
	uuid "github.com/nu7hatch/gouuid"
)

type Application struct {
	name       string
	instanceID uuid.UUID
	role       enums.RoleType
}

func NewApplication(appName string, instanceID uuid.UUID) *Application {
	result := new(Application)
	result.name = appName
	result.instanceID = instanceID
	// result.role can only be set once login has complete

	return result
}

func (a *Application) SetRole(role enums.RoleType) {
	a.role = role
}
