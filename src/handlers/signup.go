package handlers

import (
	"db"
	"general"
	"fmt"
	"net/http"
	"strings"
)

type SignupHandler struct {
	validKeys []string
	db *db.DatabaseConnection
}

// Инициализация обработчика регистрации. Обязательна к вызову.
func (handler *SignupHandler) Init(mux *http.ServeMux, db *db.DatabaseConnection) {
	handler.validKeys = []string {"first_name", "last_name", "email", "password", "password_confirmation"}
	handler.db = db
	mux.HandleFunc("/signup", handler.userSignup)
}

// Основная функция-обработчик запроса регистрации.
func (handler *SignupHandler) userSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is available.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
    if err != nil {
        w.Write(general.CreateJsonResponse(1, err.Error(), nil))
		return
    }
	if err := handler.validateData(r); err != nil {
		w.Write(general.CreateJsonResponse(2, err.Error(), nil))
		return
	}
	user, err := handler.convertToUser(r)
	if err != nil {
		w.Write(general.CreateJsonResponse(3, err.Error(), nil))
		return
	}
	if (handler.db != nil) {
		if err := handler.db.AddUser(user); err != nil {
			w.Write(general.CreateJsonResponse(4, err.Error(), nil))
			return
		}
	}
	w.Write(general.CreateJsonResponse(0, "OK", nil))
}

func (handler *SignupHandler) validateData(r *http.Request) (err error) {
	for _, key := range handler.validKeys {
		val := r.FormValue(key)
		if len(strings.TrimSpace(val)) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}
	if r.FormValue("password") != r.FormValue("password_confirmation") {
		return fmt.Errorf("Пароли не совпадают.")
	}
	return
}

func (handler *SignupHandler) convertToUser(r *http.Request) (user db.UserT, err error) {
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Password = hashPassword(r.FormValue("password"))
	return
}

func hashPassword(passw string) string {
	//log.Println("Hashing password..")
	return passw
}