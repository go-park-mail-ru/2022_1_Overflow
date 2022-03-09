package handlers

import (
	"general"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is available.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !general.IsLoggedIn(r) {
		w.Write(general.CreateJsonResponse(1, "Пользователь не выполнил вход.", nil))
		return
	}

	general.DeleteCookie(w, "email")
	general.DeleteCookie(w, "password")
	general.DeleteCookie(w, "session_token")

	w.Write(general.CreateJsonResponse(0, "OK", nil))
}
