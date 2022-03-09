package main

import (
	"general"
	"handlers"
	"net/http"
	"log"
)

var address string = ":8080"

func main() {
	mux := http.NewServeMux()
	var signin handlers.SigninHandler
	var signup handlers.SignupHandler

	signin.Init(mux)
	signup.Init(mux)
	
	log.Printf("Listening on %v", address)
	err := http.ListenAndServe(address, general.SetupCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}