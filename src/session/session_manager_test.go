package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionManager(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	//w := httptest.NewRecorder()

	data, err := GetData(r)
	if (err != nil) {
		t.Error(err)
		return
	}
	if (data != Session{}) {
		t.Errorf("Данные сессии не являются пустыми.")
		return
	}
}