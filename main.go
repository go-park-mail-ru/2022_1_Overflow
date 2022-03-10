package main

import (
	"db"
	"general"
	"handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var address string = ":8080"
var dbUrl string = "postgres://postgres:123@localhost:5432/postgres"

func main() {
	r := mux.NewRouter()
	var signin handlers.SigninHandler
	var signup handlers.SignupHandler
	var mailbox handlers.MailBox

	var conn db.DatabaseConnection
	err := conn.Create(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	signin.Init(r, &conn)
	signup.Init(r, &conn)
	mailbox.Init(r, &conn)
	r.HandleFunc("/logout", handlers.LogoutHandler)

	log.Printf("Listening on %v", address)
	err = http.ListenAndServe(address, general.SetupCORS(r))
	if err != nil {
		log.Fatal(err)
	}
}
