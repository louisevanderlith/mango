package comment

import (
	"github.com/louisevanderlith/husk"
)

type messageFilter func(obj *Message) bool

func (f messageFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Message))
}

func byItemKeyCommentType(itemKey *husk.Key, commentType CommentType) messageFilter {
	return func(obj *Message) bool {
		return obj.ItemKey == itemKey && obj.CommentType == commentType
	}
}
