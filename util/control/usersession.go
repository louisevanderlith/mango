package control

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/util/enums"
)

func GetUserID(sessionID string) (result int64, err error) {
	cookie, err := getAvoCookie(sessionID)

	if err == nil {
		result = cookie.UserID
	}

	return result, err
}

func GetRoles(sessionID string) (result []enums.RoleType, err error) {
	cookie, err := getAvoCookie(sessionID)

	if err == nil {
		result = cookie.Roles
	}

	return result, err
}

func getAvoCookie(sessionID string) (result util.Cookies, finalError error) {

	contents, statusCode := util.GETMessage("Secure.API", "login", "avo", sessionID)
	data := util.MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Println("loadRoles:", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Println("loadRoles:", err)
		}
	}

	return result, finalError
}
