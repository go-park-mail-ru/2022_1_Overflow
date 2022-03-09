package handlers

import (
	"db"
	"general"
	"net/http"
	"encoding/json"
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
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        w.Write(general.CreateJsonResponse(1, err.Error(), nil))
		return
    }
	var validators Validators
	if err := validators.CheckSignup(data["email"].(string), data["password"].(string), data["password_confirmation"].(string)); err != nil {
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