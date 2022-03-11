package handlers

import (
	session "OverflowBackend/src/session"
	db "OverflowBackend/src/db"

	"net/http"
	"github.com/gorilla/mux"
	//"encoding/json"
)

type Transfer struct {
	db *db.DatabaseConnection
}

func (t *Transfer) Init(router *mux.Router, db *db.DatabaseConnection) {
	router.HandleFunc("/send", t.sendEmail)
	t.db = db
}

func (t *Transfer) sendEmail(w http.ResponseWriter, r *http.Request) {
	if !session.IsLoggedIn(r) {
		http.Error(w, "Access denied.", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
		return
	}

	//todo
}