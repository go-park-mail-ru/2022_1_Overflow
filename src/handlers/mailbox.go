package handlers

import (
	db "OverflowBackend/src/db"
	session "OverflowBackend/src/session"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MailBox struct {
	db *db.DatabaseConnection
}

func (mb *MailBox) Init(r *mux.Router, db *db.DatabaseConnection) {
	r.HandleFunc("/income", mb.getIncome)
	r.HandleFunc("/outcome", mb.getOutcome)
	mb.db = db
}

func (mb *MailBox) getIncome(w http.ResponseWriter, r *http.Request) {
	if !session.IsLoggedIn(r) {
		http.Error(w, "Access denied.", http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
		return
	}
	var parsed []byte
	if mb.db != nil {
		data, err := session.GetData(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := mb.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id := user.Id
		mails, err := mb.db.GetIncomeMails(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		parsed, err = json.Marshal(mails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}

func (mb *MailBox) getOutcome(w http.ResponseWriter, r *http.Request) {
	if !session.IsLoggedIn(r) {
		http.Error(w, "Access denied.", http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
		return
	}
	var parsed []byte
	if mb.db != nil {
		data, err := session.GetData(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := mb.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id := user.Id
		mails, err := mb.db.GetOutcomeMails(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		parsed, err = json.Marshal(mails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}
