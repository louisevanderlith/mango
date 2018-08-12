package funds

import "github.com/louisevanderlith/husk"

type Transaction struct {
	HeroID       int64
	Total        int64
	Requisitions []Requisition
}

func (o Transaction) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func CreateTransaction() {

}

func GetTransactions(heroID int64) {

}
