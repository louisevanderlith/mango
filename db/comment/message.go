package comment

import (
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Message struct {
	db.Record
	UserID      int64
	UpVotes     int64
	DownVotes   int64
	ItemID      int64
	Text        string      `orm:"size(512)"`
	CommentType CommentType `orm:"type(int)"`
}

func (o Message) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
