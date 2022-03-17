package validation

import (
	"fmt"
	"strings"
)

var MAX_FIELD_LEN int = 20

func EmailValidator(email string) error {
	if err := CheckEmptyField(email, "email"); err != nil {
		return fmt.Errorf("Поле адреса почты является пустым.")
	}
	if strings.Contains(email, "@") {
		return fmt.Errorf("Поле почты содержит символ @.")
	}
	return nil
}

func PasswordValidator(password string) error {
	if err := CheckEmptyField(password, "password"); err != nil {
		return fmt.Errorf("Поле пароля является пустым.")
	}
	return nil
}

func SamePasswordValidator(password, passwordConfirmation string) error {
	if err := CheckEmptyField(passwordConfirmation, "password_confirmation"); err != nil {
		return fmt.Errorf("Поле повтора пароля является пустым.")
	}
	if password != passwordConfirmation {
		return fmt.Errorf("Поля пароля и повтора пароля не совпадают.")
	}
	return nil
}

func CheckSignup(data map[string]string) error {
	for k, v := range data {
		if err := CheckEmptyField(v, k); err != nil {
			return err
		}
		if err := CheckMaxField(v, k, MAX_FIELD_LEN); err != nil {
			return err
		}
	}
	if err := EmailValidator(data["email"]); err != nil {
		return err
	}
	if err := PasswordValidator(data["password"]); err != nil {
		return err
	}
	if err := SamePasswordValidator(data["password"], data["password_confirmation"]); err != nil {
		return err
	}
	return nil
}

func CheckSignin(email, password string) error {
	if err := CheckEmptyField(email, "email"); err != nil {
		return err
	}
	if err := CheckEmptyField(password, "password"); err != nil {
		return err
	}
	if err := CheckMaxField(email, "email", MAX_FIELD_LEN); err != nil {
		return err
	}
	if err := CheckMaxField(password, "password", MAX_FIELD_LEN); err != nil {
		return err
	}
	if err := EmailValidator(email); err != nil {
		return err
	}
	if err := PasswordValidator(password); err != nil {
		return err
	}
	return nil
}

func CheckEmptyField(value string, key string) error {
	if len(strings.TrimSpace(value)) == 0 {
		if len(strings.TrimSpace(key)) == 0 {
			return fmt.Errorf("Поле является пустым.")
		} else {
			return fmt.Errorf("Поле %v является пустым.", key)
		}
	}
	return nil
}

func CheckMaxField(value string, key string, max int) error {
	if len(value) > max {
		if len(strings.TrimSpace(key)) == 0 {
			return fmt.Errorf("Поле превышает максимально допустимую длину (%v).", max)
		} else {
			return fmt.Errorf("Поле %v превышает максимально допустимую длину (%v).", max)
		}
	}
	return nil
}