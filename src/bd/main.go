package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type SigninHandler struct {
	validKeys []string
}

func (handler *SigninHandler) Init() {
	handler.validKeys = []string{"email", "password"}
}

func CORSMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func (handler *SigninHandler) Handlers() http.Handler {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/signin", handler.UserSignin)
// 	handle := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedHeaders:   []string{"Version", "Authorization", "Content-Type", "csrf_token"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
// 		AllowCredentials: true,
// 	}).Handler(mux)
// 	return handle
// }

func (handler *SigninHandler) UserSignin(w http.ResponseWriter, r *http.Request) {

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
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres"
	conn, err := pgxpool.Connect(context.Background(), urlExample)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	var user UserT
	user, _ = GetUserInfoByEmail(r.FormValue("email"), conn)

	if r.FormValue("password") != user.Password {
		return fmt.Errorf("Пароли не совпадают.")
	}

	return
}

func main() {
	address := ":8080"
	log.Printf("Listening on %v", address)

	var handler SigninHandler
	mux := http.NewServeMux()
	mux.Handle("/signin", CORSMiddleWare(http.HandlerFunc(handler.UserSignin)))

	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal(err)
	}
}
