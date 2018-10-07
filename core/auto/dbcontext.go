package auto

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Adverts husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Adverts: husk.NewTable(new(Advert)),
	}
}
