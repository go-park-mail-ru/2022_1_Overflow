package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"

	"github.com/gorilla/csrf"
)

// SignIn godoc
// @Summary Выполняет аутентификацию пользователя
// @Summary Выполняет аутентификацию и выставляет сессионый cookie с названием OverflowMail
// @Success 200 {object} pkg.JsonResponse "Успешная аутентификация пользователя."
// @Failure 500 {object} pkg.JsonResponse "Пользователь не существует, ошибка БД или валидации формы."
// @Accept json
// @Param SignInForm body models.SignInForm true "Форма входа пользователя"
// @Produce json
// @Router /signin [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	if session.IsLoggedIn(r) {
		pkg.WriteJsonErrFull(w, pkg.NO_ERR)
		return
	}

	var data models.SignInForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.JSON_ERR)
		return
	}

	if err := d.uc.SignIn(data); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}

	err = session.CreateSession(w, r, data.Email)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.INTERNAL_ERR)
		return
	}
	csrf.Secure(false) // возможно стоит убрать
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// SignUp godoc
// @Summary Выполняет регистрацию пользователя
// @Description Выполняет регистрацию пользователя, НЕ выставляет сессионый cookie.
// @Success 200 {object} pkg.JsonResponse "Вход уже выполнен, либо успешная регистрация пользователя."
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или пользователь уже существует."
// @Accept json
// @Param SignUpForm body models.SignUpForm true "Форма регистрации пользователя"
// @Produce json
// @Router /signup [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	if session.IsLoggedIn(r) {
		pkg.WriteJsonErrFull(w, pkg.NO_ERR)
		return
	}

	var data models.SignUpForm

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.JSON_ERR)
		return
	}

	data.Email = xss.P.Sanitize(data.Email)
	data.FirstName = xss.P.Sanitize(data.FirstName)
	data.LastName = xss.P.Sanitize(data.LastName)
	
	if err := d.uc.SignUp(data); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}

	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// SignOut godoc
// @Summary Завершение сессии пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное завершение сессии."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует, сессия не валидна."
// @Failure 500 {object} pkg.JsonResponse
// @Produce json
// @Router /logout [get]
func (d *Delivery) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	err := session.DeleteSession(w, r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.INTERNAL_ERR)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// GetToken godoc
// @Summary СSRF проверка
// @Description Токен приходит в header ответа в поле X-CSRF-Token
// @Success 200 {object} models.TokenModel "Успешное получение CSRF токена."
// @Produce json
// @Router /get_token [get]
func (d *Delivery) CSRFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}
	csrf.Secure(false) // возможно стоит убрать
	token := models.TokenModel{Status: 0, Token: csrf.Token(r)}
	resp, _ := json.Marshal(token)
	//w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.Write(resp)
}