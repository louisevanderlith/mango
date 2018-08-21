package comment

import (
	"errors"

	"github.com/louisevanderlith/husk"
)

type Message struct {
	UserKey     husk.Key
	ItemKey     husk.Key
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

func SubmitVote(messageKey husk.Key, isUp bool, userKey husk.Key) error {
	msgRec, err := ctx.Messages.FindByKey(messageKey)

	if err != nil {
		return err
	}

	msgData := msgRec.Data()

	if _, hasVoted := msgData.Voters[userKey]; hasVoted {
		return errors.New("user has already voted")
	}

	if isUp {
		msgData.UpVotes++
	} else {
		msgData.DownVotes++
	}

	msgData.Voters[userKey] = struct{}{}

	err = ctx.Messages.Update(msgRec)

	return err
}

func SubmitMessage(msg Message) (messageRecord, error) {
	msg.UpVotes = 0
	msg.DownVotes = 0

	return ctx.Messages.Create(msg)
}

func GetCommentParts(itemKey husk.Key, commentType CommentType) (parent messageRecord, children messageSet, err error) {
	parent, err = ctx.Messages.FindFirst(func(obj Message) bool {
		return obj.ItemKey == itemKey && obj.CommentType == commentType
	})

	parentKey := parent.rec.GetKey()

	if err != nil {
		return parent, children, err
	}

	children, err = ctx.Messages.Find(1, 10, func(obj Message) bool {
		return obj.CommentType == Child && obj.ItemKey == parentKey
	})

	if err != nil {
		return parent, children, err
	}

	return parent, children, err
}
