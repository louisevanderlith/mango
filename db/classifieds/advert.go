package classifieds

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type Advert struct {
	UserID     int64
	DateListed time.Time
	Price      int
	Negotiable bool
	Tags       Tags
	Location   string `hsk:"size(128)"`
}

func (o Advert) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
