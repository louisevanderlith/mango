package logic

import (
	"time"

	"github.com/louisevanderlith/mango/db/secure"
)

type UserObject struct {
	Name        string
	Verified    bool
	DateCreated time.Time
	LastLogin   time.Time
}

func GetUsers() (result []UserObject, finalError error) {
	data, err := secure.GetUsers()

	if err != nil {
		finalError = err
	} else {
		for _, v := range data {
			item := UserObject{
				Name:        v.Name,
				DateCreated: v.CreateDate,
				LastLogin:   v.LoginDate,
				Verified:    v.Verified,
			}

			result = append(result, item)
		}
	}

	return result, finalError
}
