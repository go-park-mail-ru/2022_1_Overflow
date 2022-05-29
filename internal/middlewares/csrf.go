package middlewares

import (
	"net/http"
)

func CSRFProtectWrapper(handler http.Handler) http.Handler {
	return CsrfWrapper(handler)
}
