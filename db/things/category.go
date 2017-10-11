package things

import "github.com/louisevanderlith/mango/util"

type Category struct {
	util.BaseRecord
	Name        string
	Description string
}
