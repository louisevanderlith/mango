package control

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/enums"
)

type Cookies struct {
	UserKey   husk.Key
	Username  string
	UserRoles map[string]enums.RoleType
	IP        string
	Location  string
}

func NewCookies(userkey *husk.Key, username, ip, location string, roles map[string]enums.RoleType) *Cookies {
	return &Cookies{
		UserKey:   *userkey,
		Username:  username,
		IP:        ip,
		Location:  location,
		UserRoles: roles,
	}
}
