package handlers

import (
	"db"
	"encoding/json"
	"fmt"
	"general"
	"net/http"

	"github.com/gorilla/mux"
)

type SignupHandler struct {
	validKeys []string
	db *db.DatabaseConnection
}

// Инициализация обработчика регистрации. Обязательна к вызову.
func (handler *SignupHandler) Init(router *mux.Router, db *db.DatabaseConnection) {
	handler.validKeys = []string {"first_name", "last_name", "email", "password", "password_confirmation"}
	handler.db = db
	router.HandleFunc("/signup", handler.userSignup)
}

// Основная функция-обработчик запроса регистрации.
func (handler *SignupHandler) userSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is available.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        w.Write(general.CreateJsonResponse(1, err.Error(), nil))
		return
    }
	var validators Validators
	if err := validators.CheckSignup(data); err != nil {
		w.Write(general.CreateJsonResponse(2, err.Error(), nil))
		return
	}
	user, err := handler.convertToUser(data)
	if err != nil {
		w.Write(general.CreateJsonResponse(3, err.Error(), nil))
		return
	}
	if (handler.db != nil) {
		user, _ = handler.db.GetUserInfoByEmail(data["email"])
		if (user != db.UserT{}) {
			err = fmt.Errorf("Пользователь %v уже существует.", data["email"])
			w.Write(general.CreateJsonResponse(4, err.Error(), nil))
			return
		}
		if err := handler.db.AddUser(user); err != nil {
			w.Write(general.CreateJsonResponse(4, err.Error(), nil))
			return
		}
	}
	w.Write(general.CreateJsonResponse(0, "OK", nil))
}

func (handler *SignupHandler) convertToUser(data map[string]string) (user db.UserT, err error) {
	user.FirstName = data["first_name"]
	user.LastName = data["last_name"]
	user.Email = data["email"]
	user.Password = hashPassword(data["password"])
	return
}

func hashPassword(passw string) string {
	//log.Println("Hashing password..")
	return passw
}