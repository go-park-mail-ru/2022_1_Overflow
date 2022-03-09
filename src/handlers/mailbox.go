package handlers

import (
	"db"
	"encoding/json"
	"general"
	"net/http"
	"github.com/gorilla/mux"
)

type MailBox struct {
	db *db.DatabaseConnection
}

func (mb *MailBox) Init(r *mux.Router, db *db.DatabaseConnection) {
	r.HandleFunc("/list", mb.GetMailBox)
	mb.db = db
}

func (mb *MailBox) GetMailBox(w http.ResponseWriter, r *http.Request) {
	if !general.IsLoggedIn(r) {
		w.Write(general.CreateJsonResponse(1, "Пользователь не выполнил вход.", nil))
		return
	}
	cookie, _ := r.Cookie("email")
	user, err := mb.db.GetUserInfoByEmail(cookie.Value)
	if err != nil {
		w.Write(general.CreateJsonResponse(2, err.Error(), nil))
	}
	id := user.Id
	mails, err := mb.db.GetIncomeMails(id)
	if err != nil {
		w.Write(general.CreateJsonResponse(3, err.Error(), nil))
	}
	parsed, err := json.Marshal(mails)
	if err != nil {
		w.Write(general.CreateJsonResponse(4, err.Error(), nil))
	}
	w.Write(general.CreateJsonResponse(0, "OK", string(parsed)))
}