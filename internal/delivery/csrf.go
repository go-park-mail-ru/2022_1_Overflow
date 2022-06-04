package delivery

import (
	"OverflowBackend/pkg"
	"net/http"

	"github.com/gorilla/csrf"
)

// @Router /get_token [get]
// @Response 200 {object} pkg.JsonResponse
// @Summary Получить CSRF-Token
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags security
func (d *Delivery) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	token := csrf.Token(r)
	w.Header().Set("X-CSRF-Token", token)
	w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token, X-Csrf-Token, X-Csrf-token")
	pkg.WriteJsonErrFull(w, pkg.CreateJsonErr(pkg.STATUS_OK, token))
}