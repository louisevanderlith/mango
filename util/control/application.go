package control

import (
	uuid "github.com/nu7hatch/gouuid"
)

type Application struct {
	Name       string
	IP         string
	Location   string
	InstanceID uuid.UUID
	Roles      ActionMap
}

func NewApplication(appName, ip, location string, instanceID uuid.UUID) *Application {
	result := new(Application)
	result.Name = appName
	result.InstanceID = instanceID
	result.Roles = make(ActionMap)

	return result
}

func (a *Application) SetRoles(roles ActionMap) {
	a.Roles = roles
}
