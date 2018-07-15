package comment

import "github.com/louisevanderlith/husk"

type Message struct {
	UserID      int64
	UpVotes     int64
	DownVotes   int64
	ItemID      int64
	Text        string `hsk:"size(512)"`
	CommentType CommentType
}

func (o Message) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
