package usecase

import (
	"OverflowBackend/internal/models"
	"encoding/json"
	"net/http"
)

func (uc *UseCase) Income(w http.ResponseWriter, r *http.Request, data *models.Session) (parsed []byte, err error) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		return
	}
	id := user.Id
	mails, err := uc.db.GetIncomeMails(id)
	if err != nil {
		return
	}
	parsed, err = json.Marshal(mails)
	if err != nil {
		return
	}
	return parsed, nil
}

func (uc *UseCase) Outcome(w http.ResponseWriter, r *http.Request, data *models.Session) {
	user, err := uc.db.GetUserInfoByEmail(data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := user.Id
	mails, err := uc.db.GetOutcomeMails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsed, err := json.Marshal(mails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}

