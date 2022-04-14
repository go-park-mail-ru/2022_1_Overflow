package middlewares

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func CSRFProtectWrapper(handler http.Handler) http.Handler {
	return CsrfWrapper(handler)
}

func CSRFGetWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
			w.Header().Set("Access-Control-Expose-Headers", "*")
		}
		handler.ServeHTTP(w, r)
	})
}
