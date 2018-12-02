package logic

import (
	"github.com/astaxie/beego/context"
	"github.com/louisevanderlith/mango/pkg/control"
	uuid "github.com/nu7hatch/gouuid"
)

var jar map[string]control.Cookies

func init() {
	jar = make(map[string]control.Cookies)
}

//CreateAvo creates an Avo(cookie) & returns the session ID
func CreateAvo(ctx *context.Context, data *control.Cookies) string {
	u4, _ := uuid.NewV4()
	sessionID := u4.String()

	jar[sessionID] = *data

	return sessionID
}

func FindAvo(sessionID string) control.Cookies {
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
