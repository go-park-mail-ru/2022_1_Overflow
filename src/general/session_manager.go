package general

import (
	"time"
	"github.com/google/uuid"
	"net/http"
)

var sessions = map[string]session{}

type session struct {
	email string;
	password string;
	expiry time.Time
}

func (s session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func CreateCookies(email string, password string) []http.Cookie {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120*time.Second)

	sessions[sessionToken] = session{
		email: email,
		password: password,
		expiry: expiresAt,
	}

	return []http.Cookie{
		http.Cookie {
			Name: "session_token",
			Value: sessionToken,
			Secure: true,
			SameSite: http.SameSiteNoneMode,
			Expires: expiresAt,
		},
		http.Cookie {
			Name: "email",
			Value: email,
			Secure: true,
			SameSite: http.SameSiteNoneMode,
			Expires: expiresAt,
		},
		http.Cookie {
			Name: "password",
			Value: password,
			Secure: true,
			SameSite: http.SameSiteNoneMode,
			Expires: expiresAt,
		},
	}
}

func HandleCookie() {
	// позже
}