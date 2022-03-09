package general

import (
	"net/http"
	"github.com/rs/cors"
)

func SetupCORS(mux *http.ServeMux) http.Handler {
	options := cors.Options{
		AllowedOrigins:   []string{"http://localhost:80", "http://127.0.0.1:80", "http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.249.116:3000", "http://95.163.249.116:80"}, //"*"},
		AllowedHeaders:   []string{"Origin", "Version", "Authorization", "Accept", "Accept-Encoding", "Content-Type", "Content-Length"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}
	handler := cors.New(options).Handler(mux)
	return handler
}