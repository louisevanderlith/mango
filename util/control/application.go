package control

import (
	uuid "github.com/nu7hatch/gouuid"
)

type Application struct {
	Name       string
	InstanceID uuid.UUID
	Roles      ActionMap
}

func NewApplication(appName string, instanceID uuid.UUID) *Application {
	result := new(Application)
	result.Name = appName
	result.InstanceID = instanceID
	result.Roles = make(ActionMap)

	return result
}

func (a *Application) SetRoles(roles ActionMap) {
	a.Roles = roles
}
