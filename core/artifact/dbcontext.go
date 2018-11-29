package artifact

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Uploads husk.Tabler
}

var ctx context

func init() {
	ctx = context{
		Uploads: husk.NewTable(new(Upload)),
	}
}

func Shutdown() {
	ctx.Uploads.Save()
}
