package comment

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

type Comment struct {
	db.Record
	UserID      int64
	UpVotes     int64
	DownVotes   int64
	ItemID      int64
	Text        string            `orm:"size(512)"`
	CommentType enums.CommentType `orm:"type(int)"`
}

func (o Comment) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}
