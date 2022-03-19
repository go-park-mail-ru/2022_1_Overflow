package middlewares

import "net/http"

func Recover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}