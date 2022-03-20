package delivery

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
)

// GetInfo godoc
// @Summary Получение данных профиля пользователя
// @Produce json
// @Success 200 {object} models.User "Информация о пользователе"
// @Failure 500
// @Router /profile [get]
func (d *Delivery) GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := d.uc.GetInfo(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}
