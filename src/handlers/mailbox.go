package handlers

import (
	db "OverflowBackend/src/db"
	response "OverflowBackend/src/response"
	session "OverflowBackend/src/session"

	"encoding/json"
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
	if !session.IsLoggedIn(r) {
		w.Write(response.CreateJsonResponse(1, "Пользователь не выполнил вход.", nil))
		return
	}
	var parsed []byte
	if mb.db != nil {
		data, err := session.GetData(r)
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := mb.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			w.Write(response.CreateJsonResponse(2, err.Error(), nil))
		}
		id := user.Id
		mails, err := mb.db.GetIncomeMails(id)
		if err != nil {
			w.Write(response.CreateJsonResponse(3, err.Error(), nil))
		}
		parsed, err = json.Marshal(mails)
		if err != nil {
			w.Write(response.CreateJsonResponse(4, err.Error(), nil))
		}
	}
	w.Write(response.CreateJsonResponse(0, "OK", string(parsed)))
}
