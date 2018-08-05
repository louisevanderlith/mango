package funds

import "github.com/louisevanderlith/husk"

type Transaction struct {
	FromUserID  int64
	ToUserID    int64
	Amount      int64
	Requisition *Requisition
}

func (o Transaction) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
