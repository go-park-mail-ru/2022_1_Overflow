package middlewares

import (
	"log"
	"net/http"
)

func Recover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Fatal(err)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}