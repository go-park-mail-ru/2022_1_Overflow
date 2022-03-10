package handlers

import (
	response "OverflowBackend/src/response"
	session "OverflowBackend/src/session"

	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is available.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !session.IsLoggedIn(r) {
		w.Write(response.CreateJsonResponse(1, "Пользователь не выполнил вход.", nil))
		return
	}

	session.DeleteSession(w, r)

	w.Write(response.CreateJsonResponse(0, "OK", nil))
}
