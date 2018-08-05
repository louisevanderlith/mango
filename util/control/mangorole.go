package control

import (
	"github.com/louisevanderlith/mango/util/enums"
)

type mangorole map[string]enums.RoleType

func MakeMangorole() mangorole {
	return make(mangorole)
}

func (m mangorole) AddAppRole(appName string, role enums.RoleType) {
	m[appName] = role
}

func (m mangorole) GetAppRole(appName string) enums.RoleType {
	result := enums.Unknown

	role, ok := m[appName]

	if ok {
		result = role
	}

	return result
}
