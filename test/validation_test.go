package test

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/validation"
	"testing"
)

func TestValidation(t *testing.T) {
	err := validation.CheckSignUp(models.SignUpForm{
		FirstName: "test",
		LastName: "test",
		Email: "test",
		Password: "test",
		PasswordConf: "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadEmail(t *testing.T) {
	err := validation.CheckSignUp(models.SignUpForm{
		FirstName: "test",
		LastName: "test",
		Email: "test@",
		Password: "test",
		PasswordConf: "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}

func TestEmptyField(t *testing.T) {
	err := validation.CheckSignUp(models.SignUpForm{
		FirstName: "",
		LastName: "test",
		Email: "test",
		Password: "test",
		PasswordConf: "test",
	})
	if err == nil {
		t.Error(err)
		return
	}
}