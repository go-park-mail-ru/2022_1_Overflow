package validation

import (
	"OverflowBackend/proto/auth_proto"
	"errors"

	"gopkg.in/validator.v2"
)

func SamePassword(data *auth_proto.SignUpForm) error {
	if data.Password != data.PasswordConfirmation {
		return errors.New("Пароли не совпадают.")
	}
	return nil
}

func CheckSignUp(data *auth_proto.SignUpForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	if err := SamePassword(data); err != nil {
		return err
	}
	return nil
}

func CheckSignIn(data *auth_proto.SignInForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	return nil
}