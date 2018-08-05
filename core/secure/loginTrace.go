package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/nu7hatch/gouuid"
)

type LoginTrace struct {
	Location        string `hsk:"null;size(128)"`
	IP              string `hsk:"null;size(50)"`
	Allowed         bool   `hsk:"default(true)"`
	InstanceID      uuid.UUID
	ApplicationName string `hsk:"size(20)"`
}

func (o LoginTrace) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
