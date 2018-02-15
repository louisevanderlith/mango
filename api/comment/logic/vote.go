package logic

import "github.com/louisevanderlith/mango/db/comment"

type Vote struct {
	IsUp      bool
	CommentID int64
}

func (vote Vote) Sumbit() (finalErr error) {

	if vote.CommentID > 0 {
		filter := comment.Comment{}
		filter.ID = vote.CommentID

		result, err := comment.Ctx.Comment.ReadOne(&filter)

		if err == nil {
			record, _ := result.(comment.Comment)

			if vote.IsUp {
				record.UpVotes++
			} else {
				record.DownVotes++
			}

			comment.Ctx.Comment.Update(record)
		} else {
			finalErr = err
		}
	}

	return finalErr
}
