package comment

import (
	"github.com/louisevanderlith/husk"
)

type Message struct {
	UserKey     *husk.Key
	ItemKey     *husk.Key
	UpVotes     int64
	DownVotes   int64
	Text        string `hsk:"size(512)"`
	CommentType CommentType
	Voters      map[husk.Key]struct{}
	Children    []Message
}

func (o Message) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func SubmitMessage(msg Message) husk.CreateSet {
	msg.UpVotes = 0
	msg.DownVotes = 0

	return ctx.Messages.Create(msg)
}

func GetMessage(itemKey *husk.Key, commentType CommentType) (husk.Recorder, error) {
	return ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))
}

func UpdateMessage(key *husk.Key, data Message) error {
	rec, err := ctx.Messages.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(data)

	if err != nil {
		return err
	}

	return ctx.Messages.Update(rec)
}
