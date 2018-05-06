package funds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type Requisition struct {
	db.Record
	Reference  string
	Status     enums.RequisitionStatus
	ClientID   int64
	SupplierID int64
	Total      int64
	LineItems  []*LineItem `orm:"reverse(many)"`
}

func (o Requisition) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
