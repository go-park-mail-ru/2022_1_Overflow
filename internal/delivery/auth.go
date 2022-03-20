package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
)

// SignIn godoc
// @Summary Выполняет аутентификацию пользователя
// @Success 200
// @Failure 500
// @Router /signin [post]
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

	if err := d.uc.SignIn(data); err != nil {
		w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
		return
	}

	err = session.CreateSession(w, r, data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// SignUp godoc
// @Summary Выполняет регистрацию пользователя
// @Success 200
// @Failure 500
// @Router /signup [post]
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
	err = d.uc.SignUp(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// SignIn godoc
// @Summary Завершение сессии пользователя
// @Success 200
// @Failure 401
// @Failure 500
// @Router /logout [get]
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}