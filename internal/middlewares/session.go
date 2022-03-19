package middlewares

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
)


func CheckLogin(handler func(http.ResponseWriter, *http.Request), check bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if check && !session.IsLoggedIn(r) {
			pkg.AccessDenied(w)
			return
		}
		handler(w, r)
	}
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