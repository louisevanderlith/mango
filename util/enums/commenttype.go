package enums

import "strings"

type CommentType int

const (
	Profile   CommentType = iota
	Advert
	Child
)

var commentTypes = [...]string{
	"Profile",
	"Advert",
	"Child",
}

func (r CommentType) String() string {
	return roleTypes[r]
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
