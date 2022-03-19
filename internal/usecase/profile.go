package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
	"net/http"
)

func (uc *UseCase) GetInfo(w http.ResponseWriter, r *http.Request, data *models.Session) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}