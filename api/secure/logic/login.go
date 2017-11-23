package logic

import (
	"github.com/louisevanderlith/mango/db/secure"
)

type Login struct {
	Identifier string
	Password   string
	IP         string
	Location   string
	ReturnURL  string
}

func AttemptLogin(l Login) string {
	var token string
	loggedIn, userID, roles := secure.Login(l.Identifier, []byte(l.Password), l.IP, l.Location)

	if loggedIn {
		session := UserSession{
			IP:       l.IP,
			Location: l.Location,
			UserID:   userID,
			Roles:    roles}

		token = Set(&session)
	}

	return token
}
