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
// @Summary Выполняет аутентификацию и выставляет сессионый cookie с названием OverflowMail
// @Success 200 {string} string "Успешная аутентификация пользователя."
// @Failure 500 "Пользователь не существует, ошибка БД или валидации формы."
// @Accept json
// @Param SignInForm body models.SignInForm true "Форма входа пользователя"
// @Produce plain
// @Router /signin [post]
func (d *Delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	if session.IsLoggedIn(r) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

	var data models.SignInForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := d.uc.SignIn(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
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
// @Description Выполняет регистрацию пользователя, НЕ выставляет сессионый cookie.
// @Success 200 {string} string "Успешная регистрация пользователя."
// @Failure 500 "Ошибка валидации формы, БД или пользователь уже существует."
// @Accept json
// @Param SignUpForm body models.SignUpForm true "Форма регистрации пользователя"
// @Produce plain
// @Router /signup [post]
func (d *Delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

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
// @Success 200 {string} string "Успешное завершение сессии."
// @Failure 401 "Сессия отсутствует, сессия не валидна."
// @Failure 500
// @Produce plain
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