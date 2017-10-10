package comment

import "github.com/louisevanderlith/mango/util"

type Comment struct {
	util.Record
	UserID    int64 `orm:"null"`
	UpVotes   int
	DownVotes int
	ItemID    int64 `orm:"null"`
}
