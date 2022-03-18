package profile

import (
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
)

func (p *Profile) GetInfo(w http.ResponseWriter, r *http.Request) {
	if !p.sm.IsLoggedIn(r) {
		pkg.AccessDenied(w)
		return
	}
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	data, err := p.sm.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := p.db.GetUserInfoByEmail(data.Email)
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
