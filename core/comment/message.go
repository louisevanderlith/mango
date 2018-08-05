package comment

import (
	"errors"

	"github.com/louisevanderlith/husk"
)

type Message struct {
	UserID      int64
	UpVotes     int64
	DownVotes   int64
	ItemID      int64
	Text        string `hsk:"size(512)"`
	CommentType CommentType
	Voters      map[string]struct{}
}

func (o Message) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func SubmitVote(messageID int64, isUp bool, userID string) error {
	msgRec, err := ctx.Messages.FindByID(messageID)

	if err != nil {
		return err
	}

	msgData := msgRec.Data()

	if _, hasVoted := msgData.Voters[userID]; hasVoted {
		return errors.New("user has already voted")
	}

	if isUp {
		msgData.UpVotes++
	} else {
		msgData.DownVotes++
	}

	msgData.Voters[userID] = struct{}{}

	err = ctx.Messages.Update(msgRec)

	return err
}

func SubmitMessage(msg Message) (messageRecord, error) {
	msg.UpVotes = 0
	msg.DownVotes = 0

	return ctx.Messages.Create(msg)
}

func GetCommentChain(itemID int64, commentType CommentType) (parent messageRecord, children messageSet, err error) {
	parent, err = ctx.Messages.FindFirst(func(obj Message) bool {
		return obj.ItemID == itemID && obj.CommentType == commentType
	})

	parentID := parent.rec.GetID()

	if err != nil {
		return parent, children, err
	}

	children, err = ctx.Messages.Find(1, 10, func(obj Message) bool {
		return obj.CommentType == Child && obj.ItemID == parentID
	})

	if err != nil {
		return parent, children, err
	}

	return parent, children, err
}
