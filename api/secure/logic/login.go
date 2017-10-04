package logic

import (
	"github.com/louisevanderlith/mango/db/secure"
)

type Login struct {
	Identifier string
	Password   string
	IP string
	Location string
}

func AttemptLogin(l Login) (string, error) string{
	var token string
	loggedIn, userID := secure.Login(l.Identifier, []byte(l.Password), l.IP, l.Location)

	if loggedIn {
		session := UserSession{
				IP = l.IP,
				Location = l.Location,
				UserID = userID		}

		token = Set(&session)
	}

	return token
}