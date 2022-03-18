package pkg

import "net/http"

func MethodNotAllowed(w http.ResponseWriter, method string) {
	http.Error(w, "Only %v method is allowed.", http.StatusMethodNotAllowed)
}

func AccessDenied(w http.ResponseWriter) {
	http.Error(w, "Access denied.", http.StatusUnauthorized)
}