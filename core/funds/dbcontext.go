package funds

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Transactions husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Transactions: husk.NewTable(new(Transaction)),
	}
}
