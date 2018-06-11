package funds

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Transaction struct {
	db.Record
	FromUserID  int64
	ToUserID    int64
	Amount      int64
	Requisition *Requisition `orm:"rel(fk)"`
}

func (o Transaction) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
