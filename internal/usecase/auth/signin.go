package auth

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/add_validation"
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
)

func (a *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var data models.SignInForm

	if a.sm.IsLoggedIn(r) {
		w.Write(pkg.CreateJsonResponse(0, "OK", nil))
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := add_validation.CheckSignIn(data); err != nil {
		w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
		return
	}

	err = a.sm.CreateSession(w, r, data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(pkg.CreateJsonResponse(0, "OK", nil))
}
