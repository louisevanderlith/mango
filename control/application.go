package control

type Application struct {
	Name       string
	IP         string
	Location   string
	InstanceID string
	Roles      ActionMap
}

func NewApplication(appName, ip, location string, instanceID string) *Application {
	result := new(Application)
	result.Name = appName
	result.InstanceID = instanceID
	result.Roles = make(ActionMap)

	return result
}

func (a *Application) SetRoles(roles ActionMap) {
	a.Roles = roles
}
