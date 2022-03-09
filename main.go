package main

import (
	"db"
	"general"
	"handlers"
	"net/http"
	"log"
)

var address string = ":8080"
var dbUrl string = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	mux := http.NewServeMux()
	var signin handlers.SigninHandler
	var signup handlers.SignupHandler

	var conn db.DatabaseConnection
	err := conn.Create(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	signin.Init(mux, &conn)
	signup.Init(mux, &conn)
	
	log.Printf("Listening on %v", address)
	err = http.ListenAndServe(address, general.SetupCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}