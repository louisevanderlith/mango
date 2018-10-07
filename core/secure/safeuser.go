package secure

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type SafeUser struct {
	Key         *husk.Key
	Name        string
	Verified    bool
	DateCreated time.Time
	LastLogin   time.Time
}

func createSafeUser(user husk.Recorder) SafeUser {
	data := user.Data().(*User)
	meta := user.Meta()

	result := SafeUser{
		Key:       meta.Key,
		LastLogin: data.LoginDate,
		Name:      data.Name,
		Verified:  data.Verified,
	}

	return result
}

func GetUsers(page, size int) []SafeUser {
	var result []SafeUser
	users := getUsers(page, size)
	itor := users.GetEnumerator()
	for itor.MoveNext() {
		currUser := itor.Current()

		sfeUser := createSafeUser(currUser)
		result = append(result, sfeUser)
	}

	return result
}
