package secure

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Users husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Users: husk.NewTable(new(User)),
	}
}
