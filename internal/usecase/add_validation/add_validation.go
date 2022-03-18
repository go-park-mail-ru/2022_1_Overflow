package add_validation

import (
	"OverflowBackend/internal/models"
	"errors"

	"gopkg.in/validator.v2"
)

func SamePassword(data models.SignUpForm) error {
	if data.Password != data.PasswordConf {
		return errors.New("Пароли не совпадают.")
	}
	return nil
}

func CheckSignUp(data models.SignUpForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	if err := SamePassword(data); err != nil {
		return err
	}
	return nil
}

func CheckSignIn(data models.SignInForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	return nil
}