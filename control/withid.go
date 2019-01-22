package control

import (
	"github.com/louisevanderlith/husk"
)

type WithKey struct {
	Key *husk.Key
	Body interface{}
}
