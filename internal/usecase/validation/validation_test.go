package validation

import (
	"OverflowBackend/internal/models"
	"testing"
)

func TestValidation(t *testing.T) {
	err := CheckSignUp(models.SignUpForm{
		FirstName:    "test",
		LastName:     "test",
		Username:     "test",
		Password:     "test",
		PasswordConf: "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

/*
func TestBadEmail(t *testing.T) {
	err := validation.CheckSignUp(models.SignUpForm{
		FirstName: "test",
		LastName: "test",
		Username: "test@",
		Password: "test",
		PasswordConf: "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}
*/

func TestEmptyField(t *testing.T) {
	err := CheckSignUp(models.SignUpForm{
		FirstName:    "",
		LastName:     "test",
		Username:     "test",
		Password:     "test",
		PasswordConf: "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}
