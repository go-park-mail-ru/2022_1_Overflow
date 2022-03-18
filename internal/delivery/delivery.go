package delivery

import (
	"net/http"
)

type Delivery interface {
	handlers map[string]func(http.ResponseWriter, *http.Request)
	NewRouter() http.Handler
}