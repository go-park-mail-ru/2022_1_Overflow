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

func CreateCookie(name string, value string, expiresAt time.Time) http.Cookie {
	return http.Cookie{
		Name: name,
		Value: value,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Expires: expiresAt,
	}
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
		CreateCookie("session_token", sessionToken, expiresAt),
		CreateCookie("email", email, expiresAt),
		CreateCookie("password", password, expiresAt),
	}
}

func DeleteCookie(w http.ResponseWriter, name string) {
	cookie := CreateCookie(name, "", time.Now())
	http.SetCookie(w, &cookie)
}

func IsLoggedIn(r *http.Request) bool {
	keys := []string {"email", "password", "session_token"}
	for _, key := range keys {
		_, err := r.Cookie(key)
		if (err != nil) {
			return false
		}
	}
	return true
}

func HandleCookie() {
	// позже
}