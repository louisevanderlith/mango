package funds

import (
	"time"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

// AvoCredit is an internal system currency used to perform operations such as loading an Ad.
type AvoCredit struct {
	db.Record
	UserID      int64
	Total       int64
	LastUpdated time.Time
}

func (o AvoCredit) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
