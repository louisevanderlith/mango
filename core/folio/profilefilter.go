package folio

import (
	"github.com/louisevanderlith/husk"
)

type profileFilter func(obj *Profile) bool

func (f profileFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Profile))
}

func byName(name string) profileFilter {
	return func(obj *Profile) bool {
		return obj.Title == name
	}
}
