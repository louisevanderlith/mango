package comment

import (
	"github.com/louisevanderlith/mango/db"
)

type Comment struct {
	db.Record
	UserID    int64 `orm:"null"`
	UpVotes   int
	DownVotes int
	ItemID    int64 `orm:"null"`
}