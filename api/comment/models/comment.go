package models

import (
	"time"

	"github.com/louisevanderlith/mango/core/comment"
)

type MessageEntry struct {
	Text        string
	ParentID    int64
	CommentType comment.CommentType
}

type SimpleComment struct {
	User       string
	DatePosted time.Time
	Text       string
	UpVotes    int64
	DownVotes  int64
	Children   CommentChain
}

type CommentChain []SimpleComment

func CreateCommentChain(parentData comment.Message, children []comment.Message, username string, createdate time.Time) CommentChain {
	smplComment := SimpleComment{
		DatePosted: createdate,
		Text:       parentData.Text,
		User:       username,
		DownVotes:  parentData.DownVotes,
		UpVotes:    parentData.UpVotes,
	}

	for _, childMsg := range children {
		newChild := CreateCommentChain(childMsg)
		smplComment.Children = append(smplComment.Children)
	}
}
