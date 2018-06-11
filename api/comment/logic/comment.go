package logic

import (
	"strings"
	"time"

	"github.com/louisevanderlith/mango/db/comment"
)

type MessageEntry struct {
	Text        string
	ParentID    int64
	CommentType comment.CommentType
}

type simpleComment struct {
	User       string
	DatePosted time.Time
	Text       string
	UpVotes    int64
	DownVotes  int64
	Children   []simpleComment
}

type CommentChain []simpleComment

func SubmitComment(userID int64, entry MessageEntry) (finalErr error) {

	record := comment.Message{
		CommentType: entry.CommentType,
		Text:        strings.Trim(entry.Text, " "),
		ItemID:      entry.ParentID,
		UpVotes:     0,
		DownVotes:   0,
		UserID:      userID,
	}

	_, finalErr = comment.Ctx.Messages.Create(&record)

	return finalErr
}

func GetCommentChain(itemID int64, commentType comment.CommentType) (results CommentChain, finalErr error) {
	filter := comment.Message{}
	filter.ItemID = itemID
	filter.CommentType = commentType

	var container comment.Messages
	err := comment.Ctx.Messages.Read(&filter, &container)

	if err == nil {
		for _, v := range container {
			item := simpleComment{
				Children:   getChildren(v.Id),
				DatePosted: v.CreateDate,
				Text:       v.Text,
				DownVotes:  v.DownVotes,
				UpVotes:    v.UpVotes,
				User:       "TODO",
			}

			results = append(results, item)
		}
	} else {
		finalErr = err
	}

	return results, finalErr
}

func getChildren(itemID int64) (results CommentChain) {
	filter := comment.Message{}
	filter.ItemID = itemID
	filter.CommentType = comment.Child

	var container comment.Messages
	err := comment.Ctx.Messages.Read(&filter, &container)

	if err == nil {
		for _, v := range container {
			item := simpleComment{
				Children:   getChildren(v.Id),
				DatePosted: v.CreateDate,
				Text:       v.Text,
				DownVotes:  v.DownVotes,
				UpVotes:    v.UpVotes,
				User:       "TODO",
			}

			results = append(results, item)
		}
	}

	return results
}
