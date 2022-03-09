package handlers

import (
	"db"
	"general"
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
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
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        w.Write(general.CreateJsonResponse(1, err.Error(), nil))
		return
    }
	if err := handler.validateData(data); err != nil {
		w.Write(general.CreateJsonResponse(2, err.Error(), nil))
		return
	}
	user, err := handler.convertToUser(data)
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

func (handler *SignupHandler) validateData(data map[string]interface{}) (err error) {
	for _, key := range handler.validKeys {
		val, exists := data[key]
		if exists || len(strings.TrimSpace(val.(string))) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}
	if data["password"] != data["password_confirmation"] {
		return fmt.Errorf("Пароли не совпадают.")
	}
	return
}

func (handler *SignupHandler) convertToUser(data map[string]interface{}) (user db.UserT, err error) {
	user.FirstName = data["first_name"].(string)
	user.LastName = data["last_name"].(string)
	user.Email = data["email"].(string)
	user.Password = hashPassword(data["password"].(string))
	return
}

func hashPassword(passw string) string {
	//log.Println("Hashing password..")
	return passw
}