package middlewares

import (
	"net/http"
)

func CSRFProtectWrapper(handler http.Handler) http.Handler {
	return CsrfWrapper(handler)
}

func CSRFGetWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			token, _ := GetCSRFToken(r)
			w.Header().Set("X-CSRF-Token", token)
			w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token, X-Csrf-Token, X-Csrf-token")
		}
		handler.ServeHTTP(w, r)
	})
}
