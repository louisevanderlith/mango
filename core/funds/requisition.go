package funds

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/enums"
)

type Requisition struct {
	Reference  string
	Status     enums.RequisitionStatus
	ClientID   int64
	SupplierID int64
	Total      int64
	LineItems  []LineItem
}

func (o Requisition) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
