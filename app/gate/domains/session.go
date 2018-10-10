package domains

import (
	"net/http"
	"net/url"
	"strings"
)

func handleSession(url url.URL, w http.ResponseWriter) {
	if !strings.Contains(url.String(), "?token=") {
		return
	}

	sessionID := url.Query().Get(token)

	if sessionID != "" {
		cookie := http.Cookie{
			Name:     "avosession",
			Path:     "/",
			Value:    sessionID,
			HttpOnly: true,
			MaxAge:   0,
		}

		http.SetCookie(w, &cookie)
	}
}
