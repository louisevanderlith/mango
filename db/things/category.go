package things

import "github.com/louisevanderlith/mango/util"

type Category struct {
	util.Record
	Name        string
	Description string
}
