package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupCORS(r *mux.Router) http.Handler {
	options := cors.Options{
		AllowedOrigins:   []string{"http://localhost:80", "http://127.0.0.1:80", "http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.249.116:3000", "http://95.163.249.116:80", "http://overmail.online:3000", "http://overmail.online:80", "http://overmail.online"}, //"*"},
		AllowedHeaders:   []string{"Origin", "Version", "Authorization", "Accept", "Accept-Encoding", "Content-Type", "Content-Length", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}
	handler := cors.New(options).Handler(r)
	return handler
}
