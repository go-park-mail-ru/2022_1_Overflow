package mailbox

import (
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
)

func (mb *MailBox) Outcome(w http.ResponseWriter, r *http.Request) {
	if !mb.sm.IsLoggedIn(r) {
		pkg.AccessDenied(w)
		return
	}
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}
	var parsed []byte
	if mb.db != nil {
		data, err := mb.sm.GetData(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := mb.db.GetUserInfoByEmail(data.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id := user.Id
		mails, err := mb.db.GetOutcomeMails(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		parsed, err = json.Marshal(mails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}
