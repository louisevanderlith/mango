package funds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Experience struct {
	db.Record
	Hero   *Hero
	Type   ExperienceType
	Points int
}

func (o Experience) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
