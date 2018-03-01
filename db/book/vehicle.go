package book

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Vehicle struct {
	db.Record
	VIN            *VIN
	ManufacturerID int64
	ModelID        int64
	BodyStyle      string
	Doors          int
	EngineModel    string
	EngineSize     int
	Trim           string
	Transmission   string
	Gears          int
	Extra          string
}

func (o Vehicle) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
