package util

import (
	"time"
	"github.com/louisevanderlith/mango/util/enums"
	"encoding/json"
	"log"
	"errors"
)

type AuthCache struct {
	Roles   []enums.RoleType
	LastUse time.Time
}

type CacheCollection map[string]AuthCache

var cache CacheCollection

func init() {
	cache = make(CacheCollection)
}

func getUserRoles(token string) []enums.RoleType {
	var result []enums.RoleType

	auth, ok := cache[token]

	if ok {
		result = auth.Roles
	} else {
		roles, err := loadRoles(token)

		if err != nil {
			result = roles
			cache[token] = AuthCache{
				Roles:   result,
				LastUse: time.Now()}
		}
	}

	return result
}

func hasRole(token string, funcRole enums.RoleType) bool {
	result := false
	roles := getUserRoles(token)

	for _, val := range roles {
		if val == funcRole {
			result = true
			break
		}
	}

	return result
}

func loadRoles(token string) ([]enums.RoleType, error) {
	var result []enums.RoleType
	var finalError error

	contents, statusCode := GETMessage("Secure.API", "session", token)
	data := MarshalToMap(contents)

	if statusCode != 200 {
		var dataErr string
		err := json.Unmarshal(*data["Error"], &dataErr)

		if err != nil {
			log.Printf("loadRoles: ", err)
		}

		finalError = errors.New(dataErr)
	} else {
		err := json.Unmarshal(*data["Data"], &result)

		if err != nil {
			log.Printf("loadRoles: ", err)
		}
	}

	return result, finalError
}
