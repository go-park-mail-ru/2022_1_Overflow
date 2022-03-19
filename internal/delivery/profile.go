package delivery

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
)

// Получение данных профиля пользователя.
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
