package book

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type ServiceItem struct {
	db.Record
	Service     *Service
	Code        string
	Description string
}

func (o ServiceItem) Validate() (bool, error) {
	return util.ValidateStruct(o)
}
