package middlewares

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

var csrfWrapper func(http.Handler) http.Handler

func init() {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	csrfKey := string(b)
	csrfWrapper = csrf.Protect(
		[]byte(csrfKey),
		csrf.TrustedOrigins([]string{"localhost:3000", "95.163.249.116:3000", "localhost:8080", "95.163.249.116:8080"}),
		csrf.Path("/"),
		csrf.Secure(false),
	)
}

func CSRFProtectWrapper(handler http.Handler) http.Handler {
	return csrfWrapper(handler)
}

func CSRFGetWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
		}
		handler.ServeHTTP(w, r)
	})
}
