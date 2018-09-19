package control

import (
	"fmt"
	"log"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

//Action map[:action]:roleType
type ActionMap map[string]enums.RoleType

//ControllerMap is used to assign Priveliges to Actions
type ControllerMap struct {
	service *util.Service
	mapping map[string]ActionMap
}

var controlMap *ControllerMap

func CreateControlMap(service *util.Service, appName string) *ControllerMap {
	result := &ControllerMap{}
	result.service = service
	result.mapping = make(map[string]ActionMap)

	controlMap = result

	return result
}

// AddControllerMap is used to specify the permissions required for a controller's actions.
func (m *ControllerMap) Add(path string, actionMap ActionMap) {
	log.Printf("Add(%s, %+v)\n", path, actionMap)
	m.mapping[path] = actionMap
}

func (m *ControllerMap) GetRequiredRole(path, action string) enums.RoleType {
	fmt.Printf("path: %s, actions: %s", path, action)

	result := enums.Unknown

	if actionMap, hasCtrl := m.mapping[path]; hasCtrl {
		log.Printf("Action map: %+v\n", actionMap)
		roleType, hasAction := actionMap[action]

		if hasAction {
			result = roleType
		}
	}

	return result
}

func (m *ControllerMap) GetInstanceKey() husk.Key {
	return m.service.InstanceKey
}
