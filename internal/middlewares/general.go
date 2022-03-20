package middlewares

import "net/http"

func Middleware(handler http.Handler) http.Handler {
	return Recover(CheckLogin(handler))
}