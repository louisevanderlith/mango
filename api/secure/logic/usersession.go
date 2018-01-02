package logic

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/louisevanderlith/mango/util/enums"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

type UserSession struct {
	Timestamp time.Time
	ExpiresIn time.Duration
	UserID    int64
	IP        string
	Location  string
	Roles     []enums.RoleType
}

var sessionStore map[string]*UserSession

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

func SetAvoToken(ctx *context.Context, token string) {
	if beego.BConfig.RunMode == "dev" {
		ctx.SetCookie("avotoken", token)
	} else {
		ctx.SetCookie("avotoken", token, 600, "/", "avosa.co.za", true, true)
	}
}

func Get(token string) *UserSession {
	var result *UserSession

	if tokenValid(token) {
		result = sessionStore[token]

		if !expired(result, token) {
			result.Timestamp = time.Now()
		}
	}

	return result
}

func GetAvoToken(ctx *context.Context) string {
	return ctx.GetCookie("avotoken")
}

func GetRoles(token string) []enums.RoleType {
	var result []enums.RoleType

	session := Get(token)

	if session != nil {
		result = session.Roles
	}

	return result
}

func ExpireAvoToken(ctx *context.Context, token string) {
	ctx.SetCookie("avotoken", "expired", 0)
	delete(sessionStore, token)
}

func tokenValid(token string) bool {
	return len(token) == 16 && exists(token)
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
		delete(sessionStore, token)
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
