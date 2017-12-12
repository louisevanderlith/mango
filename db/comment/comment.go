package comment

import (
	"github.com/louisevanderlith/mango/db"
	"errors"
	"strings"
	"github.com/louisevanderlith/mango/util/enums"
	"github.com/louisevanderlith/mango/util"
)

type Comment struct {
	db.Record
	UserID    int64
	UpVotes   int
	DownVotes int
	ItemID    int64
	CommentType enums.CommentType `orm:"type(int)"`
}

func (o Comment) Validate() (bool, error) {
	return util.ValidateStruct(o)
}