package middlewares

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"crypto/rand"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
)

var allowedPaths = []string{
	"/signin",
	"/signup",
	"/swagger",
	"/get_token",
}

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
		csrf.TrustedOrigins([]string{"localhost:3000", "95.163.249.116:3000"}),
		csrf.Path("/"),
	)
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

func CSRFWrapper(handler http.Handler) http.Handler {
	return csrfWrapper(handler)
}


func CreateSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
