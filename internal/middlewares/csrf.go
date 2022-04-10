package middlewares

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
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
		lrw := NewLoggingResponseWriter(w)
		handler.ServeHTTP(lrw, r)
		if lrw.statusCode == http.StatusMethodNotAllowed {
			w.WriteHeader(http.StatusOK)
			//pkg.WriteJsonErrFull(w, pkg.NO_ERR)
		}
	})
}
