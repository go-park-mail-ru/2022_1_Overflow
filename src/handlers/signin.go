package handlers

import (
	"db"
	"encoding/json"
	"fmt"
	"general"
	"net/http"
	"strings"
)

type SigninHandler struct {
	validKeys []string
	db *db.DatabaseConnection
}

func (handler *SigninHandler) Init(mux *http.ServeMux, db *db.DatabaseConnection) {
	handler.validKeys = []string{"email", "password"}
	handler.db = db
	mux.HandleFunc("/signin", handler.userSignin)
}

func (handler *SigninHandler) userSignin(w http.ResponseWriter, r *http.Request) {
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

	cookies := general.CreateCookies(data["email"].(string), data["password"].(string))
	for _, cookie := range cookies {
		http.SetCookie(w, &cookie)
	}
	w.Write(general.CreateJsonResponse(0, "OK", nil))
}

func (handler *SigninHandler) validateData(data map[string]interface{}) (err error) {
	for _, key := range handler.validKeys {
		val, exists := data[key]
		if !exists || len(strings.TrimSpace(val.(string))) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}

	if (handler.db == nil) {
		return
	}

	user, _ := handler.db.GetUserInfoByEmail(data["email"].(string))

	if (user == db.UserT{}) {
		return fmt.Errorf("Пользователь не существует.")
	}

	if data["password"].(string) != user.Password {
		return fmt.Errorf("Пароли не совпадают.")
	}

	return
}
