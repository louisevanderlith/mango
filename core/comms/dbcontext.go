package comms

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Messages husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Messages: husk.NewTable(new(Message)),
	}
}
