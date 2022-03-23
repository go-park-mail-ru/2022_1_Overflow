package middlewares

import (
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return Recover(CreateSession(CheckLogin(CSRFWrapper(handler))))
}