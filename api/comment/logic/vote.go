package logic

import "github.com/louisevanderlith/mango/db/comment"

type Vote struct {
	IsUp      bool
	CommentID int64
}

func (vote Vote) Submit() (finalErr error) {

	if vote.CommentID > 0 {
		filter := comment.Message{}
		filter.Id = vote.CommentID

		result, err := comment.Ctx.Messages.ReadOne(&filter)

		if err == nil {
			record, _ := result.(comment.Message)

			if vote.IsUp {
				record.UpVotes++
			} else {
				record.DownVotes++
			}

			comment.Ctx.Messages.Update(&record)
		} else {
			finalErr = err
		}
	}

	return finalErr
}
