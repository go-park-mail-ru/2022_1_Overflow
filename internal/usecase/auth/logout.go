package auth

import (
	"OverflowBackend/pkg"
	"net/http"
)

func (a *Auth) SignOut(w http.ResponseWriter, r *http.Request) {
	
	if !a.sm.IsLoggedIn(r) {
		pkg.AccessDenied(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	err := a.sm.DeleteSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}