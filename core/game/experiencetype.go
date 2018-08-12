package game

import "strings"

type ExperienceType = int

const (
	PlaceAd ExperienceType = iota
	Transact
	Comment
	Vote
)

var experienceTypes = [...]string{
	"PlaceAd",
	"Transact",
	"Comment",
	"Vote",
}

func GetExperienceType(name string) ExperienceType {
	var result ExperienceType

	for k, v := range experienceTypes {
		if strings.ToUpper(name) == v {
			result = ExperienceType(k)
			break
		}
	}

	return result
}

func XPValue(xp ExperienceType) int {
	result := 0

	switch xp {
	case PlaceAd:
		result = 10
	case Transact:
		result = 5
	case Comment:
		result = 2
	case Vote:
		result = 1
	}

	return result
}
