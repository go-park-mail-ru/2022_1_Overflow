package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"context"

	"encoding/json"
	"net/http"

	"github.com/gorilla/csrf"
	log "github.com/sirupsen/logrus"
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
	log.Info("SignIn: ", "checking method")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	log.Info("SignIn: ", "checking session")
	if session.Manager.IsLoggedIn(r) {
		pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
		return
	}

	log.Info("SignIn: ", "checking data")
	var data models.SignInForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	dataBytes, _ := json.Marshal(data)
	resp, err := d.auth.SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: dataBytes,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse 
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}

	log.Info("SignIn: ", "creating session")
	err = session.Manager.CreateSession(w, r, data.Username)
	if err != nil {
		log.Errorf("SignIn: %v", err)
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	csrf.Secure(false) // возможно стоит убрать
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /signin [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SignIn() {}

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
	log.Info("SignUp: ", "checking method")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	log.Info("SignUp: ", "checking session")
	if session.Manager.IsLoggedIn(r) {
		pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
		return
	}

	log.Info("SignUp: ", "checking data")
	var data models.SignUpForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}

	log.Info("SignUp: ", "sanitizing data")
	data.Username = xss.P.Sanitize(data.Username)
	data.Firstname = xss.P.Sanitize(data.Firstname)
	data.Lastname = xss.P.Sanitize(data.Lastname)
	passSanitized := xss.P.Sanitize(data.Password)
	if passSanitized != data.Password {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, "Пароль содержит недопустимое содержимое.")
		return
	}

	dataBytes, _ := json.Marshal(data)
	resp, err := d.auth.SignUp(context.Background(), &auth_proto.SignUpRequest{
		Form: dataBytes,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse 
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}

	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /signup [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SignUp() {}

// SignOut godoc
// @Summary Завершение сессии пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное завершение сессии."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует, сессия не валидна."
// @Failure 500 {object} pkg.JsonResponse
// @Produce json
// @Router /logout [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	err := session.Manager.DeleteSession(w, r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /logout [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SignOut() {}
