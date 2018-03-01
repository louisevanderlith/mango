package logic

import (
	"strings"
	"time"

	"github.com/louisevanderlith/mango/db/comment"
	"github.com/louisevanderlith/mango/util/enums"
)

type simpleComment struct {
	User       string
	DatePosted time.Time
	Text       string
	UpVotes    int64
	DownVotes  int64
	Children   []simpleComment
}

type CommentChain []simpleComment

func SubmitComment(userID, parentID int64, text string, commentType enums.CommentType) (finalErr error) {

	record := comment.Comment{
		CommentType: commentType,
		Text:        strings.Trim(text, " "),
		ItemID:      parentID,
		UpVotes:     0,
		DownVotes:   0,
		UserID:      userID,
	}

	_, finalErr = comment.Ctx.Comment.Create(&record)

	return finalErr
}

func GetCommentChain(itemID int64, commentType enums.CommentType) (results CommentChain, finalErr error) {
	filter := comment.Comment{}
	filter.ItemID = itemID
	filter.CommentType = commentType

	var container []*comment.Comment
	err := comment.Ctx.Comment.Read(&filter, &container)

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
	filter := comment.Comment{}
	filter.ItemID = itemID
	filter.CommentType = enums.Child

	var container []*comment.Comment
	err := comment.Ctx.Comment.Read(&filter, &container)

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
