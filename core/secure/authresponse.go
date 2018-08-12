package secure

import "github.com/louisevanderlith/mango/util/control"

type AuthResponse struct {
	Passed      bool
	UserID      int64
	Username    string
	Application *control.Application
}

func NewAuthResponse(passed bool, userID int64, username string, application *control.Application) *AuthResponse {
	return &AuthResponse{passed, userID, username, application}
}
