package handlers

import (
	session "OverflowBackend/src/session"

	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !session.IsLoggedIn(r) {
		AccessDenied(w)
		return
	}

	if r.Method != http.MethodGet {
		MethodNotAllowed(w, http.MethodGet)
		return
	}

	err := session.DeleteSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
