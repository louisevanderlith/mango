package things

import "github.com/louisevanderlith/mango/util"

type Model struct {
	util.Record
	ManufacturerID *Manufacturer
	Name           string
}
