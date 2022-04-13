package middlewares

import (
	"OverflowBackend/internal/config"
	"net/http"

	"github.com/gorilla/csrf"
)

var CsrfWrapper func(http.Handler) http.Handler

func Init(config *config.Config) {
	csrfKey := config.Server.Keys.CSRFAuthKey
	CsrfWrapper = csrf.Protect(
		[]byte(csrfKey),
		csrf.TrustedOrigins([]string{"localhost:3000", "95.163.249.116:3000", "localhost:8080", "95.163.249.116:8080"}),
		csrf.Path("/"),
		csrf.Secure(false),
	)
}

func Middleware(handler http.Handler) http.Handler {
	return Recover(CSRFProtectWrapper(CSRFGetWrapper(CreateSession(CheckLogin(handler)))))
}