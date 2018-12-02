package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/pkg/enums"
)

type Role struct {
	ApplicationName string
	Description     enums.RoleType
}

func (o Role) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
