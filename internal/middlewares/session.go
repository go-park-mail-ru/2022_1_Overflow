package middlewares

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
	"strings"
)

var allowedPaths = []string{
	"/signin",
	"/signup",
	"/swagger",
}

func CheckLogin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := false
		for _, path := range allowedPaths {
			allowed = allowed || strings.Contains(r.URL.Path, path)
		}
		if !allowed && !session.IsLoggedIn(r) {
			pkg.WriteJsonErrFull(w, pkg.UNAUTHORIZED_ERR)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func CreateSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
