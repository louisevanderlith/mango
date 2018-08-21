package secure

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/util/control"
)

type AuthResponse struct {
	Passed      bool
	UserKey     husk.Key
	Username    string
	Application *control.Application
}

func NewAuthResponse(passed bool, userKey husk.Key, username string, application *control.Application) *AuthResponse {
	return &AuthResponse{passed, userKey, username, application}
}
