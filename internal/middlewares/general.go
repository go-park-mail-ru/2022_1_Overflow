package middlewares

import "net/http"

func Middleware(handler func(http.ResponseWriter, *http.Request), checkLogin bool) func(http.ResponseWriter, *http.Request) {
	return Recover(CheckLogin(handler, checkLogin))
}