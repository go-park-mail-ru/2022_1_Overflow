package pkg

import (
	"fmt"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, method string) {
	http.Error(w, fmt.Sprintf("Only %v method is allowed.", method), http.StatusMethodNotAllowed)
}

func AccessDenied(w http.ResponseWriter) {
	http.Error(w, "Access denied.", http.StatusUnauthorized)
}