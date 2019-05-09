package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/husk"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

type TinyCtx struct {
	ApplicationName string
	RequiredRole    roletype.Enum
	URL             string
	Method          string
	SessionID       string
}

const avosession = "avosession"

func NewTinyCtx(applicationName, method, url, token string, requiredrole roletype.Enum) (*TinyCtx, error) {
	if len(method) < 3 {
		return nil, errors.New("invalid method")
	}

	if len(url) < 1 {
		return nil, errors.New("invalid url")
	}

	if len(token) < 10 {
		return nil, errors.New("invalid token")
	}

	result := TinyCtx{}

	result.RequiredRole = requiredrole
	result.ApplicationName = applicationName
	result.URL = url
	result.Method = method
	result.SessionID = token

	return &result, nil
}

func (ctx *TinyCtx) allowed() (bool, error) {
	if ctx.RequiredRole == roletype.Unknown {
		return true, nil
	}

	return ctx.hasRole(ctx.RequiredRole)
}

func (ctx *TinyCtx) hasRole(required roletype.Enum) (bool, error) {
	role, err := ctx.getRole()

	if err != nil {
		return false, err
	}

	return role <= required, nil
}

func (ctx *TinyCtx) getRole() (roletype.Enum, error) {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return roletype.Unknown, err
	}

	appName := ctx.ApplicationName

	role, ok := cookie.UserRoles[appName]

	if !ok {
		msg := fmt.Errorf("application permission required. %s", appName)
		return roletype.Unknown, msg
	}

	return role, nil
}

func (ctx *TinyCtx) getUserKey() husk.Key {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return husk.CrazyKey()
	}

	return cookie.UserKey
}

func (ctx *TinyCtx) getUsername() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Unknown"
	}

	return cookie.Username
}

func (ctx *TinyCtx) getIP() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "-1.-1.-1.-1"
	}

	return cookie.IP
}

func (ctx *TinyCtx) getLocation() string {
	cookie, err := ctx.getAvoCookie()

	if err != nil {
		return "Unknown"
	}

	return cookie.Location
}

//TODO: use channels
//getAvoCookie also checks cookie validity, so repeated calls are required
func (ctx *TinyCtx) getAvoCookie() (*secure.Cookies, error) {
	if len(ctx.SessionID) == 0 {
		return nil, errors.New("SessionID empty")
	}

	token, err := jwt.Parse(ctx.SessionID, func(t *jwt.Token) (interface{}, error) {
		var rdr io.Reader
		if f, err := os.Open("/db/sign_rsa.pub"); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return "", err
		}

		bits, err := ioutil.ReadAll(rdr)

		if err != nil {
			return "", err
		}

		return jwt.ParseRSAPublicKeyFromPEM(bits)
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalid")
	}

	jClaim, err := json.Marshal(token.Claims)

	if err != nil {
		return nil, err
	}

	result := &secure.Cookies{}
	err = json.Unmarshal(jClaim, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func removeToken(url string) (string, string) {
	idx := strings.LastIndex(url, "?access_token")

	if idx == -1 {
		return url, ""
	}

	tokenIdx := strings.LastIndex(url, "=") + 1

	cleanURL := url[:idx]
	token := url[tokenIdx:]

	return cleanURL, token
}
