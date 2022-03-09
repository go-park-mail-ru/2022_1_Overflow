package handlers

import (
	"db"
	"general"
	"fmt"
	"net/http"
	"strings"
)

type SigninHandler struct {
	validKeys []string
}

func (handler *SigninHandler) Init(mux *http.ServeMux) {
	handler.validKeys = []string{"email", "password"}
	mux.HandleFunc("/signin", handler.userSignin)
}

func (handler *SigninHandler) userSignin(w http.ResponseWriter, r *http.Request) {
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
	w.Write(general.CreateJsonResponse(0, "OK", nil))
}

func (handler *SigninHandler) validateData(r *http.Request) (err error) {
	for _, key := range handler.validKeys {
		val := r.FormValue(key)
		if len(strings.TrimSpace(val)) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}

	/*

	user, _ := db.GetUserInfoByEmail(r.FormValue("email"))

	if r.FormValue("password") != user.Password {
		return fmt.Errorf("Пароли не совпадают.")
	}
	*/

	return
}
