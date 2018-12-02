package artifact

import (
	"github.com/louisevanderlith/husk"
)

type uploadFilter func(obj *Upload) bool

func (f uploadFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Upload))
}

func bySize(size int64) uploadFilter {
	return func(obj *Upload) bool {
		return obj.Size >= size
	}
}
