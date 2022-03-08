package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"https://github.com/rs/cors"
)

type UserT struct {
	LastName  string
	FirstName string
	Email     string
	Password  string
}

type SigninHandler struct {
	validKeys []string
}

func (handler *SigninHandler) Init() {
	handler.validKeys = []string{"email", "password"}
}

func (handler *SigninHandler) Handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/signin", handler.UserSignin)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string {"*"},
		AllowedHeaders:   []string {"Version", "Authorization", "Content-Type", "csrf_token"},
		AllowedMethods:   []string {"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	  }).Handler(router)
	return handler
}

func (handler *SigninHandler) UserSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is available.", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.Write(handler.createJsonResponse(1, err.Error(), nil))
		return
	}
	if err := handler.validateData(r); err != nil {
		w.Write(handler.createJsonResponse(2, err.Error(), nil))
		return
	}
	w.Write(handler.createJsonResponse(0, "OK", nil))
}

func (handler *SigninHandler) createJsonResponse(status int, message string, content interface{}) []byte {
	resp, _ := json.Marshal(
		map[string]interface{}{
			"status":  status,
			"message": message,
			"content": content,
		},
	)
	return resp
}

func (handler *SigninHandler) validateData(r *http.Request) (err error) {
	for _, key := range handler.validKeys {
		val := r.FormValue(key)
		if len(strings.TrimSpace(val)) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}
	var user UserT
	user, _ := GetUserInfoByEmail(r.FormValue("email"))

	if r.FormValue("password") != user.Password {
		return fmt.Errorf("Пароли не совпадают.")
	}
	return
}

func main() {
	address := ":8080"
	log.Printf("Listening on %v", address)
	var handler SigninHandler
	err := http.ListenAndServe(address, handler.Handlers())
	if err != nil {
		log.Fatal(err)
	}
}
