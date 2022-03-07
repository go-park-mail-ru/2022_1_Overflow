package main

import (
	"net/http"
	"log"
)

var address string = ":8080"

func main() {
	var handler SignupHandler
	handler.Init()
	
	log.Printf("Listening on %v", address)
	err := http.ListenAndServe(address, handler.Handlers())
	if err != nil {
		log.Fatal(err)
	}
}