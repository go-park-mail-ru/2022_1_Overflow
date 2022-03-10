package validation

import (
	"testing"
)

func TestValidation(t *testing.T) {
	err := CheckSignup(map[string]string{
		"first_name": "test",
		"last_name": "test",
		"email": "test",
		"password": "test",
		"password_confirmation": "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
}