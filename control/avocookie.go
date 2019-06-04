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
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func GetAvoCookie(sessionID, publickeyPath string) (*secure.Cookies, error) {
	if len(sessionID) == 0 {
		return nil, errors.New("SessionID empty")
	}

	token, err := jwt.Parse(sessionID, func(t *jwt.Token) (interface{}, error) {
		var rdr io.Reader
		if f, err := os.Open(publickeyPath); err == nil {
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

func IsAllowed(appName string, usrRoles secure.ActionMap, required roletype.Enum) (bool, error) {
	if required == roletype.Unknown {
		return true, nil
	}
	return hasRole(appName, usrRoles, required)
}

func hasRole(appName string, usrRoles secure.ActionMap, required roletype.Enum) (bool, error) {
	role, err := getRole(appName, usrRoles)

	if err != nil {
		return false, err
	}

	return role <= required, nil
}

func getRole(appName string, usrRoles secure.ActionMap) (roletype.Enum, error) {
	role, ok := usrRoles[appName]

	if !ok {
		msg := fmt.Errorf("application permission required. %s", appName)
		return roletype.Unknown, msg
	}

	return role, nil
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
