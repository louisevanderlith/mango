package util

import (
	"github.com/louisevanderlith/mango/util/enums"
	"log"
	"github.com/astaxie/beego/context"
	"encoding/json"
	"net/url"
)

const cookieName = "_avocookie"

type Cookies struct {
	UserID   int64
	IP       string
	Location string
	Roles    []enums.RoleType
}

var jar map[string]Cookies

func init() {
	jar = make(map[string]Cookies)
}

func CreateAvo(ctx *context.Context, data Cookies, sessionID string) {
	//raw := serialize(data)

	/*var host string

	if beego.BConfig.RunMode == "dev" {
		host = ""
	} else {
		host = "avosa.co.za"
	}*/

	//sessionID := ctx.GetCookie("beegosessionID")
	jar[sessionID] = data
	//ctx.SetCookie(cookieName, raw, 0, "/", host, false, false)
}

func FindAvo(sessionID string) Cookies {
	var result Cookies
	/*data := ctx.GetCookie(cookieName)

	if data != "" {
		result = deSerialize(data)
	}*/

	result, _ = jar[sessionID]

	return result
}

func DestroyAvo(sessionID string) {
	delete(jar, sessionID)
	/*empty := serialize(Cookies{})
	ctx.SetCookie(cookieName, empty, 0)*/
}

func HasAvo(sessionID string) bool {
	_, ok := jar[sessionID]
	/*data := ctx.GetCookie(cookieName)

	return data != ""*/
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
