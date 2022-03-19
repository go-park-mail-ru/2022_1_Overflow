package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
)

// Аутентификация пользователя.
func (d *Delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	if session.IsLoggedIn(r) {
		w.Write(pkg.CreateJsonResponse(0, "OK", nil))
		return
	}

	var data models.SignInForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d.uc.SignIn(w, r, data)
}

// Регистрация пользователя.
func (d *Delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var data models.SignUpForm

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d.uc.SignUp(w, r, data)
}

// Завершение сессии пользователя.
func (d *Delivery) SignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	err := session.DeleteSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}