package middlewares

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
)


func CheckLogin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/signin" && r.URL.Path != "/signup" && !session.IsLoggedIn(r) {
			pkg.AccessDenied(w)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

/*
func CreateSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		data, err := session.GetData(r)
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
		session.CreateSession(w, r, data.Email)
	})
}
*/