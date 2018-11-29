package secure

import (
	"log"

	"github.com/louisevanderlith/husk"
)

type userFilter func(obj *User) bool

func (f userFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*User))
}

func emailFilter(email string) userFilter {
	return func(obj *User) bool {
		log.Printf("emailFilter: %+v\n", obj)
		return obj.Email == email
	}
}
