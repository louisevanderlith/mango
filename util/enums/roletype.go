package enums

import (
	"strings"
)

type RoleType = int

const (
	Admin RoleType = iota
	Owner
	User
)

var roleTypes = [...]string{
	"Admin",
	"Owner",
	"User"}

/*
func (r RoleType) String() string {
	return roleTypes[r]
}*/

func GetRoleType(name string) RoleType {
	var result RoleType

	for k, v := range roleTypes {
		if strings.ToUpper(name) == v {
			result = RoleType(k)
			break
		}
	}

	return result
}
