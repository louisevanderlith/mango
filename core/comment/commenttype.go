package comment

import "strings"

type CommentType = int

const (
	Profile CommentType = iota
	Advert
	Child
)

var commentTypes = [...]string{
	"Profile",
	"Advert",
	"Child",
}

func GetCommentType(name string) CommentType {
	var result CommentType

	for k, v := range commentTypes {
		if strings.ToUpper(name) == v {
			result = CommentType(k)
			break
		}
	}

	return result
}
