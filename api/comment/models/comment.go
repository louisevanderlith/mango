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
