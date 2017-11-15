package logic

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/louisevanderlith/mango/util/enums"
)

type UserSession struct {
	Timestamp time.Time
	ExpiresIn time.Duration
	UserID    int64
	IP        string
	Location  string
	Roles     []enums.RoleType
}

var (
	sessionStore map[string]*UserSession
)

func init() {
	sessionStore = make(map[string]*UserSession)
}

func Set(session *UserSession) string {
	token := generateToken()

	session.Timestamp = time.Now()
	session.ExpiresIn = time.Minute * 15

	sessionStore[token] = session

	return token
}

func Get(token string) *UserSession {
	var result *UserSession

	if len(token) == 16 && exists(token) {
		result = sessionStore[token]

		if !expired(result, token) {
			result.Timestamp = time.Now()
		}
	}

	return result
}

func Delete(token string) {
	delete(sessionStore, token)
}

func exists(token string) bool {
	_, ok := sessionStore[token]

	return ok
}

func expired(session *UserSession, token string) bool {
	expireTime := session.Timestamp.Add(session.ExpiresIn)
	expiresIn := expireTime.Sub(session.Timestamp)
	expired := expiresIn <= 0

	if expired {
		Delete(token)
	}

	return expired
}

func generateToken() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)

	if err != nil {
		log.Printf("generateToken: ", err)
	}

	result := fmt.Sprintf("%X", b)

	if exists(result) {
		result = generateToken()
	}

	return result
}
