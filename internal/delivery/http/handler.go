package http

import (
	"OverflowBackend/internal/config"

	"net/http"

	"github.com/gorilla/mux"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request) {
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	for k,v := range(handlers) {
		router.HandleFunc(k, v)
	}
	return config.SetupCORS(router)
}
