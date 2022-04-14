package validation

import (
	"OverflowBackend/internal/models"
	"testing"
)

func TestSignUpVal(t *testing.T) {
	err := CheckSignUp(models.SignUpForm{
		FirstName:    "",
		LastName:     "test",
		Username:     "test",
		Password:     "test",
		PasswordConf: "test",
	})
	if err == nil {
		t.Errorf("Неверная валидация данных SignUp.")
		return
	}
	err = CheckSignUp(models.SignUpForm{
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

func TestSignInVal(t *testing.T) {
	err := CheckSignIn(models.SignInForm{
		Username: "bad@",
		Password: "good",
	})

	if err == nil {
		t.Errorf("Неверная валидация данных SignIn.")
		return
	}

	err = CheckSignIn(models.SignInForm{
		Username: "good",
		Password: "good",
	})

	if err != nil {
		t.Error(err)
		return
	}
}
