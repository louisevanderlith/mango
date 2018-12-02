package enums

import (
	"strings"
)

type RoleType = int

const (
	Admin RoleType = iota
	Owner
	User
	Unknown
)

var roleTypes = [...]string{
	"Admin",
	"Owner",
	"User",
	"Unknown"}

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
