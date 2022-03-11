package handlers

import (
	session "OverflowBackend/src/session"
	db "OverflowBackend/src/db"

	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Profile struct {
	db *db.DatabaseConnection
}

func (p *Profile) Init(router *mux.Router, db *db.DatabaseConnection) {
	router.HandleFunc("/profile", p.profileHandler)
	p.db = db
}

func (p *Profile) profileHandler(w http.ResponseWriter, r *http.Request) {
	if !session.IsLoggedIn(r) {
		http.Error(w, "Access denied.", http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := p.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}