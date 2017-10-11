package things

import "github.com/louisevanderlith/mango/util"

type Model struct {
	util.BaseRecord
	ManufacturerID *Manufacturer
	Name           string
}
