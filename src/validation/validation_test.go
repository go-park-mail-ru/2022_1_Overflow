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

func TestBadEmail(t *testing.T) {
	err := CheckSignup(map[string]string{
		"first_name": "test",
		"last_name": "test",
		"email": "test@",
		"password": "test",
		"password_confirmation": "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}

func TestEmptyField(t *testing.T) {
	err := CheckSignup(map[string]string{
		"first_name": "",
		"last_name": "test",
		"email": "test",
		"password": "test",
		"password_confirmation": "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}