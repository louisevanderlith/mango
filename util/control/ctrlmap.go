package control

import "github.com/louisevanderlith/mango/util/enums"

//Action map[:action]:roleType
type ActionMap map[string]enums.RoleType

//ControllerMap is used to assign Priveliges to Actions
type ControllerMap struct {
	applicationName string
	mapping         map[string]ActionMap
}

var controllerMap *ControllerMap

func CreateControllerMap(appName string) {
	result := ControllerMap{}
	result.applicationName = appName
	result.mapping = make(map[string]ActionMap)

	controllerMap = &result
}

// AddControllerMap is used to specify the permissions required for a controller's actions.
func AddControllerMap(path string, actionMap ActionMap) {
	controllerMap.mapping[path] = actionMap
}

func GetRequiredRole(path, action string) enums.RoleType {
	result := enums.Unknown

	if actionMap, hasCtrl := controllerMap.mapping[path]; hasCtrl {
		roleType, hasAction := actionMap[action]

		if hasAction {
			result = roleType
		}
	}

	return result
}
