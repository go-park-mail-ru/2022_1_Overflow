package validation

import (
	"OverflowBackend/internal/models"
	"errors"

	"gopkg.in/validator.v2"
)

func SamePassword(data *models.SignUpForm) error {
	if data.Password != data.PasswordConfirmation {
		return errors.New("Пароли не совпадают.")
	}
	return nil
}

func SameFields(field1 string, field2 string) bool {
	return field1 == field2
}

func CheckSignUp(data *models.SignUpForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	if err := SamePassword(data); err != nil {
		return err
	}
	return nil
}

func CheckSignIn(data *models.SignInForm) error {
	if err := validator.Validate(data); err != nil {
		return err
	}
	return nil
}