package handlers

import (
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
		http.Error(w, "Access denied.", http.StatusInternalServerError)
		return
	}

	err := session.DeleteSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
