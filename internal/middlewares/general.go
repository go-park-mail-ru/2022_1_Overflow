package middlewares

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/attach_proto"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
)

var CsrfWrapper func(http.Handler) http.Handler
var csrfToken string
var csrfTokenTicker *time.Ticker

func GetCSRFToken(r *http.Request) (token string, isNew bool) {
	token = csrf.Token(r)
	return
	/*
		if len(csrfToken) == 0 {
			isNew = true
			token = csrf.Token(r)
			csrfToken = token
			return
		}
		select {
			case _, ok := <-csrfTokenTicker.C: {
				isNew = ok
				if ok {
					token = csrf.Token(r)
				} else {
					token = csrfToken
				}
			}
			default: {
				isNew = false
				token = csrfToken
			}
		}
		csrfToken = token
		return
	*/
}

func Init(config *config.Config, attachDial attach_proto.AttachClient) {
	attach = attachDial
	tokenTimeout := config.Server.Timeout.CSRFTimeout
	csrfTokenTicker = time.NewTicker(tokenTimeout)
	csrfKey := config.Server.Keys.CSRFAuthKey
	CsrfWrapper = csrf.Protect(
		[]byte(csrfKey),
		csrf.TrustedOrigins([]string{"localhost:3000", "95.163.249.116:3000", "localhost:8080", "95.163.249.116:8080"}),
		csrf.Path("/"),
		csrf.Secure(false),
	)
}

func Middleware(handler http.Handler) http.Handler {
	//return Recover(CreateSession(CheckLogin(handler)))
	return Recover(CSRFProtectWrapper(CreateSession(CheckLogin(handler))))
}
