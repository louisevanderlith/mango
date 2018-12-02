package folio

import (
	"github.com/louisevanderlith/mango/util"
)

type SocialLink struct {
	Icon string `hsk:"size(25)"`
	URL  string `hsk:"size(128)"`
}

func (o SocialLink) Valid() (bool, error) {
	return util.ValidateStruct(&o)
}
