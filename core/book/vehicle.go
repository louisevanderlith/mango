package book

import (
	"github.com/louisevanderlith/husk"
)

type Vehicle struct {
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

func (o Vehicle) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
