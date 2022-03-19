package middlewares

import "net/http"

func Recover(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
		}()
		handler(w, r)
	}
}