package validation

import (
	"OverflowBackend/proto/auth_proto"
	"testing"
)

func TestSignUpVal(t *testing.T) {
	err := CheckSignUp(&auth_proto.SignUpForm{
		FirstName:    "",
		LastName:     "test",
		Username:     "test",
		Password:     "test",
		PasswordConfirmation: "test",
	})
	if err == nil {
		t.Errorf("Неверная валидация данных SignUp.")
		return
	}
	err = CheckSignUp(&auth_proto.SignUpForm{
		FirstName:    "test",
		LastName:     "test",
		Username:     "test",
		Password:     "test",
		PasswordConfirmation: "test",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignInVal(t *testing.T) {
	err := CheckSignIn(&auth_proto.SignInForm{
		Username: "bad@",
		Password: "good",
	})

	if err == nil {
		t.Errorf("Неверная валидация данных SignIn.")
		return
	}

	err = CheckSignIn(&auth_proto.SignInForm{
		Username: "good",
		Password: "good",
	})

	if err != nil {
		t.Error(err)
		return
	}
}
