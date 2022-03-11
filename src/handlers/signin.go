package handlers

import (
	db "OverflowBackend/src/db"
	response "OverflowBackend/src/response"
	validation "OverflowBackend/src/validation"
	session "OverflowBackend/src/session"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SigninHandler struct {
	validKeys []string
	db *db.DatabaseConnection
}

func (handler *SigninHandler) Init(router *mux.Router, db *db.DatabaseConnection) {
	handler.validKeys = []string{"email", "password"}
	handler.db = db
	router.HandleFunc("/signin", handler.userSignin)
}

func (handler *SigninHandler) userSignin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Write(response.CreateJsonResponse(1, err.Error(), nil))
		return
	}
	if err := handler.validateData(data); err != nil {
		w.Write(response.CreateJsonResponse(2, err.Error(), nil))
		return
	}

	err = session.CreateSession(w, r, data["email"])
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Write(response.CreateJsonResponse(0, "OK", nil))
}

func (handler *SigninHandler) validateData(data map[string]string) (err error) {

	if err = validation.CheckSignin(data["email"], data["password"]); err != nil {
		return err
	}

	if (handler.db == nil) {
		return
	}

	user, _ := handler.db.GetUserInfoByEmail(data["email"])

	if (user == db.UserT{}) {
		return fmt.Errorf("Пользователь не существует.")
	}

	if data["password"] != user.Password {
		return fmt.Errorf("Пароли не совпадают.")
	}

	return
}
