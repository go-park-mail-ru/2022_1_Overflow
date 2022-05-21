package middlewares

import (
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	log "github.com/sirupsen/logrus"
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
		cok, err := r.Cookie("OverflowMail")
		if err != nil {
			log.Info(err)
		}
		log.Info(cok.Value)

		if !allowed && !session.Manager.IsLoggedIn(r) {
			log.Info("unauthorized")
			pkg.WriteJsonErrFull(w, &pkg.UNAUTHORIZED_ERR)
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
