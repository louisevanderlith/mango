package control

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego/context"
)

const cookieName = "_avocookie"

type Cookies struct {
	UserKey   husk.Key
	Username  string
	UserRoles map[string]enums.RoleType
	IP        string
	Location  string
}

func NewCookies(userkey husk.Key, username, ip, location string) Cookies {
	result := Cookies{
		UserKey:   userkey,
		Username:  username,
		IP:        ip,
		Location:  location,
		UserRoles: make(map[string]enums.RoleType),
	}

	return result
}

var jar map[string]Cookies

func init() {
	jar = make(map[string]Cookies)
}

func CreateAvo(ctx *context.Context, data Cookies, sessionID string) {
	jar[sessionID] = data
}

func FindAvo(sessionID string) Cookies {
	result := jar[sessionID]

	return result
}

func DestroyAvo(sessionID string) {
	delete(jar, sessionID)
}

func HasAvo(sessionID string) bool {
	_, ok := jar[sessionID]

	return ok
}

func serialize(data Cookies) string {
	raw, err := json.Marshal(data)

	if err != nil {
		log.Print("Cookies Serialize:", err)
	}

	esc := url.QueryEscape(string(raw))
	return string(esc)
}

func deSerialize(data string) Cookies {
	var result Cookies

	raw, _ := url.QueryUnescape(data)
	err := json.Unmarshal([]byte(raw), &result)

	if err != nil {
		log.Print("Cookies DeSerialize:", err)
	}

	return result
}
