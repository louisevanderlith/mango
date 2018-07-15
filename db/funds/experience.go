package funds

import "github.com/louisevanderlith/husk"

type Experience struct {
	Type   ExperienceType
	Points int
}

func (o Experience) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
