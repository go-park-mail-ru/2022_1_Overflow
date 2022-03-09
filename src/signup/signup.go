package main

import (
	"general"
	"fmt"
	"net/http"
	"strings"
	"github.com/rs/cors"
)

type UserT struct {
	LastName  string
	FirstName string
	Email     string
	Password  string
}

type SignupHandler struct {
	validKeys []string
}

// Инициализация обработчика регистрации. Обязательна к вызову.
func (handler *SignupHandler) Init() {
	handler.validKeys = []string {"first_name", "last_name", "email", "password", "password_confirmation"}
}

func (handler *SignupHandler) Handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/signin", handler.UserSignup)
	options := cors.Options{
		AllowedOrigins:   []string{"http://localhost:80", "http://127.0.0.1:80", "http://localhost:3000", "http://127.0.0.1:3000"}, //"*"},
		AllowedHeaders:   []string{"Origin", "Version", "Authorization", "Accept", "Accept-Encoding", "Content-Type", "Content-Length"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}
	handle := cors.New(options).Handler(mux)
	return handle
}

// Основная функция-обработчик запроса регистрации.
func (handler *SignupHandler) UserSignup(w http.ResponseWriter, r *http.Request) {
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
	if err := writeToDatabase(user); err != nil {
		w.Write(general.CreateJsonResponse(4, err.Error(), nil))
		return
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

func (handler *SignupHandler) convertToUser(r *http.Request) (user UserT, err error) {
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Password = hashPassword(r.FormValue("password"))
	return
}

func writeToDatabase(user UserT) (err error) {
	//log.Println("Writing to database..")
	return
}

func hashPassword(passw string) string {
	//log.Println("Hashing password..")
	return passw
}